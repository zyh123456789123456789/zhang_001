// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationautoscaling

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the Application Auto Scaling scheduled actions for the specified
// service namespace. You can filter the results using the ResourceId ,
// ScalableDimension , and ScheduledActionNames parameters. For more information,
// see Scheduled scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html)
// and Managing scheduled scaling (https://docs.aws.amazon.com/autoscaling/application/userguide/scheduled-scaling-additional-cli-commands.html)
// in the Application Auto Scaling User Guide.
func (c *Client) DescribeScheduledActions(ctx context.Context, params *DescribeScheduledActionsInput, optFns ...func(*Options)) (*DescribeScheduledActionsOutput, error) {
	if params == nil {
		params = &DescribeScheduledActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScheduledActions", params, optFns, c.addOperationDescribeScheduledActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScheduledActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScheduledActionsInput struct {

	// The namespace of the Amazon Web Services service that provides the resource.
	// For a resource provided by your own application or service, use custom-resource
	// instead.
	//
	// This member is required.
	ServiceNamespace types.ServiceNamespace

	// The maximum number of scheduled action results. This value can be between 1 and
	// 50. The default value is 50. If this parameter is used, the operation returns up
	// to MaxResults results at a time, along with a NextToken value. To get the next
	// set of results, include the NextToken value in a subsequent call. If this
	// parameter is not used, the operation returns up to 50 results and a NextToken
	// value, if applicable.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// The identifier of the resource associated with the scheduled action. This
	// string consists of the resource type and unique identifier.
	//   - ECS service - The resource type is service and the unique identifier is the
	//   cluster name and service name. Example: service/default/sample-webapp .
	//   - Spot Fleet - The resource type is spot-fleet-request and the unique
	//   identifier is the Spot Fleet request ID. Example:
	//   spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE .
	//   - EMR cluster - The resource type is instancegroup and the unique identifier
	//   is the cluster ID and instance group ID. Example:
	//   instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0 .
	//   - AppStream 2.0 fleet - The resource type is fleet and the unique identifier
	//   is the fleet name. Example: fleet/sample-fleet .
	//   - DynamoDB table - The resource type is table and the unique identifier is the
	//   table name. Example: table/my-table .
	//   - DynamoDB global secondary index - The resource type is index and the unique
	//   identifier is the index name. Example: table/my-table/index/my-table-index .
	//   - Aurora DB cluster - The resource type is cluster and the unique identifier
	//   is the cluster name. Example: cluster:my-db-cluster .
	//   - SageMaker endpoint variant - The resource type is variant and the unique
	//   identifier is the resource ID. Example:
	//   endpoint/my-end-point/variant/KMeansClustering .
	//   - Custom resources are not supported with a resource type. This parameter
	//   must specify the OutputValue from the CloudFormation template stack used to
	//   access the resources. The unique identifier is defined by the service provider.
	//   More information is available in our GitHub repository (https://github.com/aws/aws-auto-scaling-custom-resource)
	//   .
	//   - Amazon Comprehend document classification endpoint - The resource type and
	//   unique identifier are specified using the endpoint ARN. Example:
	//   arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE
	//   .
	//   - Amazon Comprehend entity recognizer endpoint - The resource type and unique
	//   identifier are specified using the endpoint ARN. Example:
	//   arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE .
	//   - Lambda provisioned concurrency - The resource type is function and the
	//   unique identifier is the function name with a function version or alias name
	//   suffix that is not $LATEST . Example: function:my-function:prod or
	//   function:my-function:1 .
	//   - Amazon Keyspaces table - The resource type is table and the unique
	//   identifier is the table name. Example: keyspace/mykeyspace/table/mytable .
	//   - Amazon MSK cluster - The resource type and unique identifier are specified
	//   using the cluster ARN. Example:
	//   arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5
	//   .
	//   - Amazon ElastiCache replication group - The resource type is
	//   replication-group and the unique identifier is the replication group name.
	//   Example: replication-group/mycluster .
	//   - Neptune cluster - The resource type is cluster and the unique identifier is
	//   the cluster name. Example: cluster:mycluster .
	//   - SageMaker Serverless endpoint - The resource type is variant and the unique
	//   identifier is the resource ID. Example:
	//   endpoint/my-end-point/variant/KMeansClustering .
	ResourceId *string

	// The scalable dimension. This string consists of the service namespace, resource
	// type, and scaling property. If you specify a scalable dimension, you must also
	// specify a resource ID.
	//   - ecs:service:DesiredCount - The desired task count of an ECS service.
	//   - elasticmapreduce:instancegroup:InstanceCount - The instance count of an EMR
	//   Instance Group.
	//   - ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot Fleet.
	//   - appstream:fleet:DesiredCapacity - The desired capacity of an AppStream 2.0
	//   fleet.
	//   - dynamodb:table:ReadCapacityUnits - The provisioned read capacity for a
	//   DynamoDB table.
	//   - dynamodb:table:WriteCapacityUnits - The provisioned write capacity for a
	//   DynamoDB table.
	//   - dynamodb:index:ReadCapacityUnits - The provisioned read capacity for a
	//   DynamoDB global secondary index.
	//   - dynamodb:index:WriteCapacityUnits - The provisioned write capacity for a
	//   DynamoDB global secondary index.
	//   - rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora DB
	//   cluster. Available for Aurora MySQL-compatible edition and Aurora
	//   PostgreSQL-compatible edition.
	//   - sagemaker:variant:DesiredInstanceCount - The number of EC2 instances for a
	//   SageMaker model endpoint variant.
	//   - custom-resource:ResourceType:Property - The scalable dimension for a custom
	//   resource provided by your own application or service.
	//   - comprehend:document-classifier-endpoint:DesiredInferenceUnits - The number
	//   of inference units for an Amazon Comprehend document classification endpoint.
	//   - comprehend:entity-recognizer-endpoint:DesiredInferenceUnits - The number of
	//   inference units for an Amazon Comprehend entity recognizer endpoint.
	//   - lambda:function:ProvisionedConcurrency - The provisioned concurrency for a
	//   Lambda function.
	//   - cassandra:table:ReadCapacityUnits - The provisioned read capacity for an
	//   Amazon Keyspaces table.
	//   - cassandra:table:WriteCapacityUnits - The provisioned write capacity for an
	//   Amazon Keyspaces table.
	//   - kafka:broker-storage:VolumeSize - The provisioned volume size (in GiB) for
	//   brokers in an Amazon MSK cluster.
	//   - elasticache:replication-group:NodeGroups - The number of node groups for an
	//   Amazon ElastiCache replication group.
	//   - elasticache:replication-group:Replicas - The number of replicas per node
	//   group for an Amazon ElastiCache replication group.
	//   - neptune:cluster:ReadReplicaCount - The count of read replicas in an Amazon
	//   Neptune DB cluster.
	//   - sagemaker:variant:DesiredProvisionedConcurrency - The provisioned
	//   concurrency for a SageMaker Serverless endpoint.
	ScalableDimension types.ScalableDimension

	// The names of the scheduled actions to describe.
	ScheduledActionNames []string

	noSmithyDocumentSerde
}

type DescribeScheduledActionsOutput struct {

	// The token required to get the next set of results. This value is null if there
	// are no more results to return.
	NextToken *string

	// Information about the scheduled actions.
	ScheduledActions []types.ScheduledAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScheduledActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeScheduledActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeScheduledActions{}, middleware.After)
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
	if err = addDescribeScheduledActionsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeScheduledActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScheduledActions(options.Region), middleware.Before); err != nil {
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

// DescribeScheduledActionsAPIClient is a client that implements the
// DescribeScheduledActions operation.
type DescribeScheduledActionsAPIClient interface {
	DescribeScheduledActions(context.Context, *DescribeScheduledActionsInput, ...func(*Options)) (*DescribeScheduledActionsOutput, error)
}

var _ DescribeScheduledActionsAPIClient = (*Client)(nil)

// DescribeScheduledActionsPaginatorOptions is the paginator options for
// DescribeScheduledActions
type DescribeScheduledActionsPaginatorOptions struct {
	// The maximum number of scheduled action results. This value can be between 1 and
	// 50. The default value is 50. If this parameter is used, the operation returns up
	// to MaxResults results at a time, along with a NextToken value. To get the next
	// set of results, include the NextToken value in a subsequent call. If this
	// parameter is not used, the operation returns up to 50 results and a NextToken
	// value, if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeScheduledActionsPaginator is a paginator for DescribeScheduledActions
type DescribeScheduledActionsPaginator struct {
	options   DescribeScheduledActionsPaginatorOptions
	client    DescribeScheduledActionsAPIClient
	params    *DescribeScheduledActionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeScheduledActionsPaginator returns a new
// DescribeScheduledActionsPaginator
func NewDescribeScheduledActionsPaginator(client DescribeScheduledActionsAPIClient, params *DescribeScheduledActionsInput, optFns ...func(*DescribeScheduledActionsPaginatorOptions)) *DescribeScheduledActionsPaginator {
	if params == nil {
		params = &DescribeScheduledActionsInput{}
	}

	options := DescribeScheduledActionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeScheduledActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeScheduledActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeScheduledActions page.
func (p *DescribeScheduledActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeScheduledActionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeScheduledActions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeScheduledActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "application-autoscaling",
		OperationName: "DescribeScheduledActions",
	}
}

type opDescribeScheduledActionsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeScheduledActionsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeScheduledActionsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "application-autoscaling"
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
				signingName = "application-autoscaling"
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
				v4aScheme.SigningName = aws.String("application-autoscaling")
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

func addDescribeScheduledActionsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeScheduledActionsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
