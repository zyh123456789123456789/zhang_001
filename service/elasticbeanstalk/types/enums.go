// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActionHistoryStatus string

// Enum values for ActionHistoryStatus
const (
	ActionHistoryStatusCompleted ActionHistoryStatus = "Completed"
	ActionHistoryStatusFailed    ActionHistoryStatus = "Failed"
	ActionHistoryStatusUnknown   ActionHistoryStatus = "Unknown"
)

// Values returns all known values for ActionHistoryStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActionHistoryStatus) Values() []ActionHistoryStatus {
	return []ActionHistoryStatus{
		"Completed",
		"Failed",
		"Unknown",
	}
}

type ActionStatus string

// Enum values for ActionStatus
const (
	ActionStatusScheduled ActionStatus = "Scheduled"
	ActionStatusPending   ActionStatus = "Pending"
	ActionStatusRunning   ActionStatus = "Running"
	ActionStatusUnknown   ActionStatus = "Unknown"
)

// Values returns all known values for ActionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActionStatus) Values() []ActionStatus {
	return []ActionStatus{
		"Scheduled",
		"Pending",
		"Running",
		"Unknown",
	}
}

type ActionType string

// Enum values for ActionType
const (
	ActionTypeInstanceRefresh ActionType = "InstanceRefresh"
	ActionTypePlatformUpdate  ActionType = "PlatformUpdate"
	ActionTypeUnknown         ActionType = "Unknown"
)

// Values returns all known values for ActionType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ActionType) Values() []ActionType {
	return []ActionType{
		"InstanceRefresh",
		"PlatformUpdate",
		"Unknown",
	}
}

type ApplicationVersionStatus string

// Enum values for ApplicationVersionStatus
const (
	ApplicationVersionStatusProcessed   ApplicationVersionStatus = "Processed"
	ApplicationVersionStatusUnprocessed ApplicationVersionStatus = "Unprocessed"
	ApplicationVersionStatusFailed      ApplicationVersionStatus = "Failed"
	ApplicationVersionStatusProcessing  ApplicationVersionStatus = "Processing"
	ApplicationVersionStatusBuilding    ApplicationVersionStatus = "Building"
)

// Values returns all known values for ApplicationVersionStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationVersionStatus) Values() []ApplicationVersionStatus {
	return []ApplicationVersionStatus{
		"Processed",
		"Unprocessed",
		"Failed",
		"Processing",
		"Building",
	}
}

type ComputeType string

// Enum values for ComputeType
const (
	ComputeTypeBuildGeneral1Small  ComputeType = "BUILD_GENERAL1_SMALL"
	ComputeTypeBuildGeneral1Medium ComputeType = "BUILD_GENERAL1_MEDIUM"
	ComputeTypeBuildGeneral1Large  ComputeType = "BUILD_GENERAL1_LARGE"
)

// Values returns all known values for ComputeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ComputeType) Values() []ComputeType {
	return []ComputeType{
		"BUILD_GENERAL1_SMALL",
		"BUILD_GENERAL1_MEDIUM",
		"BUILD_GENERAL1_LARGE",
	}
}

type ConfigurationDeploymentStatus string

// Enum values for ConfigurationDeploymentStatus
const (
	ConfigurationDeploymentStatusDeployed ConfigurationDeploymentStatus = "deployed"
	ConfigurationDeploymentStatusPending  ConfigurationDeploymentStatus = "pending"
	ConfigurationDeploymentStatusFailed   ConfigurationDeploymentStatus = "failed"
)

// Values returns all known values for ConfigurationDeploymentStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ConfigurationDeploymentStatus) Values() []ConfigurationDeploymentStatus {
	return []ConfigurationDeploymentStatus{
		"deployed",
		"pending",
		"failed",
	}
}

type ConfigurationOptionValueType string

// Enum values for ConfigurationOptionValueType
const (
	ConfigurationOptionValueTypeScalar ConfigurationOptionValueType = "Scalar"
	ConfigurationOptionValueTypeList   ConfigurationOptionValueType = "List"
)

// Values returns all known values for ConfigurationOptionValueType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ConfigurationOptionValueType) Values() []ConfigurationOptionValueType {
	return []ConfigurationOptionValueType{
		"Scalar",
		"List",
	}
}

type EnvironmentHealth string

// Enum values for EnvironmentHealth
const (
	EnvironmentHealthGreen  EnvironmentHealth = "Green"
	EnvironmentHealthYellow EnvironmentHealth = "Yellow"
	EnvironmentHealthRed    EnvironmentHealth = "Red"
	EnvironmentHealthGrey   EnvironmentHealth = "Grey"
)

// Values returns all known values for EnvironmentHealth. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentHealth) Values() []EnvironmentHealth {
	return []EnvironmentHealth{
		"Green",
		"Yellow",
		"Red",
		"Grey",
	}
}

type EnvironmentHealthAttribute string

// Enum values for EnvironmentHealthAttribute
const (
	EnvironmentHealthAttributeStatus             EnvironmentHealthAttribute = "Status"
	EnvironmentHealthAttributeColor              EnvironmentHealthAttribute = "Color"
	EnvironmentHealthAttributeCauses             EnvironmentHealthAttribute = "Causes"
	EnvironmentHealthAttributeApplicationMetrics EnvironmentHealthAttribute = "ApplicationMetrics"
	EnvironmentHealthAttributeInstancesHealth    EnvironmentHealthAttribute = "InstancesHealth"
	EnvironmentHealthAttributeAll                EnvironmentHealthAttribute = "All"
	EnvironmentHealthAttributeHealthStatus       EnvironmentHealthAttribute = "HealthStatus"
	EnvironmentHealthAttributeRefreshedAt        EnvironmentHealthAttribute = "RefreshedAt"
)

// Values returns all known values for EnvironmentHealthAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentHealthAttribute) Values() []EnvironmentHealthAttribute {
	return []EnvironmentHealthAttribute{
		"Status",
		"Color",
		"Causes",
		"ApplicationMetrics",
		"InstancesHealth",
		"All",
		"HealthStatus",
		"RefreshedAt",
	}
}

type EnvironmentHealthStatus string

// Enum values for EnvironmentHealthStatus
const (
	EnvironmentHealthStatusNoData    EnvironmentHealthStatus = "NoData"
	EnvironmentHealthStatusUnknown   EnvironmentHealthStatus = "Unknown"
	EnvironmentHealthStatusPending   EnvironmentHealthStatus = "Pending"
	EnvironmentHealthStatusOk        EnvironmentHealthStatus = "Ok"
	EnvironmentHealthStatusInfo      EnvironmentHealthStatus = "Info"
	EnvironmentHealthStatusWarning   EnvironmentHealthStatus = "Warning"
	EnvironmentHealthStatusDegraded  EnvironmentHealthStatus = "Degraded"
	EnvironmentHealthStatusSevere    EnvironmentHealthStatus = "Severe"
	EnvironmentHealthStatusSuspended EnvironmentHealthStatus = "Suspended"
)

// Values returns all known values for EnvironmentHealthStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentHealthStatus) Values() []EnvironmentHealthStatus {
	return []EnvironmentHealthStatus{
		"NoData",
		"Unknown",
		"Pending",
		"Ok",
		"Info",
		"Warning",
		"Degraded",
		"Severe",
		"Suspended",
	}
}

type EnvironmentInfoType string

// Enum values for EnvironmentInfoType
const (
	EnvironmentInfoTypeTail   EnvironmentInfoType = "tail"
	EnvironmentInfoTypeBundle EnvironmentInfoType = "bundle"
)

// Values returns all known values for EnvironmentInfoType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentInfoType) Values() []EnvironmentInfoType {
	return []EnvironmentInfoType{
		"tail",
		"bundle",
	}
}

type EnvironmentStatus string

// Enum values for EnvironmentStatus
const (
	EnvironmentStatusAborting    EnvironmentStatus = "Aborting"
	EnvironmentStatusLaunching   EnvironmentStatus = "Launching"
	EnvironmentStatusUpdating    EnvironmentStatus = "Updating"
	EnvironmentStatusLinkingFrom EnvironmentStatus = "LinkingFrom"
	EnvironmentStatusLinkingTo   EnvironmentStatus = "LinkingTo"
	EnvironmentStatusReady       EnvironmentStatus = "Ready"
	EnvironmentStatusTerminating EnvironmentStatus = "Terminating"
	EnvironmentStatusTerminated  EnvironmentStatus = "Terminated"
)

// Values returns all known values for EnvironmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentStatus) Values() []EnvironmentStatus {
	return []EnvironmentStatus{
		"Aborting",
		"Launching",
		"Updating",
		"LinkingFrom",
		"LinkingTo",
		"Ready",
		"Terminating",
		"Terminated",
	}
}

type EventSeverity string

