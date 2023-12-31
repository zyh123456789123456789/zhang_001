// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the specified limit for a specified hosted zone, for example, the maximum
// number of records that you can create in the hosted zone. For the default limit,
// see Limits (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. To request a higher limit, open a case (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53)
// .
func (c *Client) GetHostedZoneLimit(ctx context.Context, params *GetHostedZoneLimitInput, optFns ...func(*Options)) (*GetHostedZoneLimitOutput, error) {
	if params == nil {
		params = &GetHostedZoneLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHostedZoneLimit", params, optFns, c.addOperationGetHostedZoneLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHostedZoneLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to create a hosted
// zone.
type GetHostedZoneLimitInput struct {

	// The ID of the hosted zone that you want to get a limit for.
	//
	// This member is required.
	HostedZoneId *string

	// The limit that you want to get. Valid values include the following:
	//   - MAX_RRSETS_BY_ZONE: The maximum number of records that you can create in
	//   the specified hosted zone.
	//   - MAX_VPCS_ASSOCIATED_BY_ZONE: The maximum number of Amazon VPCs that you can
	//   associate with the specified private hosted zone.
	//
	// This member is required.
	Type types.HostedZoneLimitType

	noSmithyDocumentSerde
}

// A complex type that contains the requested limit.
type GetHostedZoneLimitOutput struct {

	// The current number of entities that you have created of the specified type. For
	// example, if you specified MAX_RRSETS_BY_ZONE for the value of Type in the
	// request, the value of Count is the current number of records that you have
	// created in the specified hosted zone.
	//
	// This member is required.
	Count int64

	// The current setting for the specified limit. For example, if you specified
	// MAX_RRSETS_BY_ZONE for the value of Type in the request, the value of Limit is
	// the maximum number of records that you can create in the specified hosted zone.
	//
	// This member is required.
	Limit *types.HostedZoneLimit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetHostedZoneLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetHostedZoneLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetHostedZoneLimit{}, middleware.After)
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
	if err = addGetHostedZoneLimitResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetHostedZoneLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHostedZoneLimit(options.Region), middleware.Before); err != nil {
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
	if err = addSanitizeURLMiddleware(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetHostedZoneLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHostedZoneLimit",
	}
}

type opGetHostedZoneLimitResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetHostedZoneLimitResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetHostedZoneLimitResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "route53"
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
				signingName = "route53"
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
				v4aScheme.SigningName = aws.String("route53")
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

func addGetHostedZoneLimitResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetHostedZoneLimitResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
