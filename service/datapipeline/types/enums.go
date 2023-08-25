// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type OperatorType string

// Enum values for OperatorType
const (
	OperatorTypeEqual              OperatorType = "EQ"
	OperatorTypeReferenceEqual     OperatorType = "REF_EQ"
	OperatorTypeLessThanOrEqual    OperatorType = "LE"
	OperatorTypeGreaterThanOrEqual OperatorType = "GE"
	OperatorTypeBetween            OperatorType = "BETWEEN"
)

// Values returns all known values for OperatorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperatorType) Values() []OperatorType {
	return []OperatorType{
		"EQ",
		"REF_EQ",
		"LE",
		"GE",
		"BETWEEN",
	}
}

type TaskStatus string

// Enum values for TaskStatus
const (
	TaskStatusFinished TaskStatus = "FINISHED"
	TaskStatusFailed   TaskStatus = "FAILED"
	TaskStatusFalse    TaskStatus = "FALSE"
)

// Values returns all known values for TaskStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TaskStatus) Values() []TaskStatus {
	return []TaskStatus{
		"FINISHED",
		"FAILED",
		"FALSE",
	}
}
