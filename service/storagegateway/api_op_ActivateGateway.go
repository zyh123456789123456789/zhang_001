// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Activates the gateway you previously deployed on your host. In the activation
// process, you specify information such as the Amazon Web Services Region that you
// want to use for storing snapshots or tapes, the time zone for scheduled
// snapshots the gateway snapshot schedule window, an activation key, and a name
// for your gateway. The activation process also associates your gateway with your
// account. For more information, see UpdateGatewayInformation . You must turn on
// the gateway VM before you can activate your gateway.
func (c *Client) ActivateGateway(ctx context.Context, params *ActivateGatewayInput, optFns ...func(*Options)) (*ActivateGatewayOutput, error) {
	if params == nil {
		params = &ActivateGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ActivateGateway", params, optFns, c.addOperationActivateGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ActivateGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:
//   - ActivateGatewayInput$ActivationKey
//   - ActivateGatewayInput$GatewayName
//   - ActivateGatewayInput$GatewayRegion
//   - ActivateGatewayInput$GatewayTimezone
//   - ActivateGatewayInput$GatewayType
//   - ActivateGatewayInput$MediumChangerType
//   - ActivateGatewayInput$TapeDriveType
type ActivateGatewayInput struct {

	// Your gateway activation key. You can obtain the activation key by sending an
	// HTTP GET request with redirects enabled to the gateway IP address (port 80). The
	// redirect URL returned in the response provides you the activation key for your
	// gateway in the query string parameter activationKey . It may also include other
	// activation-related parameters, however, these are merely defaults -- the
	// arguments you pass to the ActivateGateway API call determine the actual
	// configuration of your gateway. For more information, see Getting activation key (https://docs.aws.amazon.com/storagegateway/latest/userguide/get-activation-key.html)
	// in the Storage Gateway User Guide.
	//
	// This member is required.
	ActivationKey *string

	// The name you configured for your gateway.
	//
	// This member is required.
	GatewayName *string

	// A value that indicates the Amazon Web Services Region where you want to store
	// your data. The gateway Amazon Web Services Region specified must be the same
	// Amazon Web Services Region as the Amazon Web Services Region in your Host
	// header in the request. For more information about available Amazon Web Services
	// Regions and endpoints for Storage Gateway, see Storage Gateway endpoints and
	// quotas (https://docs.aws.amazon.com/general/latest/gr/sg.html) in the Amazon Web
	// Services General Reference. Valid Values: See Storage Gateway endpoints and
	// quotas (https://docs.aws.amazon.com/general/latest/gr/sg.html) in the Amazon Web
	// Services General Reference.
	//
	// This member is required.
	GatewayRegion *string

	// A value that indicates the time zone you want to set for the gateway. The time
	// zone is of the format "GMT-hr:mm" or "GMT+hr:mm". For example, GMT-4:00
	// indicates the time is 4 hours behind GMT. GMT+2:00 indicates the time is 2 hours
	// ahead of GMT. The time zone is used, for example, for scheduling snapshots and
	// your gateway's maintenance schedule.
	//
	// This member is required.
	GatewayTimezone *string

	// A value that defines the type of gateway to activate. The type specified is
	// critical to all later functions of the gateway and cannot be changed after
	// activation. The default value is CACHED . Valid Values: STORED | CACHED | VTL |
	// VTL_SNOW | FILE_S3 | FILE_FSX_SMB
	GatewayType *string

	// The value that indicates the type of medium changer to use for tape gateway.
	// This field is optional. Valid Values: STK-L700 | AWS-Gateway-VTL |
	// IBM-03584L32-0402
	MediumChangerType *string

	// A list of up to 50 tags that you can assign to the gateway. Each tag is a
	// key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers that can be represented in UTF-8 format, and the following special
	// characters: + - = . _ : / @. The maximum length of a tag's key is 128
	// characters, and the maximum length for a tag's value is 256 characters.
	Tags []types.Tag

	// The value that indicates the type of tape drive to use for tape gateway. This
	// field is optional. Valid Values: IBM-ULT3580-TD5
	TapeDriveType *string

	noSmithyDocumentSerde
}

// Storage Gateway returns the Amazon Resource Name (ARN) of the activated
// gateway. It is a string made of information such as your account, gateway name,
// and Amazon Web Services Region. This ARN is used to reference the gateway in
// other API operations as well as resource-based authorization. For gateways
// activated prior to September 02, 2015, the gateway ARN contains the gateway name
// rather than the gateway ID. Changing the name of the gateway has no effect on
// the gateway ARN.
type ActivateGatewayOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationActivateGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpActivateGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpActivateGateway{}, middleware.After)
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
	if err = addActivateGatewayResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpActivateGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opActivateGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opActivateGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ActivateGateway",
	}
}

type opActivateGatewayResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opActivateGatewayResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opActivateGatewayResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "storagegateway"
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
				signingName = "storagegateway"
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
				v4aScheme.SigningName = aws.String("storagegateway")
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

func addActivateGatewayResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opActivateGatewayResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
