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

// Used by deciders to get a DecisionTask from the specified decision taskList . A
// decision task may be returned for any open workflow execution that is using the
// specified task list. The task includes a paginated view of the history of the
// workflow execution. The decider should use the workflow type and the history to
// determine how to properly handle the task. This action initiates a long poll,
// where the service holds the HTTP connection open and responds as soon a task
// becomes available. If no decision task is available in the specified task list
// before the timeout of 60 seconds expires, an empty result is returned. An empty
// result, in this context, means that a DecisionTask is returned, but that the
// value of taskToken is an empty string. Deciders should set their client side
// socket timeout to at least 70 seconds (10 seconds higher than the timeout).
// Because the number of workflow history events for a single workflow execution
// might be very large, the result returned might be split up across a number of
// pages. To retrieve subsequent pages, make additional calls to
// PollForDecisionTask using the nextPageToken returned by the initial call. Note
// that you do not call GetWorkflowExecutionHistory with this nextPageToken .
// Instead, call PollForDecisionTask again. Access Control You can use IAM
// policies to control this action's access to Amazon SWF resources as follows:
//   - Use a Resource element with the domain name to limit the action to only
//     specified domains.
//   - Use an Action element to allow or deny permission to call this action.
//   - Constrain the taskList.name parameter by using a Condition element with the
//     swf:taskList.name key to allow the action to access only certain task lists.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) PollForDecisionTask(ctx context.Context, params *PollForDecisionTaskInput, optFns ...func(*Options)) (*PollForDecisionTaskOutput, error) {
	if params == nil {
		params = &PollForDecisionTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PollForDecisionTask", params, optFns, c.addOperationPollForDecisionTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PollForDecisionTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PollForDecisionTaskInput struct {

	// The name of the domain containing the task lists to poll.
	//
	// This member is required.
	Domain *string

	// Specifies the task list to poll for decision tasks. The specified string must
	// not contain a : (colon), / (slash), | (vertical bar), or any control characters
	// ( \u0000-\u001f | \u007f-\u009f ). Also, it must not be the literal string arn .
	//
	// This member is required.
	TaskList *types.TaskList

	// Identity of the decider making the request, which is recorded in the
	// DecisionTaskStarted event in the workflow history. This enables diagnostic
	// tracing when problems arise. The form of this identity is user defined.
	Identity *string

	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results. This is an upper limit only; the actual number
	// of results returned per call may be fewer than the specified maximum.
	MaximumPageSize int32

	// If NextPageToken is returned there are more results available. The value of
	// NextPageToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return a 400 error: " Specified token has exceeded its
	// maximum lifetime ". The configured maximumPageSize determines how many results
	// can be returned in a single call. The nextPageToken returned by this action
	// cannot be used with GetWorkflowExecutionHistory to get the next page. You must
	// call PollForDecisionTask again (with the nextPageToken ) to retrieve the next
	// page of history records. Calling PollForDecisionTask with a nextPageToken
	// doesn't return a new decision task.
	NextPageToken *string

	// When set to true , returns the events in reverse order. By default the results
	// are returned in ascending order of the eventTimestamp of the events.
	ReverseOrder bool

	// When set to true , returns the events with eventTimestamp greater than or equal
	// to eventTimestamp of the most recent DecisionTaskStarted event. By default,
	// this parameter is set to false .
	StartAtPreviousStartedEvent bool

	noSmithyDocumentSerde
}

// A structure that represents a decision task. Decision tasks are sent to
// deciders in order for them to make decisions.
type PollForDecisionTaskOutput struct {

	// A paginated list of history events of the workflow execution. The decider uses
	// this during the processing of the decision task.
	//
	// This member is required.
	Events []types.HistoryEvent

	// The ID of the DecisionTaskStarted event recorded in the history.
	//
	// This member is required.
	StartedEventId int64

	// The opaque string used as a handle on the task. This token is used by workers
	// to communicate progress and response information back to the system about the
	// task.
	//
	// This member is required.
	TaskToken *string

	// The workflow execution for which this decision task was created.
	//
	// This member is required.
	WorkflowExecution *types.WorkflowExecution

	// The type of the workflow execution for which this decision task was created.
	//
	// This member is required.
	WorkflowType *types.WorkflowType

	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in nextPageToken . Keep all other arguments unchanged. The
	// configured maximumPageSize determines how many results can be returned in a
	// single call.
	NextPageToken *string

	// The ID of the DecisionTaskStarted event of the previous decision task of this
	// workflow execution that was processed by the decider. This can be used to
	// determine the events in the history new since the last decision task received by
	// the decider.
	PreviousStartedEventId int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPollForDecisionTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPollForDecisionTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPollForDecisionTask{}, middleware.After)
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
	if err = addPollForDecisionTaskResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPollForDecisionTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPollForDecisionTask(options.Region), middleware.Before); err != nil {
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

// PollForDecisionTaskAPIClient is a client that implements the
// PollForDecisionTask operation.
type PollForDecisionTaskAPIClient interface {
	PollForDecisionTask(context.Context, *PollForDecisionTaskInput, ...func(*Options)) (*PollForDecisionTaskOutput, error)
}

var _ PollForDecisionTaskAPIClient = (*Client)(nil)

// PollForDecisionTaskPaginatorOptions is the paginator options for
// PollForDecisionTask
type PollForDecisionTaskPaginatorOptions struct {
	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results. This is an upper limit only; the actual number
	// of results returned per call may be fewer than the specified maximum.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// PollForDecisionTaskPaginator is a paginator for PollForDecisionTask
type PollForDecisionTaskPaginator struct {
	options   PollForDecisionTaskPaginatorOptions
	client    PollForDecisionTaskAPIClient
	params    *PollForDecisionTaskInput
	nextToken *string
	firstPage bool
}

// NewPollForDecisionTaskPaginator returns a new PollForDecisionTaskPaginator
func NewPollForDecisionTaskPaginator(client PollForDecisionTaskAPIClient, params *PollForDecisionTaskInput, optFns ...func(*PollForDecisionTaskPaginatorOptions)) *PollForDecisionTaskPaginator {
	if params == nil {
		params = &PollForDecisionTaskInput{}
	}

	options := PollForDecisionTaskPaginatorOptions{}
	if params.MaximumPageSize != 0 {
		options.Limit = params.MaximumPageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &PollForDecisionTaskPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextPageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *PollForDecisionTaskPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next PollForDecisionTask page.
func (p *PollForDecisionTaskPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*PollForDecisionTaskOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextPageToken = p.nextToken

	params.MaximumPageSize = p.options.Limit

	result, err := p.client.PollForDecisionTask(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opPollForDecisionTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "PollForDecisionTask",
	}
}

type opPollForDecisionTaskResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPollForDecisionTaskResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPollForDecisionTaskResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addPollForDecisionTaskResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPollForDecisionTaskResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
