// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribestreaming

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribestreaming/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"sync"
	"time"
)

// AudioStreamWriter provides the interface for writing events to a stream.
//
// The writer's Close method must allow multiple concurrent calls.
type AudioStreamWriter interface {
	Send(context.Context, types.AudioStream) error
	Close() error
	Err() error
}

// CallAnalyticsTranscriptResultStreamReader provides the interface for reading
// events from a stream.
//
// The writer's Close method must allow multiple concurrent calls.
type CallAnalyticsTranscriptResultStreamReader interface {
	Events() <-chan types.CallAnalyticsTranscriptResultStream
	Close() error
	Err() error
}

// MedicalTranscriptResultStreamReader provides the interface for reading events
// from a stream.
//
// The writer's Close method must allow multiple concurrent calls.
type MedicalTranscriptResultStreamReader interface {
	Events() <-chan types.MedicalTranscriptResultStream
	Close() error
	Err() error
}

// TranscriptResultStreamReader provides the interface for reading events from a
// stream.
//
// The writer's Close method must allow multiple concurrent calls.
type TranscriptResultStreamReader interface {
	Events() <-chan types.TranscriptResultStream
	Close() error
	Err() error
}

type eventStreamSigner interface {
	GetSignature(ctx context.Context, headers, payload []byte, signingTime time.Time, optFns ...func(*v4.StreamSignerOptions)) ([]byte, error)
}

type asyncAudioStream struct {
	Event  types.AudioStream
	Result chan<- error
}

func (e asyncAudioStream) ReportResult(cancel <-chan struct{}, err error) bool {
	select {
	case e.Result <- err:
		return true

	case <-cancel:
		return false

	}
}

type audioStreamWriter struct {
	encoder             *eventstream.Encoder
	signer              eventStreamSigner
	stream              chan asyncAudioStream
	serializationBuffer *bytes.Buffer
	signingBuffer       *bytes.Buffer
	eventStream         io.WriteCloser
	done                chan struct{}
	closeOnce           sync.Once
	err                 *smithysync.OnceErr
}

func newAudioStreamWriter(stream io.WriteCloser, encoder *eventstream.Encoder, signer eventStreamSigner) *audioStreamWriter {
	w := &audioStreamWriter{
		encoder:             encoder,
		signer:              signer,
		stream:              make(chan asyncAudioStream),
		eventStream:         stream,
		done:                make(chan struct{}),
		err:                 smithysync.NewOnceErr(),
		serializationBuffer: bytes.NewBuffer(nil),
		signingBuffer:       bytes.NewBuffer(nil),
	}

	go w.writeStream()

	return w

}

func (w *audioStreamWriter) Send(ctx context.Context, event types.AudioStream) error {
	return w.send(ctx, event)
}

func (w *audioStreamWriter) send(ctx context.Context, event types.AudioStream) error {
	if err := w.err.Err(); err != nil {
		return err
	}

	resultCh := make(chan error)

	wrapped := asyncAudioStream{
		Event:  event,
		Result: resultCh,
	}

	select {
	case w.stream <- wrapped:
	case <-ctx.Done():
		return ctx.Err()
	case <-w.done:
		return fmt.Errorf("stream closed, unable to send event")

	}

	select {
	case err := <-resultCh:
		return err
	case <-ctx.Done():
		return ctx.Err()
	case <-w.done:
		return fmt.Errorf("stream closed, unable to send event")

	}

}

func (w *audioStreamWriter) writeStream() {
	defer w.Close()

	for {
		select {
		case wrapper := <-w.stream:
			err := w.writeEvent(wrapper.Event)
			wrapper.ReportResult(w.done, err)
			if err != nil {
				w.err.SetError(err)
				return
			}

		case <-w.done:
			if err := w.closeStream(); err != nil {
				w.err.SetError(err)
			}
			return

		}
	}
}

