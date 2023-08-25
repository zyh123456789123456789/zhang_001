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
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a specified stack. Once the call completes successfully, stack deletion
// starts. Deleted stacks don't show up in the DescribeStacks operation if the
// deletion has been completed successfully.
func (c *Client) DeleteStack(ctx context.Context, params *DeleteStackInput, optFns ...func(*Options)) (*DeleteStackOutput, error) {
	if params == nil {
		params = &DeleteStackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStack", params, optFns, c.addOperationDeleteStackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for DeleteStack action.
type DeleteStackInput struct {

	// The name or the unique stack ID that's associated with the stack.
	//
	// This member is required.
	StackName *string

	// A unique identifier for this DeleteStack request. Specify this token if you
	// plan to retry requests so that CloudFormation knows that you're not attempting
	// to delete a stack with the same name. You might retry DeleteStack requests to
	// ensure that CloudFormation successfully received them. All events initiated by a
	// given stack operation are assigned the same client request token, which you can
	// use to track operations. For example, if you execute a CreateStack operation
	// with the token token1 , then all the StackEvents generated by that operation
	// will have ClientRequestToken set as token1 . In the console, stack operations
	// display the client request token on the Events tab. Stack operations that are
	// initiated from the console use the token format Console-StackOperation-ID, which
	// helps you easily identify the stack operation . For example, if you create a
	// stack using the console, each stack event would be assigned the same token in
	// the following format: Console-CreateStack-7f59c3cf-00d2-40c7-b2ff-e75db0987002 .
	ClientRequestToken *string

	// For stacks in the DELETE_FAILED state, a list of resource logical IDs that are
	// associated with the resources you want to retain. During deletion,
	// CloudFormation deletes the stack but doesn't delete the retained resources.
	// Retaining resources is useful when you can't delete a resource, such as a
	// non-empty S3 bucket, but you want to delete the stack.
	RetainResources []string

	// The Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role
	// that CloudFormation assumes to delete the stack. CloudFormation uses the role's
	// credentials to make calls on your behalf. If you don't specify a value,
	// CloudFormation uses the role that was previously associated with the stack. If
	// no role is available, CloudFormation uses a temporary session that's generated
	// from your user credentials.
	RoleARN *string

	noSmithyDocumentSerde
}

type DeleteStackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteStackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteStack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteStack{}, middleware.After)
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
	if err = addDeleteStackResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteStackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteStack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DeleteStack",
	}
}

type opDeleteStackResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDeleteStackResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDeleteStackResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addDeleteStackResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDeleteStackResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}