// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the details of FlexMatch matchmaking configurations. This operation
// offers the following options: (1) retrieve all matchmaking configurations, (2)
// retrieve configurations for a specified list, or (3) retrieve all configurations
// that use a specified rule set name. When requesting multiple items, use the
// pagination parameters to retrieve results as a set of sequential pages. If
// successful, a configuration is returned for each requested name. When specifying
// a list of names, only configurations that currently exist are returned. Learn
// more Setting up FlexMatch matchmakers (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/matchmaker-build.html)
func (c *Client) DescribeMatchmakingConfigurations(ctx context.Context, params *DescribeMatchmakingConfigurationsInput, optFns ...func(*Options)) (*DescribeMatchmakingConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeMatchmakingConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMatchmakingConfigurations", params, optFns, c.addOperationDescribeMatchmakingConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMatchmakingConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMatchmakingConfigurationsInput struct {

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is limited to 10.
	Limit *int32

	// A unique identifier for the matchmaking configuration(s) to retrieve. You can
	// use either the configuration name or ARN value. To request all existing
	// configurations, leave this parameter empty.
	Names []string

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	// A unique identifier for the matchmaking rule set. You can use either the rule
	// set name or ARN value. Use this parameter to retrieve all matchmaking
	// configurations that use this rule set.
	RuleSetName *string

	noSmithyDocumentSerde
}

type DescribeMatchmakingConfigurationsOutput struct {

	// A collection of requested matchmaking configurations.
	Configurations []types.MatchmakingConfiguration

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMatchmakingConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMatchmakingConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMatchmakingConfigurations{}, middleware.After)
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
	if err = addDescribeMatchmakingConfigurationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMatchmakingConfigurations(options.Region), middleware.Before); err != nil {
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

// DescribeMatchmakingConfigurationsAPIClient is a client that implements the
// DescribeMatchmakingConfigurations operation.
type DescribeMatchmakingConfigurationsAPIClient interface {
	DescribeMatchmakingConfigurations(context.Context, *DescribeMatchmakingConfigurationsInput, ...func(*Options)) (*DescribeMatchmakingConfigurationsOutput, error)
}

var _ DescribeMatchmakingConfigurationsAPIClient = (*Client)(nil)

// DescribeMatchmakingConfigurationsPaginatorOptions is the paginator options for
// DescribeMatchmakingConfigurations
type DescribeMatchmakingConfigurationsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is limited to 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMatchmakingConfigurationsPaginator is a paginator for
// DescribeMatchmakingConfigurations
type DescribeMatchmakingConfigurationsPaginator struct {
	options   DescribeMatchmakingConfigurationsPaginatorOptions
	client    DescribeMatchmakingConfigurationsAPIClient
	params    *DescribeMatchmakingConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMatchmakingConfigurationsPaginator returns a new
// DescribeMatchmakingConfigurationsPaginator
func NewDescribeMatchmakingConfigurationsPaginator(client DescribeMatchmakingConfigurationsAPIClient, params *DescribeMatchmakingConfigurationsInput, optFns ...func(*DescribeMatchmakingConfigurationsPaginatorOptions)) *DescribeMatchmakingConfigurationsPaginator {
	if params == nil {
		params = &DescribeMatchmakingConfigurationsInput{}
	}

	options := DescribeMatchmakingConfigurationsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMatchmakingConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMatchmakingConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMatchmakingConfigurations page.
func (p *DescribeMatchmakingConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMatchmakingConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeMatchmakingConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMatchmakingConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeMatchmakingConfigurations",
	}
}

type opDescribeMatchmakingConfigurationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeMatchmakingConfigurationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeMatchmakingConfigurationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "gamelift"
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
				signingName = "gamelift"
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
				v4aScheme.SigningName = aws.String("gamelift")
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

func addDescribeMatchmakingConfigurationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeMatchmakingConfigurationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
