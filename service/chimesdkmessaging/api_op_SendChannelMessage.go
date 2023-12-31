// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a message to a particular channel that the member is a part of. The
// x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header. Also, STANDARD messages can be up to 4KB in size and contain metadata.
// Metadata is arbitrary, and you can use it in a variety of ways, such as
// containing a link to an attachment. CONTROL messages are limited to 30 bytes
// and do not contain metadata.
func (c *Client) SendChannelMessage(ctx context.Context, params *SendChannelMessageInput, optFns ...func(*Options)) (*SendChannelMessageOutput, error) {
	if params == nil {
		params = &SendChannelMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendChannelMessage", params, optFns, c.addOperationSendChannelMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendChannelMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendChannelMessageInput struct {

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the AppInstanceUser or AppInstanceBot that makes the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The Idempotency token for each client request.
	//
	// This member is required.
	ClientRequestToken *string

	// The content of the channel message.
	//
	// This member is required.
	Content *string

	// Boolean that controls whether the message is persisted on the back end.
	// Required.
	//
	// This member is required.
	Persistence types.ChannelMessagePersistenceType

	// The type of message, STANDARD or CONTROL . STANDARD messages can be up to 4KB
	// in size and contain metadata. Metadata is arbitrary, and you can use it in a
	// variety of ways, such as containing a link to an attachment. CONTROL messages
	// are limited to 30 bytes and do not contain metadata.
	//
	// This member is required.
	Type types.ChannelMessageType

	// The content type of the channel message.
	ContentType *string

	// The attributes for the message, used for message filtering along with a
	// FilterRule defined in the PushNotificationPreferences .
	MessageAttributes map[string]types.MessageAttributeValue

	// The optional metadata for each message.
	Metadata *string

	// The push notification configuration of the message.
	PushNotification *types.PushNotificationConfiguration

	// The ID of the SubChannel in the request.
	SubChannelId *string

	// The target of a message. Must be a member of the channel, such as another user,
	// a bot, or the sender. Only the target and the sender can view targeted messages.
	// Only users who can see targeted messages can take actions on them. However,
	// administrators can delete targeted messages that they can’t see.
	Target []types.Target

	noSmithyDocumentSerde
}

type SendChannelMessageOutput struct {

	// The ARN of the channel.
	ChannelArn *string

	// The ID string assigned to each message.
	MessageId *string

	// The status of the channel message.
	Status *types.ChannelMessageStatusStructure

	// The ID of the SubChannel in the response.
	SubChannelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendChannelMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendChannelMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendChannelMessage{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSendChannelMessageResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opSendChannelMessageMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSendChannelMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendChannelMessage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpSendChannelMessage struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendChannelMessage) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendChannelMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendChannelMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendChannelMessageInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSendChannelMessageMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpSendChannelMessage{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendChannelMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "SendChannelMessage",
	}
}

type opSendChannelMessageResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opSendChannelMessageResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSendChannelMessageResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "chime"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "chime"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("chime")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addSendChannelMessageResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opSendChannelMessageResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
