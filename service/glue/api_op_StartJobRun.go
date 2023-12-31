// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a job run using a job definition.
func (c *Client) StartJobRun(ctx context.Context, params *StartJobRunInput, optFns ...func(*Options)) (*StartJobRunOutput, error) {
	if params == nil {
		params = &StartJobRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartJobRun", params, optFns, c.addOperationStartJobRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartJobRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartJobRunInput struct {

	// The name of the job definition to use.
	//
	// This member is required.
	JobName *string

	// This field is deprecated. Use MaxCapacity instead. The number of Glue data
	// processing units (DPUs) to allocate to this JobRun. You can allocate a minimum
	// of 2 DPUs; the default is 10. A DPU is a relative measure of processing power
	// that consists of 4 vCPUs of compute capacity and 16 GB of memory. For more
	// information, see the Glue pricing page (https://aws.amazon.com/glue/pricing/) .
	//
	// Deprecated: This property is deprecated, use MaxCapacity instead.
	AllocatedCapacity int32

	// The job arguments associated with this run. For this job run, they replace the
	// default arguments set in the job definition itself. You can specify arguments
	// here that your own job-execution script consumes, as well as arguments that Glue
	// itself consumes. Job arguments may be logged. Do not pass plaintext secrets as
	// arguments. Retrieve secrets from a Glue Connection, Secrets Manager or other
	// secret management mechanism if you intend to keep them within the Job. For
	// information about how to specify and consume your own Job arguments, see the
	// Calling Glue APIs in Python (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html)
	// topic in the developer guide. For information about the arguments you can
	// provide to this field when configuring Spark jobs, see the Special Parameters
	// Used by Glue (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html)
	// topic in the developer guide. For information about the arguments you can
	// provide to this field when configuring Ray jobs, see Using job parameters in
	// Ray jobs (https://docs.aws.amazon.com/glue/latest/dg/author-job-ray-job-parameters.html)
	// in the developer guide.
	Arguments map[string]string

	// Indicates whether the job is run with a standard or flexible execution class.
	// The standard execution-class is ideal for time-sensitive workloads that require
	// fast job startup and dedicated resources. The flexible execution class is
	// appropriate for time-insensitive jobs whose start and completion times may vary.
	// Only jobs with Glue version 3.0 and above and command type glueetl will be
	// allowed to set ExecutionClass to FLEX . The flexible execution class is
	// available for Spark jobs.
	ExecutionClass types.ExecutionClass

	// The ID of a previous JobRun to retry.
	JobRunId *string

	// For Glue version 1.0 or earlier jobs, using the standard worker type, the
	// number of Glue data processing units (DPUs) that can be allocated when this job
	// runs. A DPU is a relative measure of processing power that consists of 4 vCPUs
	// of compute capacity and 16 GB of memory. For more information, see the Glue
	// pricing page (https://aws.amazon.com/glue/pricing/) . For Glue version 2.0+
	// jobs, you cannot specify a Maximum capacity . Instead, you should specify a
	// Worker type and the Number of workers . Do not set MaxCapacity if using
	// WorkerType and NumberOfWorkers . The value that can be allocated for MaxCapacity
	// depends on whether you are running a Python shell job, an Apache Spark ETL job,
	// or an Apache Spark streaming ETL job:
	//   - When you specify a Python shell job ( JobCommand.Name ="pythonshell"), you
	//   can allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.
	//   - When you specify an Apache Spark ETL job ( JobCommand.Name ="glueetl") or
	//   Apache Spark streaming ETL job ( JobCommand.Name ="gluestreaming"), you can
	//   allocate from 2 to 100 DPUs. The default is 10 DPUs. This job type cannot have a
	//   fractional DPU allocation.
	MaxCapacity *float64

	// Specifies configuration properties of a job run notification.
	NotificationProperty *types.NotificationProperty

	// The number of workers of a defined workerType that are allocated when a job
	// runs.
	NumberOfWorkers *int32

	// The name of the SecurityConfiguration structure to be used with this job run.
	SecurityConfiguration *string

	// The JobRun timeout in minutes. This is the maximum time that a job run can
	// consume resources before it is terminated and enters TIMEOUT status. This value
	// overrides the timeout value set in the parent job. Streaming jobs do not have a
	// timeout. The default for non-streaming jobs is 2,880 minutes (48 hours).
	Timeout *int32

	// The type of predefined worker that is allocated when a job runs. Accepts a
	// value of G.1X, G.2X, G.4X, G.8X or G.025X for Spark jobs. Accepts the value Z.2X
	// for Ray jobs.
	//   - For the G.1X worker type, each worker maps to 1 DPU (4 vCPUs, 16 GB of
	//   memory) with 84GB disk (approximately 34GB free), and provides 1 executor per
	//   worker. We recommend this worker type for workloads such as data transforms,
	//   joins, and queries, to offers a scalable and cost effective way to run most
	//   jobs.
	//   - For the G.2X worker type, each worker maps to 2 DPU (8 vCPUs, 32 GB of
	//   memory) with 128GB disk (approximately 77GB free), and provides 1 executor per
	//   worker. We recommend this worker type for workloads such as data transforms,
	//   joins, and queries, to offers a scalable and cost effective way to run most
	//   jobs.
	//   - For the G.4X worker type, each worker maps to 4 DPU (16 vCPUs, 64 GB of
	//   memory) with 256GB disk (approximately 235GB free), and provides 1 executor per
	//   worker. We recommend this worker type for jobs whose workloads contain your most
	//   demanding transforms, aggregations, joins, and queries. This worker type is
	//   available only for Glue version 3.0 or later Spark ETL jobs in the following
	//   Amazon Web Services Regions: US East (Ohio), US East (N. Virginia), US West
	//   (Oregon), Asia Pacific (Singapore), Asia Pacific (Sydney), Asia Pacific (Tokyo),
	//   Canada (Central), Europe (Frankfurt), Europe (Ireland), and Europe (Stockholm).
	//   - For the G.8X worker type, each worker maps to 8 DPU (32 vCPUs, 128 GB of
	//   memory) with 512GB disk (approximately 487GB free), and provides 1 executor per
	//   worker. We recommend this worker type for jobs whose workloads contain your most
	//   demanding transforms, aggregations, joins, and queries. This worker type is
	//   available only for Glue version 3.0 or later Spark ETL jobs, in the same Amazon
	//   Web Services Regions as supported for the G.4X worker type.
	//   - For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPUs, 4 GB of
	//   memory) with 84GB disk (approximately 34GB free), and provides 1 executor per
	//   worker. We recommend this worker type for low volume streaming jobs. This worker
	//   type is only available for Glue version 3.0 streaming jobs.
	//   - For the Z.2X worker type, each worker maps to 2 M-DPU (8vCPUs, 64 GB of
	//   memory) with 128 GB disk (approximately 120GB free), and provides up to 8 Ray
	//   workers based on the autoscaler.
	WorkerType types.WorkerType

	noSmithyDocumentSerde
}

type StartJobRunOutput struct {

	// The ID assigned to this job run.
	JobRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartJobRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartJobRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartJobRun{}, middleware.After)
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
	if err = addStartJobRunResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartJobRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartJobRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartJobRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartJobRun",
	}
}

type opStartJobRunResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartJobRunResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartJobRunResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "glue"
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
				signingName = "glue"
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
				v4aScheme.SigningName = aws.String("glue")
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

func addStartJobRunResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartJobRunResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
