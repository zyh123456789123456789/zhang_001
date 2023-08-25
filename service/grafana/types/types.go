// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A structure that defines which attributes in the IdP assertion are to be used
// to define information about the users authenticated by the IdP to use the
// workspace.
type AssertionAttributes struct {

	// The name of the attribute within the SAML assertion to use as the email names
	// for SAML users.
	Email *string

	// The name of the attribute within the SAML assertion to use as the user full
	// "friendly" names for user groups.
	Groups *string

	// The name of the attribute within the SAML assertion to use as the login names
	// for SAML users.
	Login *string

	// The name of the attribute within the SAML assertion to use as the user full
	// "friendly" names for SAML users.
	Name *string

	// The name of the attribute within the SAML assertion to use as the user full
	// "friendly" names for the users' organizations.
	Org *string

	// The name of the attribute within the SAML assertion to use as the user roles.
	Role *string

	noSmithyDocumentSerde
}

// A structure containing information about the user authentication methods used
// by the workspace.
type AuthenticationDescription struct {

	// Specifies whether this workspace uses IAM Identity Center, SAML, or both
	// methods to authenticate users to use the Grafana console in the Amazon Managed
	// Grafana workspace.
	//
	// This member is required.
	Providers []AuthenticationProviderTypes

	// A structure containing information about how this workspace works with IAM
	// Identity Center.
	AwsSso *AwsSsoAuthentication

	// A structure containing information about how this workspace works with SAML,
	// including what attributes within the assertion are to be mapped to user
	// information in the workspace.
	Saml *SamlAuthentication

	noSmithyDocumentSerde
}

// A structure that describes whether the workspace uses SAML, IAM Identity
// Center, or both methods for user authentication, and whether that authentication
// is fully configured.
type AuthenticationSummary struct {

	// Specifies whether the workspace uses SAML, IAM Identity Center, or both methods
	// for user authentication.
	//
	// This member is required.
	Providers []AuthenticationProviderTypes

	// Specifies whether the workplace's user authentication method is fully
	// configured.
	SamlConfigurationStatus SamlConfigurationStatus

	noSmithyDocumentSerde
}

// A structure containing information about how this workspace works with IAM
// Identity Center.
type AwsSsoAuthentication struct {

	// The ID of the IAM Identity Center-managed application that is created by Amazon
	// Managed Grafana.
	SsoClientId *string

	noSmithyDocumentSerde
}

// A structure containing the identity provider (IdP) metadata used to integrate
// the identity provider with this workspace. You can specify the metadata either
// by providing a URL to its location in the url parameter, or by specifying the
// full metadata in XML format in the xml parameter. Specifying both will cause an
// error.
//
// The following types satisfy this interface:
//
//	IdpMetadataMemberUrl
//	IdpMetadataMemberXml
type IdpMetadata interface {
	isIdpMetadata()
}

// The URL of the location containing the IdP metadata.
type IdpMetadataMemberUrl struct {
	Value string

	noSmithyDocumentSerde
}

func (*IdpMetadataMemberUrl) isIdpMetadata() {}

// The full IdP metadata, in XML format.
type IdpMetadataMemberXml struct {
	Value string

	noSmithyDocumentSerde
}

func (*IdpMetadataMemberXml) isIdpMetadata() {}

// The configuration settings for in-bound network access to your workspace. When
// this is configured, only listed IP addresses and VPC endpoints will be able to
// access your workspace. Standard Grafana authentication and authorization are
// still required. Access is granted to a caller that is in either the IP address
// list or the VPC endpoint list - they do not need to be in both. If this is not
// configured, or is removed, then all IP addresses and VPC endpoints are allowed.
// Standard Grafana authentication and authorization are still required. While both
// prefixListIds and vpceIds are required, you can pass in an empty array of
// strings for either parameter if you do not want to allow any of that type. If
// both are passed as empty arrays, no traffic is allowed to the workspace, because
// only explicitly allowed connections are accepted.
type NetworkAccessConfiguration struct {

	// An array of prefix list IDs. A prefix list is a list of CIDR ranges of IP
	// addresses. The IP addresses specified are allowed to access your workspace. If
	// the list is not included in the configuration (passed an empty array) then no IP
	// addresses are allowed to access the workspace. You create a prefix list using
	// the Amazon VPC console. Prefix list IDs have the format pl-1a2b3c4d . For more
	// information about prefix lists, see Group CIDR blocks using managed prefix lists (https://docs.aws.amazon.com/vpc/latest/userguide/managed-prefix-lists.html)
	// in the Amazon Virtual Private Cloud User Guide.
	//
	// This member is required.
	PrefixListIds []string

	// An array of Amazon VPC endpoint IDs for the workspace. You can create VPC
	// endpoints to your Amazon Managed Grafana workspace for access from within a VPC.
	// If a NetworkAccessConfiguration is specified then only VPC endpoints specified
	// here are allowed to access the workspace. If you pass in an empty array of
	// strings, then no VPCs are allowed to access the workspace. VPC endpoint IDs have
	// the format vpce-1a2b3c4d . For more information about creating an interface VPC
	// endpoint, see Interface VPC endpoints (https://docs.aws.amazon.com/grafana/latest/userguide/VPC-endpoints)
	// in the Amazon Managed Grafana User Guide. The only VPC endpoints that can be
	// specified here are interface VPC endpoints for Grafana workspaces (using the
	// com.amazonaws.[region].grafana-workspace service endpoint). Other VPC endpoints
	// are ignored.
	//
	// This member is required.
	VpceIds []string

	noSmithyDocumentSerde
}

