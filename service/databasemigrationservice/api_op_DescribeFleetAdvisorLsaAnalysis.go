// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides descriptions of large-scale assessment (LSA) analyses produced by your
// Fleet Advisor collectors.
func (c *Client) DescribeFleetAdvisorLsaAnalysis(ctx context.Context, params *DescribeFleetAdvisorLsaAnalysisInput, optFns ...func(*Options)) (*DescribeFleetAdvisorLsaAnalysisOutput, error) {
	if params == nil {
		params = &DescribeFleetAdvisorLsaAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetAdvisorLsaAnalysis", params, optFns, c.addOperationDescribeFleetAdvisorLsaAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetAdvisorLsaAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetAdvisorLsaAnalysisInput struct {

	// Sets the maximum number of records returned in the response.
	MaxRecords *int32

	// If NextToken is returned by a previous response, there are more results
	// available. The value of NextToken is a unique pagination token for each page.
	// Make the call again using the returned token to retrieve the next page. Keep all
	// other arguments unchanged.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFleetAdvisorLsaAnalysisOutput struct {

	// A list of FleetAdvisorLsaAnalysisResponse objects.
	Analysis []types.FleetAdvisorLsaAnalysisResponse

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetAdvisorLsaAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetAdvisorLsaAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetAdvisorLsaAnalysis{}, middleware.After)
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
	if err = addDescribeFleetAdvisorLsaAnalysisResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetAdvisorLsaAnalysis(options.Region), middleware.Before); err != nil {
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

// DescribeFleetAdvisorLsaAnalysisAPIClient is a client that implements the
// DescribeFleetAdvisorLsaAnalysis operation.
type DescribeFleetAdvisorLsaAnalysisAPIClient interface {
	DescribeFleetAdvisorLsaAnalysis(context.Context, *DescribeFleetAdvisorLsaAnalysisInput, ...func(*Options)) (*DescribeFleetAdvisorLsaAnalysisOutput, error)
}

var _ DescribeFleetAdvisorLsaAnalysisAPIClient = (*Client)(nil)

// DescribeFleetAdvisorLsaAnalysisPaginatorOptions is the paginator options for
// DescribeFleetAdvisorLsaAnalysis
type DescribeFleetAdvisorLsaAnalysisPaginatorOptions struct {
	// Sets the maximum number of records returned in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetAdvisorLsaAnalysisPaginator is a paginator for
// DescribeFleetAdvisorLsaAnalysis
type DescribeFleetAdvisorLsaAnalysisPaginator struct {
	options   DescribeFleetAdvisorLsaAnalysisPaginatorOptions
	client    DescribeFleetAdvisorLsaAnalysisAPIClient
	params    *DescribeFleetAdvisorLsaAnalysisInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetAdvisorLsaAnalysisPaginator returns a new
// DescribeFleetAdvisorLsaAnalysisPaginator
func NewDescribeFleetAdvisorLsaAnalysisPaginator(client DescribeFleetAdvisorLsaAnalysisAPIClient, params *DescribeFleetAdvisorLsaAnalysisInput, optFns ...func(*DescribeFleetAdvisorLsaAnalysisPaginatorOptions)) *DescribeFleetAdvisorLsaAnalysisPaginator {
	if params == nil {
		params = &DescribeFleetAdvisorLsaAnalysisInput{}
	}

	options := DescribeFleetAdvisorLsaAnalysisPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFleetAdvisorLsaAnalysisPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetAdvisorLsaAnalysisPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFleetAdvisorLsaAnalysis page.
func (p *DescribeFleetAdvisorLsaAnalysisPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetAdvisorLsaAnalysisOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeFleetAdvisorLsaAnalysis(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFleetAdvisorLsaAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeFleetAdvisorLsaAnalysis",
	}
}

type opDescribeFleetAdvisorLsaAnalysisResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeFleetAdvisorLsaAnalysisResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeFleetAdvisorLsaAnalysisResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "dms"
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
				signingName = "dms"
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
				v4aScheme.SigningName = aws.String("dms")
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

func addDescribeFleetAdvisorLsaAnalysisResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeFleetAdvisorLsaAnalysisResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}