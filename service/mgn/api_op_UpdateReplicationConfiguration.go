// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to update multiple ReplicationConfigurations by Source Server ID.
func (c *Client) UpdateReplicationConfiguration(ctx context.Context, params *UpdateReplicationConfigurationInput, optFns ...func(*Options)) (*UpdateReplicationConfigurationOutput, error) {
	if params == nil {
		params = &UpdateReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateReplicationConfiguration", params, optFns, c.addOperationUpdateReplicationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateReplicationConfigurationInput struct {

	// Update replication configuration Source Server ID request.
	//
	// This member is required.
	SourceServerID *string

	// Update replication configuration Account ID request.
	AccountID *string

	// Update replication configuration associate default Application Migration
	// Service Security group request.
	AssociateDefaultSecurityGroup *bool

	// Update replication configuration bandwidth throttling request.
	BandwidthThrottling int64

	// Update replication configuration create Public IP request.
	CreatePublicIP *bool

	// Update replication configuration data plane routing request.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// Update replication configuration use default large Staging Disk type request.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// Update replication configuration EBS encryption request.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// Update replication configuration EBS encryption key ARN request.
	EbsEncryptionKeyArn *string

	// Update replication configuration name request.
	Name *string

	// Update replication configuration replicated disks request.
	ReplicatedDisks []types.ReplicationConfigurationReplicatedDisk

	// Update replication configuration Replication Server instance type request.
	ReplicationServerInstanceType *string

	// Update replication configuration Replication Server Security Groups IDs request.
	ReplicationServersSecurityGroupsIDs []string

	// Update replication configuration Staging Area subnet request.
	StagingAreaSubnetId *string

	// Update replication configuration Staging Area Tags request.
	StagingAreaTags map[string]string

	// Update replication configuration use dedicated Replication Server request.
	UseDedicatedReplicationServer *bool

	// Update replication configuration use Fips Endpoint.
	UseFipsEndpoint *bool

	noSmithyDocumentSerde
}

type UpdateReplicationConfigurationOutput struct {

	// Replication Configuration associate default Application Migration Service
	// Security Group.
	AssociateDefaultSecurityGroup *bool

	// Replication Configuration set bandwidth throttling.
	BandwidthThrottling int64

	// Replication Configuration create Public IP.
	CreatePublicIP *bool

	// Replication Configuration data plane routing.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// Replication Configuration use default large Staging Disks.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// Replication Configuration EBS encryption.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// Replication Configuration EBS encryption key ARN.
	EbsEncryptionKeyArn *string

	// Replication Configuration name.
	Name *string

	// Replication Configuration replicated disks.
	ReplicatedDisks []types.ReplicationConfigurationReplicatedDisk

	// Replication Configuration Replication Server instance type.
	ReplicationServerInstanceType *string

	// Replication Configuration Replication Server Security Group IDs.
	ReplicationServersSecurityGroupsIDs []string

	// Replication Configuration Source Server ID.
	SourceServerID *string

	// Replication Configuration Staging Area subnet ID.
	StagingAreaSubnetId *string

	// Replication Configuration Staging Area tags.
	StagingAreaTags map[string]string

	// Replication Configuration use Dedicated Replication Server.
	UseDedicatedReplicationServer *bool

	// Replication Configuration use Fips Endpoint.
	UseFipsEndpoint *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateReplicationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateReplicationConfiguration{}, middleware.After)
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
	if err = addUpdateReplicationConfigurationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateReplicationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReplicationConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateReplicationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "UpdateReplicationConfiguration",
	}
}

type opUpdateReplicationConfigurationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateReplicationConfigurationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateReplicationConfigurationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "mgn"
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
				signingName = "mgn"
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
				v4aScheme.SigningName = aws.String("mgn")
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

func addUpdateReplicationConfigurationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateReplicationConfigurationResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}