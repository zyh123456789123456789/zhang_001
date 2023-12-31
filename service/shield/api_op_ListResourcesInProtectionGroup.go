// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the resources that are included in the protection group.
func (c *Client) ListResourcesInProtectionGroup(ctx context.Context, params *ListResourcesInProtectionGroupInput, optFns ...func(*Options)) (*ListResourcesInProtectionGroupOutput, error) {
	if params == nil {
		params = &ListResourcesInProtectionGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourcesInProtectionGroup", params, optFns, c.addOperationListResourcesInProtectionGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourcesInProtectionGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcesInProtectionGroupInput struct {

	// The name of the protection group. You use this to identify the protection group
	// in lists and to manage the protection group, for example to update, delete, or
	// describe it.
	//
	// This member is required.
	ProtectionGroupId *string

	// The greatest number of objects that you want Shield Advanced to return to the
	// list request. Shield Advanced might return fewer objects than you indicate in
	// this setting, even if more objects are available. If there are more objects
	// remaining, Shield Advanced will always also return a NextToken value in the
	// response. The default setting is 20.
	MaxResults *int32

	// When you request a list of objects from Shield Advanced, if the response does
	// not include all of the remaining available objects, Shield Advanced includes a
	// NextToken value in the response. You can retrieve the next batch of objects by
	// requesting the list again and providing the token that was returned by the prior
	// call in your request. You can indicate the maximum number of objects that you
	// want Shield Advanced to return for a single call with the MaxResults setting.
	// Shield Advanced will not return more than MaxResults objects, but may return
	// fewer, even if more objects are still available. Whenever more objects remain
	// that Shield Advanced has not yet returned to you, the response will include a
	// NextToken value. On your first call to a list operation, leave this setting
	// empty.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourcesInProtectionGroupOutput struct {

	// The Amazon Resource Names (ARNs) of the resources that are included in the
	// protection group.
	//
	// This member is required.
	ResourceArns []string

	// When you request a list of objects from Shield Advanced, if the response does
	// not include all of the remaining available objects, Shield Advanced includes a
	// NextToken value in the response. You can retrieve the next batch of objects by
	// requesting the list again and providing the token that was returned by the prior
	// call in your request. You can indicate the maximum number of objects that you
	// want Shield Advanced to return for a single call with the MaxResults setting.
	// Shield Advanced will not return more than MaxResults objects, but may return
	// fewer, even if more objects are still available. Whenever more objects remain
	// that Shield Advanced has not yet returned to you, the response will include a
	// NextToken value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourcesInProtectionGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourcesInProtectionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourcesInProtectionGroup{}, middleware.After)
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
	if err = addListResourcesInProtectionGroupResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListResourcesInProtectionGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourcesInProtectionGroup(options.Region), middleware.Before); err != nil {
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

// ListResourcesInProtectionGroupAPIClient is a client that implements the
// ListResourcesInProtectionGroup operation.
type ListResourcesInProtectionGroupAPIClient interface {
	ListResourcesInProtectionGroup(context.Context, *ListResourcesInProtectionGroupInput, ...func(*Options)) (*ListResourcesInProtectionGroupOutput, error)
}

var _ ListResourcesInProtectionGroupAPIClient = (*Client)(nil)

// ListResourcesInProtectionGroupPaginatorOptions is the paginator options for
// ListResourcesInProtectionGroup
type ListResourcesInProtectionGroupPaginatorOptions struct {
	// The greatest number of objects that you want Shield Advanced to return to the
	// list request. Shield Advanced might return fewer objects than you indicate in
	// this setting, even if more objects are available. If there are more objects
	// remaining, Shield Advanced will always also return a NextToken value in the
	// response. The default setting is 20.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourcesInProtectionGroupPaginator is a paginator for
// ListResourcesInProtectionGroup
type ListResourcesInProtectionGroupPaginator struct {
	options   ListResourcesInProtectionGroupPaginatorOptions
	client    ListResourcesInProtectionGroupAPIClient
	params    *ListResourcesInProtectionGroupInput
	nextToken *string
	firstPage bool
}

// NewListResourcesInProtectionGroupPaginator returns a new
// ListResourcesInProtectionGroupPaginator
func NewListResourcesInProtectionGroupPaginator(client ListResourcesInProtectionGroupAPIClient, params *ListResourcesInProtectionGroupInput, optFns ...func(*ListResourcesInProtectionGroupPaginatorOptions)) *ListResourcesInProtectionGroupPaginator {
	if params == nil {
		params = &ListResourcesInProtectionGroupInput{}
	}

	options := ListResourcesInProtectionGroupPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourcesInProtectionGroupPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourcesInProtectionGroupPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourcesInProtectionGroup page.
func (p *ListResourcesInProtectionGroupPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourcesInProtectionGroupOutput, error) {
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

	result, err := p.client.ListResourcesInProtectionGroup(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourcesInProtectionGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "ListResourcesInProtectionGroup",
	}
}

type opListResourcesInProtectionGroupResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListResourcesInProtectionGroupResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListResourcesInProtectionGroupResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "shield"
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
				signingName = "shield"
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
				v4aScheme.SigningName = aws.String("shield")
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

func addListResourcesInProtectionGroupResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListResourcesInProtectionGroupResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