func (w *audioStreamWriter) writeEvent(event types.AudioStream) error {
	// serializedEvent returned bytes refers to an underlying byte buffer and must not
	// escape this writeEvent scope without first copying. Any previous bytes stored in
	// the buffer are cleared by this call.
	serializedEvent, err := w.serializeEvent(event)
	if err != nil {
		return err
	}

	// signedEvent returned bytes refers to an underlying byte buffer and must not
	// escape this writeEvent scope without first copying. Any previous bytes stored in
	// the buffer are cleared by this call.
	signedEvent, err := w.signEvent(serializedEvent)
	if err != nil {
		return err
	}

	// bytes are now copied to the underlying stream writer
	_, err = io.Copy(w.eventStream, bytes.NewReader(signedEvent))
	return err
}

func (w *audioStreamWriter) serializeEvent(event types.AudioStream) ([]byte, error) {
	w.serializationBuffer.Reset()

	eventMessage := eventstream.Message{}

	if err := awsRestjson1_serializeEventStreamAudioStream(event, &eventMessage); err != nil {
		return nil, err
	}

	if err := w.encoder.Encode(w.serializationBuffer, eventMessage); err != nil {
		return nil, err
	}

	return w.serializationBuffer.Bytes(), nil
}

func (w *audioStreamWriter) signEvent(payload []byte) ([]byte, error) {
	w.signingBuffer.Reset()

	date := time.Now().UTC()

	var msg eventstream.Message
	msg.Headers.Set(eventstreamapi.DateHeader, eventstream.TimestampValue(date))
	msg.Payload = payload

	var headers bytes.Buffer
	if err := eventstream.EncodeHeaders(&headers, msg.Headers); err != nil {
		return nil, err
	}

	sig, err := w.signer.GetSignature(context.Background(), headers.Bytes(), msg.Payload, date)
	if err != nil {
		return nil, err
	}

	msg.Headers.Set(eventstreamapi.ChunkSignatureHeader, eventstream.BytesValue(sig))

	if err := w.encoder.Encode(w.signingBuffer, msg); err != nil {
		return nil, err
	}

	return w.signingBuffer.Bytes(), nil
}

func (w *audioStreamWriter) closeStream() (err error) {
	defer func() {
		if cErr := w.eventStream.Close(); cErr != nil && err == nil {
			err = cErr
		}
	}()

	// Per the protocol, a signed empty message is used to indicate the end of the stream,
	// and that no subsequent events will be sent.
	signedEvent, err := w.signEvent([]byte{})
	if err != nil {
		return err
	}

	_, err = io.Copy(w.eventStream, bytes.NewReader(signedEvent))
	return err
}

func (w *audioStreamWriter) ErrorSet() <-chan struct{} {
	return w.err.ErrorSet()
}

func (w *audioStreamWriter) Close() error {
	w.closeOnce.Do(w.safeClose)
	return w.Err()
}

func (w *audioStreamWriter) safeClose() {
	close(w.done)
}

func (w *audioStreamWriter) Err() error {
	return w.err.Err()
}

type callAnalyticsTranscriptResultStreamReader struct {
	stream      chan types.CallAnalyticsTranscriptResultStream
	decoder     *eventstream.Decoder
	eventStream io.ReadCloser
	err         *smithysync.OnceErr
	payloadBuf  []byte
	done        chan struct{}
	closeOnce   sync.Once
}

func newCallAnalyticsTranscriptResultStreamReader(readCloser io.ReadCloser, decoder *eventstream.Decoder) *callAnalyticsTranscriptResultStreamReader {
	w := &callAnalyticsTranscriptResultStreamReader{
		stream:      make(chan types.CallAnalyticsTranscriptResultStream),
		decoder:     decoder,
		eventStream: readCloser,
		err:         smithysync.NewOnceErr(),
		done:        make(chan struct{}),
		payloadBuf:  make([]byte, 10*1024),
	}

	go w.readEventStream()

	return w
}

func (r *callAnalyticsTranscriptResultStreamReader) Events() <-chan types.CallAnalyticsTranscriptResultStream {
	return r.stream
}

func (r *callAnalyticsTranscriptResultStreamReader) readEventStream() {
	defer r.Close()
	defer close(r.stream)

	for {
		r.payloadBuf = r.payloadBuf[0:0]
		decodedMessage, err := r.decoder.Decode(r.eventStream, r.payloadBuf)
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case <-r.done:
				return
			default:
				r.err.SetError(err)
				return
			}
		}

		event, err := r.deserializeEventMessage(&decodedMessage)
		if err != nil {
			r.err.SetError(err)
			return
		}

		select {
		case r.stream <- event:
		case <-r.done:
			return
		}

	}
}

