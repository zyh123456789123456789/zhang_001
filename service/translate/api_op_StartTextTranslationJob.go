// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous batch translation job. Use batch translation jobs to
// translate large volumes of text across multiple documents at once. For batch
// translation, you can input documents with different source languages (specify
// auto as the source language). You can specify one or more target languages.
// Batch translation translates each input document into each of the target
// languages. For more information, see Asynchronous batch processing (https://docs.aws.amazon.com/translate/latest/dg/async.html)
// . Batch translation jobs can be described with the DescribeTextTranslationJob
// operation, listed with the ListTextTranslationJobs operation, and stopped with
// the StopTextTranslationJob operation.
func (c *Client) StartTextTranslationJob(ctx context.Context, params *StartTextTranslationJobInput, optFns ...func(*Options)) (*StartTextTranslationJobOutput, error) {
	if params == nil {
		params = &StartTextTranslationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTextTranslationJob", params, optFns, c.addOperationStartTextTranslationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTextTranslationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTextTranslationJobInput struct {

	// A unique identifier for the request. This token is generated for you when using
	// the Amazon Translate SDK.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of an AWS Identity Access and Management (IAM)
	// role that grants Amazon Translate read access to your input data. For more
	// information, see Identity and access management  (https://docs.aws.amazon.com/translate/latest/dg/identity-and-access-management.html)
	// .
	//
	// This member is required.
	DataAccessRoleArn *string

	// Specifies the format and location of the input documents for the translation
	// job.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// Specifies the S3 folder to which your job output will be saved.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// The language code of the input language. Specify the language if all input
	// documents share the same language. If you don't know the language of the source
	// files, or your input documents contains different source languages, select auto
	// . Amazon Translate auto detects the source language for each input document. For
	// a list of supported language codes, see Supported languages (https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html)
	// .
	//
	// This member is required.
	SourceLanguageCode *string

	// The target languages of the translation job. Enter up to 10 language codes.
	// Each input file is translated into each target language. Each language code is 2
	// or 5 characters long. For a list of language codes, see Supported languages (https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html)
	// .
	//
	// This member is required.
	TargetLanguageCodes []string

	// The name of the batch translation job to be performed.
	JobName *string

	// The name of a parallel data resource to add to the translation job. This
	// resource consists of examples that show how you want segments of text to be
	// translated. If you specify multiple target languages for the job, the parallel
	// data file must include translations for all the target languages. When you add
	// parallel data to a translation job, you create an Active Custom Translation job.
	// This parameter accepts only one parallel data resource. Active Custom
	// Translation jobs are priced at a higher rate than other jobs that don't use
	// parallel data. For more information, see Amazon Translate pricing (http://aws.amazon.com/translate/pricing/)
	// . For a list of available parallel data resources, use the ListParallelData
	// operation. For more information, see Customizing your translations with
	// parallel data (https://docs.aws.amazon.com/translate/latest/dg/customizing-translations-parallel-data.html)
	// .
	ParallelDataNames []string

	// Settings to configure your translation output, including the option to set the
	// formality level of the output text and the option to mask profane words and
	// phrases.
	Settings *types.TranslationSettings

	// The name of a custom terminology resource to add to the translation job. This
	// resource lists examples source terms and the desired translation for each term.
	// This parameter accepts only one custom terminology resource. If you specify
	// multiple target languages for the job, translate uses the designated terminology
	// for each requested target language that has an entry for the source term in the
	// terminology file. For a list of available custom terminology resources, use the
	// ListTerminologies operation. For more information, see Custom terminology (https://docs.aws.amazon.com/translate/latest/dg/how-custom-terminology.html)
	// .
	TerminologyNames []string

	noSmithyDocumentSerde
}

type StartTextTranslationJobOutput struct {

	// The identifier generated for the job. To get the status of a job, use this ID
	// with the DescribeTextTranslationJob operation.
	JobId *string

	// The status of the job. Possible values include:
	//   - SUBMITTED - The job has been received and is queued for processing.
	//   - IN_PROGRESS - Amazon Translate is processing the job.
	//   - COMPLETED - The job was successfully completed and the output is available.
	//   - COMPLETED_WITH_ERROR - The job was completed with errors. The errors can be
	//   analyzed in the job's output.
	//   - FAILED - The job did not complete. To get details, use the
	//   DescribeTextTranslationJob operation.
	//   - STOP_REQUESTED - The user who started the job has requested that it be
	//   stopped.
	//   - STOPPED - The job has been stopped.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartTextTranslationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTextTranslationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTextTranslationJob{}, middleware.After)
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
	if err = addStartTextTranslationJobResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opStartTextTranslationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartTextTranslationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTextTranslationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartTextTranslationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartTextTranslationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartTextTranslationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartTextTranslationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartTextTranslationJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartTextTranslationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartTextTranslationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartTextTranslationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "StartTextTranslationJob",
	}
}

type opStartTextTranslationJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartTextTranslationJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartTextTranslationJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "translate"
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
				signingName = "translate"
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
				v4aScheme.SigningName = aws.String("translate")
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

func addStartTextTranslationJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartTextTranslationJobResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}