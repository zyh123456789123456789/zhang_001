// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a specific domain.
func (c *Client) GetDomain(ctx context.Context, params *GetDomainInput, optFns ...func(*Options)) (*GetDomainOutput, error) {
	if params == nil {
		params = &GetDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomain", params, optFns, c.addOperationGetDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDomainInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

type GetDomainOutput struct {

	// The timestamp of when the domain was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The timestamp of when the domain was most recently edited.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The URL of the SQS dead letter queue, which is used for reporting errors
	// associated with ingesting data from third party applications.
	DeadLetterQueueUrl *string

	// The default encryption key, which is an AWS managed key, is used when no
	// specific type of encryption key is specified. It is used to encrypt all data
	// before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey *string

	// The default number of days until the data within the domain expires.
	DefaultExpirationDays *int32

	// The process of matching duplicate profiles. If Matching = true , Amazon Connect
	// Customer Profiles starts a weekly batch process called Identity Resolution Job.
	// If you do not specify a date and time for Identity Resolution Job to run, by
	// default it runs every Saturday at 12AM UTC to detect duplicate profiles in your
	// domains. After the Identity Resolution Job completes, use the GetMatches (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html)
	// API to return and review the results. Or, if you have configured ExportingConfig
	// in the MatchingRequest , you can download the results from S3.
	Matching *types.MatchingResponse

	// The process of matching duplicate profiles using the Rule-Based matching. If
	// RuleBasedMatching = true, Amazon Connect Customer Profiles will start to match
	// and merge your profiles according to your configuration in the
	// RuleBasedMatchingRequest . You can use the ListRuleBasedMatches and
	// GetSimilarProfiles API to return and review the results. Also, if you have
	// configured ExportingConfig in the RuleBasedMatchingRequest , you can download
	// the results from S3.
	RuleBasedMatching *types.RuleBasedMatchingResponse

	// Usage-specific statistics about the domain.
	Stats *types.DomainStats

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomain{}, middleware.After)
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
	if err = addGetDomainResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "GetDomain",
	}
}

type opGetDomainResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetDomainResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetDomainResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "profile"
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
				signingName = "profile"
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
				v4aScheme.SigningName = aws.String("profile")
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

func addGetDomainResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetDomainResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}