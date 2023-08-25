// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// A second request to use or change an object was not allowed. This can result
// from retrying a request using a parameter that was not present in the original
// request.
type IdempotentParameterMismatchException struct {
	Message *string

	ErrorCodeOverride *string

	Code int32

	noSmithyDocumentSerde
}

func (e *IdempotentParameterMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotentParameterMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotentParameterMismatchException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "IdempotentParameterMismatchException"
	}
	return *e.ErrorCodeOverride
}
func (e *IdempotentParameterMismatchException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// An error on the server occurred when trying to process a request.
type InternalServerException struct {
	Message *string

	ErrorCodeOverride *string

	Code int32

	noSmithyDocumentSerde
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// An error on the client occurred. Typically, the cause is an invalid input value.
type InvalidInputException struct {
	Message *string

	ErrorCodeOverride *string

	Code int32

	noSmithyDocumentSerde
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidInputException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidTagException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidTagException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTagException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTagException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidTagException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidTagException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The subscriber exceeded the maximum number of operations. This exception can
// occur when listing objects such as DataSource .
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	Code int32

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The exception is thrown when a predict request is made to an unmounted MLModel .
type PredictorNotMountedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *PredictorNotMountedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PredictorNotMountedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PredictorNotMountedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PredictorNotMountedException"
	}
	return *e.ErrorCodeOverride
}
func (e *PredictorNotMountedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A specified resource cannot be located.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	Code int32

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type TagLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TagLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TagLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *TagLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
