// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an instance refresh. During an instance refresh, Amazon EC2 Auto Scaling
// performs a rolling update of instances in an Auto Scaling group. Instances are
// terminated first and then replaced, which temporarily reduces the capacity
// available within your Auto Scaling group. This operation is part of the
// instance refresh feature (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html)
// in Amazon EC2 Auto Scaling, which helps you update instances in your Auto
// Scaling group. This feature is helpful, for example, when you have a new AMI or
// a new user data script. You just need to create a new launch template that
// specifies the new AMI or user data script. Then start an instance refresh to
// immediately begin the process of updating instances in the group. If successful,
// the request's response contains a unique ID that you can use to track the
// progress of the instance refresh. To query its status, call the
// DescribeInstanceRefreshes API. To describe the instance refreshes that have
// already run, call the DescribeInstanceRefreshes API. To cancel an instance
// refresh that is in progress, use the CancelInstanceRefresh API. An instance
// refresh might fail for several reasons, such as EC2 launch failures,
// misconfigured health checks, or not ignoring or allowing the termination of
// instances that are in Standby state or protected from scale in. You can monitor
// for failed EC2 launches using the scaling activities. To find the scaling
// activities, call the DescribeScalingActivities API. If you enable auto
// rollback, your Auto Scaling group will be rolled back automatically when the
// instance refresh fails. You can enable this feature before starting an instance
// refresh by specifying the AutoRollback property in the instance refresh
// preferences. Otherwise, to roll back an instance refresh before it finishes, use
// the RollbackInstanceRefresh API.
func (c *Client) StartInstanceRefresh(ctx context.Context, params *StartInstanceRefreshInput, optFns ...func(*Options)) (*StartInstanceRefreshOutput, error) {
	if params == nil {
		params = &StartInstanceRefreshInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartInstanceRefresh", params, optFns, c.addOperationStartInstanceRefreshMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartInstanceRefreshOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartInstanceRefreshInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The desired configuration. For example, the desired configuration can specify a
	// new launch template or a new version of the current launch template. Once the
	// instance refresh succeeds, Amazon EC2 Auto Scaling updates the settings of the
	// Auto Scaling group to reflect the new desired configuration. When you specify a
	// new launch template or a new version of the current launch template for your
	// desired configuration, consider enabling the SkipMatching property in
	// preferences. If it's enabled, Amazon EC2 Auto Scaling skips replacing instances
	// that already use the specified launch template and instance types. This can help
	// you reduce the number of replacements that are required to apply updates.
	DesiredConfiguration *types.DesiredConfiguration

	// Sets your preferences for the instance refresh so that it performs as expected
	// when you start it. Includes the instance warmup time, the minimum healthy
	// percentage, and the behaviors that you want Amazon EC2 Auto Scaling to use if
	// instances that are in Standby state or protected from scale in are found. You
	// can also choose to enable additional features, such as the following:
	//   - Auto rollback
	//   - Checkpoints
	//   - CloudWatch alarms
	//   - Skip matching
	Preferences *types.RefreshPreferences

	// The strategy to use for the instance refresh. The only valid value is Rolling .
	Strategy types.RefreshStrategy

	noSmithyDocumentSerde
}

type StartInstanceRefreshOutput struct {

	// A unique ID for tracking the progress of the instance refresh.
	InstanceRefreshId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartInstanceRefreshMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStartInstanceRefresh{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStartInstanceRefresh{}, middleware.After)
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
	if err = addStartInstanceRefreshResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartInstanceRefreshValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartInstanceRefresh(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartInstanceRefresh(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "StartInstanceRefresh",
	}
}

type opStartInstanceRefreshResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartInstanceRefreshResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartInstanceRefreshResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "autoscaling"
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
				signingName = "autoscaling"
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
				v4aScheme.SigningName = aws.String("autoscaling")
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

func addStartInstanceRefreshResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartInstanceRefreshResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
