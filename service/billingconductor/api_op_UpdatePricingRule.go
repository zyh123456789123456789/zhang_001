// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing pricing rule.
func (c *Client) UpdatePricingRule(ctx context.Context, params *UpdatePricingRuleInput, optFns ...func(*Options)) (*UpdatePricingRuleOutput, error) {
	if params == nil {
		params = &UpdatePricingRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePricingRule", params, optFns, c.addOperationUpdatePricingRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePricingRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePricingRuleInput struct {

	// The Amazon Resource Name (ARN) of the pricing rule to update.
	//
	// This member is required.
	Arn *string

	// The new description for the pricing rule.
	Description *string

	// The new modifier to show pricing plan rates as a percentage.
	ModifierPercentage *float64

	// The new name of the pricing rule. The name must be unique to each pricing rule.
	Name *string

	// The set of tiering configurations for the pricing rule.
	Tiering *types.UpdateTieringInput

	// The new pricing rule type.
	Type types.PricingRuleType

	noSmithyDocumentSerde
}

type UpdatePricingRuleOutput struct {

	// The Amazon Resource Name (ARN) of the successfully updated pricing rule.
	Arn *string

	// The pricing plans count that this pricing rule is associated with.
	AssociatedPricingPlanCount int64

	// The seller of services provided by Amazon Web Services, their affiliates, or
	// third-party providers selling services via Amazon Web Services Marketplace.
	BillingEntity *string

	// The new description for the pricing rule.
	Description *string

	// The most recent time the pricing rule was modified.
	LastModifiedTime int64

	// The new modifier to show pricing plan rates as a percentage.
	ModifierPercentage *float64

	// The new name of the pricing rule. The name must be unique to each pricing rule.
	Name *string

	// Operation refers to the specific Amazon Web Services covered by this line item.
	// This describes the specific usage of the line item. If the Scope attribute is
	// set to SKU , this attribute indicates which operation the PricingRule is
	// modifying. For example, a value of RunInstances:0202 indicates the operation of
	// running an Amazon EC2 instance.
	Operation *string

	// The scope of pricing rule that indicates if it's globally applicable, or it's
	// service-specific.
	Scope types.PricingRuleScope

	// If the Scope attribute is set to SERVICE , the attribute indicates which service
	// the PricingRule is applicable for.
	Service *string

	// The set of tiering configurations for the pricing rule.
	Tiering *types.UpdateTieringInput

	// The new pricing rule type.
	Type types.PricingRuleType

	// Usage type is the unit that each service uses to measure the usage of a
	// specific type of resource. If the Scope attribute is set to SKU , this attribute
	// indicates which usage type the PricingRule is modifying. For example,
	// USW2-BoxUsage:m2.2xlarge describes an M2 High Memory Double Extra Large
	// instance in the US West (Oregon) Region.
	UsageType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePricingRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePricingRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePricingRule{}, middleware.After)
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
	if err = addUpdatePricingRuleResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdatePricingRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePricingRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePricingRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "UpdatePricingRule",
	}
}

type opUpdatePricingRuleResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdatePricingRuleResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdatePricingRuleResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "billingconductor"
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
				signingName = "billingconductor"
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
				v4aScheme.SigningName = aws.String("billingconductor")
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

func addUpdatePricingRuleResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdatePricingRuleResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