func (r *callAnalyticsTranscriptResultStreamReader) deserializeEventMessage(msg *eventstream.Message) (types.CallAnalyticsTranscriptResultStream, error) {
	messageType := msg.Headers.Get(eventstreamapi.MessageTypeHeader)
	if messageType == nil {
		return nil, fmt.Errorf("%s event header not present", eventstreamapi.MessageTypeHeader)
	}

	switch messageType.String() {
	case eventstreamapi.EventMessageType:
		var v types.CallAnalyticsTranscriptResultStream
		if err := awsRestjson1_deserializeEventStreamCallAnalyticsTranscriptResultStream(&v, msg); err != nil {
			return nil, err
		}
		return v, nil

	case eventstreamapi.ExceptionMessageType:
		return nil, awsRestjson1_deserializeEventStreamExceptionCallAnalyticsTranscriptResultStream(msg)

	case eventstreamapi.ErrorMessageType:
		errorCode := "UnknownError"
		errorMessage := errorCode
		if header := msg.Headers.Get(eventstreamapi.ErrorCodeHeader); header != nil {
			errorCode = header.String()
		}
		if header := msg.Headers.Get(eventstreamapi.ErrorMessageHeader); header != nil {
			errorMessage = header.String()
		}
		return nil, &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}

	default:
		mc := msg.Clone()
		return nil, &UnknownEventMessageError{
			Type:    messageType.String(),
			Message: &mc,
		}

	}
}

func (r *callAnalyticsTranscriptResultStreamReader) ErrorSet() <-chan struct{} {
	return r.err.ErrorSet()
}

func (r *callAnalyticsTranscriptResultStreamReader) Close() error {
	r.closeOnce.Do(r.safeClose)
	return r.Err()
}

func (r *callAnalyticsTranscriptResultStreamReader) safeClose() {
	close(r.done)
	r.eventStream.Close()

}

func (r *callAnalyticsTranscriptResultStreamReader) Err() error {
	return r.err.Err()
}

func (r *callAnalyticsTranscriptResultStreamReader) Closed() <-chan struct{} {
	return r.done
}

type medicalTranscriptResultStreamReader struct {
	stream      chan types.MedicalTranscriptResultStream
	decoder     *eventstream.Decoder
	eventStream io.ReadCloser
	err         *smithysync.OnceErr
	payloadBuf  []byte
	done        chan struct{}
	closeOnce   sync.Once
}

func newMedicalTranscriptResultStreamReader(readCloser io.ReadCloser, decoder *eventstream.Decoder) *medicalTranscriptResultStreamReader {
	w := &medicalTranscriptResultStreamReader{
		stream:      make(chan types.MedicalTranscriptResultStream),
		decoder:     decoder,
		eventStream: readCloser,
		err:         smithysync.NewOnceErr(),
		done:        make(chan struct{}),
		payloadBuf:  make([]byte, 10*1024),
	}

	go w.readEventStream()

	return w
}

func (r *medicalTranscriptResultStreamReader) Events() <-chan types.MedicalTranscriptResultStream {
	return r.stream
}

func (r *medicalTranscriptResultStreamReader) readEventStream() {
	defer r.Close()
	defer close(r.stream)

	for {
		r.payloadBuf = r.payloadBuf[0:0]
		decodedMessage, err := r.decoder.Decode(r.eventStream, r.payloadBuf)
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case <-r.done:
				return
			default:
				r.err.SetError(err)
				return
			}
		}

		event, err := r.deserializeEventMessage(&decodedMessage)
		if err != nil {
			r.err.SetError(err)
			return
		}

		select {
		case r.stream <- event:
		case <-r.done:
			return
		}

	}
}

