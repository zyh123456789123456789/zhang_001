// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAcceptResourceShareInvitation struct {
}

func (*validateOpAcceptResourceShareInvitation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAcceptResourceShareInvitation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AcceptResourceShareInvitationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAcceptResourceShareInvitationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateResourceShare struct {
}

func (*validateOpAssociateResourceShare) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateResourceShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateResourceShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateResourceShareInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateResourceSharePermission struct {
}

func (*validateOpAssociateResourceSharePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateResourceSharePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateResourceSharePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateResourceSharePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePermission struct {
}

func (*validateOpCreatePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePermissionVersion struct {
}

func (*validateOpCreatePermissionVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePermissionVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePermissionVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePermissionVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateResourceShare struct {
}

func (*validateOpCreateResourceShare) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateResourceShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateResourceShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateResourceShareInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePermission struct {
}

func (*validateOpDeletePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePermissionVersion struct {
}

func (*validateOpDeletePermissionVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePermissionVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePermissionVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePermissionVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteResourceShare struct {
}

func (*validateOpDeleteResourceShare) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteResourceShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteResourceShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteResourceShareInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateResourceShare struct {
}

func (*validateOpDisassociateResourceShare) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateResourceShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateResourceShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateResourceShareInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateResourceSharePermission struct {
}

func (*validateOpDisassociateResourceSharePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateResourceSharePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateResourceSharePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateResourceSharePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPermission struct {
}

func (*validateOpGetPermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourcePolicies struct {
}

func (*validateOpGetResourcePolicies) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourcePolicies) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourcePoliciesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourcePoliciesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourceShareAssociations struct {
}

func (*validateOpGetResourceShareAssociations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourceShareAssociations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourceShareAssociationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourceShareAssociationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourceShares struct {
}

func (*validateOpGetResourceShares) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourceShares) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourceSharesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourceSharesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPendingInvitationResources struct {
}

func (*validateOpListPendingInvitationResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPendingInvitationResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPendingInvitationResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPendingInvitationResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPermissionVersions struct {
}

func (*validateOpListPermissionVersions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPermissionVersions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPermissionVersionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPermissionVersionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPrincipals struct {
}

func (*validateOpListPrincipals) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPrincipals) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPrincipalsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPrincipalsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListResourceSharePermissions struct {
}

func (*validateOpListResourceSharePermissions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListResourceSharePermissions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListResourceSharePermissionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListResourceSharePermissionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListResources struct {
}

func (*validateOpListResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPromotePermissionCreatedFromPolicy struct {
}

func (*validateOpPromotePermissionCreatedFromPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPromotePermissionCreatedFromPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PromotePermissionCreatedFromPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPromotePermissionCreatedFromPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPromoteResourceShareCreatedFromPolicy struct {
}

func (*validateOpPromoteResourceShareCreatedFromPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPromoteResourceShareCreatedFromPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PromoteResourceShareCreatedFromPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPromoteResourceShareCreatedFromPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRejectResourceShareInvitation struct {
}

func (*validateOpRejectResourceShareInvitation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRejectResourceShareInvitation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RejectResourceShareInvitationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRejectResourceShareInvitationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpReplacePermissionAssociations struct {
}

func (*validateOpReplacePermissionAssociations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpReplacePermissionAssociations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ReplacePermissionAssociationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpReplacePermissionAssociationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetDefaultPermissionVersion struct {
}

func (*validateOpSetDefaultPermissionVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetDefaultPermissionVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetDefaultPermissionVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetDefaultPermissionVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateResourceShare struct {
}

func (*validateOpUpdateResourceShare) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateResourceShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateResourceShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateResourceShareInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAcceptResourceShareInvitationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAcceptResourceShareInvitation{}, middleware.After)
}

func addOpAssociateResourceShareValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateResourceShare{}, middleware.After)
}

func addOpAssociateResourceSharePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateResourceSharePermission{}, middleware.After)
}

func addOpCreatePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePermission{}, middleware.After)
}

func addOpCreatePermissionVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePermissionVersion{}, middleware.After)
}

func addOpCreateResourceShareValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateResourceShare{}, middleware.After)
}

func addOpDeletePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePermission{}, middleware.After)
}

func addOpDeletePermissionVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePermissionVersion{}, middleware.After)
}

func addOpDeleteResourceShareValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteResourceShare{}, middleware.After)
}

func addOpDisassociateResourceShareValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateResourceShare{}, middleware.After)
}

func addOpDisassociateResourceSharePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateResourceSharePermission{}, middleware.After)
}

func addOpGetPermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPermission{}, middleware.After)
}

func addOpGetResourcePoliciesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResourcePolicies{}, middleware.After)
}

func addOpGetResourceShareAssociationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResourceShareAssociations{}, middleware.After)
}

func addOpGetResourceSharesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResourceShares{}, middleware.After)
}

func addOpListPendingInvitationResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListPendingInvitationResources{}, middleware.After)
}

func addOpListPermissionVersionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListPermissionVersions{}, middleware.After)
}

func addOpListPrincipalsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListPrincipals{}, middleware.After)
}

func addOpListResourceSharePermissionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListResourceSharePermissions{}, middleware.After)
}

func addOpListResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListResources{}, middleware.After)
}

func addOpPromotePermissionCreatedFromPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPromotePermissionCreatedFromPolicy{}, middleware.After)
}

func addOpPromoteResourceShareCreatedFromPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPromoteResourceShareCreatedFromPolicy{}, middleware.After)
}

func addOpRejectResourceShareInvitationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRejectResourceShareInvitation{}, middleware.After)
}

func addOpReplacePermissionAssociationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpReplacePermissionAssociations{}, middleware.After)
}

func addOpSetDefaultPermissionVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetDefaultPermissionVersion{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateResourceShareValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateResourceShare{}, middleware.After)
}

func validateOpAcceptResourceShareInvitationInput(v *AcceptResourceShareInvitationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AcceptResourceShareInvitationInput"}
	if v.ResourceShareInvitationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareInvitationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateResourceShareInput(v *AssociateResourceShareInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateResourceShareInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateResourceSharePermissionInput(v *AssociateResourceSharePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateResourceSharePermissionInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePermissionInput(v *CreatePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePermissionInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ResourceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if v.PolicyTemplate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyTemplate"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePermissionVersionInput(v *CreatePermissionVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePermissionVersionInput"}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if v.PolicyTemplate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyTemplate"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateResourceShareInput(v *CreateResourceShareInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateResourceShareInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePermissionInput(v *DeletePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePermissionInput"}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePermissionVersionInput(v *DeletePermissionVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePermissionVersionInput"}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if v.PermissionVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteResourceShareInput(v *DeleteResourceShareInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteResourceShareInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateResourceShareInput(v *DisassociateResourceShareInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateResourceShareInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateResourceSharePermissionInput(v *DisassociateResourceSharePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateResourceSharePermissionInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPermissionInput(v *GetPermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPermissionInput"}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourcePoliciesInput(v *GetResourcePoliciesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourcePoliciesInput"}
	if v.ResourceArns == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArns"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourceShareAssociationsInput(v *GetResourceShareAssociationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourceShareAssociationsInput"}
	if len(v.AssociationType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AssociationType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourceSharesInput(v *GetResourceSharesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourceSharesInput"}
	if len(v.ResourceOwner) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceOwner"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPendingInvitationResourcesInput(v *ListPendingInvitationResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPendingInvitationResourcesInput"}
	if v.ResourceShareInvitationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareInvitationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPermissionVersionsInput(v *ListPermissionVersionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPermissionVersionsInput"}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPrincipalsInput(v *ListPrincipalsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPrincipalsInput"}
	if len(v.ResourceOwner) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceOwner"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListResourceSharePermissionsInput(v *ListResourceSharePermissionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListResourceSharePermissionsInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListResourcesInput(v *ListResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListResourcesInput"}
	if len(v.ResourceOwner) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceOwner"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPromotePermissionCreatedFromPolicyInput(v *PromotePermissionCreatedFromPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PromotePermissionCreatedFromPolicyInput"}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPromoteResourceShareCreatedFromPolicyInput(v *PromoteResourceShareCreatedFromPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PromoteResourceShareCreatedFromPolicyInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRejectResourceShareInvitationInput(v *RejectResourceShareInvitationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RejectResourceShareInvitationInput"}
	if v.ResourceShareInvitationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareInvitationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpReplacePermissionAssociationsInput(v *ReplacePermissionAssociationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ReplacePermissionAssociationsInput"}
	if v.FromPermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FromPermissionArn"))
	}
	if v.ToPermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ToPermissionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetDefaultPermissionVersionInput(v *SetDefaultPermissionVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetDefaultPermissionVersionInput"}
	if v.PermissionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionArn"))
	}
	if v.PermissionVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateResourceShareInput(v *UpdateResourceShareInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateResourceShareInput"}
	if v.ResourceShareArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceShareArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
