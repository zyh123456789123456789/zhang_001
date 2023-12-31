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

// Creates an SageMaker notebook instance. A notebook instance is a machine
// learning (ML) compute instance running on a Jupyter notebook. In a
// CreateNotebookInstance request, specify the type of ML compute instance that you
// want to run. SageMaker launches the instance, installs common libraries that you
// can use to explore datasets for model training, and attaches an ML storage
// volume to the notebook instance. SageMaker also provides a set of example
// notebooks. Each notebook demonstrates how to use SageMaker with a specific
// algorithm or with a machine learning framework. After receiving the request,
// SageMaker does the following:
//   - Creates a network interface in the SageMaker VPC.
//   - (Option) If you specified SubnetId , SageMaker creates a network interface
//     in your own VPC, which is inferred from the subnet ID that you provide in the
//     input. When creating this network interface, SageMaker attaches the security
//     group that you specified in the request to the network interface that it creates
//     in your VPC.
//   - Launches an EC2 instance of the type specified in the request in the
//     SageMaker VPC. If you specified SubnetId of your VPC, SageMaker specifies both
//     network interfaces when launching this instance. This enables inbound traffic
//     from your own VPC to the notebook instance, assuming that the security groups
//     allow it.
//
// After creating the notebook instance, SageMaker returns its Amazon Resource
// Name (ARN). You can't change the name of a notebook instance after you create
// it. After SageMaker creates the notebook instance, you can connect to the
// Jupyter server and work in Jupyter notebooks. For example, you can write code to
// explore a dataset that you can use for model training, train a model, host
// models by creating SageMaker endpoints, and validate hosted models. For more
// information, see How It Works (https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html)
// .
func (c *Client) CreateNotebookInstance(ctx context.Context, params *CreateNotebookInstanceInput, optFns ...func(*Options)) (*CreateNotebookInstanceOutput, error) {
	if params == nil {
		params = &CreateNotebookInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNotebookInstance", params, optFns, c.addOperationCreateNotebookInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNotebookInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNotebookInstanceInput struct {

	// The type of ML compute instance to launch for the notebook instance.
	//
	// This member is required.
	InstanceType types.InstanceType

	// The name of the new notebook instance.
	//
	// This member is required.
	NotebookInstanceName *string

	// When you send any requests to Amazon Web Services resources from the notebook
	// instance, SageMaker assumes this role to perform tasks on your behalf. You must
	// grant this role necessary permissions so SageMaker can perform these tasks. The
	// policy must allow the SageMaker service principal (sagemaker.amazonaws.com)
	// permissions to assume this role. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html)
	// . To be able to pass this role to SageMaker, the caller of this API must have
	// the iam:PassRole permission.
	//
	// This member is required.
	RoleArn *string

	// A list of Elastic Inference (EI) instance types to associate with this notebook
	// instance. Currently, only one instance type can be associated with a notebook
	// instance. For more information, see Using Elastic Inference in Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html)
	// .
	AcceleratorTypes []types.NotebookInstanceAcceleratorType

	// An array of up to three Git repositories to associate with the notebook
	// instance. These can be either the names of Git repositories stored as resources
	// in your account, or the URL of Git repositories in Amazon Web Services
	// CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. These repositories are cloned at the same level
	// as the default repository of your notebook instance. For more information, see
	// Associating Git Repositories with SageMaker Notebook Instances (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html)
	// .
	AdditionalCodeRepositories []string

	// A Git repository to associate with the notebook instance as its default code
	// repository. This can be either the name of a Git repository stored as a resource
	// in your account, or the URL of a Git repository in Amazon Web Services
	// CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. When you open a notebook instance, it opens in
	// the directory that contains this repository. For more information, see
	// Associating Git Repositories with SageMaker Notebook Instances (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html)
	// .
	DefaultCodeRepository *string

	// Sets whether SageMaker provides internet access to the notebook instance. If
	// you set this to Disabled this notebook instance is able to access resources
	// only in your VPC, and is not be able to connect to SageMaker training and
	// endpoint services unless you configure a NAT Gateway in your VPC. For more
	// information, see Notebook Instances Are Internet-Enabled by Default (https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access)
	// . You can set the value of this parameter to Disabled only if you set a value
	// for the SubnetId parameter.
	DirectInternetAccess types.DirectInternetAccess

	// Information on the IMDS configuration of the notebook instance
	InstanceMetadataServiceConfiguration *types.InstanceMetadataServiceConfiguration

	// The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service
	// key that SageMaker uses to encrypt data on the storage volume attached to your
	// notebook instance. The KMS key you provide must be enabled. For information, see
	// Enabling and Disabling Keys (https://docs.aws.amazon.com/kms/latest/developerguide/enabling-keys.html)
	// in the Amazon Web Services Key Management Service Developer Guide.
	KmsKeyId *string

	// The name of a lifecycle configuration to associate with the notebook instance.
	// For information about lifestyle configurations, see Step 2.1: (Optional)
	// Customize a Notebook Instance (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html)
	// .
	LifecycleConfigName *string

	// The platform identifier of the notebook instance runtime environment.
	PlatformIdentifier *string

	// Whether root access is enabled or disabled for users of the notebook instance.
	// The default value is Enabled . Lifecycle configurations need root access to be
	// able to set up a notebook instance. Because of this, lifecycle configurations
	// associated with a notebook instance always run with root access even if you
	// disable root access for users.
	RootAccess types.RootAccess

	// The VPC security group IDs, in the form sg-xxxxxxxx. The security groups must
	// be for the same VPC as specified in the subnet.
	SecurityGroupIds []string

	// The ID of the subnet in a VPC to which you would like to have a connectivity
	// from your ML compute instance.
	SubnetId *string

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	Tags []types.Tag

	// The size, in GB, of the ML storage volume to attach to the notebook instance.
	// The default value is 5 GB.
	VolumeSizeInGB *int32

	noSmithyDocumentSerde
}

type CreateNotebookInstanceOutput struct {

	// The Amazon Resource Name (ARN) of the notebook instance.
	NotebookInstanceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNotebookInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNotebookInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNotebookInstance{}, middleware.After)
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
	if err = addCreateNotebookInstanceResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateNotebookInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNotebookInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateNotebookInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateNotebookInstance",
	}
}

type opCreateNotebookInstanceResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateNotebookInstanceResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateNotebookInstanceResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addCreateNotebookInstanceResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateNotebookInstanceResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