func (r *medicalTranscriptResultStreamReader) deserializeEventMessage(msg *eventstream.Message) (types.MedicalTranscriptResultStream, error) {
	messageType := msg.Headers.Get(eventstreamapi.MessageTypeHeader)
	if messageType == nil {
		return nil, fmt.Errorf("%s event header not present", eventstreamapi.MessageTypeHeader)
	}

	switch messageType.String() {
	case eventstreamapi.EventMessageType:
		var v types.MedicalTranscriptResultStream
		if err := awsRestjson1_deserializeEventStreamMedicalTranscriptResultStream(&v, msg); err != nil {
			return nil, err
		}
		return v, nil

	case eventstreamapi.ExceptionMessageType:
		return nil, awsRestjson1_deserializeEventStreamExceptionMedicalTranscriptResultStream(msg)

	case eventstreamapi.ErrorMessageType:
		errorCode := "UnknownError"
		errorMessage := errorCode
		if header := msg.Headers.Get(eventstreamapi.ErrorCodeHeader); header != nil {
			errorCode = header.String()
		}
		if header := msg.Headers.Get(eventstreamapi.ErrorMessageHeader); header != nil {
			errorMessage = header.String()
		}
		return nil, &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}

	default:
		mc := msg.Clone()
		return nil, &UnknownEventMessageError{
			Type:    messageType.String(),
			Message: &mc,
		}

	}
}

func (r *medicalTranscriptResultStreamReader) ErrorSet() <-chan struct{} {
	return r.err.ErrorSet()
}

func (r *medicalTranscriptResultStreamReader) Close() error {
	r.closeOnce.Do(r.safeClose)
	return r.Err()
}

func (r *medicalTranscriptResultStreamReader) safeClose() {
	close(r.done)
	r.eventStream.Close()

}

func (r *medicalTranscriptResultStreamReader) Err() error {
	return r.err.Err()
}

func (r *medicalTranscriptResultStreamReader) Closed() <-chan struct{} {
	return r.done
}

type transcriptResultStreamReader struct {
	stream      chan types.TranscriptResultStream
	decoder     *eventstream.Decoder
	eventStream io.ReadCloser
	err         *smithysync.OnceErr
	payloadBuf  []byte
	done        chan struct{}
	closeOnce   sync.Once
}

func newTranscriptResultStreamReader(readCloser io.ReadCloser, decoder *eventstream.Decoder) *transcriptResultStreamReader {
	w := &transcriptResultStreamReader{
		stream:      make(chan types.TranscriptResultStream),
		decoder:     decoder,
		eventStream: readCloser,
		err:         smithysync.NewOnceErr(),
		done:        make(chan struct{}),
		payloadBuf:  make([]byte, 10*1024),
	}

	go w.readEventStream()

	return w
}

func (r *transcriptResultStreamReader) Events() <-chan types.TranscriptResultStream {
	return r.stream
}

func (r *transcriptResultStreamReader) readEventStream() {
	defer r.Close()
	defer close(r.stream)

	for {
		r.payloadBuf = r.payloadBuf[0:0]
		decodedMessage, err := r.decoder.Decode(r.eventStream, r.payloadBuf)
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case <-r.done:
				return
			default:
				r.err.SetError(err)
				return
			}
		}

		event, err := r.deserializeEventMessage(&decodedMessage)
		if err != nil {
			r.err.SetError(err)
			return
		}

		select {
		case r.stream <- event:
		case <-r.done:
			return
		}

	}
}

func (r *transcriptResultStreamReader) deserializeEventMessage(msg *eventstream.Message) (types.TranscriptResultStream, error) {
	messageType := msg.Headers.Get(eventstreamapi.MessageTypeHeader)
	if messageType == nil {
		return nil, fmt.Errorf("%s event header not present", eventstreamapi.MessageTypeHeader)
	}

	switch messageType.String() {
	case eventstreamapi.EventMessageType:
		var v types.TranscriptResultStream
		if err := awsRestjson1_deserializeEventStreamTranscriptResultStream(&v, msg); err != nil {
			return nil, err
		}
		return v, nil

	case eventstreamapi.ExceptionMessageType:
		return nil, awsRestjson1_deserializeEventStreamExceptionTranscriptResultStream(msg)

	case eventstreamapi.ErrorMessageType:
		errorCode := "UnknownError"
		errorMessage := errorCode
		if header := msg.Headers.Get(eventstreamapi.ErrorCodeHeader); header != nil {
			errorCode = header.String()
		}
		if header := msg.Headers.Get(eventstreamapi.ErrorMessageHeader); header != nil {
			errorMessage = header.String()
		}
		return nil, &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}

	default:
		mc := msg.Clone()
		return nil, &UnknownEventMessageError{
			Type:    messageType.String(),
			Message: &mc,
		}

	}
}

