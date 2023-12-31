// Code generated by smithy-go-codegen DO NOT EDIT.

package ivsrealtime

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateParticipantToken struct {
}

func (*validateOpCreateParticipantToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateParticipantToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateParticipantTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateParticipantTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteStage struct {
}

func (*validateOpDeleteStage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteStage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteStageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteStageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisconnectParticipant struct {
}

func (*validateOpDisconnectParticipant) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisconnectParticipant) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisconnectParticipantInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisconnectParticipantInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetParticipant struct {
}

func (*validateOpGetParticipant) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetParticipant) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetParticipantInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetParticipantInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetStage struct {
}

func (*validateOpGetStage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetStage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetStageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetStageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetStageSession struct {
}

func (*validateOpGetStageSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetStageSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetStageSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetStageSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListParticipantEvents struct {
}

func (*validateOpListParticipantEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListParticipantEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListParticipantEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListParticipantEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListParticipants struct {
}

func (*validateOpListParticipants) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListParticipants) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListParticipantsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListParticipantsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListStageSessions struct {
}

func (*validateOpListStageSessions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListStageSessions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListStageSessionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListStageSessionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
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

type validateOpUpdateStage struct {
}

func (*validateOpUpdateStage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateStage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateStageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateStageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateParticipantTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateParticipantToken{}, middleware.After)
}

func addOpDeleteStageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteStage{}, middleware.After)
}

func addOpDisconnectParticipantValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisconnectParticipant{}, middleware.After)
}

func addOpGetParticipantValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetParticipant{}, middleware.After)
}

func addOpGetStageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetStage{}, middleware.After)
}

func addOpGetStageSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetStageSession{}, middleware.After)
}

func addOpListParticipantEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListParticipantEvents{}, middleware.After)
}

func addOpListParticipantsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListParticipants{}, middleware.After)
}

func addOpListStageSessionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListStageSessions{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateStageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateStage{}, middleware.After)
}

func validateOpCreateParticipantTokenInput(v *CreateParticipantTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateParticipantTokenInput"}
	if v.StageArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StageArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteStageInput(v *DeleteStageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteStageInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisconnectParticipantInput(v *DisconnectParticipantInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisconnectParticipantInput"}
	if v.StageArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StageArn"))
	}
	if v.ParticipantId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParticipantId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetParticipantInput(v *GetParticipantInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetParticipantInput"}
	if v.StageArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StageArn"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if v.ParticipantId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParticipantId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetStageInput(v *GetStageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetStageInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetStageSessionInput(v *GetStageSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetStageSessionInput"}
	if v.StageArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StageArn"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListParticipantEventsInput(v *ListParticipantEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListParticipantEventsInput"}
	if v.StageArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StageArn"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if v.ParticipantId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ParticipantId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListParticipantsInput(v *ListParticipantsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListParticipantsInput"}
	if v.StageArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StageArn"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListStageSessionsInput(v *ListStageSessionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListStageSessionsInput"}
	if v.StageArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StageArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
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
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateStageInput(v *UpdateStageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateStageInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
