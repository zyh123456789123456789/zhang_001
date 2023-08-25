// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves cost and usage metrics with resources for your account. You can
// specify which cost and usage-related metric, such as BlendedCosts or
// UsageQuantity , that you want the request to return. You can also filter and
// group your data by various dimensions, such as SERVICE or AZ , in a specific
// time range. For a complete list of valid dimensions, see the GetDimensionValues (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetDimensionValues.html)
// operation. Management account in an organization in Organizations have access to
// all member accounts. This API is currently available for the Amazon Elastic
// Compute Cloud – Compute service only. This is an opt-in only feature. You can
// enable this feature from the Cost Explorer Settings page. For information about
// how to access the Settings page, see Controlling Access for Cost Explorer (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/ce-access.html)
// in the Billing and Cost Management User Guide.
func (c *Client) GetCostAndUsageWithResources(ctx context.Context, params *GetCostAndUsageWithResourcesInput, optFns ...func(*Options)) (*GetCostAndUsageWithResourcesOutput, error) {
	if params == nil {
		params = &GetCostAndUsageWithResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCostAndUsageWithResources", params, optFns, c.addOperationGetCostAndUsageWithResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCostAndUsageWithResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCostAndUsageWithResourcesInput struct {

	// Filters Amazon Web Services costs by different dimensions. For example, you can
	// specify SERVICE and LINKED_ACCOUNT and get the costs that are associated with
	// that account's usage of that service. You can nest Expression objects to define
	// any combination of dimension filters. For more information, see Expression (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// . The GetCostAndUsageWithResources operation requires that you either group by
	// or filter by a ResourceId . It requires the Expression (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// "SERVICE = Amazon Elastic Compute Cloud - Compute" in the filter. Valid values
	// for MatchOptions for Dimensions are EQUALS and CASE_SENSITIVE . Valid values for
	// MatchOptions for CostCategories and Tags are EQUALS , ABSENT , and
	// CASE_SENSITIVE . Default values are EQUALS and CASE_SENSITIVE .
	//
	// This member is required.
	Filter *types.Expression

	// Sets the Amazon Web Services cost granularity to MONTHLY , DAILY , or HOURLY .
	// If Granularity isn't set, the response object doesn't include the Granularity ,
	// MONTHLY , DAILY , or HOURLY .
	//
	// This member is required.
	Granularity types.Granularity

	// Sets the start and end dates for retrieving Amazon Web Services costs. The
	// range must be within the last 14 days (the start date cannot be earlier than 14
	// days ago). The start date is inclusive, but the end date is exclusive. For
	// example, if start is 2017-01-01 and end is 2017-05-01 , then the cost and usage
	// data is retrieved from 2017-01-01 up to and including 2017-04-30 but not
	// including 2017-05-01 .
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// You can group Amazon Web Services costs using up to two different groups:
	// DIMENSION , TAG , COST_CATEGORY .
	GroupBy []types.GroupDefinition

	// Which metrics are returned in the query. For more information about blended and
	// unblended rates, see Why does the "blended" annotation appear on some line
	// items in my bill? (http://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/)
	// . Valid values are AmortizedCost , BlendedCost , NetAmortizedCost ,
	// NetUnblendedCost , NormalizedUsageAmount , UnblendedCost , and UsageQuantity .
	// If you return the UsageQuantity metric, the service aggregates all usage
	// numbers without taking the units into account. For example, if you aggregate
	// usageQuantity across all of Amazon EC2, the results aren't meaningful because
	// Amazon EC2 compute hours and data transfer are measured in different units (for
	// example, hour or GB). To get more meaningful UsageQuantity metrics, filter by
	// UsageType or UsageTypeGroups . Metrics is required for
	// GetCostAndUsageWithResources requests.
	Metrics []string

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	noSmithyDocumentSerde
}

type GetCostAndUsageWithResourcesOutput struct {

	// The attributes that apply to a specific dimension value. For example, if the
	// value is a linked account, the attribute is that account name.
	DimensionValueAttributes []types.DimensionValuesWithAttributes

	// The groups that are specified by the Filter or GroupBy parameters in the
	// request.
	GroupDefinitions []types.GroupDefinition

	// The token for the next set of retrievable results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextPageToken *string

	// The time period that's covered by the results in the response.
	ResultsByTime []types.ResultByTime

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCostAndUsageWithResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCostAndUsageWithResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCostAndUsageWithResources{}, middleware.After)
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
	if err = addGetCostAndUsageWithResourcesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetCostAndUsageWithResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCostAndUsageWithResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCostAndUsageWithResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetCostAndUsageWithResources",
	}
}

type opGetCostAndUsageWithResourcesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetCostAndUsageWithResourcesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetCostAndUsageWithResourcesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ce"
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
				signingName = "ce"
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
				v4aScheme.SigningName = aws.String("ce")
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

func addGetCostAndUsageWithResourcesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetCostAndUsageWithResourcesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}