func (r *transcriptResultStreamReader) ErrorSet() <-chan struct{} {
	return r.err.ErrorSet()
}

func (r *transcriptResultStreamReader) Close() error {
	r.closeOnce.Do(r.safeClose)
	return r.Err()
}

func (r *transcriptResultStreamReader) safeClose() {
	close(r.done)
	r.eventStream.Close()

}

func (r *transcriptResultStreamReader) Err() error {
	return r.err.Err()
}

func (r *transcriptResultStreamReader) Closed() <-chan struct{} {
	return r.done
}

type awsRestjson1_deserializeOpEventStreamStartCallAnalyticsStreamTranscription struct {
	LogEventStreamWrites bool
	LogEventStreamReads  bool
}

func (*awsRestjson1_deserializeOpEventStreamStartCallAnalyticsStreamTranscription) ID() string {
	return "OperationEventStreamDeserializer"
}

func (m *awsRestjson1_deserializeOpEventStreamStartCallAnalyticsStreamTranscription) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	defer func() {
		if err == nil {
			return
		}
		m.closeResponseBody(out)
	}()

	logger := middleware.GetLogger(ctx)

	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", in.Request)
	}
	_ = request

	if err := eventstreamapi.ApplyHTTPTransportFixes(request); err != nil {
		return out, metadata, err
	}

	requestSignature, err := v4.GetSignedRequestSignature(request.Request)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to get event stream seed signature: %v", err)
	}

	signer := v4.NewStreamSigner(
		awsmiddleware.GetSigningCredentials(ctx),
		awsmiddleware.GetSigningName(ctx),
		awsmiddleware.GetSigningRegion(ctx),
		requestSignature,
	)

	eventWriter := newAudioStreamWriter(
		eventstreamapi.GetInputStreamWriter(ctx),
		eventstream.NewEncoder(func(options *eventstream.EncoderOptions) {
			options.Logger = logger
			options.LogMessages = m.LogEventStreamWrites

		}),
		signer,
	)
	defer func() {
		if err == nil {
			return
		}
		_ = eventWriter.Close()
	}()

	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	deserializeOutput, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", out.RawResponse)
	}
	_ = deserializeOutput

	output, ok := out.Result.(*StartCallAnalyticsStreamTranscriptionOutput)
	if out.Result != nil && !ok {
		return out, metadata, fmt.Errorf("unexpected output result type: %T", out.Result)
	} else if out.Result == nil {
		output = &StartCallAnalyticsStreamTranscriptionOutput{}
		out.Result = output
	}

	eventReader := newCallAnalyticsTranscriptResultStreamReader(
		deserializeOutput.Body,
		eventstream.NewDecoder(func(options *eventstream.DecoderOptions) {
			options.Logger = logger
			options.LogMessages = m.LogEventStreamReads

		}),
	)
	defer func() {
		if err == nil {
			return
		}
		_ = eventReader.Close()
	}()

	output.eventStream = NewStartCallAnalyticsStreamTranscriptionEventStream(func(stream *StartCallAnalyticsStreamTranscriptionEventStream) {
		stream.Writer = eventWriter
		stream.Reader = eventReader
	})

	go output.eventStream.waitStreamClose()

	return out, metadata, nil
}

