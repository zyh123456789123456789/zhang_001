// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// To use this API operation, your IAM role must have permissions to perform the
// ListAggregatedUtterances (https://docs.aws.amazon.com/lexv2/latest/APIReference/API_ListAggregatedUtterances.html)
// operation, which provides access to utterance-related analytics. See Viewing
// utterance statistics (https://docs.aws.amazon.com/lexv2/latest/dg/monitoring-utterances.html)
// for the IAM policy to apply to the IAM role. Retrieves a list of metadata for
// individual user utterances to your bot. The following fields are required:
//   - startDateTime and endDateTime – Define a time range for which you want to
//     retrieve results.
//
// Of the optional fields, you can organize the results in the following ways:
//   - Use the filters field to filter the results and the sortBy field to specify
//     the values by which to sort the results.
//   - Use the maxResults field to limit the number of results to return in a
//     single response and the nextToken field to return the next batch of results if
//     the response does not return the full set of results.
func (c *Client) ListUtteranceAnalyticsData(ctx context.Context, params *ListUtteranceAnalyticsDataInput, optFns ...func(*Options)) (*ListUtteranceAnalyticsDataOutput, error) {
	if params == nil {
		params = &ListUtteranceAnalyticsDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUtteranceAnalyticsData", params, optFns, c.addOperationListUtteranceAnalyticsDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUtteranceAnalyticsDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUtteranceAnalyticsDataInput struct {

	// The identifier for the bot for which you want to retrieve utterance analytics.
	//
	// This member is required.
	BotId *string

	// The date and time that marks the end of the range of time for which you want to
	// see utterance analytics.
	//
	// This member is required.
	EndDateTime *time.Time

	// The date and time that marks the beginning of the range of time for which you
	// want to see utterance analytics.
	//
	// This member is required.
	StartDateTime *time.Time

	// A list of objects, each of which describes a condition by which you want to
	// filter the results.
	Filters []types.AnalyticsUtteranceFilter

	// The maximum number of results to return in each page of results. If there are
	// fewer results than the maximum page size, only the actual number of results are
	// returned.
	MaxResults *int32

	// If the response from the ListUtteranceAnalyticsData operation contains more
	// results than specified in the maxResults parameter, a token is returned in the
	// response. Use the returned token in the nextToken parameter of a
	// ListUtteranceAnalyticsData request to return the next page of results. For a
	// complete set of results, call the ListUtteranceAnalyticsData operation until the
	// nextToken returned in the response is null.
	NextToken *string

	// An object specifying the measure and method by which to sort the utterance
	// analytics data.
	SortBy *types.UtteranceDataSortBy

	noSmithyDocumentSerde
}

type ListUtteranceAnalyticsDataOutput struct {

	// The unique identifier of the bot that the utterances belong to.
	BotId *string

	// If the response from the ListUtteranceAnalyticsData operation contains more
	// results than specified in the maxResults parameter, a token is returned in the
	// response. Use the returned token in the nextToken parameter of a
	// ListUtteranceAnalyticsData request to return the next page of results. For a
	// complete set of results, call the ListUtteranceAnalyticsData operation until the
	// nextToken returned in the response is null.
	NextToken *string

	// A list of objects, each of which contains information about an utterance in a
	// user session with your bot.
	Utterances []types.UtteranceSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUtteranceAnalyticsDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUtteranceAnalyticsData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUtteranceAnalyticsData{}, middleware.After)
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
	if err = addListUtteranceAnalyticsDataResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListUtteranceAnalyticsDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUtteranceAnalyticsData(options.Region), middleware.Before); err != nil {
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

// ListUtteranceAnalyticsDataAPIClient is a client that implements the
// ListUtteranceAnalyticsData operation.
type ListUtteranceAnalyticsDataAPIClient interface {
	ListUtteranceAnalyticsData(context.Context, *ListUtteranceAnalyticsDataInput, ...func(*Options)) (*ListUtteranceAnalyticsDataOutput, error)
}

var _ ListUtteranceAnalyticsDataAPIClient = (*Client)(nil)

// ListUtteranceAnalyticsDataPaginatorOptions is the paginator options for
// ListUtteranceAnalyticsData
type ListUtteranceAnalyticsDataPaginatorOptions struct {
	// The maximum number of results to return in each page of results. If there are
	// fewer results than the maximum page size, only the actual number of results are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUtteranceAnalyticsDataPaginator is a paginator for
// ListUtteranceAnalyticsData
type ListUtteranceAnalyticsDataPaginator struct {
	options   ListUtteranceAnalyticsDataPaginatorOptions
	client    ListUtteranceAnalyticsDataAPIClient
	params    *ListUtteranceAnalyticsDataInput
	nextToken *string
	firstPage bool
}

// NewListUtteranceAnalyticsDataPaginator returns a new
// ListUtteranceAnalyticsDataPaginator
func NewListUtteranceAnalyticsDataPaginator(client ListUtteranceAnalyticsDataAPIClient, params *ListUtteranceAnalyticsDataInput, optFns ...func(*ListUtteranceAnalyticsDataPaginatorOptions)) *ListUtteranceAnalyticsDataPaginator {
	if params == nil {
		params = &ListUtteranceAnalyticsDataInput{}
	}

	options := ListUtteranceAnalyticsDataPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUtteranceAnalyticsDataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUtteranceAnalyticsDataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUtteranceAnalyticsData page.
func (p *ListUtteranceAnalyticsDataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUtteranceAnalyticsDataOutput, error) {
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

	result, err := p.client.ListUtteranceAnalyticsData(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUtteranceAnalyticsData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "ListUtteranceAnalyticsData",
	}
}

type opListUtteranceAnalyticsDataResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListUtteranceAnalyticsDataResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListUtteranceAnalyticsDataResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "lex"
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
				signingName = "lex"
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
				v4aScheme.SigningName = aws.String("lex")
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

func addListUtteranceAnalyticsDataResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListUtteranceAnalyticsDataResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}