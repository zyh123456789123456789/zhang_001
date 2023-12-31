// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a model in SageMaker. In the request, you name the model and describe a
// primary container. For the primary container, you specify the Docker image that
// contains inference code, artifacts (from prior training), and a custom
// environment map that the inference code uses when you deploy the model for
// predictions. Use this API to create a model if you want to use SageMaker hosting
// services or run a batch transform job. To host your model, you create an
// endpoint configuration with the CreateEndpointConfig API, and then create an
// endpoint with the CreateEndpoint API. SageMaker then deploys all of the
// containers that you defined for the model in the hosting environment. For an
// example that calls this method when deploying a model to SageMaker hosting
// services, see Create a Model (Amazon Web Services SDK for Python (Boto 3)). (https://docs.aws.amazon.com/sagemaker/latest/dg/realtime-endpoints-deployment.html#realtime-endpoints-deployment-create-model)
// To run a batch transform using your model, you start a job with the
// CreateTransformJob API. SageMaker uses your model and your dataset to get
// inferences which are then saved to a specified S3 location. In the request, you
// also provide an IAM role that SageMaker can assume to access model artifacts and
// docker image for deployment on ML compute hosting instances or for batch
// transform jobs. In addition, you also use the IAM role to manage permissions the
// inference code needs. For example, if the inference code access any other Amazon
// Web Services resources, you grant necessary permissions via this role.
func (c *Client) CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) {
	if params == nil {
		params = &CreateModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModel", params, optFns, c.addOperationCreateModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelInput struct {

	// The Amazon Resource Name (ARN) of the IAM role that SageMaker can assume to
	// access model artifacts and docker image for deployment on ML compute instances
	// or for batch transform jobs. Deploying on ML compute instances is part of model
	// hosting. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html)
	// . To be able to pass this role to SageMaker, the caller of this API must have
	// the iam:PassRole permission.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The name of the new model.
	//
	// This member is required.
	ModelName *string

	// Specifies the containers in the inference pipeline.
	Containers []types.ContainerDefinition

	// Isolates the model container. No inbound or outbound network calls can be made
	// to or from the model container.
	EnableNetworkIsolation bool

	// Specifies details of how containers in a multi-container endpoint are called.
	InferenceExecutionConfig *types.InferenceExecutionConfig

	// The location of the primary docker image containing inference code, associated
	// artifacts, and custom environment map that the inference code uses when the
	// model is deployed for predictions.
	PrimaryContainer *types.ContainerDefinition

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	Tags []types.Tag

	// A VpcConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html)
	// object that specifies the VPC that you want your model to connect to. Control
	// access to and from your model container by configuring the VPC. VpcConfig is
	// used in hosting services and in batch transform. For more information, see
	// Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html)
	// and Protect Data in Batch Transform Jobs by Using an Amazon Virtual Private
	// Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-vpc.html) .
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type CreateModelOutput struct {

	// The ARN of the model created in SageMaker.
	//
	// This member is required.
	ModelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateModel{}, middleware.After)
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
	if err = addCreateModelResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateModel",
	}
}

type opCreateModelResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateModelResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateModelResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "sagemaker"
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
				signingName = "sagemaker"
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
				v4aScheme.SigningName = aws.String("sagemaker")
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

func addCreateModelResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateModelResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
