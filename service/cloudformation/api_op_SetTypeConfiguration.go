// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specifies the configuration data for a registered CloudFormation extension, in
// the given account and Region. To view the current configuration data for an
// extension, refer to the ConfigurationSchema element of DescribeType (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeType.html)
// . For more information, see Configuring extensions at the account level (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-register.html#registry-set-configuration)
// in the CloudFormation User Guide. It's strongly recommended that you use dynamic
// references to restrict sensitive configuration definitions, such as third-party
// credentials. For more details on dynamic references, see Using dynamic
// references to specify template values (https://docs.aws.amazon.com/) in the
// CloudFormation User Guide.
func (c *Client) SetTypeConfiguration(ctx context.Context, params *SetTypeConfigurationInput, optFns ...func(*Options)) (*SetTypeConfigurationOutput, error) {
	if params == nil {
		params = &SetTypeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTypeConfiguration", params, optFns, c.addOperationSetTypeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTypeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetTypeConfigurationInput struct {

	// The configuration data for the extension, in this account and Region. The
	// configuration data must be formatted as JSON, and validate against the schema
	// returned in the ConfigurationSchema response element of DescribeType (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeType.html)
	// . For more information, see Defining account-level configuration data for an
	// extension (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-model.html#resource-type-howto-configuration)
	// in the CloudFormation CLI User Guide.
	//
	// This member is required.
	Configuration *string

	// An alias by which to refer to this extension configuration data. Conditional:
	// Specifying a configuration alias is required when setting a configuration for a
	// resource type extension.
	ConfigurationAlias *string

	// The type of extension. Conditional: You must specify ConfigurationArn , or Type
	// and TypeName .
	Type types.ThirdPartyType

	// The Amazon Resource Name (ARN) for the extension, in this account and Region.
	// For public extensions, this will be the ARN assigned when you activate the type (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ActivateType.html)
	// in this account and Region. For private extensions, this will be the ARN
	// assigned when you register the type (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterType.html)
	// in this account and Region. Do not include the extension versions suffix at the
	// end of the ARN. You can set the configuration for an extension, but not for a
	// specific extension version.
	TypeArn *string

	// The name of the extension. Conditional: You must specify ConfigurationArn , or
	// Type and TypeName .
	TypeName *string

	noSmithyDocumentSerde
}

type SetTypeConfigurationOutput struct {

	// The Amazon Resource Name (ARN) for the configuration data, in this account and
	// Region. Conditional: You must specify ConfigurationArn , or Type and TypeName .
	ConfigurationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetTypeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetTypeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetTypeConfiguration{}, middleware.After)
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
	if err = addSetTypeConfigurationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSetTypeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetTypeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetTypeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "SetTypeConfiguration",
	}
}

type opSetTypeConfigurationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opSetTypeConfigurationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSetTypeConfigurationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "cloudformation"
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
				signingName = "cloudformation"
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
				v4aScheme.SigningName = aws.String("cloudformation")
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

func addSetTypeConfigurationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opSetTypeConfigurationResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
