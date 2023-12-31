// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a previously provisioned cluster without its final snapshot being
// created. A successful response from the web service indicates that the request
// was received correctly. Use DescribeClusters to monitor the status of the
// deletion. The delete operation cannot be canceled or reverted once submitted.
// For more information about managing clusters, go to Amazon Redshift Clusters (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html)
// in the Amazon Redshift Cluster Management Guide. If you want to shut down the
// cluster and retain it for future use, set SkipFinalClusterSnapshot to false and
// specify a name for FinalClusterSnapshotIdentifier. You can later restore this
// snapshot to resume using the cluster. If a final cluster snapshot is requested,
// the status of the cluster will be "final-snapshot" while the snapshot is being
// taken, then it's "deleting" once Amazon Redshift begins deleting the cluster.
// For more information about managing clusters, go to Amazon Redshift Clusters (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) DeleteCluster(ctx context.Context, params *DeleteClusterInput, optFns ...func(*Options)) (*DeleteClusterOutput, error) {
	if params == nil {
		params = &DeleteClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCluster", params, optFns, c.addOperationDeleteClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteClusterInput struct {

	// The identifier of the cluster to be deleted. Constraints:
	//   - Must contain lowercase characters.
	//   - Must contain from 1 to 63 alphanumeric characters or hyphens.
	//   - First character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	ClusterIdentifier *string

	// The identifier of the final snapshot that is to be created immediately before
	// deleting the cluster. If this parameter is provided, SkipFinalClusterSnapshot
	// must be false . Constraints:
	//   - Must be 1 to 255 alphanumeric characters.
	//   - First character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	FinalClusterSnapshotIdentifier *string

	// The number of days that a manual snapshot is retained. If the value is -1, the
	// manual snapshot is retained indefinitely. The value must be either -1 or an
	// integer between 1 and 3,653. The default value is -1.
	FinalClusterSnapshotRetentionPeriod *int32

	// Determines whether a final snapshot of the cluster is created before Amazon
	// Redshift deletes the cluster. If true , a final cluster snapshot is not created.
	// If false , a final cluster snapshot is created before the cluster is deleted.
	// The FinalClusterSnapshotIdentifier parameter must be specified if
	// SkipFinalClusterSnapshot is false . Default: false
	SkipFinalClusterSnapshot bool

	noSmithyDocumentSerde
}

type DeleteClusterOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteCluster{}, middleware.After)
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
	if err = addDeleteClusterResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DeleteCluster",
	}
}

type opDeleteClusterResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDeleteClusterResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDeleteClusterResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "redshift"
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
				signingName = "redshift"
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
				v4aScheme.SigningName = aws.String("redshift")
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

func addDeleteClusterResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDeleteClusterResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
