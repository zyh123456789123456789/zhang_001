// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// You start a zonal shift to temporarily move load balancer traffic away from an
// Availability Zone in a AWS Region, to help your application recover immediately,
// for example, from a developer's bad code deployment or from an AWS
// infrastructure failure in a single Availability Zone. You can start a zonal
// shift in Route 53 ARC only for managed resources in your account in an AWS
// Region. Resources are automatically registered with Route 53 ARC by AWS
// services. At this time, you can only start a zonal shift for Network Load
// Balancers and Application Load Balancers with cross-zone load balancing turned
// off. When you start a zonal shift, traffic for the resource is no longer routed
// to the Availability Zone. The zonal shift is created immediately in Route 53
// ARC. However, it can take a short time, typically up to a few minutes, for
// existing, in-progress connections in the Availability Zone to complete. For more
// information, see Zonal shift (https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.html)
// in the Amazon Route 53 Application Recovery Controller Developer Guide.
func (c *Client) StartZonalShift(ctx context.Context, params *StartZonalShiftInput, optFns ...func(*Options)) (*StartZonalShiftOutput, error) {
	if params == nil {
		params = &StartZonalShiftInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartZonalShift", params, optFns, c.addOperationStartZonalShiftMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartZonalShiftOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartZonalShiftInput struct {

	// The Availability Zone that traffic is moved away from for a resource when you
	// start a zonal shift. Until the zonal shift expires or you cancel it, traffic for
	// the resource is instead moved to other Availability Zones in the AWS Region.
	//
	// This member is required.
	AwayFrom *string

	// A comment that you enter about the zonal shift. Only the latest comment is
	// retained; no comment history is maintained. A new comment overwrites any
	// existing comment string.
	//
	// This member is required.
	Comment *string

	// The length of time that you want a zonal shift to be active, which Route 53 ARC
	// converts to an expiry time (expiration time). Zonal shifts are temporary. You
	// can set a zonal shift to be active initially for up to three days (72 hours). If
	// you want to still keep traffic away from an Availability Zone, you can update
	// the zonal shift and set a new expiration. You can also cancel a zonal shift,
	// before it expires, for example, if you're ready to restore traffic to the
	// Availability Zone. To set a length of time for a zonal shift to be active,
	// specify a whole number, and then one of the following, with no space:
	//   - A lowercase letter m: To specify that the value is in minutes.
	//   - A lowercase letter h: To specify that the value is in hours.
	// For example: 20h means the zonal shift expires in 20 hours. 120m means the
	// zonal shift expires in 120 minutes (2 hours).
	//
	// This member is required.
	ExpiresIn *string

	// The identifier for the resource to include in a zonal shift. The identifier is
	// the Amazon Resource Name (ARN) for the resource. At this time, you can only
	// start a zonal shift for Network Load Balancers and Application Load Balancers
	// with cross-zone load balancing turned off.
	//
	// This member is required.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type StartZonalShiftOutput struct {

	// The Availability Zone that traffic is moved away from for a resource when you
	// start a zonal shift. Until the zonal shift expires or you cancel it, traffic for
	// the resource is instead moved to other Availability Zones in the AWS Region.
	//
	// This member is required.
	AwayFrom *string

	// A comment that you enter about the zonal shift. Only the latest comment is
	// retained; no comment history is maintained. A new comment overwrites any
	// existing comment string.
	//
	// This member is required.
	Comment *string

	// The expiry time (expiration time) for the zonal shift. A zonal shift is
	// temporary and must be set to expire when you start the zonal shift. You can
	// initially set a zonal shift to expire in a maximum of three days (72 hours).
	// However, you can update a zonal shift to set a new expiration at any time. When
	// you start a zonal shift, you specify how long you want it to be active, which
	// Route 53 ARC converts to an expiry time (expiration time). You can cancel a
	// zonal shift, for example, if you're ready to restore traffic to the Availability
	// Zone. Or you can update the zonal shift to specify another length of time to
	// expire in.
	//
	// This member is required.
	ExpiryTime *time.Time

	// The identifier for the resource to include in a zonal shift. The identifier is
	// the Amazon Resource Name (ARN) for the resource. At this time, you can only
	// start a zonal shift for Network Load Balancers and Application Load Balancers
	// with cross-zone load balancing turned off.
	//
	// This member is required.
	ResourceIdentifier *string

	// The time (UTC) when the zonal shift is started.
	//
	// This member is required.
	StartTime *time.Time

	// A status for a zonal shift. The Status for a zonal shift can have one of the
	// following values:
	//   - ACTIVE: The zonal shift is started and active.
	//   - EXPIRED: The zonal shift has expired (the expiry time was exceeded).
	//   - CANCELED: The zonal shift was canceled.
	//
	// This member is required.
	Status types.ZonalShiftStatus

	// The identifier of a zonal shift.
	//
	// This member is required.
	ZonalShiftId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartZonalShiftMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartZonalShift{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartZonalShift{}, middleware.After)
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
	if err = addStartZonalShiftResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartZonalShiftValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartZonalShift(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartZonalShift(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "arc-zonal-shift",
		OperationName: "StartZonalShift",
	}
}

type opStartZonalShiftResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartZonalShiftResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartZonalShiftResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "arc-zonal-shift"
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
				signingName = "arc-zonal-shift"
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
				v4aScheme.SigningName = aws.String("arc-zonal-shift")
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

func addStartZonalShiftResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartZonalShiftResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
