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

// Lists information about associations between Amazon VPCs and query logging
// configurations.
func (c *Client) ListResolverQueryLogConfigAssociations(ctx context.Context, params *ListResolverQueryLogConfigAssociationsInput, optFns ...func(*Options)) (*ListResolverQueryLogConfigAssociationsOutput, error) {
	if params == nil {
		params = &ListResolverQueryLogConfigAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResolverQueryLogConfigAssociations", params, optFns, c.addOperationListResolverQueryLogConfigAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResolverQueryLogConfigAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverQueryLogConfigAssociationsInput struct {

	// An optional specification to return a subset of query logging associations. If
	// you submit a second or subsequent ListResolverQueryLogConfigAssociations
	// request and specify the NextToken parameter, you must use the same values for
	// Filters , if any, as in the previous request.
	Filters []types.Filter

	// The maximum number of query logging associations that you want to return in the
	// response to a ListResolverQueryLogConfigAssociations request. If you don't
	// specify a value for MaxResults , Resolver returns up to 100 query logging
	// associations.
	MaxResults *int32

	// For the first ListResolverQueryLogConfigAssociations request, omit this value.
	// If there are more than MaxResults query logging associations that match the
	// values that you specify for Filters , you can submit another
	// ListResolverQueryLogConfigAssociations request to get the next group of
	// associations. In the next request, specify the value of NextToken from the
	// previous response.
	NextToken *string

	// The element that you want Resolver to sort query logging associations by. If
	// you submit a second or subsequent ListResolverQueryLogConfigAssociations
	// request and specify the NextToken parameter, you must use the same value for
	// SortBy , if any, as in the previous request. Valid values include the following
	// elements:
	//   - CreationTime : The ID of the query logging association.
	//   - Error : If the value of Status is FAILED , the value of Error indicates the
	//   cause:
	//   - DESTINATION_NOT_FOUND : The specified destination (for example, an Amazon S3
	//   bucket) was deleted.
	//   - ACCESS_DENIED : Permissions don't allow sending logs to the destination. If
	//   Status is a value other than FAILED , ERROR is null.
	//   - Id : The ID of the query logging association
	//   - ResolverQueryLogConfigId : The ID of the query logging configuration
	//   - ResourceId : The ID of the VPC that is associated with the query logging
	//   configuration
	//   - Status : The current status of the configuration. Valid values include the
	//   following:
	//   - CREATING : Resolver is creating an association between an Amazon VPC and a
	//   query logging configuration.
	//   - CREATED : The association between an Amazon VPC and a query logging
	//   configuration was successfully created. Resolver is logging queries that
	//   originate in the specified VPC.
	//   - DELETING : Resolver is deleting this query logging association.
	//   - FAILED : Resolver either couldn't create or couldn't delete the query
	//   logging association. Here are two common causes:
	//   - The specified destination (for example, an Amazon S3 bucket) was deleted.
	//   - Permissions don't allow sending logs to the destination.
	SortBy *string

	// If you specified a value for SortBy , the order that you want query logging
	// associations to be listed in, ASCENDING or DESCENDING . If you submit a second
	// or subsequent ListResolverQueryLogConfigAssociations request and specify the
	// NextToken parameter, you must use the same value for SortOrder , if any, as in
	// the previous request.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListResolverQueryLogConfigAssociationsOutput struct {

	// If there are more than MaxResults query logging associations, you can submit
	// another ListResolverQueryLogConfigAssociations request to get the next group of
	// associations. In the next request, specify the value of NextToken from the
	// previous response.
	NextToken *string

	// A list that contains one ResolverQueryLogConfigAssociations element for each
	// query logging association that matches the values that you specified for Filter .
	ResolverQueryLogConfigAssociations []types.ResolverQueryLogConfigAssociation

	// The total number of query logging associations that were created by the current
	// account in the specified Region. This count can differ from the number of
	// associations that are returned in a ListResolverQueryLogConfigAssociations
	// response, depending on the values that you specify in the request.
	TotalCount int32

	// The total number of query logging associations that were created by the current
	// account in the specified Region and that match the filters that were specified
	// in the ListResolverQueryLogConfigAssociations request. For the total number of
	// associations that were created by the current account in the specified Region,
	// see TotalCount .
	TotalFilteredCount int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResolverQueryLogConfigAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverQueryLogConfigAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverQueryLogConfigAssociations{}, middleware.After)
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
	if err = addListResolverQueryLogConfigAssociationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverQueryLogConfigAssociations(options.Region), middleware.Before); err != nil {
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

// ListResolverQueryLogConfigAssociationsAPIClient is a client that implements the
// ListResolverQueryLogConfigAssociations operation.
type ListResolverQueryLogConfigAssociationsAPIClient interface {
	ListResolverQueryLogConfigAssociations(context.Context, *ListResolverQueryLogConfigAssociationsInput, ...func(*Options)) (*ListResolverQueryLogConfigAssociationsOutput, error)
}

var _ ListResolverQueryLogConfigAssociationsAPIClient = (*Client)(nil)

// ListResolverQueryLogConfigAssociationsPaginatorOptions is the paginator options
// for ListResolverQueryLogConfigAssociations
type ListResolverQueryLogConfigAssociationsPaginatorOptions struct {
	// The maximum number of query logging associations that you want to return in the
	// response to a ListResolverQueryLogConfigAssociations request. If you don't
	// specify a value for MaxResults , Resolver returns up to 100 query logging
	// associations.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResolverQueryLogConfigAssociationsPaginator is a paginator for
// ListResolverQueryLogConfigAssociations
type ListResolverQueryLogConfigAssociationsPaginator struct {
	options   ListResolverQueryLogConfigAssociationsPaginatorOptions
	client    ListResolverQueryLogConfigAssociationsAPIClient
	params    *ListResolverQueryLogConfigAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListResolverQueryLogConfigAssociationsPaginator returns a new
// ListResolverQueryLogConfigAssociationsPaginator
func NewListResolverQueryLogConfigAssociationsPaginator(client ListResolverQueryLogConfigAssociationsAPIClient, params *ListResolverQueryLogConfigAssociationsInput, optFns ...func(*ListResolverQueryLogConfigAssociationsPaginatorOptions)) *ListResolverQueryLogConfigAssociationsPaginator {
	if params == nil {
		params = &ListResolverQueryLogConfigAssociationsInput{}
	}

	options := ListResolverQueryLogConfigAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResolverQueryLogConfigAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResolverQueryLogConfigAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResolverQueryLogConfigAssociations page.
func (p *ListResolverQueryLogConfigAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResolverQueryLogConfigAssociationsOutput, error) {
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

	result, err := p.client.ListResolverQueryLogConfigAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResolverQueryLogConfigAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "ListResolverQueryLogConfigAssociations",
	}
}

type opListResolverQueryLogConfigAssociationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListResolverQueryLogConfigAssociationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListResolverQueryLogConfigAssociationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addListResolverQueryLogConfigAssociationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListResolverQueryLogConfigAssociationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}