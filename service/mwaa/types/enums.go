// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EnvironmentStatus string

// Enum values for EnvironmentStatus
const (
	EnvironmentStatusCreating         EnvironmentStatus = "CREATING"
	EnvironmentStatusCreateFailed     EnvironmentStatus = "CREATE_FAILED"
	EnvironmentStatusAvailable        EnvironmentStatus = "AVAILABLE"
	EnvironmentStatusUpdating         EnvironmentStatus = "UPDATING"
	EnvironmentStatusDeleting         EnvironmentStatus = "DELETING"
	EnvironmentStatusDeleted          EnvironmentStatus = "DELETED"
	EnvironmentStatusUnavailable      EnvironmentStatus = "UNAVAILABLE"
	EnvironmentStatusUpdateFailed     EnvironmentStatus = "UPDATE_FAILED"
	EnvironmentStatusRollingBack      EnvironmentStatus = "ROLLING_BACK"
	EnvironmentStatusCreatingSnapshot EnvironmentStatus = "CREATING_SNAPSHOT"
)

// Values returns all known values for EnvironmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentStatus) Values() []EnvironmentStatus {
	return []EnvironmentStatus{
		"CREATING",
		"CREATE_FAILED",
		"AVAILABLE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"UNAVAILABLE",
		"UPDATE_FAILED",
		"ROLLING_BACK",
		"CREATING_SNAPSHOT",
	}
}

type LoggingLevel string

// Enum values for LoggingLevel
const (
	LoggingLevelCritical LoggingLevel = "CRITICAL"
	LoggingLevelError    LoggingLevel = "ERROR"
	LoggingLevelWarning  LoggingLevel = "WARNING"
	LoggingLevelInfo     LoggingLevel = "INFO"
	LoggingLevelDebug    LoggingLevel = "DEBUG"
)

// Values returns all known values for LoggingLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LoggingLevel) Values() []LoggingLevel {
	return []LoggingLevel{
		"CRITICAL",
		"ERROR",
		"WARNING",
		"INFO",
		"DEBUG",
	}
}

type Unit string

// Enum values for Unit
const (
	UnitSeconds            Unit = "Seconds"
	UnitMicroseconds       Unit = "Microseconds"
	UnitMilliseconds       Unit = "Milliseconds"
	UnitBytes              Unit = "Bytes"
	UnitKilobytes          Unit = "Kilobytes"
	UnitMegabytes          Unit = "Megabytes"
	UnitGigabytes          Unit = "Gigabytes"
	UnitTerabytes          Unit = "Terabytes"
	UnitBits               Unit = "Bits"
	UnitKilobits           Unit = "Kilobits"
	UnitMegabits           Unit = "Megabits"
	UnitGigabits           Unit = "Gigabits"
	UnitTerabits           Unit = "Terabits"
	UnitPercent            Unit = "Percent"
	UnitCount              Unit = "Count"
	UnitBytesPerSecond     Unit = "Bytes/Second"
	UnitKilobytesPerSecond Unit = "Kilobytes/Second"
	UnitMegabytesPerSecond Unit = "Megabytes/Second"
	UnitGigabytesPerSecond Unit = "Gigabytes/Second"
	UnitTerabytesPerSecond Unit = "Terabytes/Second"
	UnitBitsPerSecond      Unit = "Bits/Second"
	UnitKilobitsPerSecond  Unit = "Kilobits/Second"
	UnitMegabitsPerSecond  Unit = "Megabits/Second"
	UnitGigabitsPerSecond  Unit = "Gigabits/Second"
	UnitTerabitsPerSecond  Unit = "Terabits/Second"
	UnitCountPerSecond     Unit = "Count/Second"
	UnitNone               Unit = "None"
)

// Values returns all known values for Unit. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Unit) Values() []Unit {
	return []Unit{
		"Seconds",
		"Microseconds",
		"Milliseconds",
		"Bytes",
		"Kilobytes",
		"Megabytes",
		"Gigabytes",
		"Terabytes",
		"Bits",
		"Kilobits",
		"Megabits",
		"Gigabits",
		"Terabits",
		"Percent",
		"Count",
		"Bytes/Second",
		"Kilobytes/Second",
		"Megabytes/Second",
		"Gigabytes/Second",
		"Terabytes/Second",
		"Bits/Second",
		"Kilobits/Second",
		"Megabits/Second",
		"Gigabits/Second",
		"Terabits/Second",
		"Count/Second",
		"None",
	}
}

type UpdateStatus string

// Enum values for UpdateStatus
const (
	UpdateStatusSuccess UpdateStatus = "SUCCESS"
	UpdateStatusPending UpdateStatus = "PENDING"
	UpdateStatusFailed  UpdateStatus = "FAILED"
)

// Values returns all known values for UpdateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UpdateStatus) Values() []UpdateStatus {
	return []UpdateStatus{
		"SUCCESS",
		"PENDING",
		"FAILED",
	}
}

type WebserverAccessMode string

// Enum values for WebserverAccessMode
const (
	WebserverAccessModePrivateOnly WebserverAccessMode = "PRIVATE_ONLY"
	WebserverAccessModePublicOnly  WebserverAccessMode = "PUBLIC_ONLY"
)

// Values returns all known values for WebserverAccessMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WebserverAccessMode) Values() []WebserverAccessMode {
	return []WebserverAccessMode{
		"PRIVATE_ONLY",
		"PUBLIC_ONLY",
	}
}