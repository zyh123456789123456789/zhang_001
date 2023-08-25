// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a simulation job. After 90 days, simulation jobs expire and will be
// deleted. They will no longer be accessible.
func (c *Client) CreateSimulationJob(ctx context.Context, params *CreateSimulationJobInput, optFns ...func(*Options)) (*CreateSimulationJobOutput, error) {
	if params == nil {
		params = &CreateSimulationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSimulationJob", params, optFns, c.addOperationCreateSimulationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSimulationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSimulationJobInput struct {

	// The IAM role name that allows the simulation instance to call the AWS APIs that
	// are specified in its associated policies on your behalf. This is how credentials
	// are passed in to your simulation job.
	//
	// This member is required.
	IamRole *string

	// The maximum simulation job duration in seconds (up to 14 days or 1,209,600
	// seconds. When maxJobDurationInSeconds is reached, the simulation job will
	// status will transition to Completed .
	//
	// This member is required.
	MaxJobDurationInSeconds int64

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// Compute information for the simulation job.
	Compute *types.Compute

	// Specify data sources to mount read-only files from S3 into your simulation.
	// These files are available under /opt/robomaker/datasources/data_source_name .
	// There is a limit of 100 files and a combined size of 25GB for all
	// DataSourceConfig objects.
	DataSources []types.DataSourceConfig

	// The failure behavior the simulation job. Continue Leaves the instance running
	// for its maximum timeout duration after a 4XX error code. Fail Stop the
	// simulation job and terminate the instance.
	FailureBehavior types.FailureBehavior

	// The logging configuration.
	LoggingConfig *types.LoggingConfig

	// Location for output files generated by the simulation job.
	OutputLocation *types.OutputLocation

	// The robot application to use in the simulation job.
	RobotApplications []types.RobotApplicationConfig

	// The simulation application to use in the simulation job.
	SimulationApplications []types.SimulationApplicationConfig

	// A map that contains tag keys and tag values that are attached to the simulation
	// job.
	Tags map[string]string

	// If your simulation job accesses resources in a VPC, you provide this parameter
	// identifying the list of security group IDs and subnet IDs. These must belong to
	// the same VPC. You must provide at least one security group and one subnet ID.
	VpcConfig *types.VPCConfig

	noSmithyDocumentSerde
}

type CreateSimulationJobOutput struct {

	// The Amazon Resource Name (ARN) of the simulation job.
	Arn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// Compute information for the simulation job.
	Compute *types.ComputeResponse

	// The data sources for the simulation job.
	DataSources []types.DataSource

	// the failure behavior for the simulation job.
	FailureBehavior types.FailureBehavior

	// The failure code of the simulation job if it failed: InternalServiceError
	// Internal service error. RobotApplicationCrash Robot application exited
	// abnormally. SimulationApplicationCrash Simulation application exited abnormally.
	// BadPermissionsRobotApplication Robot application bundle could not be downloaded.
	// BadPermissionsSimulationApplication Simulation application bundle could not be
	// downloaded. BadPermissionsS3Output Unable to publish outputs to
	// customer-provided S3 bucket. BadPermissionsCloudwatchLogs Unable to publish logs
	// to customer-provided CloudWatch Logs resource. SubnetIpLimitExceeded Subnet IP
	// limit exceeded. ENILimitExceeded ENI limit exceeded.
	// BadPermissionsUserCredentials Unable to use the Role provided.
	// InvalidBundleRobotApplication Robot bundle cannot be extracted (invalid format,
	// bundling error, or other issue). InvalidBundleSimulationApplication Simulation
	// bundle cannot be extracted (invalid format, bundling error, or other issue).
	// RobotApplicationVersionMismatchedEtag Etag for RobotApplication does not match
	// value during version creation. SimulationApplicationVersionMismatchedEtag Etag
	// for SimulationApplication does not match value during version creation.
	FailureCode types.SimulationJobErrorCode

	// The IAM role that allows the simulation job to call the AWS APIs that are
	// specified in its associated policies on your behalf.
	IamRole *string

	// The time, in milliseconds since the epoch, when the simulation job was last
	// started.
	LastStartedAt *time.Time

	// The time, in milliseconds since the epoch, when the simulation job was last
	// updated.
	LastUpdatedAt *time.Time

	// The logging configuration.
	LoggingConfig *types.LoggingConfig

	// The maximum simulation job duration in seconds.
	MaxJobDurationInSeconds int64

	// Simulation job output files location.
	OutputLocation *types.OutputLocation

	// The robot application used by the simulation job.
	RobotApplications []types.RobotApplicationConfig

	// The simulation application used by the simulation job.
	SimulationApplications []types.SimulationApplicationConfig

	// The simulation job execution duration in milliseconds.
	SimulationTimeMillis int64

	// The status of the simulation job.
	Status types.SimulationJobStatus

	// The list of all tags added to the simulation job.
	Tags map[string]string

	// Information about the vpc configuration.
	VpcConfig *types.VPCConfigResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSimulationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSimulationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSimulationJob{}, middleware.After)
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
	if err = addCreateSimulationJobResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateSimulationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSimulationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSimulationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateSimulationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateSimulationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateSimulationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateSimulationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateSimulationJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateSimulationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateSimulationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateSimulationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateSimulationJob",
	}
}

type opCreateSimulationJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateSimulationJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateSimulationJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "robomaker"
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
				signingName = "robomaker"
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
				v4aScheme.SigningName = aws.String("robomaker")
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

func addCreateSimulationJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateSimulationJobResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}