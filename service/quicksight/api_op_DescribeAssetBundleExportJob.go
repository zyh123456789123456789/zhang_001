// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes an existing export job. Poll job descriptions after a job starts to
// know the status of the job. When a job succeeds, a URL is provided to download
// the exported assets' data from. Download URLs are valid for five minutes after
// they are generated. You can call the DescribeAssetBundleExportJob API for a new
// download URL as needed. Job descriptions are available for 14 days after the job
// starts.
func (c *Client) DescribeAssetBundleExportJob(ctx context.Context, params *DescribeAssetBundleExportJobInput, optFns ...func(*Options)) (*DescribeAssetBundleExportJobOutput, error) {
	if params == nil {
		params = &DescribeAssetBundleExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssetBundleExportJob", params, optFns, c.addOperationDescribeAssetBundleExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssetBundleExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssetBundleExportJobInput struct {

	// The ID of the job that you want described. The job ID is set when you start a
	// new job with a StartAssetBundleExportJob API call.
	//
	// This member is required.
	AssetBundleExportJobId *string

	// The ID of the Amazon Web Services account the export job is executed in.
	//
	// This member is required.
	AwsAccountId *string

	noSmithyDocumentSerde
}

type DescribeAssetBundleExportJobOutput struct {

	// The Amazon Resource Name (ARN) for the export job.
	Arn *string

	// The ID of the job. The job ID is set when you start a new job with a
	// StartAssetBundleExportJob API call.
	AssetBundleExportJobId *string

	// The ID of the Amazon Web Services account that the export job was executed in.
	AwsAccountId *string

	// The CloudFormation override property configuration for the export job.
	CloudFormationOverridePropertyConfiguration *types.AssetBundleCloudFormationOverridePropertyConfiguration

	// The time that the export job was created.
	CreatedTime *time.Time

	// The URL to download the exported asset bundle data from. This URL is available
	// only after the job has succeeded. This URL is valid for 5 minutes after
	// issuance. Call DescribeAssetBundleExportJob again for a fresh URL if needed.
	// The downloaded asset bundle is a zip file named assetbundle-{jobId}.qs . The
	// file has a .qs extension. This URL can't be used in a StartAssetBundleImportJob
	// API call and should only be used for download purposes.
	DownloadUrl *string

	// An array of error records that describes any failures that occurred during the
	// export job processing. Error records accumulate while the job runs. The complete
	// set of error records is available after the job has completed and failed.
	Errors []types.AssetBundleExportJobError

	// The format of the exported asset bundle. A QUICKSIGHT_JSON formatted file can
	// be used to make a StartAssetBundleImportJob API call. A CLOUDFORMATION_JSON
	// formatted file can be used in the CloudFormation console and with the
	// CloudFormation APIs.
	ExportFormat types.AssetBundleExportFormat

	// The include dependencies flag.
	IncludeAllDependencies bool

	// Indicates the status of a job through its queuing and execution. Poll this
	// DescribeAssetBundleExportApi until JobStatus is either SUCCESSFUL or FAILED .
	JobStatus types.AssetBundleExportJobStatus

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// A list of resource ARNs that exported with the job.
	ResourceArns []string

	// The HTTP status of the response.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAssetBundleExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAssetBundleExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAssetBundleExportJob{}, middleware.After)
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
	if err = addDescribeAssetBundleExportJobResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeAssetBundleExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssetBundleExportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAssetBundleExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeAssetBundleExportJob",
	}
}

type opDescribeAssetBundleExportJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeAssetBundleExportJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeAssetBundleExportJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "quicksight"
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
				signingName = "quicksight"
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
				v4aScheme.SigningName = aws.String("quicksight")
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

func addDescribeAssetBundleExportJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeAssetBundleExportJobResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
