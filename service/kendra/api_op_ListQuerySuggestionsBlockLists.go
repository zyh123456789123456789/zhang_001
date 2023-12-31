// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the block lists used for query suggestions for an index. For information
// on the current quota limits for block lists, see Quotas for Amazon Kendra (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html)
// . ListQuerySuggestionsBlockLists is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
func (c *Client) ListQuerySuggestionsBlockLists(ctx context.Context, params *ListQuerySuggestionsBlockListsInput, optFns ...func(*Options)) (*ListQuerySuggestionsBlockListsOutput, error) {
	if params == nil {
		params = &ListQuerySuggestionsBlockListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQuerySuggestionsBlockLists", params, optFns, c.addOperationListQuerySuggestionsBlockListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQuerySuggestionsBlockListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQuerySuggestionsBlockListsInput struct {

	// The identifier of the index for a list of all block lists that exist for that
	// index. For information on the current quota limits for block lists, see Quotas
	// for Amazon Kendra (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html) .
	//
	// This member is required.
	IndexId *string

	// The maximum number of block lists to return.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Kendra returns a pagination token in the response. You can use
	// this pagination token to retrieve the next set of block lists (
	// BlockListSummaryItems ).
	NextToken *string

	noSmithyDocumentSerde
}

type ListQuerySuggestionsBlockListsOutput struct {

	// Summary items for a block list. This includes summary items on the block list
	// ID, block list name, when the block list was created, when the block list was
	// last updated, and the count of block words/phrases in the block list. For
	// information on the current quota limits for block lists, see Quotas for Amazon
	// Kendra (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html) .
	BlockListSummaryItems []types.QuerySuggestionsBlockListSummary

	// If the response is truncated, Amazon Kendra returns this token that you can use
	// in the subsequent request to retrieve the next set of block lists.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQuerySuggestionsBlockListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListQuerySuggestionsBlockLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListQuerySuggestionsBlockLists{}, middleware.After)
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
	if err = addListQuerySuggestionsBlockListsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListQuerySuggestionsBlockListsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQuerySuggestionsBlockLists(options.Region), middleware.Before); err != nil {
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

// ListQuerySuggestionsBlockListsAPIClient is a client that implements the
// ListQuerySuggestionsBlockLists operation.
type ListQuerySuggestionsBlockListsAPIClient interface {
	ListQuerySuggestionsBlockLists(context.Context, *ListQuerySuggestionsBlockListsInput, ...func(*Options)) (*ListQuerySuggestionsBlockListsOutput, error)
}

var _ ListQuerySuggestionsBlockListsAPIClient = (*Client)(nil)

// ListQuerySuggestionsBlockListsPaginatorOptions is the paginator options for
// ListQuerySuggestionsBlockLists
type ListQuerySuggestionsBlockListsPaginatorOptions struct {
	// The maximum number of block lists to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQuerySuggestionsBlockListsPaginator is a paginator for
// ListQuerySuggestionsBlockLists
type ListQuerySuggestionsBlockListsPaginator struct {
	options   ListQuerySuggestionsBlockListsPaginatorOptions
	client    ListQuerySuggestionsBlockListsAPIClient
	params    *ListQuerySuggestionsBlockListsInput
	nextToken *string
	firstPage bool
}

// NewListQuerySuggestionsBlockListsPaginator returns a new
// ListQuerySuggestionsBlockListsPaginator
func NewListQuerySuggestionsBlockListsPaginator(client ListQuerySuggestionsBlockListsAPIClient, params *ListQuerySuggestionsBlockListsInput, optFns ...func(*ListQuerySuggestionsBlockListsPaginatorOptions)) *ListQuerySuggestionsBlockListsPaginator {
	if params == nil {
		params = &ListQuerySuggestionsBlockListsInput{}
	}

	options := ListQuerySuggestionsBlockListsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQuerySuggestionsBlockListsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQuerySuggestionsBlockListsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQuerySuggestionsBlockLists page.
func (p *ListQuerySuggestionsBlockListsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQuerySuggestionsBlockListsOutput, error) {
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

	result, err := p.client.ListQuerySuggestionsBlockLists(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListQuerySuggestionsBlockLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "ListQuerySuggestionsBlockLists",
	}
}

type opListQuerySuggestionsBlockListsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListQuerySuggestionsBlockListsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListQuerySuggestionsBlockListsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "kendra"
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
				signingName = "kendra"
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
				v4aScheme.SigningName = aws.String("kendra")
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

func addListQuerySuggestionsBlockListsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListQuerySuggestionsBlockListsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
