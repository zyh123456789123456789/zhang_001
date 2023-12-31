// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this action to manage EFS lifecycle management and EFS Intelligent-Tiering.
// A LifecycleConfiguration consists of one or more LifecyclePolicy objects that
// define the following:
//   - EFS Lifecycle management - When Amazon EFS automatically transitions files
//     in a file system into the lower-cost EFS Infrequent Access (IA) storage class.
//     To enable EFS Lifecycle management, set the value of TransitionToIA to one of
//     the available options.
//   - EFS Intelligent-Tiering - When Amazon EFS automatically transitions files
//     from IA back into the file system's primary storage class (EFS Standard or EFS
//     One Zone Standard). To enable EFS Intelligent-Tiering, set the value of
//     TransitionToPrimaryStorageClass to AFTER_1_ACCESS .
//
// For more information, see EFS Lifecycle Management (https://docs.aws.amazon.com/efs/latest/ug/lifecycle-management-efs.html)
// . Each Amazon EFS file system supports one lifecycle configuration, which
// applies to all files in the file system. If a LifecycleConfiguration object
// already exists for the specified file system, a PutLifecycleConfiguration call
// modifies the existing configuration. A PutLifecycleConfiguration call with an
// empty LifecyclePolicies array in the request body deletes any existing
// LifecycleConfiguration and turns off lifecycle management and EFS
// Intelligent-Tiering for the file system. In the request, specify the following:
//   - The ID for the file system for which you are enabling, disabling, or
//     modifying lifecycle management and EFS Intelligent-Tiering.
//   - A LifecyclePolicies array of LifecyclePolicy objects that define when files
//     are moved into IA storage, and when they are moved back to Standard storage.
//     Amazon EFS requires that each LifecyclePolicy object have only have a single
//     transition, so the LifecyclePolicies array needs to be structured with
//     separate LifecyclePolicy objects. See the example requests in the following
//     section for more information.
//
// This operation requires permissions for the
// elasticfilesystem:PutLifecycleConfiguration operation. To apply a
// LifecycleConfiguration object to an encrypted file system, you need the same Key
// Management Service permissions as when you created the encrypted file system.
func (c *Client) PutLifecycleConfiguration(ctx context.Context, params *PutLifecycleConfigurationInput, optFns ...func(*Options)) (*PutLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &PutLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLifecycleConfiguration", params, optFns, c.addOperationPutLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLifecycleConfigurationInput struct {

	// The ID of the file system for which you are creating the LifecycleConfiguration
	// object (String).
	//
	// This member is required.
	FileSystemId *string

	// An array of LifecyclePolicy objects that define the file system's
	// LifecycleConfiguration object. A LifecycleConfiguration object informs EFS
	// lifecycle management and EFS Intelligent-Tiering of the following:
	//   - When to move files in the file system from primary storage to the IA
	//   storage class.
	//   - When to move files that are in IA storage to primary storage.
	// When using the put-lifecycle-configuration CLI command or the
	// PutLifecycleConfiguration API action, Amazon EFS requires that each
	// LifecyclePolicy object have only a single transition. This means that in a
	// request body, LifecyclePolicies must be structured as an array of
	// LifecyclePolicy objects, one object for each transition, TransitionToIA ,
	// TransitionToPrimaryStorageClass . See the example requests in the following
	// section for more information.
	//
	// This member is required.
	LifecyclePolicies []types.LifecyclePolicy

	noSmithyDocumentSerde
}

type PutLifecycleConfigurationOutput struct {

	// An array of lifecycle management policies. EFS supports a maximum of one policy
	// per file system.
	LifecyclePolicies []types.LifecyclePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLifecycleConfiguration{}, middleware.After)
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
	if err = addPutLifecycleConfigurationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutLifecycleConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLifecycleConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "PutLifecycleConfiguration",
	}
}

type opPutLifecycleConfigurationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPutLifecycleConfigurationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutLifecycleConfigurationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "elasticfilesystem"
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
				signingName = "elasticfilesystem"
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
				v4aScheme.SigningName = aws.String("elasticfilesystem")
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

func addPutLifecycleConfigurationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutLifecycleConfigurationResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
