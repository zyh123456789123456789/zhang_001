// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A document that defines an entity.
type DefinitionDocument struct {

	// The language used to define the entity. GRAPHQL is the only valid value.
	//
	// This member is required.
	Language DefinitionLanguage

	// The GraphQL text that defines the entity.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

// An object that contains the ID and revision number of a workflow or system that
// is part of a deployment.
type DependencyRevision struct {

	// The ID of the workflow or system.
	Id *string

	// The revision number of the workflow or system.
	RevisionNumber *int64

	noSmithyDocumentSerde
}

// Describes the properties of an entity.
type EntityDescription struct {

	// The entity ARN.
	Arn *string

	// The time at which the entity was created.
	CreatedAt *time.Time

	// The definition document of the entity.
	Definition *DefinitionDocument

	// The entity ID.
	Id *string

	// The entity type.
	Type EntityType

	noSmithyDocumentSerde
}

// An object that filters an entity search. Multiple filters function as OR
// criteria in the search. For example a search that includes a NAMESPACE and a
// REFERENCED_ENTITY_ID filter searches for entities in the specified namespace
// that use the entity specified by the value of REFERENCED_ENTITY_ID .
type EntityFilter struct {

	// The name of the entity search filter field. REFERENCED_ENTITY_ID filters on
	// entities that are used by the entity in the result set. For example, you can
	// filter on the ID of a property that is used in a state.
	Name EntityFilterName

	// An array of string values for the search filter field. Multiple values function
	// as AND criteria in the search.
	Value []string

	noSmithyDocumentSerde
}

// An object that contains information about a flow event.
type FlowExecutionMessage struct {

	// The type of flow event .
	EventType FlowExecutionEventType

	// The unique identifier of the message.
	MessageId *string

	// A string containing information about the flow event.
	Payload *string

	// The date and time when the message was last updated.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// An object that contains summary information about a flow execution.
type FlowExecutionSummary struct {

	// The date and time when the flow execution summary was created.
	CreatedAt *time.Time

	// The ID of the flow execution.
	FlowExecutionId *string

	// The ID of the flow.
	FlowTemplateId *string

	// The current status of the flow execution.
	Status FlowExecutionStatus

	// The ID of the system instance that contains the flow.
	SystemInstanceId *string

	// The date and time when the flow execution summary was last updated.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// An object that contains a workflow's definition and summary information.
type FlowTemplateDescription struct {

	// A workflow's definition document.
	Definition *DefinitionDocument

	// An object that contains summary information about a workflow.
	Summary *FlowTemplateSummary

	// The version of the user's namespace against which the workflow was validated.
	// Use this value in your system instance.
	ValidatedNamespaceVersion *int64

	noSmithyDocumentSerde
}

// An object that filters a workflow search.
type FlowTemplateFilter struct {

	// The name of the search filter field.
	//
	// This member is required.
	Name FlowTemplateFilterName

	// An array of string values for the search filter field. Multiple values function
	// as AND criteria in the search.
	//
	// This member is required.
	Value []string

	noSmithyDocumentSerde
}

// An object that contains summary information about a workflow.
type FlowTemplateSummary struct {

	// The ARN of the workflow.
	Arn *string

	// The date when the workflow was created.
	CreatedAt *time.Time

	// The ID of the workflow.
	Id *string

	// The revision number of the workflow.
	RevisionNumber *int64

	noSmithyDocumentSerde
}

// An object that specifies whether cloud metrics are collected in a deployment
// and, if so, what role is used to collect metrics.
type MetricsConfiguration struct {

	// A Boolean that specifies whether cloud metrics are collected.
	CloudMetricEnabled bool

	// The ARN of the role that is used to collect cloud metrics.
	MetricRuleRoleArn *string

	noSmithyDocumentSerde
}

// An object that contains a system instance definition and summary information.
type SystemInstanceDescription struct {

	// A document that defines an entity.
	Definition *DefinitionDocument

	// The AWS Identity and Access Management (IAM) role that AWS IoT Things Graph
	// assumes during flow execution in a cloud deployment. This role must have read
	// and write permissionss to AWS Lambda and AWS IoT and to any other AWS services
	// that the flow uses.
	FlowActionsRoleArn *string

	// An object that specifies whether cloud metrics are collected in a deployment
	// and, if so, what role is used to collect metrics.
	MetricsConfiguration *MetricsConfiguration

	// The Amazon Simple Storage Service bucket where information about a system
	// instance is stored.
	S3BucketName *string

	// An object that contains summary information about a system instance.
	Summary *SystemInstanceSummary

	// A list of objects that contain all of the IDs and revision numbers of workflows
	// and systems that are used in a system instance.
	ValidatedDependencyRevisions []DependencyRevision

	// The version of the user's namespace against which the system instance was
	// validated.
	ValidatedNamespaceVersion *int64

	noSmithyDocumentSerde
}

// An object that filters a system instance search. Multiple filters function as
// OR criteria in the search. For example a search that includes a
// GREENGRASS_GROUP_NAME and a STATUS filter searches for system instances in the
// specified Greengrass group that have the specified status.
type SystemInstanceFilter struct {

	// The name of the search filter field.
	Name SystemInstanceFilterName

	// An array of string values for the search filter field. Multiple values function
	// as AND criteria in the search.
	Value []string

	noSmithyDocumentSerde
}

// An object that contains summary information about a system instance.
type SystemInstanceSummary struct {

	// The ARN of the system instance.
	Arn *string

	// The date when the system instance was created.
	CreatedAt *time.Time

	// The ID of the Greengrass group where the system instance is deployed.
	GreengrassGroupId *string

	// The ID of the Greengrass group where the system instance is deployed.
	GreengrassGroupName *string

	// The version of the Greengrass group where the system instance is deployed.
	GreengrassGroupVersionId *string

	// The ID of the system instance.
	Id *string

	// The status of the system instance.
	Status SystemInstanceDeploymentStatus

	// The target of the system instance.
	Target DeploymentTarget

	// The date and time when the system instance was last updated.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// An object that contains a system's definition document and summary information.
type SystemTemplateDescription struct {

	// The definition document of a system.
	Definition *DefinitionDocument

	// An object that contains summary information about a system.
	Summary *SystemTemplateSummary

	// The namespace version against which the system was validated. Use this value in
	// your system instance.
	ValidatedNamespaceVersion *int64

	noSmithyDocumentSerde
}

// An object that filters a system search.
type SystemTemplateFilter struct {

	// The name of the system search filter field.
	//
	// This member is required.
	Name SystemTemplateFilterName

	// An array of string values for the search filter field. Multiple values function
	// as AND criteria in the search.
	//
	// This member is required.
	Value []string

	noSmithyDocumentSerde
}

// An object that contains information about a system.
type SystemTemplateSummary struct {

	// The ARN of the system.
	Arn *string

	// The date when the system was created.
	CreatedAt *time.Time

	// The ID of the system.
	Id *string

	// The revision number of the system.
	RevisionNumber *int64

	noSmithyDocumentSerde
}

// Metadata assigned to an AWS IoT Things Graph resource consisting of a key-value
// pair.
type Tag struct {

	// The required name of the tag. The string value can be from 1 to 128 Unicode
	// characters in length.
	//
	// This member is required.
	Key *string

	// The optional value of the tag. The string value can be from 1 to 256 Unicode
	// characters in length.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// An AWS IoT thing.
type Thing struct {

	// The ARN of the thing.
	ThingArn *string

	// The name of the thing.
	ThingName *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde