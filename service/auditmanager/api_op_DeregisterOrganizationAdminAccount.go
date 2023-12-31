// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified Amazon Web Services account as a delegated administrator
// for Audit Manager. When you remove a delegated administrator from your Audit
// Manager settings, you continue to have access to the evidence that you
// previously collected under that account. This is also the case when you
// deregister a delegated administrator from Organizations. However, Audit Manager
// stops collecting and attaching evidence to that delegated administrator account
// moving forward. Keep in mind the following cleanup task if you use evidence
// finder: Before you use your management account to remove a delegated
// administrator, make sure that the current delegated administrator account signs
// in to Audit Manager and disables evidence finder first. Disabling evidence
// finder automatically deletes the event data store that was created in their
// account when they enabled evidence finder. If this task isn’t completed, the
// event data store remains in their account. In this case, we recommend that the
// original delegated administrator goes to CloudTrail Lake and manually deletes
// the event data store (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-eds-disable-termination.html)
// . This cleanup task is necessary to ensure that you don't end up with multiple
// event data stores. Audit Manager ignores an unused event data store after you
// remove or change a delegated administrator account. However, the unused event
// data store continues to incur storage costs from CloudTrail Lake if you don't
// delete it. When you deregister a delegated administrator account for Audit
// Manager, the data for that account isn’t deleted. If you want to delete resource
// data for a delegated administrator account, you must perform that task
// separately before you deregister the account. Either, you can do this in the
// Audit Manager console. Or, you can use one of the delete API operations that are
// provided by Audit Manager. To delete your Audit Manager resource data, see the
// following instructions:
//   - DeleteAssessment (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessment.html)
//     (see also: Deleting an assessment (https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-assessment.html)
//     in the Audit Manager User Guide)
//   - DeleteAssessmentFramework (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentFramework.html)
//     (see also: Deleting a custom framework (https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-custom-framework.html)
//     in the Audit Manager User Guide)
//   - DeleteAssessmentFrameworkShare (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentFrameworkShare.html)
//     (see also: Deleting a share request (https://docs.aws.amazon.com/audit-manager/latest/userguide/deleting-shared-framework-requests.html)
//     in the Audit Manager User Guide)
//   - DeleteAssessmentReport (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentReport.html)
//     (see also: Deleting an assessment report (https://docs.aws.amazon.com/audit-manager/latest/userguide/generate-assessment-report.html#delete-assessment-report-steps)
//     in the Audit Manager User Guide)
//   - DeleteControl (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteControl.html)
//     (see also: Deleting a custom control (https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-controls.html)
//     in the Audit Manager User Guide)
//
// At this time, Audit Manager doesn't provide an option to delete evidence for a
// specific delegated administrator. Instead, when your management account
// deregisters Audit Manager, we perform a cleanup for the current delegated
// administrator account at the time of deregistration.
func (c *Client) DeregisterOrganizationAdminAccount(ctx context.Context, params *DeregisterOrganizationAdminAccountInput, optFns ...func(*Options)) (*DeregisterOrganizationAdminAccountOutput, error) {
	if params == nil {
		params = &DeregisterOrganizationAdminAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterOrganizationAdminAccount", params, optFns, c.addOperationDeregisterOrganizationAdminAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterOrganizationAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterOrganizationAdminAccountInput struct {

	// The identifier for the administrator account.
	AdminAccountId *string

	noSmithyDocumentSerde
}

type DeregisterOrganizationAdminAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterOrganizationAdminAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeregisterOrganizationAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeregisterOrganizationAdminAccount{}, middleware.After)
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
	if err = addDeregisterOrganizationAdminAccountResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterOrganizationAdminAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterOrganizationAdminAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "DeregisterOrganizationAdminAccount",
	}
}

type opDeregisterOrganizationAdminAccountResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDeregisterOrganizationAdminAccountResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDeregisterOrganizationAdminAccountResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "auditmanager"
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
				signingName = "auditmanager"
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
				v4aScheme.SigningName = aws.String("auditmanager")
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

func addDeregisterOrganizationAdminAccountResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDeregisterOrganizationAdminAccountResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
