// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the current handshakes that are associated with the account of the
// requesting user. Handshakes that are ACCEPTED , DECLINED , CANCELED , or EXPIRED
// appear in the results of this API for only 30 days after changing to that state.
// After that, they're deleted and no longer accessible. Always check the NextToken
// response parameter for a null value when calling a List* operation. These
// operations can occasionally return an empty set of results even when there are
// more results available. The NextToken response parameter value is null only
// when there are no more results to display. This operation can be called from any
// account in the organization.
func (c *Client) ListHandshakesForAccount(ctx context.Context, params *ListHandshakesForAccountInput, optFns ...func(*Options)) (*ListHandshakesForAccountOutput, error) {
	if params == nil {
		params = &ListHandshakesForAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHandshakesForAccount", params, optFns, c.addOperationListHandshakesForAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHandshakesForAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHandshakesForAccountInput struct {

	// Filters the handshakes that you want included in the response. The default is
	// all types. Use the ActionType element to limit the output to only a specified
	// type, such as INVITE , ENABLE_ALL_FEATURES , or APPROVE_ALL_FEATURES .
	// Alternatively, for the ENABLE_ALL_FEATURES handshake that generates a separate
	// child handshake for each member account, you can specify ParentHandshakeId to
	// see only the handshakes that were generated by that parent request.
	Filter *types.HandshakeFilter

	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	noSmithyDocumentSerde
}

type ListHandshakesForAccountOutput struct {

	// A list of Handshake objects with details about each of the handshakes that is
	// associated with the specified account.
	Handshakes []types.Handshake

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null .
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHandshakesForAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListHandshakesForAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHandshakesForAccount{}, middleware.After)
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
	if err = addListHandshakesForAccountResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHandshakesForAccount(options.Region), middleware.Before); err != nil {
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

// ListHandshakesForAccountAPIClient is a client that implements the
// ListHandshakesForAccount operation.
type ListHandshakesForAccountAPIClient interface {
	ListHandshakesForAccount(context.Context, *ListHandshakesForAccountInput, ...func(*Options)) (*ListHandshakesForAccountOutput, error)
}

var _ ListHandshakesForAccountAPIClient = (*Client)(nil)

// ListHandshakesForAccountPaginatorOptions is the paginator options for
// ListHandshakesForAccount
type ListHandshakesForAccountPaginatorOptions struct {
	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHandshakesForAccountPaginator is a paginator for ListHandshakesForAccount
type ListHandshakesForAccountPaginator struct {
	options   ListHandshakesForAccountPaginatorOptions
	client    ListHandshakesForAccountAPIClient
	params    *ListHandshakesForAccountInput
	nextToken *string
	firstPage bool
}

// NewListHandshakesForAccountPaginator returns a new
// ListHandshakesForAccountPaginator
func NewListHandshakesForAccountPaginator(client ListHandshakesForAccountAPIClient, params *ListHandshakesForAccountInput, optFns ...func(*ListHandshakesForAccountPaginatorOptions)) *ListHandshakesForAccountPaginator {
	if params == nil {
		params = &ListHandshakesForAccountInput{}
	}

	options := ListHandshakesForAccountPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHandshakesForAccountPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHandshakesForAccountPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHandshakesForAccount page.
func (p *ListHandshakesForAccountPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHandshakesForAccountOutput, error) {
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

	result, err := p.client.ListHandshakesForAccount(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListHandshakesForAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListHandshakesForAccount",
	}
}

type opListHandshakesForAccountResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListHandshakesForAccountResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListHandshakesForAccountResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "organizations"
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
				signingName = "organizations"
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
				v4aScheme.SigningName = aws.String("organizations")
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

func addListHandshakesForAccountResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListHandshakesForAccountResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}