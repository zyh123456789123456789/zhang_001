// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type LinuxSubscriptionsDiscovery string

// Enum values for LinuxSubscriptionsDiscovery
const (
	// Enabled LinuxSubscriptionsDiscovery
	LinuxSubscriptionsDiscoveryEnabled LinuxSubscriptionsDiscovery = "Enabled"
	// Disabled LinuxSubscriptionsDiscovery
	LinuxSubscriptionsDiscoveryDisabled LinuxSubscriptionsDiscovery = "Disabled"
)

// Values returns all known values for LinuxSubscriptionsDiscovery. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LinuxSubscriptionsDiscovery) Values() []LinuxSubscriptionsDiscovery {
	return []LinuxSubscriptionsDiscovery{
		"Enabled",
		"Disabled",
	}
}

type Operator string

// Enum values for Operator
const (
	// Equal operator
	OperatorEqual Operator = "Equal"
	// Not equal operator
	OperatorNotEqual Operator = "NotEqual"
	// Contains operator
	OperatorContains Operator = "Contains"
)

// Values returns all known values for Operator. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Operator) Values() []Operator {
	return []Operator{
		"Equal",
		"NotEqual",
		"Contains",
	}
}

type OrganizationIntegration string

// Enum values for OrganizationIntegration
const (
	// Enabled OrganizationIntegration
	OrganizationIntegrationEnabled OrganizationIntegration = "Enabled"
	// Disabled OrganizationIntegration
	OrganizationIntegrationDisabled OrganizationIntegration = "Disabled"
)

// Values returns all known values for OrganizationIntegration. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OrganizationIntegration) Values() []OrganizationIntegration {
	return []OrganizationIntegration{
		"Enabled",
		"Disabled",
	}
}

type Status string

// Enum values for Status
const (
	// InProgress status
	StatusInProgress Status = "InProgress"
	// Completed status
	StatusCompleted Status = "Completed"
	// Successful status
	StatusSuccessful Status = "Successful"
	// Failed status
	StatusFailed Status = "Failed"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"InProgress",
		"Completed",
		"Successful",
		"Failed",
	}
}
