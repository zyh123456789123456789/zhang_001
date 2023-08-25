// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DataType string

// Enum values for DataType
const (
	DataTypeString     DataType = "STRING"
	DataTypeInteger    DataType = "INTEGER"
	DataTypeStringlist DataType = "STRINGLIST"
	DataTypeStringmap  DataType = "STRINGMAP"
)

// Values returns all known values for DataType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (DataType) Values() []DataType {
	return []DataType{
		"STRING",
		"INTEGER",
		"STRINGLIST",
		"STRINGMAP",
	}
}

type MigrationWorkflowStatusEnum string

// Enum values for MigrationWorkflowStatusEnum
const (
	MigrationWorkflowStatusEnumCreating              MigrationWorkflowStatusEnum = "CREATING"
	MigrationWorkflowStatusEnumNotStarted            MigrationWorkflowStatusEnum = "NOT_STARTED"
	MigrationWorkflowStatusEnumCreationFailed        MigrationWorkflowStatusEnum = "CREATION_FAILED"
	MigrationWorkflowStatusEnumStarting              MigrationWorkflowStatusEnum = "STARTING"
	MigrationWorkflowStatusEnumInProgress            MigrationWorkflowStatusEnum = "IN_PROGRESS"
	MigrationWorkflowStatusEnumWorkflowFailed        MigrationWorkflowStatusEnum = "WORKFLOW_FAILED"
	MigrationWorkflowStatusEnumPaused                MigrationWorkflowStatusEnum = "PAUSED"
	MigrationWorkflowStatusEnumPausing               MigrationWorkflowStatusEnum = "PAUSING"
	MigrationWorkflowStatusEnumPausingFailed         MigrationWorkflowStatusEnum = "PAUSING_FAILED"
	MigrationWorkflowStatusEnumUserAttentionRequired MigrationWorkflowStatusEnum = "USER_ATTENTION_REQUIRED"
	MigrationWorkflowStatusEnumDeleting              MigrationWorkflowStatusEnum = "DELETING"
	MigrationWorkflowStatusEnumDeletionFailed        MigrationWorkflowStatusEnum = "DELETION_FAILED"
	MigrationWorkflowStatusEnumDeleted               MigrationWorkflowStatusEnum = "DELETED"
	MigrationWorkflowStatusEnumCompleted             MigrationWorkflowStatusEnum = "COMPLETED"
)

// Values returns all known values for MigrationWorkflowStatusEnum. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (MigrationWorkflowStatusEnum) Values() []MigrationWorkflowStatusEnum {
	return []MigrationWorkflowStatusEnum{
		"CREATING",
		"NOT_STARTED",
		"CREATION_FAILED",
		"STARTING",
		"IN_PROGRESS",
		"WORKFLOW_FAILED",
		"PAUSED",
		"PAUSING",
		"PAUSING_FAILED",
		"USER_ATTENTION_REQUIRED",
		"DELETING",
		"DELETION_FAILED",
		"DELETED",
		"COMPLETED",
	}
}

type Owner string

// Enum values for Owner
const (
	OwnerAWSManaged Owner = "AWS_MANAGED"
	OwnerCustom     Owner = "CUSTOM"
)

// Values returns all known values for Owner. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Owner) Values() []Owner {
	return []Owner{
		"AWS_MANAGED",
		"CUSTOM",
	}
}

type PluginHealth string

// Enum values for PluginHealth
const (
	PluginHealthPluginHealthy   PluginHealth = "HEALTHY"
	PluginHealthPluginUnhealthy PluginHealth = "UNHEALTHY"
)

// Values returns all known values for PluginHealth. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PluginHealth) Values() []PluginHealth {
	return []PluginHealth{
		"HEALTHY",
		"UNHEALTHY",
	}
}

type RunEnvironment string

// Enum values for RunEnvironment
const (
	RunEnvironmentAws       RunEnvironment = "AWS"
	RunEnvironmentOnpremise RunEnvironment = "ONPREMISE"
)

// Values returns all known values for RunEnvironment. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RunEnvironment) Values() []RunEnvironment {
	return []RunEnvironment{
		"AWS",
		"ONPREMISE",
	}
}

type StepActionType string

// Enum values for StepActionType
const (
	StepActionTypeManual    StepActionType = "MANUAL"
	StepActionTypeAutomated StepActionType = "AUTOMATED"
)

// Values returns all known values for StepActionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StepActionType) Values() []StepActionType {
	return []StepActionType{
		"MANUAL",
		"AUTOMATED",
	}
}

type StepGroupStatus string

// Enum values for StepGroupStatus
const (
	StepGroupStatusAwaitingDependencies  StepGroupStatus = "AWAITING_DEPENDENCIES"
	StepGroupStatusReady                 StepGroupStatus = "READY"
	StepGroupStatusInProgress            StepGroupStatus = "IN_PROGRESS"
	StepGroupStatusCompleted             StepGroupStatus = "COMPLETED"
	StepGroupStatusFailed                StepGroupStatus = "FAILED"
	StepGroupStatusPaused                StepGroupStatus = "PAUSED"
	StepGroupStatusPausing               StepGroupStatus = "PAUSING"
	StepGroupStatusUserAttentionRequired StepGroupStatus = "USER_ATTENTION_REQUIRED"
)

// Values returns all known values for StepGroupStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StepGroupStatus) Values() []StepGroupStatus {
	return []StepGroupStatus{
		"AWAITING_DEPENDENCIES",
		"READY",
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
		"PAUSED",
		"PAUSING",
		"USER_ATTENTION_REQUIRED",
	}
}

type StepStatus string

// Enum values for StepStatus
const (
	StepStatusAwaitingDependencies  StepStatus = "AWAITING_DEPENDENCIES"
	StepStatusReady                 StepStatus = "READY"
	StepStatusInProgress            StepStatus = "IN_PROGRESS"
	StepStatusCompleted             StepStatus = "COMPLETED"
	StepStatusFailed                StepStatus = "FAILED"
	StepStatusPaused                StepStatus = "PAUSED"
	StepStatusUserAttentionRequired StepStatus = "USER_ATTENTION_REQUIRED"
)

// Values returns all known values for StepStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StepStatus) Values() []StepStatus {
	return []StepStatus{
		"AWAITING_DEPENDENCIES",
		"READY",
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
		"PAUSED",
		"USER_ATTENTION_REQUIRED",
	}
}

type TargetType string

// Enum values for TargetType
const (
	TargetTypeSingle TargetType = "SINGLE"
	TargetTypeAll    TargetType = "ALL"
	TargetTypeNone   TargetType = "NONE"
)

// Values returns all known values for TargetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetType) Values() []TargetType {
	return []TargetType{
		"SINGLE",
		"ALL",
		"NONE",
	}
}

type TemplateStatus string

// Enum values for TemplateStatus
const (
	TemplateStatusCreated TemplateStatus = "CREATED"
)

// Values returns all known values for TemplateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TemplateStatus) Values() []TemplateStatus {
	return []TemplateStatus{
		"CREATED",
	}
}
