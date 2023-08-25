// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves IDs and annotations for traces available for a specified time frame
// using an optional filter. To get the full traces, pass the trace IDs to
// BatchGetTraces . A filter expression can target traced requests that hit
// specific service nodes or edges, have errors, or come from a known user. For
// example, the following filter expression targets traces that pass through
// api.example.com : service("api.example.com") This filter expression finds
// traces that have an annotation named account with the value 12345 :
// annotation.account = "12345" For a full list of indexed fields and keywords that
// you can use in filter expressions, see Using Filter Expressions (https://docs.aws.amazon.com/xray/latest/devguide/xray-console-filters.html)
// in the Amazon Web Services X-Ray Developer Guide.
func (c *Client) GetTraceSummaries(ctx context.Context, params *GetTraceSummariesInput, optFns ...func(*Options)) (*GetTraceSummariesOutput, error) {
	if params == nil {
		params = &GetTraceSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTraceSummaries", params, optFns, c.addOperationGetTraceSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTraceSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTraceSummariesInput struct {

	// The end of the time frame for which to retrieve traces.
	//
	// This member is required.
	EndTime *time.Time

	// The start of the time frame for which to retrieve traces.
	//
	// This member is required.
	StartTime *time.Time

	// Specify a filter expression to retrieve trace summaries for services or
	// requests that meet certain requirements.
	FilterExpression *string

	// Specify the pagination token returned by a previous request to retrieve the
	// next page of results.
	NextToken *string

	// Set to true to get summaries for only a subset of available traces.
	Sampling *bool

	// A parameter to indicate whether to enable sampling on trace summaries. Input
	// parameters are Name and Value.
	SamplingStrategy *types.SamplingStrategy

	// A parameter to indicate whether to query trace summaries by TraceId or Event
	// time.
	TimeRangeType types.TimeRangeType

	noSmithyDocumentSerde
}

type GetTraceSummariesOutput struct {

	// The start time of this page of results.
	ApproximateTime *time.Time

	// If the requested time frame contained more than one page of results, you can
	// use this token to retrieve the next page. The first page contains the most
	// recent results, closest to the end of the time frame.
	NextToken *string

	// Trace IDs and annotations for traces that were found in the specified time
	// frame.
	TraceSummaries []types.TraceSummary

	// The total number of traces processed, including traces that did not match the
	// specified filter expression.
	TracesProcessedCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTraceSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTraceSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTraceSummaries{}, middleware.After)
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
	if err = addGetTraceSummariesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetTraceSummariesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTraceSummaries(options.Region), middleware.Before); err != nil {
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

// GetTraceSummariesAPIClient is a client that implements the GetTraceSummaries
// operation.
type GetTraceSummariesAPIClient interface {
	GetTraceSummaries(context.Context, *GetTraceSummariesInput, ...func(*Options)) (*GetTraceSummariesOutput, error)
}

var _ GetTraceSummariesAPIClient = (*Client)(nil)

// GetTraceSummariesPaginatorOptions is the paginator options for GetTraceSummaries
type GetTraceSummariesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTraceSummariesPaginator is a paginator for GetTraceSummaries
type GetTraceSummariesPaginator struct {
	options   GetTraceSummariesPaginatorOptions
	client    GetTraceSummariesAPIClient
	params    *GetTraceSummariesInput
	nextToken *string
	firstPage bool
}

// NewGetTraceSummariesPaginator returns a new GetTraceSummariesPaginator
func NewGetTraceSummariesPaginator(client GetTraceSummariesAPIClient, params *GetTraceSummariesInput, optFns ...func(*GetTraceSummariesPaginatorOptions)) *GetTraceSummariesPaginator {
	if params == nil {
		params = &GetTraceSummariesInput{}
	}

	options := GetTraceSummariesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTraceSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTraceSummariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTraceSummaries page.
func (p *GetTraceSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTraceSummariesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetTraceSummaries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTraceSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetTraceSummaries",
	}
}

type opGetTraceSummariesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetTraceSummariesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetTraceSummariesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "xray"
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
				signingName = "xray"
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
				v4aScheme.SigningName = aws.String("xray")
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

func addGetTraceSummariesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetTraceSummariesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}