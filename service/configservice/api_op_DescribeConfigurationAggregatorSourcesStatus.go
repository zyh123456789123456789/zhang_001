// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns status information for sources within an aggregator. The status
// includes information about the last time Config verified authorization between
// the source account and an aggregator account. In case of a failure, the status
// contains the related error code or message.
func (c *Client) DescribeConfigurationAggregatorSourcesStatus(ctx context.Context, params *DescribeConfigurationAggregatorSourcesStatusInput, optFns ...func(*Options)) (*DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	if params == nil {
		params = &DescribeConfigurationAggregatorSourcesStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationAggregatorSourcesStatus", params, optFns, c.addOperationDescribeConfigurationAggregatorSourcesStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationAggregatorSourcesStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationAggregatorSourcesStatusInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// The maximum number of AggregatorSourceStatus returned on each page. The default
	// is maximum. If you specify 0, Config uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Filters the status type.
	//   - Valid value FAILED indicates errors while moving data.
	//   - Valid value SUCCEEDED indicates the data was successfully moved.
	//   - Valid value OUTDATED indicates the data is not the most recent.
	UpdateStatus []types.AggregatedSourceStatusType

	noSmithyDocumentSerde
}

type DescribeConfigurationAggregatorSourcesStatusOutput struct {

	// Returns an AggregatedSourceStatus object.
	AggregatedSourceStatusList []types.AggregatedSourceStatus

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigurationAggregatorSourcesStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigurationAggregatorSourcesStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigurationAggregatorSourcesStatus{}, middleware.After)
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
	if err = addDescribeConfigurationAggregatorSourcesStatusResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeConfigurationAggregatorSourcesStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationAggregatorSourcesStatus(options.Region), middleware.Before); err != nil {
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

// DescribeConfigurationAggregatorSourcesStatusAPIClient is a client that
// implements the DescribeConfigurationAggregatorSourcesStatus operation.
type DescribeConfigurationAggregatorSourcesStatusAPIClient interface {
	DescribeConfigurationAggregatorSourcesStatus(context.Context, *DescribeConfigurationAggregatorSourcesStatusInput, ...func(*Options)) (*DescribeConfigurationAggregatorSourcesStatusOutput, error)
}

var _ DescribeConfigurationAggregatorSourcesStatusAPIClient = (*Client)(nil)

// DescribeConfigurationAggregatorSourcesStatusPaginatorOptions is the paginator
// options for DescribeConfigurationAggregatorSourcesStatus
type DescribeConfigurationAggregatorSourcesStatusPaginatorOptions struct {
	// The maximum number of AggregatorSourceStatus returned on each page. The default
	// is maximum. If you specify 0, Config uses the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeConfigurationAggregatorSourcesStatusPaginator is a paginator for
// DescribeConfigurationAggregatorSourcesStatus
type DescribeConfigurationAggregatorSourcesStatusPaginator struct {
	options   DescribeConfigurationAggregatorSourcesStatusPaginatorOptions
	client    DescribeConfigurationAggregatorSourcesStatusAPIClient
	params    *DescribeConfigurationAggregatorSourcesStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeConfigurationAggregatorSourcesStatusPaginator returns a new
// DescribeConfigurationAggregatorSourcesStatusPaginator
func NewDescribeConfigurationAggregatorSourcesStatusPaginator(client DescribeConfigurationAggregatorSourcesStatusAPIClient, params *DescribeConfigurationAggregatorSourcesStatusInput, optFns ...func(*DescribeConfigurationAggregatorSourcesStatusPaginatorOptions)) *DescribeConfigurationAggregatorSourcesStatusPaginator {
	if params == nil {
		params = &DescribeConfigurationAggregatorSourcesStatusInput{}
	}

	options := DescribeConfigurationAggregatorSourcesStatusPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeConfigurationAggregatorSourcesStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeConfigurationAggregatorSourcesStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeConfigurationAggregatorSourcesStatus page.
func (p *DescribeConfigurationAggregatorSourcesStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeConfigurationAggregatorSourcesStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeConfigurationAggregatorSourcesStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeConfigurationAggregatorSourcesStatus",
	}
}

type opDescribeConfigurationAggregatorSourcesStatusResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeConfigurationAggregatorSourcesStatusResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeConfigurationAggregatorSourcesStatusResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "config"
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
				signingName = "config"
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
				v4aScheme.SigningName = aws.String("config")
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

func addDescribeConfigurationAggregatorSourcesStatusResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeConfigurationAggregatorSourcesStatusResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}