// A structure containing the identity of one user or group and the Admin , Editor
// , or Viewer role that they have.
type PermissionEntry struct {

	// Specifies whether the user or group has the Admin , Editor , or Viewer role.
	//
	// This member is required.
	Role Role

	// A structure with the ID of the user or group with this role.
	//
	// This member is required.
	User *User

	noSmithyDocumentSerde
}

// This structure defines which groups defined in the SAML assertion attribute are
// to be mapped to the Grafana Admin and Editor roles in the workspace. SAML
// authenticated users not part of Admin or Editor role groups have Viewer
// permission over the workspace.
type RoleValues struct {

	// A list of groups from the SAML assertion attribute to grant the Grafana Admin
	// role to.
	Admin []string

	// A list of groups from the SAML assertion attribute to grant the Grafana Editor
	// role to.
	Editor []string

	noSmithyDocumentSerde
}

// A structure containing information about how this workspace works with SAML.
type SamlAuthentication struct {

	// Specifies whether the workspace's SAML configuration is complete.
	//
	// This member is required.
	Status SamlConfigurationStatus

	// A structure containing details about how this workspace works with SAML.
	Configuration *SamlConfiguration

	noSmithyDocumentSerde
}

// A structure containing information about how this workspace works with SAML.
type SamlConfiguration struct {

	// A structure containing the identity provider (IdP) metadata used to integrate
	// the identity provider with this workspace.
	//
	// This member is required.
	IdpMetadata IdpMetadata

	// Lists which organizations defined in the SAML assertion are allowed to use the
	// Amazon Managed Grafana workspace. If this is empty, all organizations in the
	// assertion attribute have access.
	AllowedOrganizations []string

	// A structure that defines which attributes in the SAML assertion are to be used
	// to define information about the users authenticated by that IdP to use the
	// workspace.
	AssertionAttributes *AssertionAttributes

	// How long a sign-on session by a SAML user is valid, before the user has to sign
	// on again.
	LoginValidityDuration int32

	// A structure containing arrays that map group names in the SAML assertion to the
	// Grafana Admin and Editor roles in the workspace.
	RoleValues *RoleValues

	noSmithyDocumentSerde
}

// A structure containing information about one error encountered while performing
// an UpdatePermissions (https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdatePermissions.html)
// operation.
type UpdateError struct {

	// Specifies which permission update caused the error.
	//
	// This member is required.
	CausedBy *UpdateInstruction

	// The error code.
	//
	// This member is required.
	Code *int32

	// The message for this error.
	//
	// This member is required.
	Message *string

	noSmithyDocumentSerde
}

// Contains the instructions for one Grafana role permission update in a
// UpdatePermissions (https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdatePermissions.html)
// operation.
type UpdateInstruction struct {

	// Specifies whether this update is to add or revoke role permissions.
	//
	// This member is required.
	Action UpdateAction

	// The role to add or revoke for the user or the group specified in users .
	//
	// This member is required.
	Role Role

	// A structure that specifies the user or group to add or revoke the role for.
	//
	// This member is required.
	Users []User

	noSmithyDocumentSerde
}

// A structure that specifies one user or group in the workspace.
type User struct {

	// The ID of the user or group. Pattern:
	// ^([0-9a-fA-F]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$
	//
	// This member is required.
	Id *string

	// Specifies whether this is a single user or a group.
	//
	// This member is required.
	Type UserType

	noSmithyDocumentSerde
}

