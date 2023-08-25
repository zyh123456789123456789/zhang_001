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

// Starts the asynchronous analysis of invoices or receipts for data like contact
// information, items purchased, and vendor names. StartExpenseAnalysis can
// analyze text in documents that are in JPEG, PNG, and PDF format. The documents
// must be stored in an Amazon S3 bucket. Use the DocumentLocation parameter to
// specify the name of your S3 bucket and the name of the document in that bucket.
// StartExpenseAnalysis returns a job identifier ( JobId ) that you will provide to
// GetExpenseAnalysis to retrieve the results of the operation. When the analysis
// of the input invoices/receipts is finished, Amazon Textract publishes a
// completion status to the Amazon Simple Notification Service (Amazon SNS) topic
// that you provide to the NotificationChannel . To obtain the results of the
// invoice and receipt analysis operation, ensure that the status value published
// to the Amazon SNS topic is SUCCEEDED . If so, call GetExpenseAnalysis , and pass
// the job identifier ( JobId ) that was returned by your call to
// StartExpenseAnalysis . For more information, see Analyzing Invoices and Receipts (https://docs.aws.amazon.com/textract/latest/dg/invoice-receipts.html)
// .
func (c *Client) StartExpenseAnalysis(ctx context.Context, params *StartExpenseAnalysisInput, optFns ...func(*Options)) (*StartExpenseAnalysisOutput, error) {
	if params == nil {
		params = &StartExpenseAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartExpenseAnalysis", params, optFns, c.addOperationStartExpenseAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartExpenseAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartExpenseAnalysisInput struct {

	// The location of the document to be processed.
	//
	// This member is required.
	DocumentLocation *types.DocumentLocation

	// The idempotent token that's used to identify the start request. If you use the
	// same token with multiple StartDocumentTextDetection requests, the same JobId is
	// returned. Use ClientRequestToken to prevent the same job from being
	// accidentally started more than once. For more information, see Calling Amazon
	// Textract Asynchronous Operations (https://docs.aws.amazon.com/textract/latest/dg/api-async.html)
	ClientRequestToken *string

	// An identifier you specify that's included in the completion notification
	// published to the Amazon SNS topic. For example, you can use JobTag to identify
	// the type of document that the completion notification corresponds to (such as a
	// tax form or a receipt).
	JobTag *string

	// The KMS key used to encrypt the inference results. This can be in either Key ID
	// or Key Alias format. When a KMS key is provided, the KMS key will be used for
	// server-side encryption of the objects in the customer bucket. When this
	// parameter is not enabled, the result will be encrypted server side,using SSE-S3.
	KMSKeyId *string

	// The Amazon SNS topic ARN that you want Amazon Textract to publish the
	// completion status of the operation to.
	NotificationChannel *types.NotificationChannel

	// Sets if the output will go to a customer defined bucket. By default, Amazon
	// Textract will save the results internally to be accessed by the
	// GetExpenseAnalysis operation.
	OutputConfig *types.OutputConfig

	noSmithyDocumentSerde
}

type StartExpenseAnalysisOutput struct {

	// A unique identifier for the text detection job. The JobId is returned from
	// StartExpenseAnalysis . A JobId value is only valid for 7 days.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartExpenseAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartExpenseAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartExpenseAnalysis{}, middleware.After)
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
	if err = addStartExpenseAnalysisResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartExpenseAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartExpenseAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartExpenseAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "StartExpenseAnalysis",
	}
}

type opStartExpenseAnalysisResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartExpenseAnalysisResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartExpenseAnalysisResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addStartExpenseAnalysisResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartExpenseAnalysisResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}