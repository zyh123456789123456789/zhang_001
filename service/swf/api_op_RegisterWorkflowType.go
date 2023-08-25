// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a new workflow type and its configuration settings in the specified
// domain. The retention period for the workflow history is set by the
// RegisterDomain action. If the type already exists, then a TypeAlreadyExists
// fault is returned. You cannot change the configuration settings of a workflow
// type once it is registered and it must be registered as a new version. Access
// Control You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//   - Use a Resource element with the domain name to limit the action to only
//     specified domains.
//   - Use an Action element to allow or deny permission to call this action.
//   - Constrain the following parameters by using a Condition element with the
//     appropriate keys.
//   - defaultTaskList.name : String constraint. The key is
//     swf:defaultTaskList.name .
//   - name : String constraint. The key is swf:name .
//   - version : String constraint. The key is swf:version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) RegisterWorkflowType(ctx context.Context, params *RegisterWorkflowTypeInput, optFns ...func(*Options)) (*RegisterWorkflowTypeOutput, error) {
	if params == nil {
		params = &RegisterWorkflowTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterWorkflowType", params, optFns, c.addOperationRegisterWorkflowTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterWorkflowTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterWorkflowTypeInput struct {

	// The name of the domain in which to register the workflow type.
	//
	// This member is required.
	Domain *string

	// The name of the workflow type. The specified string must not contain a :
	// (colon), / (slash), | (vertical bar), or any control characters ( \u0000-\u001f
	// | \u007f-\u009f ). Also, it must not be the literal string arn .
	//
	// This member is required.
	Name *string

	// The version of the workflow type. The workflow type consists of the name and
	// version, the combination of which must be unique within the domain. To get a
	// list of all currently registered workflow types, use the ListWorkflowTypes
	// action. The specified string must not contain a : (colon), / (slash), |
	// (vertical bar), or any control characters ( \u0000-\u001f | \u007f-\u009f ).
	// Also, it must not be the literal string arn .
	//
	// This member is required.
	Version *string

	// If set, specifies the default policy to use for the child workflow executions
	// when a workflow execution of this type is terminated, by calling the
	// TerminateWorkflowExecution action explicitly or due to an expired timeout. This
	// default can be overridden when starting a workflow execution using the
	// StartWorkflowExecution action or the StartChildWorkflowExecution Decision . The
	// supported child policies are:
	//   - TERMINATE – The child executions are terminated.
	//   - REQUEST_CANCEL – A request to cancel is attempted for each child execution
	//   by recording a WorkflowExecutionCancelRequested event in its history. It is up
	//   to the decider to take appropriate actions when it receives an execution history
	//   with this event.
	//   - ABANDON – No action is taken. The child executions continue to run.
	DefaultChildPolicy types.ChildPolicy

	// If set, specifies the default maximum duration for executions of this workflow
	// type. You can override this default when starting an execution through the
	// StartWorkflowExecution Action or StartChildWorkflowExecution Decision . The
	// duration is specified in seconds; an integer greater than or equal to 0. Unlike
	// some of the other timeout parameters in Amazon SWF, you cannot specify a value
	// of "NONE" for defaultExecutionStartToCloseTimeout ; there is a one-year max
	// limit on the time that a workflow execution can run. Exceeding this limit always
	// causes the workflow execution to time out.
	DefaultExecutionStartToCloseTimeout *string

	// The default IAM role attached to this workflow type. Executions of this
	// workflow type need IAM roles to invoke Lambda functions. If you don't specify an
	// IAM role when you start this workflow type, the default Lambda role is attached
	// to the execution. For more information, see
	// https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html (https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html)
	// in the Amazon SWF Developer Guide.
	DefaultLambdaRole *string

	// If set, specifies the default task list to use for scheduling decision tasks
	// for executions of this workflow type. This default is used only if a task list
	// isn't provided when starting the execution through the StartWorkflowExecution
	// Action or StartChildWorkflowExecution Decision .
	DefaultTaskList *types.TaskList

	// The default task priority to assign to the workflow type. If not assigned, then
	// 0 is used. Valid values are integers that range from Java's Integer.MIN_VALUE
	// (-2147483648) to Integer.MAX_VALUE (2147483647). Higher numbers indicate higher
	// priority. For more information about setting task priority, see Setting Task
	// Priority (https://docs.aws.amazon.com/amazonswf/latest/developerguide/programming-priority.html)
	// in the Amazon SWF Developer Guide.
	DefaultTaskPriority *string

	// If set, specifies the default maximum duration of decision tasks for this
	// workflow type. This default can be overridden when starting a workflow execution
	// using the StartWorkflowExecution action or the StartChildWorkflowExecution
	// Decision . The duration is specified in seconds, an integer greater than or
	// equal to 0 . You can use NONE to specify unlimited duration.
	DefaultTaskStartToCloseTimeout *string

	// Textual description of the workflow type.
	Description *string

	noSmithyDocumentSerde
}

type RegisterWorkflowTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterWorkflowTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRegisterWorkflowType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRegisterWorkflowType{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addRegisterWorkflowTypeResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRegisterWorkflowTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterWorkflowType(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opRegisterWorkflowType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "RegisterWorkflowType",
	}
}

type opRegisterWorkflowTypeResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opRegisterWorkflowTypeResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opRegisterWorkflowTypeResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "swf"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "swf"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("swf")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addRegisterWorkflowTypeResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opRegisterWorkflowTypeResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}