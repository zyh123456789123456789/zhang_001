// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List lens review improvements.
func (c *Client) ListLensReviewImprovements(ctx context.Context, params *ListLensReviewImprovementsInput, optFns ...func(*Options)) (*ListLensReviewImprovementsOutput, error) {
	if params == nil {
		params = &ListLensReviewImprovementsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLensReviewImprovements", params, optFns, c.addOperationListLensReviewImprovementsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLensReviewImprovementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to list lens review improvements.
type ListLensReviewImprovementsInput struct {

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// . Each lens is identified by its LensSummary$LensAlias .
	//
	// This member is required.
	LensAlias *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	//
	// This member is required.
	WorkloadId *string

	// The maximum number of results to return for this request.
	MaxResults int32

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ID used to identify a pillar, for example, security . A pillar is identified
	// by its PillarReviewSummary$PillarId .
	PillarId *string

	// The priority of the question.
	QuestionPriority types.QuestionPriority

	noSmithyDocumentSerde
}

// Output of a list lens review improvements call.
type ListLensReviewImprovementsOutput struct {

	// List of improvement summaries of lens review in a workload.
	ImprovementSummaries []types.ImprovementSummary

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// . Each lens is identified by its LensSummary$LensAlias .
	LensAlias *string

	// The ARN for the lens.
	LensArn *string

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLensReviewImprovementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLensReviewImprovements{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLensReviewImprovements{}, middleware.After)
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
	if err = addListLensReviewImprovementsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListLensReviewImprovementsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLensReviewImprovements(options.Region), middleware.Before); err != nil {
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

// ListLensReviewImprovementsAPIClient is a client that implements the
// ListLensReviewImprovements operation.
type ListLensReviewImprovementsAPIClient interface {
	ListLensReviewImprovements(context.Context, *ListLensReviewImprovementsInput, ...func(*Options)) (*ListLensReviewImprovementsOutput, error)
}

var _ ListLensReviewImprovementsAPIClient = (*Client)(nil)

// ListLensReviewImprovementsPaginatorOptions is the paginator options for
// ListLensReviewImprovements
type ListLensReviewImprovementsPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLensReviewImprovementsPaginator is a paginator for
// ListLensReviewImprovements
type ListLensReviewImprovementsPaginator struct {
	options   ListLensReviewImprovementsPaginatorOptions
	client    ListLensReviewImprovementsAPIClient
	params    *ListLensReviewImprovementsInput
	nextToken *string
	firstPage bool
}

// NewListLensReviewImprovementsPaginator returns a new
// ListLensReviewImprovementsPaginator
func NewListLensReviewImprovementsPaginator(client ListLensReviewImprovementsAPIClient, params *ListLensReviewImprovementsInput, optFns ...func(*ListLensReviewImprovementsPaginatorOptions)) *ListLensReviewImprovementsPaginator {
	if params == nil {
		params = &ListLensReviewImprovementsInput{}
	}

	options := ListLensReviewImprovementsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLensReviewImprovementsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLensReviewImprovementsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLensReviewImprovements page.
func (p *ListLensReviewImprovementsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLensReviewImprovementsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListLensReviewImprovements(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLensReviewImprovements(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "ListLensReviewImprovements",
	}
}

type opListLensReviewImprovementsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListLensReviewImprovementsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListLensReviewImprovementsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "wellarchitected"
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
				signingName = "wellarchitected"
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
				v4aScheme.SigningName = aws.String("wellarchitected")
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

func addListLensReviewImprovementsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListLensReviewImprovementsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
