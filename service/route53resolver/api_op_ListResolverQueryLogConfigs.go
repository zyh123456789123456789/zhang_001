// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists information about the specified query logging configurations. Each
// configuration defines where you want Resolver to save DNS query logs and
// specifies the VPCs that you want to log queries for.
func (c *Client) ListResolverQueryLogConfigs(ctx context.Context, params *ListResolverQueryLogConfigsInput, optFns ...func(*Options)) (*ListResolverQueryLogConfigsOutput, error) {
	if params == nil {
		params = &ListResolverQueryLogConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResolverQueryLogConfigs", params, optFns, c.addOperationListResolverQueryLogConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResolverQueryLogConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverQueryLogConfigsInput struct {

	// An optional specification to return a subset of query logging configurations.
	// If you submit a second or subsequent ListResolverQueryLogConfigs request and
	// specify the NextToken parameter, you must use the same values for Filters , if
	// any, as in the previous request.
	Filters []types.Filter

	// The maximum number of query logging configurations that you want to return in
	// the response to a ListResolverQueryLogConfigs request. If you don't specify a
	// value for MaxResults , Resolver returns up to 100 query logging configurations.
	MaxResults *int32

	// For the first ListResolverQueryLogConfigs request, omit this value. If there
	// are more than MaxResults query logging configurations that match the values
	// that you specify for Filters , you can submit another
	// ListResolverQueryLogConfigs request to get the next group of configurations. In
	// the next request, specify the value of NextToken from the previous response.
	NextToken *string

	// The element that you want Resolver to sort query logging configurations by. If
	// you submit a second or subsequent ListResolverQueryLogConfigs request and
	// specify the NextToken parameter, you must use the same value for SortBy , if
	// any, as in the previous request. Valid values include the following elements:
	//   - Arn : The ARN of the query logging configuration
	//   - AssociationCount : The number of VPCs that are associated with the specified
	//   configuration
	//   - CreationTime : The date and time that Resolver returned when the
	//   configuration was created
	//   - CreatorRequestId : The value that was specified for CreatorRequestId when
	//   the configuration was created
	//   - DestinationArn : The location that logs are sent to
	//   - Id : The ID of the configuration
	//   - Name : The name of the configuration
	//   - OwnerId : The Amazon Web Services account number of the account that created
	//   the configuration
	//   - ShareStatus : Whether the configuration is shared with other Amazon Web
	//   Services accounts or shared with the current account by another Amazon Web
	//   Services account. Sharing is configured through Resource Access Manager (RAM).
	//   - Status : The current status of the configuration. Valid values include the
	//   following:
	//   - CREATING : Resolver is creating the query logging configuration.
	//   - CREATED : The query logging configuration was successfully created. Resolver
	//   is logging queries that originate in the specified VPC.
	//   - DELETING : Resolver is deleting this query logging configuration.
	//   - FAILED : Resolver either couldn't create or couldn't delete the query
	//   logging configuration. Here are two common causes:
	//   - The specified destination (for example, an Amazon S3 bucket) was deleted.
	//   - Permissions don't allow sending logs to the destination.
	SortBy *string

	// If you specified a value for SortBy , the order that you want query logging
	// configurations to be listed in, ASCENDING or DESCENDING . If you submit a second
	// or subsequent ListResolverQueryLogConfigs request and specify the NextToken
	// parameter, you must use the same value for SortOrder , if any, as in the
	// previous request.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListResolverQueryLogConfigsOutput struct {

	// If there are more than MaxResults query logging configurations, you can submit
	// another ListResolverQueryLogConfigs request to get the next group of
	// configurations. In the next request, specify the value of NextToken from the
	// previous response.
	NextToken *string

	// A list that contains one ResolverQueryLogConfig element for each query logging
	// configuration that matches the values that you specified for Filter .
	ResolverQueryLogConfigs []types.ResolverQueryLogConfig

	// The total number of query logging configurations that were created by the
	// current account in the specified Region. This count can differ from the number
	// of query logging configurations that are returned in a
	// ListResolverQueryLogConfigs response, depending on the values that you specify
	// in the request.
	TotalCount int32

	// The total number of query logging configurations that were created by the
	// current account in the specified Region and that match the filters that were
	// specified in the ListResolverQueryLogConfigs request. For the total number of
	// query logging configurations that were created by the current account in the
	// specified Region, see TotalCount .
	TotalFilteredCount int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResolverQueryLogConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverQueryLogConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverQueryLogConfigs{}, middleware.After)
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
	if err = addListResolverQueryLogConfigsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverQueryLogConfigs(options.Region), middleware.Before); err != nil {
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

// ListResolverQueryLogConfigsAPIClient is a client that implements the
// ListResolverQueryLogConfigs operation.
type ListResolverQueryLogConfigsAPIClient interface {
	ListResolverQueryLogConfigs(context.Context, *ListResolverQueryLogConfigsInput, ...func(*Options)) (*ListResolverQueryLogConfigsOutput, error)
}

var _ ListResolverQueryLogConfigsAPIClient = (*Client)(nil)

// ListResolverQueryLogConfigsPaginatorOptions is the paginator options for
// ListResolverQueryLogConfigs
type ListResolverQueryLogConfigsPaginatorOptions struct {
	// The maximum number of query logging configurations that you want to return in
	// the response to a ListResolverQueryLogConfigs request. If you don't specify a
	// value for MaxResults , Resolver returns up to 100 query logging configurations.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResolverQueryLogConfigsPaginator is a paginator for
// ListResolverQueryLogConfigs
type ListResolverQueryLogConfigsPaginator struct {
	options   ListResolverQueryLogConfigsPaginatorOptions
	client    ListResolverQueryLogConfigsAPIClient
	params    *ListResolverQueryLogConfigsInput
	nextToken *string
	firstPage bool
}

// NewListResolverQueryLogConfigsPaginator returns a new
// ListResolverQueryLogConfigsPaginator
func NewListResolverQueryLogConfigsPaginator(client ListResolverQueryLogConfigsAPIClient, params *ListResolverQueryLogConfigsInput, optFns ...func(*ListResolverQueryLogConfigsPaginatorOptions)) *ListResolverQueryLogConfigsPaginator {
	if params == nil {
		params = &ListResolverQueryLogConfigsInput{}
	}

	options := ListResolverQueryLogConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResolverQueryLogConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResolverQueryLogConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResolverQueryLogConfigs page.
func (p *ListResolverQueryLogConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResolverQueryLogConfigsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListResolverQueryLogConfigs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListResolverQueryLogConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "ListResolverQueryLogConfigs",
	}
}

type opListResolverQueryLogConfigsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListResolverQueryLogConfigsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListResolverQueryLogConfigsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "route53resolver"
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
				signingName = "route53resolver"
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
				v4aScheme.SigningName = aws.String("route53resolver")
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

func addListResolverQueryLogConfigsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListResolverQueryLogConfigsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
