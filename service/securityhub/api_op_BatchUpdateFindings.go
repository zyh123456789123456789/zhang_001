// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used by Security Hub customers to update information about their investigation
// into a finding. Requested by administrator accounts or member accounts.
// Administrator accounts can update findings for their account and their member
// accounts. Member accounts can update findings for their account. Updates from
// BatchUpdateFindings do not affect the value of UpdatedAt for a finding.
// Administrator and member accounts can use BatchUpdateFindings to update the
// following finding fields and objects.
//   - Confidence
//   - Criticality
//   - Note
//   - RelatedFindings
//   - Severity
//   - Types
//   - UserDefinedFields
//   - VerificationState
//   - Workflow
//
// You can configure IAM policies to restrict access to fields and field values.
// For example, you might not want member accounts to be able to suppress findings
// or change the finding severity. See Configuring access to BatchUpdateFindings (https://docs.aws.amazon.com/securityhub/latest/userguide/finding-update-batchupdatefindings.html#batchupdatefindings-configure-access)
// in the Security Hub User Guide.
func (c *Client) BatchUpdateFindings(ctx context.Context, params *BatchUpdateFindingsInput, optFns ...func(*Options)) (*BatchUpdateFindingsOutput, error) {
	if params == nil {
		params = &BatchUpdateFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateFindings", params, optFns, c.addOperationBatchUpdateFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateFindingsInput struct {

	// The list of findings to update. BatchUpdateFindings can be used to update up to
	// 100 findings at a time. For each finding, the list provides the finding
	// identifier and the ARN of the finding provider.
	//
	// This member is required.
	FindingIdentifiers []types.AwsSecurityFindingIdentifier

	// The updated value for the finding confidence. Confidence is defined as the
	// likelihood that a finding accurately identifies the behavior or issue that it
	// was intended to identify. Confidence is scored on a 0-100 basis using a ratio
	// scale, where 0 means zero percent confidence and 100 means 100 percent
	// confidence.
	Confidence int32

	// The updated value for the level of importance assigned to the resources
	// associated with the findings. A score of 0 means that the underlying resources
	// have no criticality, and a score of 100 is reserved for the most critical
	// resources.
	Criticality int32

	// The updated note.
	Note *types.NoteUpdate

	// A list of findings that are related to the updated findings.
	RelatedFindings []types.RelatedFinding

	// Used to update the finding severity.
	Severity *types.SeverityUpdate

	// One or more finding types in the format of namespace/category/classifier that
	// classify a finding. Valid namespace values are as follows.
	//   - Software and Configuration Checks
	//   - TTPs
	//   - Effects
	//   - Unusual Behaviors
	//   - Sensitive Data Identifications
	Types []string

	// A list of name/value string pairs associated with the finding. These are
	// custom, user-defined fields added to a finding.
	UserDefinedFields map[string]string

	// Indicates the veracity of a finding. The available values for VerificationState
	// are as follows.
	//   - UNKNOWN – The default disposition of a security finding
	//   - TRUE_POSITIVE – The security finding is confirmed
	//   - FALSE_POSITIVE – The security finding was determined to be a false alarm
	//   - BENIGN_POSITIVE – A special case of TRUE_POSITIVE where the finding doesn't
	//   pose any threat, is expected, or both
	VerificationState types.VerificationState

	// Used to update the workflow status of a finding. The workflow status indicates
	// the progress of the investigation into the finding.
	Workflow *types.WorkflowUpdate

	noSmithyDocumentSerde
}

type BatchUpdateFindingsOutput struct {

	// The list of findings that were updated successfully.
	//
	// This member is required.
	ProcessedFindings []types.AwsSecurityFindingIdentifier

	// The list of findings that were not updated.
	//
	// This member is required.
	UnprocessedFindings []types.BatchUpdateFindingsUnprocessedFinding

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateFindings{}, middleware.After)
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
	if err = addBatchUpdateFindingsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpBatchUpdateFindingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateFindings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdateFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchUpdateFindings",
	}
}

type opBatchUpdateFindingsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opBatchUpdateFindingsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opBatchUpdateFindingsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "securityhub"
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
				signingName = "securityhub"
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
				v4aScheme.SigningName = aws.String("securityhub")
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

func addBatchUpdateFindingsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opBatchUpdateFindingsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}