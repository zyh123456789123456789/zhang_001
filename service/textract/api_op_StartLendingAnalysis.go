// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the classification and analysis of an input document.
// StartLendingAnalysis initiates the classification and analysis of a packet of
// lending documents. StartLendingAnalysis operates on a document file located in
// an Amazon S3 bucket. StartLendingAnalysis can analyze text in documents that
// are in one of the following formats: JPEG, PNG, TIFF, PDF. Use DocumentLocation
// to specify the bucket name and the file name of the document.
// StartLendingAnalysis returns a job identifier ( JobId ) that you use to get the
// results of the operation. When the text analysis is finished, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that you specify in NotificationChannel . To get the results of the
// text analysis operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED. If the status is SUCCEEDED you can call either
// GetLendingAnalysis or GetLendingAnalysisSummary and provide the JobId to obtain
// the results of the analysis. If using OutputConfig to specify an Amazon S3
// bucket, the output will be contained within the specified prefix in a directory
// labeled with the job-id. In the directory there are 3 sub-directories:
//   - detailedResponse (contains the GetLendingAnalysis response)
//   - summaryResponse (for the GetLendingAnalysisSummary response)
//   - splitDocuments (documents split across logical boundaries)
func (c *Client) StartLendingAnalysis(ctx context.Context, params *StartLendingAnalysisInput, optFns ...func(*Options)) (*StartLendingAnalysisOutput, error) {
	if params == nil {
		params = &StartLendingAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartLendingAnalysis", params, optFns, c.addOperationStartLendingAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartLendingAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartLendingAnalysisInput struct {

	// The Amazon S3 bucket that contains the document to be processed. It's used by
	// asynchronous operations. The input document can be an image file in JPEG or PNG
	// format. It can also be a file in PDF format.
	//
	// This member is required.
	DocumentLocation *types.DocumentLocation

	// The idempotent token that you use to identify the start request. If you use the
	// same token with multiple StartLendingAnalysis requests, the same JobId is
	// returned. Use ClientRequestToken to prevent the same job from being
	// accidentally started more than once. For more information, see Calling Amazon
	// Textract Asynchronous Operations (https://docs.aws.amazon.com/textract/latest/dg/api-sync.html)
	// .
	ClientRequestToken *string

	// An identifier that you specify to be included in the completion notification
	// published to the Amazon SNS topic. For example, you can use JobTag to identify
	// the type of document that the completion notification corresponds to (such as a
	// tax form or a receipt).
	JobTag *string

	// The KMS key used to encrypt the inference results. This can be in either Key ID
	// or Key Alias format. When a KMS key is provided, the KMS key will be used for
	// server-side encryption of the objects in the customer bucket. When this
	// parameter is not enabled, the result will be encrypted server side, using
	// SSE-S3.
	KMSKeyId *string

	// The Amazon Simple Notification Service (Amazon SNS) topic to which Amazon
	// Textract publishes the completion status of an asynchronous document operation.
	NotificationChannel *types.NotificationChannel

	// Sets whether or not your output will go to a user created bucket. Used to set
	// the name of the bucket, and the prefix on the output file. OutputConfig is an
	// optional parameter which lets you adjust where your output will be placed. By
	// default, Amazon Textract will store the results internally and can only be
	// accessed by the Get API operations. With OutputConfig enabled, you can set the
	// name of the bucket the output will be sent to the file prefix of the results
	// where you can download your results. Additionally, you can set the KMSKeyID
	// parameter to a customer master key (CMK) to encrypt your output. Without this
	// parameter set Amazon Textract will encrypt server-side using the AWS managed CMK
	// for Amazon S3. Decryption of Customer Content is necessary for processing of the
	// documents by Amazon Textract. If your account is opted out under an AI services
	// opt out policy then all unencrypted Customer Content is immediately and
	// permanently deleted after the Customer Content has been processed by the
	// service. No copy of of the output is retained by Amazon Textract. For
	// information about how to opt out, see Managing AI services opt-out policy.  (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	// For more information on data privacy, see the Data Privacy FAQ (https://aws.amazon.com/compliance/data-privacy-faq/)
	// .
	OutputConfig *types.OutputConfig

	noSmithyDocumentSerde
}

type StartLendingAnalysisOutput struct {

	// A unique identifier for the lending or text-detection job. The JobId is
	// returned from StartLendingAnalysis . A JobId value is only valid for 7 days.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartLendingAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartLendingAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartLendingAnalysis{}, middleware.After)
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
	if err = addStartLendingAnalysisResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartLendingAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartLendingAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartLendingAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "StartLendingAnalysis",
	}
}

type opStartLendingAnalysisResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartLendingAnalysisResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartLendingAnalysisResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "textract"
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
				signingName = "textract"
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
				v4aScheme.SigningName = aws.String("textract")
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

func addStartLendingAnalysisResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartLendingAnalysisResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
