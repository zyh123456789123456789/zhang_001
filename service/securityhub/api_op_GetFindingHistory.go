// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns history for a Security Hub finding in the last 90 days. The history
// includes changes made to any fields in the Amazon Web Services Security Finding
// Format (ASFF).
func (c *Client) GetFindingHistory(ctx context.Context, params *GetFindingHistoryInput, optFns ...func(*Options)) (*GetFindingHistoryOutput, error) {
	if params == nil {
		params = &GetFindingHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFindingHistory", params, optFns, c.addOperationGetFindingHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFindingHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFindingHistoryInput struct {

	// Identifies which finding to get the finding history for.
	//
	// This member is required.
	FindingIdentifier *types.AwsSecurityFindingIdentifier

	// An ISO 8601-formatted timestamp that indicates the end time of the requested
	// finding history. A correctly formatted example is 2020-05-21T20:16:34.724Z . The
	// value cannot contain spaces, and date and time should be separated by T . For
	// more information, see RFC 3339 section 5.6, Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
	// . If you provide values for both StartTime and EndTime , Security Hub returns
	// finding history for the specified time period. If you provide a value for
	// StartTime but not for EndTime , Security Hub returns finding history from the
	// StartTime to the time at which the API is called. If you provide a value for
	// EndTime but not for StartTime , Security Hub returns finding history from the
	// CreatedAt (https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_AwsSecurityFindingFilters.html#securityhub-Type-AwsSecurityFindingFilters-CreatedAt)
	// timestamp of the finding to the EndTime . If you provide neither StartTime nor
	// EndTime , Security Hub returns finding history from the CreatedAt timestamp of
	// the finding to the time at which the API is called. In all of these scenarios,
	// the response is limited to 100 results, and the maximum time period is limited
	// to 90 days.
	EndTime *time.Time

	// The maximum number of results to be returned. If you don’t provide it, Security
	// Hub returns up to 100 results of finding history.
	MaxResults int32

	// A token for pagination purposes. Provide NULL as the initial value. In
	// subsequent requests, provide the token included in the response to get up to an
	// additional 100 results of finding history. If you don’t provide NextToken ,
	// Security Hub returns up to 100 results of finding history for each request.
	NextToken *string

	// An ISO 8601-formatted timestamp that indicates the start time of the requested
	// finding history. A correctly formatted example is 2020-05-21T20:16:34.724Z . The
	// value cannot contain spaces, and date and time should be separated by T . For
	// more information, see RFC 3339 section 5.6, Internet Date/Time Format (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
	// . If you provide values for both StartTime and EndTime , Security Hub returns
	// finding history for the specified time period. If you provide a value for
	// StartTime but not for EndTime , Security Hub returns finding history from the
	// StartTime to the time at which the API is called. If you provide a value for
	// EndTime but not for StartTime , Security Hub returns finding history from the
	// CreatedAt (https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_AwsSecurityFindingFilters.html#securityhub-Type-AwsSecurityFindingFilters-CreatedAt)
	// timestamp of the finding to the EndTime . If you provide neither StartTime nor
	// EndTime , Security Hub returns finding history from the CreatedAt timestamp of
	// the finding to the time at which the API is called. In all of these scenarios,
	// the response is limited to 100 results, and the maximum time period is limited
	// to 90 days.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type GetFindingHistoryOutput struct {

	// A token for pagination purposes. Provide this token in the subsequent request
	// to GetFindingsHistory to get up to an additional 100 results of history for the
	// same finding that you specified in your initial request.
	NextToken *string

	// A list of events that altered the specified finding during the specified time
	// period.
	Records []types.FindingHistoryRecord

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFindingHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFindingHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindingHistory{}, middleware.After)
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
	if err = addGetFindingHistoryResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetFindingHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindingHistory(options.Region), middleware.Before); err != nil {
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

// GetFindingHistoryAPIClient is a client that implements the GetFindingHistory
// operation.
type GetFindingHistoryAPIClient interface {
	GetFindingHistory(context.Context, *GetFindingHistoryInput, ...func(*Options)) (*GetFindingHistoryOutput, error)
}

var _ GetFindingHistoryAPIClient = (*Client)(nil)

// GetFindingHistoryPaginatorOptions is the paginator options for GetFindingHistory
type GetFindingHistoryPaginatorOptions struct {
	// The maximum number of results to be returned. If you don’t provide it, Security
	// Hub returns up to 100 results of finding history.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetFindingHistoryPaginator is a paginator for GetFindingHistory
type GetFindingHistoryPaginator struct {
	options   GetFindingHistoryPaginatorOptions
	client    GetFindingHistoryAPIClient
	params    *GetFindingHistoryInput
	nextToken *string
	firstPage bool
}

// NewGetFindingHistoryPaginator returns a new GetFindingHistoryPaginator
func NewGetFindingHistoryPaginator(client GetFindingHistoryAPIClient, params *GetFindingHistoryInput, optFns ...func(*GetFindingHistoryPaginatorOptions)) *GetFindingHistoryPaginator {
	if params == nil {
		params = &GetFindingHistoryInput{}
	}

	options := GetFindingHistoryPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetFindingHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetFindingHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetFindingHistory page.
func (p *GetFindingHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetFindingHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetFindingHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetFindingHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "GetFindingHistory",
	}
}

type opGetFindingHistoryResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetFindingHistoryResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetFindingHistoryResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "securityhub"
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
				signingName = "securityhub"
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
				v4aScheme.SigningName = aws.String("securityhub")
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

func addGetFindingHistoryResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetFindingHistoryResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
