// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all journal export jobs for a specified ledger. This action returns a
// maximum of MaxResults items, and is paginated so that you can retrieve all the
// items by calling ListJournalS3ExportsForLedger multiple times. This action does
// not return any expired export jobs. For more information, see Export job
// expiration (https://docs.aws.amazon.com/qldb/latest/developerguide/export-journal.request.html#export-journal.request.expiration)
// in the Amazon QLDB Developer Guide.
func (c *Client) ListJournalS3ExportsForLedger(ctx context.Context, params *ListJournalS3ExportsForLedgerInput, optFns ...func(*Options)) (*ListJournalS3ExportsForLedgerOutput, error) {
	if params == nil {
		params = &ListJournalS3ExportsForLedgerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJournalS3ExportsForLedger", params, optFns, c.addOperationListJournalS3ExportsForLedgerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJournalS3ExportsForLedgerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJournalS3ExportsForLedgerInput struct {

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The maximum number of results to return in a single
	// ListJournalS3ExportsForLedger request. (The actual number of results returned
	// might be fewer.)
	MaxResults *int32

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListJournalS3ExportsForLedger call, then you should use that value as input here.
	NextToken *string

	noSmithyDocumentSerde
}

type ListJournalS3ExportsForLedgerOutput struct {

	// The journal export jobs that are currently associated with the specified ledger.
	JournalS3Exports []types.JournalS3ExportDescription

	//   - If NextToken is empty, then the last page of results has been processed and
	//   there are no more results to be retrieved.
	//   - If NextToken is not empty, then there are more results available. To
	//   retrieve the next page of results, use the value of NextToken in a subsequent
	//   ListJournalS3ExportsForLedger call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListJournalS3ExportsForLedgerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJournalS3ExportsForLedger{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJournalS3ExportsForLedger{}, middleware.After)
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
	if err = addListJournalS3ExportsForLedgerResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListJournalS3ExportsForLedgerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJournalS3ExportsForLedger(options.Region), middleware.Before); err != nil {
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

// ListJournalS3ExportsForLedgerAPIClient is a client that implements the
// ListJournalS3ExportsForLedger operation.
type ListJournalS3ExportsForLedgerAPIClient interface {
	ListJournalS3ExportsForLedger(context.Context, *ListJournalS3ExportsForLedgerInput, ...func(*Options)) (*ListJournalS3ExportsForLedgerOutput, error)
}

var _ ListJournalS3ExportsForLedgerAPIClient = (*Client)(nil)

// ListJournalS3ExportsForLedgerPaginatorOptions is the paginator options for
// ListJournalS3ExportsForLedger
type ListJournalS3ExportsForLedgerPaginatorOptions struct {
	// The maximum number of results to return in a single
	// ListJournalS3ExportsForLedger request. (The actual number of results returned
	// might be fewer.)
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJournalS3ExportsForLedgerPaginator is a paginator for
// ListJournalS3ExportsForLedger
type ListJournalS3ExportsForLedgerPaginator struct {
	options   ListJournalS3ExportsForLedgerPaginatorOptions
	client    ListJournalS3ExportsForLedgerAPIClient
	params    *ListJournalS3ExportsForLedgerInput
	nextToken *string
	firstPage bool
}

// NewListJournalS3ExportsForLedgerPaginator returns a new
// ListJournalS3ExportsForLedgerPaginator
func NewListJournalS3ExportsForLedgerPaginator(client ListJournalS3ExportsForLedgerAPIClient, params *ListJournalS3ExportsForLedgerInput, optFns ...func(*ListJournalS3ExportsForLedgerPaginatorOptions)) *ListJournalS3ExportsForLedgerPaginator {
	if params == nil {
		params = &ListJournalS3ExportsForLedgerInput{}
	}

	options := ListJournalS3ExportsForLedgerPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJournalS3ExportsForLedgerPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJournalS3ExportsForLedgerPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListJournalS3ExportsForLedger page.
func (p *ListJournalS3ExportsForLedgerPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJournalS3ExportsForLedgerOutput, error) {
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

	result, err := p.client.ListJournalS3ExportsForLedger(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJournalS3ExportsForLedger(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ListJournalS3ExportsForLedger",
	}
}

type opListJournalS3ExportsForLedgerResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListJournalS3ExportsForLedgerResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListJournalS3ExportsForLedgerResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "qldb"
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
				signingName = "qldb"
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
				v4aScheme.SigningName = aws.String("qldb")
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

func addListJournalS3ExportsForLedgerResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListJournalS3ExportsForLedgerResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
