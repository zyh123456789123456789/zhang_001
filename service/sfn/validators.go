// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateActivity struct {
}

func (*validateOpCreateActivity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateActivity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateActivityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateActivityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateStateMachineAlias struct {
}

func (*validateOpCreateStateMachineAlias) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateStateMachineAlias) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateStateMachineAliasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateStateMachineAliasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateStateMachine struct {
}

func (*validateOpCreateStateMachine) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateStateMachine) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateStateMachineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateStateMachineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteActivity struct {
}

func (*validateOpDeleteActivity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteActivity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteActivityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteActivityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteStateMachineAlias struct {
}

func (*validateOpDeleteStateMachineAlias) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteStateMachineAlias) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteStateMachineAliasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteStateMachineAliasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteStateMachine struct {
}

func (*validateOpDeleteStateMachine) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteStateMachine) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteStateMachineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteStateMachineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteStateMachineVersion struct {
}

func (*validateOpDeleteStateMachineVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteStateMachineVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteStateMachineVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteStateMachineVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeActivity struct {
}

func (*validateOpDescribeActivity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeActivity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeActivityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeActivityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeExecution struct {
}

func (*validateOpDescribeExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeMapRun struct {
}

func (*validateOpDescribeMapRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeMapRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeMapRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeMapRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeStateMachineAlias struct {
}

func (*validateOpDescribeStateMachineAlias) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeStateMachineAlias) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeStateMachineAliasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeStateMachineAliasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeStateMachineForExecution struct {
}

func (*validateOpDescribeStateMachineForExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeStateMachineForExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeStateMachineForExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeStateMachineForExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeStateMachine struct {
}

func (*validateOpDescribeStateMachine) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeStateMachine) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeStateMachineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeStateMachineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetActivityTask struct {
}

func (*validateOpGetActivityTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetActivityTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetActivityTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetActivityTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetExecutionHistory struct {
}

func (*validateOpGetExecutionHistory) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetExecutionHistory) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetExecutionHistoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetExecutionHistoryInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListMapRuns struct {
}

func (*validateOpListMapRuns) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListMapRuns) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListMapRunsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListMapRunsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListStateMachineAliases struct {
}

func (*validateOpListStateMachineAliases) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListStateMachineAliases) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListStateMachineAliasesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListStateMachineAliasesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListStateMachineVersions struct {
}

func (*validateOpListStateMachineVersions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListStateMachineVersions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListStateMachineVersionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListStateMachineVersionsInput(input); err != nil {
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

type validateOpPublishStateMachineVersion struct {
}

func (*validateOpPublishStateMachineVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPublishStateMachineVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PublishStateMachineVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPublishStateMachineVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendTaskFailure struct {
}

func (*validateOpSendTaskFailure) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendTaskFailure) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendTaskFailureInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendTaskFailureInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendTaskHeartbeat struct {
}

func (*validateOpSendTaskHeartbeat) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendTaskHeartbeat) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendTaskHeartbeatInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendTaskHeartbeatInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendTaskSuccess struct {
}

func (*validateOpSendTaskSuccess) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendTaskSuccess) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendTaskSuccessInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendTaskSuccessInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartExecution struct {
}

func (*validateOpStartExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartSyncExecution struct {
}

func (*validateOpStartSyncExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartSyncExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartSyncExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartSyncExecutionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopExecution struct {
}

func (*validateOpStopExecution) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopExecutionInput(input); err != nil {
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

type validateOpUpdateMapRun struct {
}

func (*validateOpUpdateMapRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateMapRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateMapRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateMapRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateStateMachineAlias struct {
}

func (*validateOpUpdateStateMachineAlias) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateStateMachineAlias) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateStateMachineAliasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateStateMachineAliasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateStateMachine struct {
}

func (*validateOpUpdateStateMachine) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateStateMachine) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateStateMachineInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateStateMachineInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateActivityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateActivity{}, middleware.After)
}

func addOpCreateStateMachineAliasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateStateMachineAlias{}, middleware.After)
}

func addOpCreateStateMachineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateStateMachine{}, middleware.After)
}

func addOpDeleteActivityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteActivity{}, middleware.After)
}

func addOpDeleteStateMachineAliasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteStateMachineAlias{}, middleware.After)
}

func addOpDeleteStateMachineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteStateMachine{}, middleware.After)
}

func addOpDeleteStateMachineVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteStateMachineVersion{}, middleware.After)
}

func addOpDescribeActivityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeActivity{}, middleware.After)
}

func addOpDescribeExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeExecution{}, middleware.After)
}

func addOpDescribeMapRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeMapRun{}, middleware.After)
}

func addOpDescribeStateMachineAliasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeStateMachineAlias{}, middleware.After)
}

func addOpDescribeStateMachineForExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeStateMachineForExecution{}, middleware.After)
}

func addOpDescribeStateMachineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeStateMachine{}, middleware.After)
}

func addOpGetActivityTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetActivityTask{}, middleware.After)
}

func addOpGetExecutionHistoryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetExecutionHistory{}, middleware.After)
}

func addOpListMapRunsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListMapRuns{}, middleware.After)
}

func addOpListStateMachineAliasesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListStateMachineAliases{}, middleware.After)
}

func addOpListStateMachineVersionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListStateMachineVersions{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPublishStateMachineVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPublishStateMachineVersion{}, middleware.After)
}

func addOpSendTaskFailureValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendTaskFailure{}, middleware.After)
}

func addOpSendTaskHeartbeatValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendTaskHeartbeat{}, middleware.After)
}

func addOpSendTaskSuccessValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendTaskSuccess{}, middleware.After)
}

func addOpStartExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartExecution{}, middleware.After)
}

func addOpStartSyncExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartSyncExecution{}, middleware.After)
}

func addOpStopExecutionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopExecution{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateMapRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateMapRun{}, middleware.After)
}

func addOpUpdateStateMachineAliasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateStateMachineAlias{}, middleware.After)
}

func addOpUpdateStateMachineValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateStateMachine{}, middleware.After)
}

func validateRoutingConfigurationList(v []types.RoutingConfigurationListItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RoutingConfigurationList"}
	for i := range v {
		if err := validateRoutingConfigurationListItem(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRoutingConfigurationListItem(v *types.RoutingConfigurationListItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RoutingConfigurationListItem"}
	if v.StateMachineVersionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineVersionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateActivityInput(v *CreateActivityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateActivityInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateStateMachineAliasInput(v *CreateStateMachineAliasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateStateMachineAliasInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.RoutingConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoutingConfiguration"))
	} else if v.RoutingConfiguration != nil {
		if err := validateRoutingConfigurationList(v.RoutingConfiguration); err != nil {
			invalidParams.AddNested("RoutingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateStateMachineInput(v *CreateStateMachineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateStateMachineInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Definition == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Definition"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteActivityInput(v *DeleteActivityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteActivityInput"}
	if v.ActivityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActivityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteStateMachineAliasInput(v *DeleteStateMachineAliasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteStateMachineAliasInput"}
	if v.StateMachineAliasArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineAliasArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteStateMachineInput(v *DeleteStateMachineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteStateMachineInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteStateMachineVersionInput(v *DeleteStateMachineVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteStateMachineVersionInput"}
	if v.StateMachineVersionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineVersionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeActivityInput(v *DescribeActivityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeActivityInput"}
	if v.ActivityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActivityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeExecutionInput(v *DescribeExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeExecutionInput"}
	if v.ExecutionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeMapRunInput(v *DescribeMapRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeMapRunInput"}
	if v.MapRunArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MapRunArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeStateMachineAliasInput(v *DescribeStateMachineAliasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeStateMachineAliasInput"}
	if v.StateMachineAliasArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineAliasArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeStateMachineForExecutionInput(v *DescribeStateMachineForExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeStateMachineForExecutionInput"}
	if v.ExecutionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeStateMachineInput(v *DescribeStateMachineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeStateMachineInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetActivityTaskInput(v *GetActivityTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetActivityTaskInput"}
	if v.ActivityArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActivityArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetExecutionHistoryInput(v *GetExecutionHistoryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetExecutionHistoryInput"}
	if v.ExecutionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListMapRunsInput(v *ListMapRunsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListMapRunsInput"}
	if v.ExecutionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListStateMachineAliasesInput(v *ListStateMachineAliasesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListStateMachineAliasesInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListStateMachineVersionsInput(v *ListStateMachineVersionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListStateMachineVersionsInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
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

func validateOpPublishStateMachineVersionInput(v *PublishStateMachineVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PublishStateMachineVersionInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendTaskFailureInput(v *SendTaskFailureInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendTaskFailureInput"}
	if v.TaskToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendTaskHeartbeatInput(v *SendTaskHeartbeatInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendTaskHeartbeatInput"}
	if v.TaskToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendTaskSuccessInput(v *SendTaskSuccessInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendTaskSuccessInput"}
	if v.TaskToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskToken"))
	}
	if v.Output == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Output"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartExecutionInput(v *StartExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartExecutionInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartSyncExecutionInput(v *StartSyncExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartSyncExecutionInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopExecutionInput(v *StopExecutionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopExecutionInput"}
	if v.ExecutionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionArn"))
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

func validateOpUpdateMapRunInput(v *UpdateMapRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateMapRunInput"}
	if v.MapRunArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MapRunArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateStateMachineAliasInput(v *UpdateStateMachineAliasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateStateMachineAliasInput"}
	if v.StateMachineAliasArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineAliasArn"))
	}
	if v.RoutingConfiguration != nil {
		if err := validateRoutingConfigurationList(v.RoutingConfiguration); err != nil {
			invalidParams.AddNested("RoutingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateStateMachineInput(v *UpdateStateMachineInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateStateMachineInput"}
	if v.StateMachineArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateMachineArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}