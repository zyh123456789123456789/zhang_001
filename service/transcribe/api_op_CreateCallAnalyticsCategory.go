// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Call Analytics category. All categories are automatically applied
// to your Call Analytics transcriptions. Note that in order to apply categories to
// your transcriptions, you must create them before submitting your transcription
// request, as categories cannot be applied retroactively. When creating a new
// category, you can use the InputType parameter to label the category as a
// POST_CALL or a REAL_TIME category. POST_CALL categories can only be applied to
// post-call transcriptions and REAL_TIME categories can only be applied to
// real-time transcriptions. If you do not include InputType , your category is
// created as a POST_CALL category by default. Call Analytics categories are
// composed of rules. For each category, you must create between 1 and 20 rules.
// Rules can include these parameters: , , , and . To update an existing category,
// see . To learn more about Call Analytics categories, see Creating categories
// for post-call transcriptions (https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-batch.html)
// and Creating categories for real-time transcriptions (https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-stream.html)
// .
func (c *Client) CreateCallAnalyticsCategory(ctx context.Context, params *CreateCallAnalyticsCategoryInput, optFns ...func(*Options)) (*CreateCallAnalyticsCategoryOutput, error) {
	if params == nil {
		params = &CreateCallAnalyticsCategoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCallAnalyticsCategory", params, optFns, c.addOperationCreateCallAnalyticsCategoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCallAnalyticsCategoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCallAnalyticsCategoryInput struct {

	// A unique name, chosen by you, for your Call Analytics category. It's helpful to
	// use a detailed naming system that will make sense to you in the future. For
	// example, it's better to use sentiment-positive-last30seconds for a category
	// over a generic name like test-category . Category names are case sensitive.
	//
	// This member is required.
	CategoryName *string

	// Rules define a Call Analytics category. When creating a new category, you must
	// create between 1 and 20 rules for that category. For each rule, you specify a
	// filter you want applied to the attributes of a call. For example, you can choose
	// a sentiment filter that detects if a customer's sentiment was positive during
	// the last 30 seconds of the call.
	//
	// This member is required.
	Rules []types.Rule

	// Choose whether you want to create a real-time or a post-call category for your
	// Call Analytics transcription. Specifying POST_CALL assigns your category to
	// post-call transcriptions; categories with this input type cannot be applied to
	// streaming (real-time) transcriptions. Specifying REAL_TIME assigns your
	// category to streaming transcriptions; categories with this input type cannot be
	// applied to post-call transcriptions. If you do not include InputType , your
	// category is created as a post-call category by default.
	InputType types.InputType

	noSmithyDocumentSerde
}

type CreateCallAnalyticsCategoryOutput struct {

	// Provides you with the properties of your new category, including its associated
	// rules.
	CategoryProperties *types.CategoryProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCallAnalyticsCategoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCallAnalyticsCategory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCallAnalyticsCategory{}, middleware.After)
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
	if err = addCreateCallAnalyticsCategoryResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCallAnalyticsCategoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCallAnalyticsCategory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCallAnalyticsCategory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "CreateCallAnalyticsCategory",
	}
}

type opCreateCallAnalyticsCategoryResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateCallAnalyticsCategoryResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateCallAnalyticsCategoryResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "transcribe"
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
				signingName = "transcribe"
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
				v4aScheme.SigningName = aws.String("transcribe")
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

func addCreateCallAnalyticsCategoryResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateCallAnalyticsCategoryResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
