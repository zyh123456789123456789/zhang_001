// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of RepositoryAssociationSummary (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html)
// objects that contain summary information about a repository association. You can
// filter the returned list by ProviderType (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-ProviderType)
// , Name (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-Name)
// , State (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-State)
// , and Owner (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-Owner)
// .
func (c *Client) ListRepositoryAssociations(ctx context.Context, params *ListRepositoryAssociationsInput, optFns ...func(*Options)) (*ListRepositoryAssociationsOutput, error) {
	if params == nil {
		params = &ListRepositoryAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRepositoryAssociations", params, optFns, c.addOperationListRepositoryAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRepositoryAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRepositoryAssociationsInput struct {

	// The maximum number of repository association results returned by
	// ListRepositoryAssociations in paginated output. When this parameter is used,
	// ListRepositoryAssociations only returns maxResults results in a single page
	// with a nextToken response element. The remaining results of the initial request
	// can be seen by sending another ListRepositoryAssociations request with the
	// returned nextToken value. This value can be between 1 and 100. If this
	// parameter is not used, ListRepositoryAssociations returns up to 100 results and
	// a nextToken value if applicable.
	MaxResults *int32

	// List of repository names to use as a filter.
	Names []string

	// The nextToken value returned from a previous paginated
	// ListRepositoryAssociations request where maxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the nextToken value. Treat this token as an
	// opaque identifier that is only used to retrieve the next items in a list and not
	// for other programmatic purposes.
	NextToken *string

	// List of owners to use as a filter. For Amazon Web Services CodeCommit, it is
	// the name of the CodeCommit account that was used to associate the repository.
	// For other repository source providers, such as Bitbucket and GitHub Enterprise
	// Server, this is name of the account that was used to associate the repository.
	Owners []string

	// List of provider types to use as a filter.
	ProviderTypes []types.ProviderType

	// List of repository association states to use as a filter. The valid repository
	// association states are:
	//   - Associated: The repository association is complete.
	//   - Associating: CodeGuru Reviewer is:
	//   - Setting up pull request notifications. This is required for pull requests
	//   to trigger a CodeGuru Reviewer review. If your repository ProviderType is
	//   GitHub , GitHub Enterprise Server , or Bitbucket , CodeGuru Reviewer creates
	//   webhooks in your repository to trigger CodeGuru Reviewer reviews. If you delete
	//   these webhooks, reviews of code in your repository cannot be triggered.
	//   - Setting up source code access. This is required for CodeGuru Reviewer to
	//   securely clone code in your repository.
	//   - Failed: The repository failed to associate or disassociate.
	//   - Disassociating: CodeGuru Reviewer is removing the repository's pull request
	//   notifications and source code access.
	//   - Disassociated: CodeGuru Reviewer successfully disassociated the repository.
	//   You can create a new association with this repository if you want to review
	//   source code in it later. You can control access to code reviews created in
	//   anassociated repository with tags after it has been disassociated. For more
	//   information, see Using tags to control access to associated repositories (https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/auth-and-access-control-using-tags.html)
	//   in the Amazon CodeGuru Reviewer User Guide.
	States []types.RepositoryAssociationState

	noSmithyDocumentSerde
}

type ListRepositoryAssociationsOutput struct {

	// The nextToken value to include in a future ListRecommendations request. When
	// the results of a ListRecommendations request exceed maxResults , this value can
	// be used to retrieve the next page of results. This value is null when there are
	// no more results to return.
	NextToken *string

	// A list of repository associations that meet the criteria of the request.
	RepositoryAssociationSummaries []types.RepositoryAssociationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRepositoryAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRepositoryAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRepositoryAssociations{}, middleware.After)
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
	if err = addListRepositoryAssociationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositoryAssociations(options.Region), middleware.Before); err != nil {
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

// ListRepositoryAssociationsAPIClient is a client that implements the
// ListRepositoryAssociations operation.
type ListRepositoryAssociationsAPIClient interface {
	ListRepositoryAssociations(context.Context, *ListRepositoryAssociationsInput, ...func(*Options)) (*ListRepositoryAssociationsOutput, error)
}

var _ ListRepositoryAssociationsAPIClient = (*Client)(nil)

// ListRepositoryAssociationsPaginatorOptions is the paginator options for
// ListRepositoryAssociations
type ListRepositoryAssociationsPaginatorOptions struct {
	// The maximum number of repository association results returned by
	// ListRepositoryAssociations in paginated output. When this parameter is used,
	// ListRepositoryAssociations only returns maxResults results in a single page
	// with a nextToken response element. The remaining results of the initial request
	// can be seen by sending another ListRepositoryAssociations request with the
	// returned nextToken value. This value can be between 1 and 100. If this
	// parameter is not used, ListRepositoryAssociations returns up to 100 results and
	// a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRepositoryAssociationsPaginator is a paginator for
// ListRepositoryAssociations
type ListRepositoryAssociationsPaginator struct {
	options   ListRepositoryAssociationsPaginatorOptions
	client    ListRepositoryAssociationsAPIClient
	params    *ListRepositoryAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListRepositoryAssociationsPaginator returns a new
// ListRepositoryAssociationsPaginator
func NewListRepositoryAssociationsPaginator(client ListRepositoryAssociationsAPIClient, params *ListRepositoryAssociationsInput, optFns ...func(*ListRepositoryAssociationsPaginatorOptions)) *ListRepositoryAssociationsPaginator {
	if params == nil {
		params = &ListRepositoryAssociationsInput{}
	}

	options := ListRepositoryAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRepositoryAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRepositoryAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRepositoryAssociations page.
func (p *ListRepositoryAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRepositoryAssociationsOutput, error) {
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

	result, err := p.client.ListRepositoryAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRepositoryAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "ListRepositoryAssociations",
	}
}

type opListRepositoryAssociationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListRepositoryAssociationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListRepositoryAssociationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "codeguru-reviewer"
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
				signingName = "codeguru-reviewer"
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
				v4aScheme.SigningName = aws.String("codeguru-reviewer")
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

func addListRepositoryAssociationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListRepositoryAssociationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}