// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specifies and starts a remote access session.
func (c *Client) CreateRemoteAccessSession(ctx context.Context, params *CreateRemoteAccessSessionInput, optFns ...func(*Options)) (*CreateRemoteAccessSessionOutput, error) {
	if params == nil {
		params = &CreateRemoteAccessSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRemoteAccessSession", params, optFns, c.addOperationCreateRemoteAccessSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRemoteAccessSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates and submits a request to start a remote access session.
type CreateRemoteAccessSessionInput struct {

	// The ARN of the device for which you want to create a remote access session.
	//
	// This member is required.
	DeviceArn *string

	// The Amazon Resource Name (ARN) of the project for which you want to create a
	// remote access session.
	//
	// This member is required.
	ProjectArn *string

	// Unique identifier for the client. If you want access to multiple devices on the
	// same client, you should pass the same clientId value in each call to
	// CreateRemoteAccessSession . This identifier is required only if
	// remoteDebugEnabled is set to true . Remote debugging is no longer supported (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html)
	// .
	ClientId *string

	// The configuration information for the remote access session request.
	Configuration *types.CreateRemoteAccessSessionConfiguration

	// The Amazon Resource Name (ARN) of the device instance for which you want to
	// create a remote access session.
	InstanceArn *string

	// The interaction mode of the remote access session. Valid values are:
	//   - INTERACTIVE: You can interact with the iOS device by viewing, touching, and
	//   rotating the screen. You cannot run XCUITest framework-based tests in this mode.
	//
	//   - NO_VIDEO: You are connected to the device, but cannot interact with it or
	//   view the screen. This mode has the fastest test execution speed. You can run
	//   XCUITest framework-based tests in this mode.
	//   - VIDEO_ONLY: You can view the screen, but cannot touch or rotate it. You can
	//   run XCUITest framework-based tests and watch the screen in this mode.
	InteractionMode types.InteractionMode

	// The name of the remote access session to create.
	Name *string

	// Set to true if you want to access devices remotely for debugging in your remote
	// access session. Remote debugging is no longer supported (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html)
	// .
	RemoteDebugEnabled *bool

	// The Amazon Resource Name (ARN) for the app to be recorded in the remote access
	// session.
	RemoteRecordAppArn *string

	// Set to true to enable remote recording for the remote access session.
	RemoteRecordEnabled *bool

	// When set to true , for private devices, Device Farm does not sign your app
	// again. For public devices, Device Farm always signs your apps again. For more
	// information on how Device Farm modifies your uploads during tests, see Do you
	// modify my app? (http://aws.amazon.com/device-farm/faqs/)
	SkipAppResign *bool

	// Ignored. The public key of the ssh key pair you want to use for connecting to
	// remote devices in your remote debugging session. This key is required only if
	// remoteDebugEnabled is set to true . Remote debugging is no longer supported (https://docs.aws.amazon.com/devicefarm/latest/developerguide/history.html)
	// .
	SshPublicKey *string

	noSmithyDocumentSerde
}

// Represents the server response from a request to create a remote access session.
type CreateRemoteAccessSessionOutput struct {

	// A container that describes the remote access session when the request to create
	// a remote access session is sent.
	RemoteAccessSession *types.RemoteAccessSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRemoteAccessSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRemoteAccessSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRemoteAccessSession{}, middleware.After)
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
	if err = addCreateRemoteAccessSessionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateRemoteAccessSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRemoteAccessSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRemoteAccessSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateRemoteAccessSession",
	}
}

type opCreateRemoteAccessSessionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateRemoteAccessSessionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateRemoteAccessSessionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "devicefarm"
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
				signingName = "devicefarm"
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
				v4aScheme.SigningName = aws.String("devicefarm")
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

func addCreateRemoteAccessSessionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateRemoteAccessSessionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