func (*awsRestjson1_deserializeOpEventStreamStartCallAnalyticsStreamTranscription) closeResponseBody(out middleware.DeserializeOutput) {
	if resp, ok := out.RawResponse.(*smithyhttp.Response); ok && resp != nil && resp.Body != nil {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}

func addEventStreamStartCallAnalyticsStreamTranscriptionMiddleware(stack *middleware.Stack, options Options) error {
	if err := stack.Deserialize.Insert(&awsRestjson1_deserializeOpEventStreamStartCallAnalyticsStreamTranscription{
		LogEventStreamWrites: options.ClientLogMode.IsRequestEventMessage(),
		LogEventStreamReads:  options.ClientLogMode.IsResponseEventMessage(),
	}, "OperationDeserializer", middleware.Before); err != nil {
		return err
	}
	return nil

}

type awsRestjson1_deserializeOpEventStreamStartMedicalStreamTranscription struct {
	LogEventStreamWrites bool
	LogEventStreamReads  bool
}

func (*awsRestjson1_deserializeOpEventStreamStartMedicalStreamTranscription) ID() string {
	return "OperationEventStreamDeserializer"
}

func (m *awsRestjson1_deserializeOpEventStreamStartMedicalStreamTranscription) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	defer func() {
		if err == nil {
			return
		}
		m.closeResponseBody(out)
	}()

	logger := middleware.GetLogger(ctx)

	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", in.Request)
	}
	_ = request

	if err := eventstreamapi.ApplyHTTPTransportFixes(request); err != nil {
		return out, metadata, err
	}

	requestSignature, err := v4.GetSignedRequestSignature(request.Request)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to get event stream seed signature: %v", err)
	}

	signer := v4.NewStreamSigner(
		awsmiddleware.GetSigningCredentials(ctx),
		awsmiddleware.GetSigningName(ctx),
		awsmiddleware.GetSigningRegion(ctx),
		requestSignature,
	)

	eventWriter := newAudioStreamWriter(
		eventstreamapi.GetInputStreamWriter(ctx),
		eventstream.NewEncoder(func(options *eventstream.EncoderOptions) {
			options.Logger = logger
			options.LogMessages = m.LogEventStreamWrites

		}),
		signer,
	)
	defer func() {
		if err == nil {
			return
		}
		_ = eventWriter.Close()
	}()

	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	deserializeOutput, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", out.RawResponse)
	}
	_ = deserializeOutput

	output, ok := out.Result.(*StartMedicalStreamTranscriptionOutput)
	if out.Result != nil && !ok {
		return out, metadata, fmt.Errorf("unexpected output result type: %T", out.Result)
	} else if out.Result == nil {
		output = &StartMedicalStreamTranscriptionOutput{}
		out.Result = output
	}

	eventReader := newMedicalTranscriptResultStreamReader(
		deserializeOutput.Body,
		eventstream.NewDecoder(func(options *eventstream.DecoderOptions) {
			options.Logger = logger
			options.LogMessages = m.LogEventStreamReads

		}),
	)
	defer func() {
		if err == nil {
			return
		}
		_ = eventReader.Close()
	}()

	output.eventStream = NewStartMedicalStreamTranscriptionEventStream(func(stream *StartMedicalStreamTranscriptionEventStream) {
		stream.Writer = eventWriter
		stream.Reader = eventReader
	})

	go output.eventStream.waitStreamClose()

	return out, metadata, nil
}