// A structure that contains information about a request parameter that caused an
// error.
type ValidationExceptionField struct {

	// A message describing why this field couldn't be validated.
	//
	// This member is required.
	Message *string

	// The name of the field that caused the validation error.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// The configuration settings for an Amazon VPC that contains data sources for
// your Grafana workspace to connect to. Provided securityGroupIds and subnetIds
// must be part of the same VPC. Connecting to a private VPC is not yet available
// in the Asia Pacific (Seoul) Region (ap-northeast-2).
type VpcConfiguration struct {

	// The list of Amazon EC2 security group IDs attached to the Amazon VPC for your
	// Grafana workspace to connect. Duplicates not allowed.
	//
	// This member is required.
	SecurityGroupIds []string

	// The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana
	// workspace to connect. Duplicates not allowed.
	//
	// This member is required.
	SubnetIds []string

	noSmithyDocumentSerde
}

// A structure containing information about an Amazon Managed Grafana workspace in
// your account.
type WorkspaceDescription struct {

	// A structure that describes whether the workspace uses SAML, IAM Identity
	// Center, or both methods for user authentication.
	//
	// This member is required.
	Authentication *AuthenticationSummary

	// The date that the workspace was created.
	//
	// This member is required.
	Created *time.Time

	// Specifies the Amazon Web Services data sources that have been configured to
	// have IAM roles and permissions created to allow Amazon Managed Grafana to read
	// data from these sources. This list is only used when the workspace was created
	// through the Amazon Web Services console, and the permissionType is
	// SERVICE_MANAGED .
	//
	// This member is required.
	DataSources []DataSourceType

	// The URL that users can use to access the Grafana console in the workspace.
	//
	// This member is required.
	Endpoint *string

	// The version of Grafana supported in this workspace.
	//
	// This member is required.
	GrafanaVersion *string

	// The unique ID of this workspace.
	//
	// This member is required.
	Id *string

	// The most recent date that the workspace was modified.
	//
	// This member is required.
	Modified *time.Time

	// The current status of the workspace.
	//
	// This member is required.
	Status WorkspaceStatus

	// Specifies whether the workspace can access Amazon Web Services resources in
	// this Amazon Web Services account only, or whether it can also access Amazon Web
	// Services resources in other accounts in the same organization. If this is
	// ORGANIZATION , the workspaceOrganizationalUnits parameter specifies which
	// organizational units the workspace can access.
	AccountAccessType AccountAccessType

	// The user-defined description of the workspace.
	Description *string

	// Specifies whether this workspace has already fully used its free trial for
	// Grafana Enterprise.
	FreeTrialConsumed *bool

	// If this workspace is currently in the free trial period for Grafana Enterprise,
	// this value specifies when that free trial ends.
	FreeTrialExpiration *time.Time

	// If this workspace has a full Grafana Enterprise license, this specifies when
	// the license ends and will need to be renewed.
	LicenseExpiration *time.Time

	// Specifies whether this workspace has a full Grafana Enterprise license or a
	// free trial license.
	LicenseType LicenseType

	// The name of the workspace.
	Name *string

	// The configuration settings for network access to your workspace.
	NetworkAccessControl *NetworkAccessConfiguration

	// The Amazon Web Services notification channels that Amazon Managed Grafana can
	// automatically create IAM roles and permissions for, to allow Amazon Managed
	// Grafana to use these channels.
	NotificationDestinations []NotificationDestinationType

	// The name of the IAM role that is used to access resources through Organizations.
	OrganizationRoleName *string

	// Specifies the organizational units that this workspace is allowed to use data
	// sources from, if this workspace is in an account that is part of an
	// organization.
	OrganizationalUnits []string

	// If this is SERVICE_MANAGED , and the workplace was created through the Amazon
	// Managed Grafana console, then Amazon Managed Grafana automatically creates the
	// IAM roles and provisions the permissions that the workspace needs to use Amazon
	// Web Services data sources and notification channels. If this is CUSTOMER_MANAGED
	// , you must manage those roles and permissions yourself. If you are working with
	// a workspace in a member account of an organization and that account is not a
	// delegated administrator account, and you want the workspace to access data
	// sources in other Amazon Web Services accounts in the organization, this
	// parameter must be set to CUSTOMER_MANAGED . For more information about
	// converting between customer and service managed, see Managing permissions for
	// data sources and notification channels (https://docs.aws.amazon.com/grafana/latest/userguide/AMG-datasource-and-notification.html)
	// . For more information about the roles and permissions that must be managed for
	// customer managed workspaces, see Amazon Managed Grafana permissions and
	// policies for Amazon Web Services data sources and notification channels (https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-permissions.html)
	PermissionType PermissionType

	// The name of the CloudFormation stack set that is used to generate IAM roles to
	// be used for this workspace.
	StackSetName *string

	// The list of tags associated with the workspace.
	Tags map[string]string

	// The configuration for connecting to data sources in a private VPC (Amazon
	// Virtual Private Cloud).
	VpcConfiguration *VpcConfiguration

	// The IAM role that grants permissions to the Amazon Web Services resources that
	// the workspace will view data from. This role must already exist.
	WorkspaceRoleArn *string

	noSmithyDocumentSerde
}

// A structure that contains some information about one workspace in the account.
type WorkspaceSummary struct {

	// A structure containing information about the authentication methods used in the
	// workspace.
	//
	// This member is required.
	Authentication *AuthenticationSummary

	// The date that the workspace was created.
	//
	// This member is required.
	Created *time.Time

	// The URL endpoint to use to access the Grafana console in the workspace.
	//
	// This member is required.
	Endpoint *string

	// The Grafana version that the workspace is running.
	//
	// This member is required.
	GrafanaVersion *string

	// The unique ID of the workspace.
	//
	// This member is required.
	Id *string

	// The most recent date that the workspace was modified.
	//
	// This member is required.
	Modified *time.Time

	// The current status of the workspace.
	//
	// This member is required.
	Status WorkspaceStatus

	// The customer-entered description of the workspace.
	Description *string

	// The name of the workspace.
	Name *string

	// The Amazon Web Services notification channels that Amazon Managed Grafana can
	// automatically create IAM roles and permissions for, which allows Amazon Managed
	// Grafana to use these channels.
	NotificationDestinations []NotificationDestinationType

	// The list of tags associated with the workspace.
	Tags map[string]string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isIdpMetadata() {}