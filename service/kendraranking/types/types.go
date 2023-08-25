// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Sets additional capacity units configured for your rescore execution plan. A
// rescore execution plan is an Amazon Kendra Intelligent Ranking resource used for
// provisioning the Rescore API. You can add and remove capacity units to fit your
// usage requirements.
type CapacityUnitsConfiguration struct {

	// The amount of extra capacity for your rescore execution plan. A single extra
	// capacity unit for a rescore execution plan provides 0.01 rescore requests per
	// second. You can add up to 1000 extra capacity units.
	//
	// This member is required.
	RescoreCapacityUnits *int32

	noSmithyDocumentSerde
}

// Information about a document from a search service such as OpenSearch (self
// managed). Amazon Kendra Intelligent Ranking uses this information to rank and
// score on.
type Document struct {

	// The identifier of the document from the search service.
	//
	// This member is required.
	Id *string

	// The original document score or rank from the search service. Amazon Kendra
	// Intelligent Ranking gives the document a new score or rank based on its
	// intelligent search algorithms.
	//
	// This member is required.
	OriginalScore *float32

	// The body text of the search service's document.
	Body *string

	// The optional group identifier of the document from the search service.
	// Documents with the same group identifier are grouped together and processed as
	// one document within the service.
	GroupId *string

	// The title of the search service's document.
	Title *string

	// The body text of the search service's document represented as a list of tokens
	// or words. You must choose to provide Body or TokenizedBody . You cannot provide
	// both.
	TokenizedBody []string

	// The title of the search service's document represented as a list of tokens or
	// words. You must choose to provide Title or TokenizedTitle . You cannot provide
	// both.
	TokenizedTitle []string

	noSmithyDocumentSerde
}

// Summary information for a rescore execution plan. A rescore execution plan is
// an Amazon Kendra Intelligent Ranking resource used for provisioning the Rescore
// API.
type RescoreExecutionPlanSummary struct {

	// The Unix timestamp when the rescore execution plan was created.
	CreatedAt *time.Time

	// The identifier of the rescore execution plan.
	Id *string

	// The name of the rescore execution plan.
	Name *string

	// The current status of the rescore execution plan. When the value is ACTIVE , the
	// rescore execution plan is ready for use.
	Status RescoreExecutionPlanStatus

	// The Unix timestamp when the rescore execution plan was last updated.
	UpdatedAt *time.Time

	noSmithyDocumentSerde
}

// A result item for a document with a new relevancy score.
type RescoreResultItem struct {

	// The identifier of the document from the search service.
	DocumentId *string

	// The relevancy score or rank that Amazon Kendra Intelligent Ranking gives to the
	// result.
	Score *float32

	noSmithyDocumentSerde
}

// A key-value pair that identifies or categorizes a rescore execution plan. A
// rescore execution plan is an Amazon Kendra Intelligent Ranking resource used for
// provisioning the Rescore API. You can also use a tag to help control access to
// a rescore execution plan. A tag key and value can consist of Unicode letters,
// digits, white space, and any of the following symbols: _ . : / = + - @.
type Tag struct {

	// The key for the tag. Keys are not case sensitive and must be unique.
	//
	// This member is required.
	Key *string

	// The value associated with the tag. The value can be an empty string but it
	// can't be null.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde