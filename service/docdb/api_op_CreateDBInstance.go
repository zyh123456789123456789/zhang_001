// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new instance.
func (c *Client) CreateDBInstance(ctx context.Context, params *CreateDBInstanceInput, optFns ...func(*Options)) (*CreateDBInstanceOutput, error) {
	if params == nil {
		params = &CreateDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBInstance", params, optFns, c.addOperationCreateDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to CreateDBInstance .
type CreateDBInstanceInput struct {

	// The identifier of the cluster that the instance will belong to.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The compute and memory capacity of the instance; for example, db.r5.large .
	//
	// This member is required.
	DBInstanceClass *string

	// The instance identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//   - The first character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	// Example: mydbinstance
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The name of the database engine to be used for this instance. Valid value: docdb
	//
	// This member is required.
	Engine *string

	// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not
	// perform minor version upgrades regardless of the value set. Default: false
	AutoMinorVersionUpgrade *bool

	// The Amazon EC2 Availability Zone that the instance is created in. Default: A
	// random, system-chosen Availability Zone in the endpoint's Amazon Web Services
	// Region. Example: us-east-1d
	AvailabilityZone *string

	// A value that indicates whether to copy tags from the DB instance to snapshots
	// of the DB instance. By default, tags are not copied.
	CopyTagsToSnapshot *bool

	// A value that indicates whether to enable Performance Insights for the DB
	// Instance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html)
	// .
	EnablePerformanceInsights *bool

	// The KMS key identifier for encryption of Performance Insights data. The KMS key
	// identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If
	// you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon
	// DocumentDB uses your default KMS key. There is a default KMS key for your Amazon
	// Web Services account. Your Amazon Web Services account has a different default
	// KMS key for each Amazon Web Services region.
	PerformanceInsightsKMSKeyId *string

	// The time range each week during which system maintenance can occur, in
	// Universal Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default
	// is a 30-minute window selected at random from an 8-hour block of time for each
	// Amazon Web Services Region, occurring on a random day of the week. Valid days:
	// Mon, Tue, Wed, Thu, Fri, Sat, Sun Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// A value that specifies the order in which an Amazon DocumentDB replica is
	// promoted to the primary instance after a failure of the existing primary
	// instance. Default: 1 Valid values: 0-15
	PromotionTier *int32

	// The tags to be assigned to the instance. You can assign up to 10 tags to an
	// instance.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDBInstanceOutput struct {

	// Detailed information about an instance.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBInstance{}, middleware.After)
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
	if err = addCreateDBInstanceResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBInstance",
	}
}

type opCreateDBInstanceResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateDBInstanceResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateDBInstanceResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "rds"
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
				signingName = "rds"
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
				v4aScheme.SigningName = aws.String("rds")
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

func addCreateDBInstanceResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateDBInstanceResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
