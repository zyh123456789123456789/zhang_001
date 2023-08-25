// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

type ACLAlreadyExistsFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ACLAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ACLAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ACLAlreadyExistsFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ACLAlreadyExistsFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ACLAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ACLNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ACLNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ACLNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ACLNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ACLNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ACLNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ACLQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ACLQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ACLQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ACLQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ACLQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ACLQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type APICallRateForCustomerExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *APICallRateForCustomerExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *APICallRateForCustomerExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *APICallRateForCustomerExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "APICallRateForCustomerExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *APICallRateForCustomerExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

type ClusterAlreadyExistsFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ClusterAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterAlreadyExistsFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ClusterAlreadyExistsFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ClusterAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ClusterNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ClusterNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ClusterNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ClusterNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ClusterQuotaForCustomerExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ClusterQuotaForCustomerExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClusterQuotaForCustomerExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClusterQuotaForCustomerExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ClusterQuotaForCustomerExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ClusterQuotaForCustomerExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

type DefaultUserRequired struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DefaultUserRequired) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DefaultUserRequired) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DefaultUserRequired) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DefaultUserRequired"
	}
	return *e.ErrorCodeOverride
}
func (e *DefaultUserRequired) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type DuplicateUserNameFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DuplicateUserNameFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateUserNameFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateUserNameFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DuplicateUserNameFault"
	}
	return *e.ErrorCodeOverride
}
func (e *DuplicateUserNameFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InsufficientClusterCapacityFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InsufficientClusterCapacityFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientClusterCapacityFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientClusterCapacityFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InsufficientClusterCapacityFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InsufficientClusterCapacityFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidACLStateFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidACLStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidACLStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidACLStateFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidACLStateFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidACLStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidARNFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidARNFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidARNFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidARNFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidARNFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidARNFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidClusterStateFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidClusterStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidClusterStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidClusterStateFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidClusterStateFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidClusterStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidCredentialsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidCredentialsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCredentialsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCredentialsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidCredentialsException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidCredentialsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidKMSKeyFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidKMSKeyFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidKMSKeyFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidKMSKeyFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidKMSKeyFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidKMSKeyFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidNodeStateFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidNodeStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNodeStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNodeStateFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidNodeStateFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidNodeStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidParameterCombinationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterCombinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterCombinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterCombinationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterCombinationException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterCombinationException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

type InvalidParameterGroupStateFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterGroupStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterGroupStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterGroupStateFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterGroupStateFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterGroupStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidParameterValueException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterValueException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidSnapshotStateFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidSnapshotStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSnapshotStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSnapshotStateFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSnapshotStateFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSnapshotStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidSubnet struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidSubnet) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSubnet) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSubnet) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSubnet"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSubnet) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidUserStateFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidUserStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidUserStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidUserStateFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidUserStateFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidUserStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidVPCNetworkStateFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidVPCNetworkStateFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidVPCNetworkStateFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidVPCNetworkStateFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidVPCNetworkStateFault"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidVPCNetworkStateFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type NodeQuotaForClusterExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NodeQuotaForClusterExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NodeQuotaForClusterExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NodeQuotaForClusterExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NodeQuotaForClusterExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *NodeQuotaForClusterExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type NodeQuotaForCustomerExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NodeQuotaForCustomerExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NodeQuotaForCustomerExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NodeQuotaForCustomerExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NodeQuotaForCustomerExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *NodeQuotaForCustomerExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type NoOperationFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NoOperationFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoOperationFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoOperationFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NoOperationFault"
	}
	return *e.ErrorCodeOverride
}
func (e *NoOperationFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ParameterGroupAlreadyExistsFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ParameterGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ParameterGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ParameterGroupAlreadyExistsFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ParameterGroupAlreadyExistsFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ParameterGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ParameterGroupNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ParameterGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ParameterGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ParameterGroupNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ParameterGroupNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ParameterGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ParameterGroupQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ParameterGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ParameterGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ParameterGroupQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ParameterGroupQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ParameterGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You already have a reservation with the given identifier.
type ReservedNodeAlreadyExistsFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ReservedNodeAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReservedNodeAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReservedNodeAlreadyExistsFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ReservedNodeAlreadyExistsFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ReservedNodeAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested node does not exist.
type ReservedNodeNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ReservedNodeNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReservedNodeNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReservedNodeNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ReservedNodeNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ReservedNodeNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request cannot be processed because it would exceed the user's node quota.
type ReservedNodeQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ReservedNodeQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReservedNodeQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReservedNodeQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ReservedNodeQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ReservedNodeQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested node offering does not exist.
type ReservedNodesOfferingNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ReservedNodesOfferingNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReservedNodesOfferingNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReservedNodesOfferingNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ReservedNodesOfferingNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ReservedNodesOfferingNotFoundFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

type ServiceLinkedRoleNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceLinkedRoleNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceLinkedRoleNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceLinkedRoleNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceLinkedRoleNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceLinkedRoleNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ServiceUpdateNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceUpdateNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUpdateNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUpdateNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceUpdateNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUpdateNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ShardNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ShardNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ShardNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ShardNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ShardNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ShardNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ShardsPerClusterQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ShardsPerClusterQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ShardsPerClusterQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ShardsPerClusterQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ShardsPerClusterQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *ShardsPerClusterQuotaExceededFault) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

type SnapshotAlreadyExistsFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SnapshotAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotAlreadyExistsFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SnapshotAlreadyExistsFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SnapshotAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type SnapshotNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SnapshotNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SnapshotNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SnapshotNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type SnapshotQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SnapshotQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SnapshotQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SnapshotQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SnapshotQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SnapshotQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type SubnetGroupAlreadyExistsFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetGroupAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetGroupAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetGroupAlreadyExistsFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetGroupAlreadyExistsFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetGroupAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type SubnetGroupInUseFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetGroupInUseFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetGroupInUseFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetGroupInUseFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetGroupInUseFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetGroupInUseFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type SubnetGroupNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetGroupNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetGroupNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetGroupNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetGroupNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetGroupNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type SubnetGroupQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetGroupQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetGroupQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetGroupQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetGroupQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetGroupQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type SubnetInUse struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetInUse) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetInUse"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type SubnetNotAllowedFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetNotAllowedFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetNotAllowedFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetNotAllowedFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetNotAllowedFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetNotAllowedFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type SubnetQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SubnetQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type TagNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TagNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TagNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *TagNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type TagQuotaPerResourceExceeded struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TagQuotaPerResourceExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagQuotaPerResourceExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagQuotaPerResourceExceeded) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TagQuotaPerResourceExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *TagQuotaPerResourceExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type TestFailoverNotAvailableFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TestFailoverNotAvailableFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TestFailoverNotAvailableFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TestFailoverNotAvailableFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TestFailoverNotAvailableFault"
	}
	return *e.ErrorCodeOverride
}
func (e *TestFailoverNotAvailableFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type UserAlreadyExistsFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UserAlreadyExistsFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserAlreadyExistsFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserAlreadyExistsFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UserAlreadyExistsFault"
	}
	return *e.ErrorCodeOverride
}
func (e *UserAlreadyExistsFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type UserNotFoundFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UserNotFoundFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserNotFoundFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserNotFoundFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UserNotFoundFault"
	}
	return *e.ErrorCodeOverride
}
func (e *UserNotFoundFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type UserQuotaExceededFault struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UserQuotaExceededFault) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserQuotaExceededFault) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserQuotaExceededFault) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UserQuotaExceededFault"
	}
	return *e.ErrorCodeOverride
}
func (e *UserQuotaExceededFault) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }