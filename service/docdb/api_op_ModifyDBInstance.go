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

// Modifies settings for an instance. You can change one or more database
// configuration parameters by specifying these parameters and the new values in
// the request.
func (c *Client) ModifyDBInstance(ctx context.Context, params *ModifyDBInstanceInput, optFns ...func(*Options)) (*ModifyDBInstanceOutput, error) {
	if params == nil {
		params = &ModifyDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBInstance", params, optFns, c.addOperationModifyDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ModifyDBInstance .
type ModifyDBInstanceInput struct {

	// The instance identifier. This value is stored as a lowercase string.
	// Constraints:
	//   - Must match the identifier of an existing DBInstance .
	//
	// This member is required.
	DBInstanceIdentifier *string

	// Specifies whether the modifications in this request and any pending
	// modifications are asynchronously applied as soon as possible, regardless of the
	// PreferredMaintenanceWindow setting for the instance. If this parameter is set to
	// false , changes to the instance are applied during the next maintenance window.
	// Some parameter changes can cause an outage and are applied on the next reboot.
	// Default: false
	ApplyImmediately bool

	// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not
	// perform minor version upgrades regardless of the value set.
	AutoMinorVersionUpgrade *bool

	// Indicates the certificate that needs to be associated with the instance.
	CACertificateIdentifier *string

	// A value that indicates whether to copy all tags from the DB instance to
	// snapshots of the DB instance. By default, tags are not copied.
	CopyTagsToSnapshot *bool

	// The new compute and memory capacity of the instance; for example, db.r5.large .
	// Not all instance classes are available in all Amazon Web Services Regions. If
	// you modify the instance class, an outage occurs during the change. The change is
	// applied during the next maintenance window, unless ApplyImmediately is
	// specified as true for this request. Default: Uses existing setting.
	DBInstanceClass *string

	// A value that indicates whether to enable Performance Insights for the DB
	// Instance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html)
	// .
	EnablePerformanceInsights *bool

	// The new instance identifier for the instance when renaming an instance. When
	// you change the instance identifier, an instance reboot occurs immediately if you
	// set Apply Immediately to true . It occurs during the next maintenance window if
	// you set Apply Immediately to false . This value is stored as a lowercase string.
	// Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//   - The first character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	// Example: mydbinstance
	NewDBInstanceIdentifier *string

	// The KMS key identifier for encryption of Performance Insights data. The KMS key
	// identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If
	// you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon
	// DocumentDB uses your default KMS key. There is a default KMS key for your Amazon
	// Web Services account. Your Amazon Web Services account has a different default
	// KMS key for each Amazon Web Services region.
	PerformanceInsightsKMSKeyId *string

	// The weekly time range (in UTC) during which system maintenance can occur, which
	// might result in an outage. Changing this parameter doesn't result in an outage
	// except in the following situation, and the change is asynchronously applied as
	// soon as possible. If there are pending actions that cause a reboot, and the
	// maintenance window is changed to include the current time, changing this
	// parameter causes a reboot of the instance. If you are moving this window to the
	// current time, there must be at least 30 minutes between the current time and end
	// of the window to ensure that pending changes are applied. Default: Uses existing
	// setting. Format: ddd:hh24:mi-ddd:hh24:mi Valid days: Mon, Tue, Wed, Thu, Fri,
	// Sat, Sun Constraints: Must be at least 30 minutes.
	PreferredMaintenanceWindow *string

	// A value that specifies the order in which an Amazon DocumentDB replica is
	// promoted to the primary instance after a failure of the existing primary
	// instance. Default: 1 Valid values: 0-15
	PromotionTier *int32

	noSmithyDocumentSerde
}

type ModifyDBInstanceOutput struct {

	// Detailed information about an instance.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBInstance{}, middleware.After)
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
	if err = addModifyDBInstanceResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpModifyDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBInstance",
	}
}

type opModifyDBInstanceResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opModifyDBInstanceResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opModifyDBInstanceResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addModifyDBInstanceResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opModifyDBInstanceResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
