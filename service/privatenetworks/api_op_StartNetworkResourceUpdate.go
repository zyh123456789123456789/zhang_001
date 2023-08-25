// Code generated by smithy-go-codegen DO NOT EDIT.

package privatenetworks

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/privatenetworks/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this action to do the following tasks:
//   - Update the duration and renewal status of the commitment period for a radio
//     unit. The update goes into effect immediately.
//   - Request a replacement for a network resource.
//   - Request that you return a network resource.
//
// After you submit a request to replace or return a network resource, the status
// of the network resource changes to CREATING_SHIPPING_LABEL . The shipping label
// is available when the status of the network resource is PENDING_RETURN . After
// the network resource is successfully returned, its status changes to DELETED .
// For more information, see Return a radio unit (https://docs.aws.amazon.com/private-networks/latest/userguide/radio-units.html#return-radio-unit)
// .
func (c *Client) StartNetworkResourceUpdate(ctx context.Context, params *StartNetworkResourceUpdateInput, optFns ...func(*Options)) (*StartNetworkResourceUpdateOutput, error) {
	if params == nil {
		params = &StartNetworkResourceUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartNetworkResourceUpdate", params, optFns, c.addOperationStartNetworkResourceUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartNetworkResourceUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartNetworkResourceUpdateInput struct {

	// The Amazon Resource Name (ARN) of the network resource.
	//
	// This member is required.
	NetworkResourceArn *string

	// The update type.
	//   - REPLACE - Submits a request to replace a defective radio unit. We provide a
	//   shipping label that you can use for the return process and we ship a replacement
	//   radio unit to you.
	//   - RETURN - Submits a request to return a radio unit that you no longer need.
	//   We provide a shipping label that you can use for the return process.
	//   - COMMITMENT - Submits a request to change or renew the commitment period. If
	//   you choose this value, then you must set commitmentConfiguration (https://docs.aws.amazon.com/private-networks/latest/APIReference/API_StartNetworkResourceUpdate.html#privatenetworks-StartNetworkResourceUpdate-request-commitmentConfiguration)
	//   .
	//
	// This member is required.
	UpdateType types.UpdateType

	// Use this action to extend and automatically renew the commitment period for the
	// radio unit. You can do the following:
	//   - Change a 60-day commitment to a 1-year or 3-year commitment. The change is
	//   immediate and the hourly rate decreases to the rate for the new commitment
	//   period.
	//   - Change a 1-year commitment to a 3-year commitment. The change is immediate
	//   and the hourly rate decreases to the rate for the 3-year commitment period.
	//   - Set a 1-year commitment to automatically renew for an additional 1 year.
	//   The hourly rate for the additional year will continue to be the same as your
	//   existing 1-year rate.
	//   - Set a 3-year commitment to automatically renew for an additional 1 year.
	//   The hourly rate for the additional year will continue to be the same as your
	//   existing 3-year rate.
	//   - Turn off a previously-enabled automatic renewal on a 1-year or 3-year
	//   commitment. You cannot use the automatic-renewal option for a 60-day commitment.
	//
	// For pricing, see Amazon Web Services Private 5G Pricing (http://aws.amazon.com/private5g/pricing)
	// .
	CommitmentConfiguration *types.CommitmentConfiguration

	// The reason for the return. Providing a reason for a return is optional.
	ReturnReason *string

	// The shipping address. If you don't provide a shipping address when replacing or
	// returning a network resource, we use the address from the original order for the
	// network resource.
	ShippingAddress *types.Address

	noSmithyDocumentSerde
}

type StartNetworkResourceUpdateOutput struct {

	// The network resource.
	NetworkResource *types.NetworkResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartNetworkResourceUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartNetworkResourceUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartNetworkResourceUpdate{}, middleware.After)
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
	if err = addStartNetworkResourceUpdateResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartNetworkResourceUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartNetworkResourceUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartNetworkResourceUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "private-networks",
		OperationName: "StartNetworkResourceUpdate",
	}
}

type opStartNetworkResourceUpdateResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartNetworkResourceUpdateResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartNetworkResourceUpdateResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "private-networks"
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
				signingName = "private-networks"
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
				v4aScheme.SigningName = aws.String("private-networks")
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

func addStartNetworkResourceUpdateResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartNetworkResourceUpdateResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}