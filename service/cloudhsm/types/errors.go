// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Indicates that an internal error occurred.
type CloudHsmInternalException struct {
	Message *string

	ErrorCodeOverride *string

	Retryable bool

	noSmithyDocumentSerde
}

func (e *CloudHsmInternalException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CloudHsmInternalException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CloudHsmInternalException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CloudHsmInternalException"
	}
	return *e.ErrorCodeOverride
}
func (e *CloudHsmInternalException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Indicates that an exception occurred in the AWS CloudHSM service.
type CloudHsmServiceException struct {
	Message *string

	ErrorCodeOverride *string

	Retryable bool

	noSmithyDocumentSerde
}

func (e *CloudHsmServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CloudHsmServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CloudHsmServiceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CloudHsmServiceException"
	}
	return *e.ErrorCodeOverride
}
func (e *CloudHsmServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that one or more of the request parameters are not valid.
type InvalidRequestException struct {
	Message *string

	ErrorCodeOverride *string

	Retryable bool

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }