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
	"time"
)

// Creates a hosted connection on the specified interconnect or a link aggregation
// group (LAG) of interconnects. Allocates a VLAN number and a specified amount of
// capacity (bandwidth) for use by a hosted connection on the specified
// interconnect or LAG of interconnects. Amazon Web Services polices the hosted
// connection for the specified capacity and the Direct Connect Partner must also
// police the hosted connection for the specified capacity. Intended for use by
// Direct Connect Partners only.
func (c *Client) AllocateHostedConnection(ctx context.Context, params *AllocateHostedConnectionInput, optFns ...func(*Options)) (*AllocateHostedConnectionOutput, error) {
	if params == nil {
		params = &AllocateHostedConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AllocateHostedConnection", params, optFns, c.addOperationAllocateHostedConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AllocateHostedConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AllocateHostedConnectionInput struct {

	// The bandwidth of the connection. The possible values are 50Mbps, 100Mbps,
	// 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, and 10Gbps. Note that
	// only those Direct Connect Partners who have met specific requirements are
	// allowed to create a 1Gbps, 2Gbps, 5Gbps or 10Gbps hosted connection.
	//
	// This member is required.
	Bandwidth *string

	// The ID of the interconnect or LAG.
	//
	// This member is required.
	ConnectionId *string

	// The name of the hosted connection.
	//
	// This member is required.
	ConnectionName *string

	// The ID of the Amazon Web Services account ID of the customer for the connection.
	//
	// This member is required.
	OwnerAccount *string

	// The dedicated VLAN provisioned to the hosted connection.
	//
	// This member is required.
	Vlan int32

	// The tags associated with the connection.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Information about an Direct Connect connection.
type AllocateHostedConnectionOutput struct {

	// The Direct Connect endpoint on which the physical connection terminates.
	//
	// Deprecated: This member has been deprecated.
	AwsDevice *string

	// The Direct Connect endpoint that terminates the physical connection.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The bandwidth of the connection.
	Bandwidth *string

	// The ID of the connection.
	ConnectionId *string

	// The name of the connection.
	ConnectionName *string

	// The state of the connection. The following are the possible values:
	//   - ordering : The initial state of a hosted connection provisioned on an
	//   interconnect. The connection stays in the ordering state until the owner of the
	//   hosted connection confirms or declines the connection order.
	//   - requested : The initial state of a standard connection. The connection stays
	//   in the requested state until the Letter of Authorization (LOA) is sent to the
	//   customer.
	//   - pending : The connection has been approved and is being initialized.
	//   - available : The network link is up and the connection is ready for use.
	//   - down : The network link is down.
	//   - deleting : The connection is being deleted.
	//   - deleted : The connection has been deleted.
	//   - rejected : A hosted connection in the ordering state enters the rejected
	//   state if it is deleted by the customer.
	//   - unknown : The state of the connection is not available.
	ConnectionState types.ConnectionState

	// The MAC Security (MACsec) connection encryption mode. The valid values are
	// no_encrypt , should_encrypt , and must_encrypt .
	EncryptionMode *string

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// Indicates whether jumbo frames are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time

	// The location of the connection.
	Location *string

	// Indicates whether the connection supports MAC Security (MACsec).
	MacSecCapable *bool

	// The MAC Security (MACsec) security keys associated with the connection.
	MacSecKeys []types.MacSecKey

	// The ID of the Amazon Web Services account that owns the connection.
	OwnerAccount *string

	// The name of the Direct Connect service provider associated with the connection.
	PartnerName *string

	// The MAC Security (MACsec) port link status of the connection. The valid values
	// are Encryption Up , which means that there is an active Connection Key Name, or
	// Encryption Down .
	PortEncryptionStatus *string

	// The name of the service provider associated with the connection.
	ProviderName *string

	// The Amazon Web Services Region where the connection is located.
	Region *string

	// The tags associated with the connection.
	Tags []types.Tag

	// The ID of the VLAN.
	Vlan int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAllocateHostedConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAllocateHostedConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAllocateHostedConnection{}, middleware.After)
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
	if err = addAllocateHostedConnectionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAllocateHostedConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAllocateHostedConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAllocateHostedConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "AllocateHostedConnection",
	}
}

type opAllocateHostedConnectionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opAllocateHostedConnectionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opAllocateHostedConnectionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addAllocateHostedConnectionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opAllocateHostedConnectionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
