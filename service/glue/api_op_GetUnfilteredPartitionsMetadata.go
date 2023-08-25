// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves partition metadata from the Data Catalog that contains unfiltered
// metadata. For IAM authorization, the public IAM action associated with this API
// is glue:GetPartitions .
func (c *Client) GetUnfilteredPartitionsMetadata(ctx context.Context, params *GetUnfilteredPartitionsMetadataInput, optFns ...func(*Options)) (*GetUnfilteredPartitionsMetadataOutput, error) {
	if params == nil {
		params = &GetUnfilteredPartitionsMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUnfilteredPartitionsMetadata", params, optFns, c.addOperationGetUnfilteredPartitionsMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUnfilteredPartitionsMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUnfilteredPartitionsMetadataInput struct {

	// The ID of the Data Catalog where the partitions in question reside. If none is
	// provided, the AWS account ID is used by default.
	//
	// This member is required.
	CatalogId *string

	// The name of the catalog database where the partitions reside.
	//
	// This member is required.
	DatabaseName *string

	// A list of supported permission types.
	//
	// This member is required.
	SupportedPermissionTypes []types.PermissionType

	// The name of the table that contains the partition.
	//
	// This member is required.
	TableName *string

	// A structure containing Lake Formation audit context information.
	AuditContext *types.AuditContext

	// An expression that filters the partitions to be returned. The expression uses
	// SQL syntax similar to the SQL WHERE filter clause. The SQL statement parser
	// JSQLParser (http://jsqlparser.sourceforge.net/home.php) parses the expression.
	// Operators: The following are the operators that you can use in the Expression
	// API call: = Checks whether the values of the two operands are equal; if yes,
	// then the condition becomes true. Example: Assume 'variable a' holds 10 and
	// 'variable b' holds 20. (a = b) is not true. < > Checks whether the values of two
	// operands are equal; if the values are not equal, then the condition becomes
	// true. Example: (a < > b) is true. > Checks whether the value of the left operand
	// is greater than the value of the right operand; if yes, then the condition
	// becomes true. Example: (a > b) is not true. < Checks whether the value of the
	// left operand is less than the value of the right operand; if yes, then the
	// condition becomes true. Example: (a < b) is true. >= Checks whether the value of
	// the left operand is greater than or equal to the value of the right operand; if
	// yes, then the condition becomes true. Example: (a >= b) is not true. <= Checks
	// whether the value of the left operand is less than or equal to the value of the
	// right operand; if yes, then the condition becomes true. Example: (a <= b) is
	// true. AND, OR, IN, BETWEEN, LIKE, NOT, IS NULL Logical operators. Supported
	// Partition Key Types: The following are the supported partition keys.
	//   - string
	//   - date
	//   - timestamp
	//   - int
	//   - bigint
	//   - long
	//   - tinyint
	//   - smallint
	//   - decimal
	// If an type is encountered that is not valid, an exception is thrown.
	Expression *string

	// The maximum number of partitions to return in a single response.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve these
	// partitions.
	NextToken *string

	// The segment of the table's partitions to scan in this request.
	Segment *types.Segment

	noSmithyDocumentSerde
}

type GetUnfilteredPartitionsMetadataOutput struct {

	// A continuation token, if the returned list of partitions does not include the
	// last one.
	NextToken *string

	// A list of requested partitions.
	UnfilteredPartitions []types.UnfilteredPartition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUnfilteredPartitionsMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUnfilteredPartitionsMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUnfilteredPartitionsMetadata{}, middleware.After)
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
	if err = addGetUnfilteredPartitionsMetadataResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetUnfilteredPartitionsMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUnfilteredPartitionsMetadata(options.Region), middleware.Before); err != nil {
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

// GetUnfilteredPartitionsMetadataAPIClient is a client that implements the
// GetUnfilteredPartitionsMetadata operation.
type GetUnfilteredPartitionsMetadataAPIClient interface {
	GetUnfilteredPartitionsMetadata(context.Context, *GetUnfilteredPartitionsMetadataInput, ...func(*Options)) (*GetUnfilteredPartitionsMetadataOutput, error)
}

var _ GetUnfilteredPartitionsMetadataAPIClient = (*Client)(nil)

// GetUnfilteredPartitionsMetadataPaginatorOptions is the paginator options for
// GetUnfilteredPartitionsMetadata
type GetUnfilteredPartitionsMetadataPaginatorOptions struct {
	// The maximum number of partitions to return in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetUnfilteredPartitionsMetadataPaginator is a paginator for
// GetUnfilteredPartitionsMetadata
type GetUnfilteredPartitionsMetadataPaginator struct {
	options   GetUnfilteredPartitionsMetadataPaginatorOptions
	client    GetUnfilteredPartitionsMetadataAPIClient
	params    *GetUnfilteredPartitionsMetadataInput
	nextToken *string
	firstPage bool
}

// NewGetUnfilteredPartitionsMetadataPaginator returns a new
// GetUnfilteredPartitionsMetadataPaginator
func NewGetUnfilteredPartitionsMetadataPaginator(client GetUnfilteredPartitionsMetadataAPIClient, params *GetUnfilteredPartitionsMetadataInput, optFns ...func(*GetUnfilteredPartitionsMetadataPaginatorOptions)) *GetUnfilteredPartitionsMetadataPaginator {
	if params == nil {
		params = &GetUnfilteredPartitionsMetadataInput{}
	}

	options := GetUnfilteredPartitionsMetadataPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetUnfilteredPartitionsMetadataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetUnfilteredPartitionsMetadataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetUnfilteredPartitionsMetadata page.
func (p *GetUnfilteredPartitionsMetadataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetUnfilteredPartitionsMetadataOutput, error) {
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

	result, err := p.client.GetUnfilteredPartitionsMetadata(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetUnfilteredPartitionsMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetUnfilteredPartitionsMetadata",
	}
}

type opGetUnfilteredPartitionsMetadataResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetUnfilteredPartitionsMetadataResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetUnfilteredPartitionsMetadataResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "glue"
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
				signingName = "glue"
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
				v4aScheme.SigningName = aws.String("glue")
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

func addGetUnfilteredPartitionsMetadataResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetUnfilteredPartitionsMetadataResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}