// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all skills in the Alexa skill store by category.
//
// Deprecated: Alexa For Business is no longer supported
func (c *Client) ListSkillsStoreSkillsByCategory(ctx context.Context, params *ListSkillsStoreSkillsByCategoryInput, optFns ...func(*Options)) (*ListSkillsStoreSkillsByCategoryOutput, error) {
	if params == nil {
		params = &ListSkillsStoreSkillsByCategoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSkillsStoreSkillsByCategory", params, optFns, c.addOperationListSkillsStoreSkillsByCategoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSkillsStoreSkillsByCategoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSkillsStoreSkillsByCategoryInput struct {

	// The category ID for which the skills are being retrieved from the skill store.
	//
	// This member is required.
	CategoryId *int64

	// The maximum number of skills returned per paginated calls.
	MaxResults *int32

	// The tokens used for pagination.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSkillsStoreSkillsByCategoryOutput struct {

	// The tokens used for pagination.
	NextToken *string

	// The skill store skills.
	SkillsStoreSkills []types.SkillsStoreSkill

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSkillsStoreSkillsByCategoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSkillsStoreSkillsByCategory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSkillsStoreSkillsByCategory{}, middleware.After)
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
	if err = addListSkillsStoreSkillsByCategoryResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListSkillsStoreSkillsByCategoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSkillsStoreSkillsByCategory(options.Region), middleware.Before); err != nil {
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

// ListSkillsStoreSkillsByCategoryAPIClient is a client that implements the
// ListSkillsStoreSkillsByCategory operation.
type ListSkillsStoreSkillsByCategoryAPIClient interface {
	ListSkillsStoreSkillsByCategory(context.Context, *ListSkillsStoreSkillsByCategoryInput, ...func(*Options)) (*ListSkillsStoreSkillsByCategoryOutput, error)
}

var _ ListSkillsStoreSkillsByCategoryAPIClient = (*Client)(nil)

// ListSkillsStoreSkillsByCategoryPaginatorOptions is the paginator options for
// ListSkillsStoreSkillsByCategory
type ListSkillsStoreSkillsByCategoryPaginatorOptions struct {
	// The maximum number of skills returned per paginated calls.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSkillsStoreSkillsByCategoryPaginator is a paginator for
// ListSkillsStoreSkillsByCategory
type ListSkillsStoreSkillsByCategoryPaginator struct {
	options   ListSkillsStoreSkillsByCategoryPaginatorOptions
	client    ListSkillsStoreSkillsByCategoryAPIClient
	params    *ListSkillsStoreSkillsByCategoryInput
	nextToken *string
	firstPage bool
}

// NewListSkillsStoreSkillsByCategoryPaginator returns a new
// ListSkillsStoreSkillsByCategoryPaginator
func NewListSkillsStoreSkillsByCategoryPaginator(client ListSkillsStoreSkillsByCategoryAPIClient, params *ListSkillsStoreSkillsByCategoryInput, optFns ...func(*ListSkillsStoreSkillsByCategoryPaginatorOptions)) *ListSkillsStoreSkillsByCategoryPaginator {
	if params == nil {
		params = &ListSkillsStoreSkillsByCategoryInput{}
	}

	options := ListSkillsStoreSkillsByCategoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSkillsStoreSkillsByCategoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSkillsStoreSkillsByCategoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSkillsStoreSkillsByCategory page.
func (p *ListSkillsStoreSkillsByCategoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSkillsStoreSkillsByCategoryOutput, error) {
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

	result, err := p.client.ListSkillsStoreSkillsByCategory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSkillsStoreSkillsByCategory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ListSkillsStoreSkillsByCategory",
	}
}

type opListSkillsStoreSkillsByCategoryResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListSkillsStoreSkillsByCategoryResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListSkillsStoreSkillsByCategoryResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "a4b"
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
				signingName = "a4b"
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
				v4aScheme.SigningName = aws.String("a4b")
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

func addListSkillsStoreSkillsByCategoryResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListSkillsStoreSkillsByCategoryResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}