// Enum values for EventSeverity
const (
	EventSeverityTrace EventSeverity = "TRACE"
	EventSeverityDebug EventSeverity = "DEBUG"
	EventSeverityInfo  EventSeverity = "INFO"
	EventSeverityWarn  EventSeverity = "WARN"
	EventSeverityError EventSeverity = "ERROR"
	EventSeverityFatal EventSeverity = "FATAL"
)

// Values returns all known values for EventSeverity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EventSeverity) Values() []EventSeverity {
	return []EventSeverity{
		"TRACE",
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
		"FATAL",
	}
}

type FailureType string

// Enum values for FailureType
const (
	FailureTypeUpdateCancelled         FailureType = "UpdateCancelled"
	FailureTypeCancellationFailed      FailureType = "CancellationFailed"
	FailureTypeRollbackFailed          FailureType = "RollbackFailed"
	FailureTypeRollbackSuccessful      FailureType = "RollbackSuccessful"
	FailureTypeInternalFailure         FailureType = "InternalFailure"
	FailureTypeInvalidEnvironmentState FailureType = "InvalidEnvironmentState"
	FailureTypePermissionsError        FailureType = "PermissionsError"
)

// Values returns all known values for FailureType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FailureType) Values() []FailureType {
	return []FailureType{
		"UpdateCancelled",
		"CancellationFailed",
		"RollbackFailed",
		"RollbackSuccessful",
		"InternalFailure",
		"InvalidEnvironmentState",
		"PermissionsError",
	}
}

type InstancesHealthAttribute string

// Enum values for InstancesHealthAttribute
const (
	InstancesHealthAttributeHealthStatus       InstancesHealthAttribute = "HealthStatus"
	InstancesHealthAttributeColor              InstancesHealthAttribute = "Color"
	InstancesHealthAttributeCauses             InstancesHealthAttribute = "Causes"
	InstancesHealthAttributeApplicationMetrics InstancesHealthAttribute = "ApplicationMetrics"
	InstancesHealthAttributeRefreshedAt        InstancesHealthAttribute = "RefreshedAt"
	InstancesHealthAttributeLaunchedAt         InstancesHealthAttribute = "LaunchedAt"
	InstancesHealthAttributeSystem             InstancesHealthAttribute = "System"
	InstancesHealthAttributeDeployment         InstancesHealthAttribute = "Deployment"
	InstancesHealthAttributeAvailabilityZone   InstancesHealthAttribute = "AvailabilityZone"
	InstancesHealthAttributeInstanceType       InstancesHealthAttribute = "InstanceType"
	InstancesHealthAttributeAll                InstancesHealthAttribute = "All"
)

// Values returns all known values for InstancesHealthAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (InstancesHealthAttribute) Values() []InstancesHealthAttribute {
	return []InstancesHealthAttribute{
		"HealthStatus",
		"Color",
		"Causes",
		"ApplicationMetrics",
		"RefreshedAt",
		"LaunchedAt",
		"System",
		"Deployment",
		"AvailabilityZone",
		"InstanceType",
		"All",
	}
}

type PlatformStatus string

// Enum values for PlatformStatus
const (
	PlatformStatusCreating PlatformStatus = "Creating"
	PlatformStatusFailed   PlatformStatus = "Failed"
	PlatformStatusReady    PlatformStatus = "Ready"
	PlatformStatusDeleting PlatformStatus = "Deleting"
	PlatformStatusDeleted  PlatformStatus = "Deleted"
)

// Values returns all known values for PlatformStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlatformStatus) Values() []PlatformStatus {
	return []PlatformStatus{
		"Creating",
		"Failed",
		"Ready",
		"Deleting",
		"Deleted",
	}
}

type SourceRepository string

// Enum values for SourceRepository
const (
	SourceRepositoryCodeCommit SourceRepository = "CodeCommit"
	SourceRepositoryS3         SourceRepository = "S3"
)

// Values returns all known values for SourceRepository. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SourceRepository) Values() []SourceRepository {
	return []SourceRepository{
		"CodeCommit",
		"S3",
	}
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeGit SourceType = "Git"
	SourceTypeZip SourceType = "Zip"
)

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"Git",
		"Zip",
	}
}

type ValidationSeverity string

// Enum values for ValidationSeverity
const (
	ValidationSeverityError   ValidationSeverity = "error"
	ValidationSeverityWarning ValidationSeverity = "warning"
)

// Values returns all known values for ValidationSeverity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ValidationSeverity) Values() []ValidationSeverity {
	return []ValidationSeverity{
		"error",
		"warning",
	}
}