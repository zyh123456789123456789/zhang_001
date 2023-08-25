// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies settings for a DB instance. You can change one or more database
// configuration parameters by specifying these parameters and the new values in
// the request. To learn what modifications you can make to your DB instance, call
// DescribeValidDBInstanceModifications before you call ModifyDBInstance .
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

type ModifyDBInstanceInput struct {

	// The DB instance identifier. This value is stored as a lowercase string.
	// Constraints:
	//   - Must match the identifier of an existing DBInstance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// Not supported by Neptune.
	AllocatedStorage *int32

	// Indicates that major version upgrades are allowed. Changing this parameter
	// doesn't result in an outage and the change is asynchronously applied as soon as
	// possible.
	AllowMajorVersionUpgrade bool

	// Specifies whether the modifications in this request and any pending
	// modifications are asynchronously applied as soon as possible, regardless of the
	// PreferredMaintenanceWindow setting for the DB instance. If this parameter is set
	// to false , changes to the DB instance are applied during the next maintenance
	// window. Some parameter changes can cause an outage and are applied on the next
	// call to RebootDBInstance , or the next failure reboot. Default: false
	ApplyImmediately bool

	// Indicates that minor version upgrades are applied automatically to the DB
	// instance during the maintenance window. Changing this parameter doesn't result
	// in an outage except in the following case and the change is asynchronously
	// applied as soon as possible. An outage will result if this parameter is set to
	// true during the maintenance window, and a newer minor version is available, and
	// Neptune has enabled auto patching for that engine version.
	AutoMinorVersionUpgrade *bool

	// Not applicable. The retention period for automated backups is managed by the DB
	// cluster. For more information, see ModifyDBCluster . Default: Uses existing
	// setting
	BackupRetentionPeriod *int32

	// Indicates the certificate that needs to be associated with the instance.
	CACertificateIdentifier *string

	// The configuration setting for the log types to be enabled for export to
	// CloudWatch Logs for a specific DB instance or DB cluster.
	CloudwatchLogsExportConfiguration *types.CloudwatchLogsExportConfiguration

	// True to copy all tags from the DB instance to snapshots of the DB instance, and
	// otherwise false. The default is false.
	CopyTagsToSnapshot *bool

	// The new compute and memory capacity of the DB instance, for example, db.m4.large
	// . Not all DB instance classes are available in all Amazon Regions. If you modify
	// the DB instance class, an outage occurs during the change. The change is applied
	// during the next maintenance window, unless ApplyImmediately is specified as true
	// for this request. Default: Uses existing setting
	DBInstanceClass *string

	// The name of the DB parameter group to apply to the DB instance. Changing this
	// setting doesn't result in an outage. The parameter group name itself is changed
	// immediately, but the actual parameter changes are not applied until you reboot
	// the instance without failover. The db instance will NOT be rebooted
	// automatically and the parameter changes will NOT be applied during the next
	// maintenance window. Default: Uses existing setting Constraints: The DB parameter
	// group must be in the same DB parameter group family as this DB instance.
	DBParameterGroupName *string

	// The port number on which the database accepts connections. The value of the
	// DBPortNumber parameter must not match any of the port values specified for
	// options in the option group for the DB instance. Your database will restart when
	// you change the DBPortNumber value regardless of the value of the
	// ApplyImmediately parameter. Default: 8182
	DBPortNumber *int32

	// A list of DB security groups to authorize on this DB instance. Changing this
	// setting doesn't result in an outage and the change is asynchronously applied as
	// soon as possible. Constraints:
	//   - If supplied, must match existing DBSecurityGroups.
	DBSecurityGroups []string

	// The new DB subnet group for the DB instance. You can use this parameter to move
	// your DB instance to a different VPC. Changing the subnet group causes an outage
	// during the change. The change is applied during the next maintenance window,
	// unless you specify true for the ApplyImmediately parameter. Constraints: If
	// supplied, must match the name of an existing DBSubnetGroup. Example:
	// mySubnetGroup
	DBSubnetGroupName *string

	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. See Deleting a DB Instance (https://docs.aws.amazon.com/neptune/latest/userguide/manage-console-instances-delete.html)
	// .
	DeletionProtection *bool

	// Not supported.
	Domain *string

	// Not supported
	DomainIAMRoleName *string

	// True to enable mapping of Amazon Identity and Access Management (IAM) accounts
	// to database accounts, and otherwise false. You can enable IAM database
	// authentication for the following database engines Not applicable. Mapping Amazon
	// IAM accounts to database accounts is managed by the DB cluster. For more
	// information, see ModifyDBCluster . Default: false
	EnableIAMDatabaseAuthentication *bool

	// (Not supported by Neptune)
	EnablePerformanceInsights *bool

	// The version number of the database engine to upgrade to. Currently, setting
	// this parameter has no effect. To upgrade your database engine to the most recent
	// release, use the ApplyPendingMaintenanceAction API.
	EngineVersion *string

	// The new Provisioned IOPS (I/O operations per second) value for the instance.
	// Changing this setting doesn't result in an outage and the change is applied
	// during the next maintenance window unless the ApplyImmediately parameter is set
	// to true for this request. Default: Uses existing setting
	Iops *int32

	// Not supported by Neptune.
	LicenseModel *string

	// Not supported by Neptune.
	MasterUserPassword *string

	// The interval, in seconds, between points when Enhanced Monitoring metrics are
	// collected for the DB instance. To disable collecting Enhanced Monitoring
	// metrics, specify 0. The default is 0. If MonitoringRoleArn is specified, then
	// you must also set MonitoringInterval to a value other than 0. Valid Values: 0,
	// 1, 5, 10, 15, 30, 60
	MonitoringInterval *int32

	// The ARN for the IAM role that permits Neptune to send enhanced monitoring
	// metrics to Amazon CloudWatch Logs. For example,
	// arn:aws:iam:123456789012:role/emaccess . If MonitoringInterval is set to a
	// value other than 0, then you must supply a MonitoringRoleArn value.
	MonitoringRoleArn *string

	// Specifies if the DB instance is a Multi-AZ deployment. Changing this parameter
	// doesn't result in an outage and the change is applied during the next
	// maintenance window unless the ApplyImmediately parameter is set to true for
	// this request.
	MultiAZ *bool

	// The new DB instance identifier for the DB instance when renaming a DB instance.
	// When you change the DB instance identifier, an instance reboot will occur
	// immediately if you set Apply Immediately to true, or will occur during the next
	// maintenance window if Apply Immediately to false. This value is stored as a
	// lowercase string. Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//   - The first character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	// Example: mydbinstance
	NewDBInstanceIdentifier *string

	// (Not supported by Neptune)
	OptionGroupName *string

	// (Not supported by Neptune)
	PerformanceInsightsKMSKeyId *string

	// The daily time range during which automated backups are created if automated
	// backups are enabled. Not applicable. The daily time range for creating automated
	// backups is managed by the DB cluster. For more information, see ModifyDBCluster
	// . Constraints:
	//   - Must be in the format hh24:mi-hh24:mi
	//   - Must be in Universal Time Coordinated (UTC)
	//   - Must not conflict with the preferred maintenance window
	//   - Must be at least 30 minutes
	PreferredBackupWindow *string

	// The weekly time range (in UTC) during which system maintenance can occur, which
	// might result in an outage. Changing this parameter doesn't result in an outage,
	// except in the following situation, and the change is asynchronously applied as
	// soon as possible. If there are pending actions that cause a reboot, and the
	// maintenance window is changed to include the current time, then changing this
	// parameter will cause a reboot of the DB instance. If moving this window to the
	// current time, there must be at least 30 minutes between the current time and end
	// of the window to ensure pending changes are applied. Default: Uses existing
	// setting Format: ddd:hh24:mi-ddd:hh24:mi Valid Days: Mon | Tue | Wed | Thu | Fri
	// | Sat | Sun Constraints: Must be at least 30 minutes
	PreferredMaintenanceWindow *string

	// A value that specifies the order in which a Read Replica is promoted to the
	// primary instance after a failure of the existing primary instance. Default: 1
	// Valid Values: 0 - 15
	PromotionTier *int32

	// This flag should no longer be used.
	//
	// Deprecated: This member has been deprecated.
	PubliclyAccessible *bool

	// Not supported.
	StorageType *string

	// The ARN from the key store with which to associate the instance for TDE
	// encryption.
	TdeCredentialArn *string

	// The password for the given ARN from the key store in order to access the device.
	TdeCredentialPassword *string

	// A list of EC2 VPC security groups to authorize on this DB instance. This change
	// is asynchronously applied as soon as possible. Not applicable. The associated
	// list of EC2 VPC security groups is managed by the DB cluster. For more
	// information, see ModifyDBCluster . Constraints:
	//   - If supplied, must match existing VpcSecurityGroupIds.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type ModifyDBInstanceOutput struct {

	// Contains the details of an Amazon Neptune DB instance. This data type is used
	// as a response element in the DescribeDBInstances action.
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