func (*awsRestjson1_deserializeOpEventStreamStartMedicalStreamTranscription) closeResponseBody(out middleware.DeserializeOutput) {
	if resp, ok := out.RawResponse.(*smithyhttp.Response); ok && resp != nil && resp.Body != nil {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}

func addEventStreamStartMedicalStreamTranscriptionMiddleware(stack *middleware.Stack, options Options) error {
	if err := stack.Deserialize.Insert(&awsRestjson1_deserializeOpEventStreamStartMedicalStreamTranscription{
		LogEventStreamWrites: options.ClientLogMode.IsRequestEventMessage(),
		LogEventStreamReads:  options.ClientLogMode.IsResponseEventMessage(),
	}, "OperationDeserializer", middleware.Before); err != nil {
		return err
	}
	return nil

}

type awsRestjson1_deserializeOpEventStreamStartStreamTranscription struct {
	LogEventStreamWrites bool
	LogEventStreamReads  bool
}

func (*awsRestjson1_deserializeOpEventStreamStartStreamTranscription) ID() string {
	return "OperationEventStreamDeserializer"
}

func (m *awsRestjson1_deserializeOpEventStreamStartStreamTranscription) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	defer func() {
		if err == nil {
			return
		}
		m.closeResponseBody(out)
	}()

	logger := middleware.GetLogger(ctx)

	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", in.Request)
	}
	_ = request

	if err := eventstreamapi.ApplyHTTPTransportFixes(request); err != nil {
		return out, metadata, err
	}

	requestSignature, err := v4.GetSignedRequestSignature(request.Request)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to get event stream seed signature: %v", err)
	}

	signer := v4.NewStreamSigner(
		awsmiddleware.GetSigningCredentials(ctx),
		awsmiddleware.GetSigningName(ctx),
		awsmiddleware.GetSigningRegion(ctx),
		requestSignature,
	)

	eventWriter := newAudioStreamWriter(
		eventstreamapi.GetInputStreamWriter(ctx),
		eventstream.NewEncoder(func(options *eventstream.EncoderOptions) {
			options.Logger = logger
			options.LogMessages = m.LogEventStreamWrites

		}),
		signer,
	)
	defer func() {
		if err == nil {
			return
		}
		_ = eventWriter.Close()
	}()

	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	deserializeOutput, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", out.RawResponse)
	}
	_ = deserializeOutput

	output, ok := out.Result.(*StartStreamTranscriptionOutput)
	if out.Result != nil && !ok {
		return out, metadata, fmt.Errorf("unexpected output result type: %T", out.Result)
	} else if out.Result == nil {
		output = &StartStreamTranscriptionOutput{}
		out.Result = output
	}

	eventReader := newTranscriptResultStreamReader(
		deserializeOutput.Body,
		eventstream.NewDecoder(func(options *eventstream.DecoderOptions) {
			options.Logger = logger
			options.LogMessages = m.LogEventStreamReads

		}),
	)
	defer func() {
		if err == nil {
			return
		}
		_ = eventReader.Close()
	}()

	output.eventStream = NewStartStreamTranscriptionEventStream(func(stream *StartStreamTranscriptionEventStream) {
		stream.Writer = eventWriter
		stream.Reader = eventReader
	})

	go output.eventStream.waitStreamClose()

	return out, metadata, nil
}

func (*awsRestjson1_deserializeOpEventStreamStartStreamTranscription) closeResponseBody(out middleware.DeserializeOutput) {
	if resp, ok := out.RawResponse.(*smithyhttp.Response); ok && resp != nil && resp.Body != nil {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}

func addEventStreamStartStreamTranscriptionMiddleware(stack *middleware.Stack, options Options) error {
	if err := stack.Deserialize.Insert(&awsRestjson1_deserializeOpEventStreamStartStreamTranscription{
		LogEventStreamWrites: options.ClientLogMode.IsRequestEventMessage(),
		LogEventStreamReads:  options.ClientLogMode.IsResponseEventMessage(),
	}, "OperationDeserializer", middleware.Before); err != nil {
		return err
	}
	return nil

}

// UnknownEventMessageError provides an error when a message is received from the stream,
// but the reader is unable to determine what kind of message it is.
type UnknownEventMessageError struct {
	Type    string
	Message *eventstream.Message
}

// Error retruns the error message string.
func (e *UnknownEventMessageError) Error() string {
	return "unknown event stream message type, " + e.Type
}

func setSafeEventStreamClientLogMode(o *Options, operation string) {
	switch operation {
	case "StartCallAnalyticsStreamTranscription":
		toggleEventStreamClientLogMode(o, true, true)
		return

	case "StartMedicalStreamTranscription":
		toggleEventStreamClientLogMode(o, true, true)
		return

	case "StartStreamTranscription":
		toggleEventStreamClientLogMode(o, true, true)
		return

	default:
		return

	}
}
func toggleEventStreamClientLogMode(o *Options, request, response bool) {
	mode := o.ClientLogMode

	if request && mode.IsRequestWithBody() {
		mode.ClearRequestWithBody()
		mode |= aws.LogRequest
	}

	if response && mode.IsResponseWithBody() {
		mode.ClearResponseWithBody()
		mode |= aws.LogResponse
	}

	o.ClientLogMode = mode

}
