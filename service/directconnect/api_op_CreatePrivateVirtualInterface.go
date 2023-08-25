// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a private virtual interface. A virtual interface is the VLAN that
// transports Direct Connect traffic. A private virtual interface can be connected
// to either a Direct Connect gateway or a Virtual Private Gateway (VGW).
// Connecting the private virtual interface to a Direct Connect gateway enables the
// possibility for connecting to multiple VPCs, including VPCs in different Amazon
// Web Services Regions. Connecting the private virtual interface to a VGW only
// provides access to a single VPC within the same Region. Setting the MTU of a
// virtual interface to 9001 (jumbo frames) can cause an update to the underlying
// physical connection if it wasn't updated to support jumbo frames. Updating the
// connection disrupts network connectivity for all virtual interfaces associated
// with the connection for up to 30 seconds. To check whether your connection
// supports jumbo frames, call DescribeConnections . To check whether your virtual
// interface supports jumbo frames, call DescribeVirtualInterfaces .
func (c *Client) CreatePrivateVirtualInterface(ctx context.Context, params *CreatePrivateVirtualInterfaceInput, optFns ...func(*Options)) (*CreatePrivateVirtualInterfaceOutput, error) {
	if params == nil {
		params = &CreatePrivateVirtualInterfaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePrivateVirtualInterface", params, optFns, c.addOperationCreatePrivateVirtualInterfaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePrivateVirtualInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePrivateVirtualInterfaceInput struct {

	// The ID of the connection.
	//
	// This member is required.
	ConnectionId *string

	// Information about the private virtual interface.
	//
	// This member is required.
	NewPrivateVirtualInterface *types.NewPrivateVirtualInterface

	noSmithyDocumentSerde
}

// Information about a virtual interface.
type CreatePrivateVirtualInterfaceOutput struct {

	// The address family for the BGP peer.
	AddressFamily types.AddressFamily

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string

	// The autonomous system number (ASN) for the Amazon side of the connection.
	AmazonSideAsn *int64

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration. The valid values are 1-2147483647.
	Asn int32

	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string

	// The Direct Connect endpoint that terminates the physical connection.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The BGP peers configured on this virtual interface.
	BgpPeers []types.BGPPeer

	// The ID of the connection.
	ConnectionId *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The customer router configuration.
	CustomerRouterConfig *string

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// Indicates whether jumbo frames are supported.
	JumboFrameCapable *bool

	// The location of the connection.
	Location *string

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 8500. The default value is 1500
	Mtu *int32

	// The ID of the Amazon Web Services account that owns the virtual interface.
	OwnerAccount *string

	// The Amazon Web Services Region where the virtual interface is located.
	Region *string

	// The routes to be advertised to the Amazon Web Services network in this Region.
	// Applies to public virtual interfaces.
	RouteFilterPrefixes []types.RouteFilterPrefix

	// Indicates whether SiteLink is enabled.
	SiteLinkEnabled *bool

	// The tags associated with the virtual interface.
	Tags []types.Tag

	// The ID of the virtual private gateway. Applies only to private virtual
	// interfaces.
	VirtualGatewayId *string

	// The ID of the virtual interface.
	VirtualInterfaceId *string

	// The name of the virtual interface assigned by the customer network. The name
	// has a maximum of 100 characters. The following are valid characters: a-z, 0-9
	// and a hyphen (-).
	VirtualInterfaceName *string

	// The state of the virtual interface. The following are the possible values:
	//   - confirming : The creation of the virtual interface is pending confirmation
	//   from the virtual interface owner. If the owner of the virtual interface is
	//   different from the owner of the connection on which it is provisioned, then the
	//   virtual interface will remain in this state until it is confirmed by the virtual
	//   interface owner.
	//   - verifying : This state only applies to public virtual interfaces. Each
	//   public virtual interface needs validation before the virtual interface can be
	//   created.
	//   - pending : A virtual interface is in this state from the time that it is
	//   created until the virtual interface is ready to forward traffic.
	//   - available : A virtual interface that is able to forward traffic.
	//   - down : A virtual interface that is BGP down.
	//   - deleting : A virtual interface is in this state immediately after calling
	//   DeleteVirtualInterface until it can no longer forward traffic.
	//   - deleted : A virtual interface that cannot forward traffic.
	//   - rejected : The virtual interface owner has declined creation of the virtual
	//   interface. If a virtual interface in the Confirming state is deleted by the
	//   virtual interface owner, the virtual interface enters the Rejected state.
	//   - unknown : The state of the virtual interface is not available.
	VirtualInterfaceState types.VirtualInterfaceState

	// The type of virtual interface. The possible values are private and public .
	VirtualInterfaceType *string

	// The ID of the VLAN.
	Vlan int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePrivateVirtualInterfaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePrivateVirtualInterface{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePrivateVirtualInterface{}, middleware.After)
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
	if err = addCreatePrivateVirtualInterfaceResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePrivateVirtualInterfaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePrivateVirtualInterface(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePrivateVirtualInterface(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "CreatePrivateVirtualInterface",
	}
}

type opCreatePrivateVirtualInterfaceResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreatePrivateVirtualInterfaceResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreatePrivateVirtualInterfaceResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "directconnect"
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
				signingName = "directconnect"
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
				v4aScheme.SigningName = aws.String("directconnect")
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

func addCreatePrivateVirtualInterfaceResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreatePrivateVirtualInterfaceResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}