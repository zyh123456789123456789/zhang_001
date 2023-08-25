// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

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

// Resets the specified user's password in a user pool as an administrator. Works
// on any user. This action might generate an SMS text message. Starting June 1,
// 2021, US telecom carriers require you to register an origination phone number
// before you can send SMS messages to US phone numbers. If you use SMS text
// messages in Amazon Cognito, you must register a phone number with Amazon
// Pinpoint (https://console.aws.amazon.com/pinpoint/home/) . Amazon Cognito uses
// the registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in. If you have never used SMS text messages with Amazon Cognito or any
// other Amazon Web Service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In sandbox mode (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html)
// , you can send messages only to verified phone numbers. After you test your app
// while in the sandbox environment, you can move out of the sandbox and into
// production. For more information, see SMS message settings for Amazon Cognito
// user pools (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html)
// in the Amazon Cognito Developer Guide. Deactivates a user's password, requiring
// them to change it. If a user tries to sign in after the API is called, Amazon
// Cognito responds with a PasswordResetRequiredException error. Your app must
// then perform the actions that reset your user's password: the forgot-password
// flow. In addition, if the user pool has phone verification selected and a
// verified phone number exists for the user, or if email verification is selected
// and a verified email exists for the user, calling this API will also result in
// sending a message to the end user with the code to change their password. Amazon
// Cognito evaluates Identity and Access Management (IAM) policies in requests for
// this API operation. For this operation, you must use IAM credentials to
// authorize requests, and you must grant yourself the corresponding IAM permission
// in a policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) AdminResetUserPassword(ctx context.Context, params *AdminResetUserPasswordInput, optFns ...func(*Options)) (*AdminResetUserPasswordOutput, error) {
	if params == nil {
		params = &AdminResetUserPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminResetUserPassword", params, optFns, c.addOperationAdminResetUserPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminResetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to reset a user's password as an administrator.
type AdminResetUserPasswordInput struct {

	// The user pool ID for the user pool where you want to reset the user's password.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user whose password you want to reset.
	//
	// This member is required.
	Username *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// Lambda functions to user pool triggers. When you use the AdminResetUserPassword
	// API action, Amazon Cognito invokes the function that is assigned to the custom
	// message trigger. When Amazon Cognito invokes this function, it passes a JSON
	// payload, which the function receives as input. This payload contains a
	// clientMetadata attribute, which provides the data that you assigned to the
	// ClientMetadata parameter in your AdminResetUserPassword request. In your
	// function code in Lambda, you can process the clientMetadata value to enhance
	// your workflow for your specific needs. For more information, see Customizing
	// user pool Workflows with Lambda Triggers (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. When you use the ClientMetadata
	// parameter, remember that Amazon Cognito won't do the following:
	//   - Store the ClientMetadata value. This data is available only to Lambda
	//   triggers that are assigned to a user pool to support custom workflows. If your
	//   user pool configuration doesn't include triggers, the ClientMetadata parameter
	//   serves no purpose.
	//   - Validate the ClientMetadata value.
	//   - Encrypt the ClientMetadata value. Don't use Amazon Cognito to provide
	//   sensitive information.
	ClientMetadata map[string]string

	noSmithyDocumentSerde
}

// Represents the response from the server to reset a user password as an
// administrator.
type AdminResetUserPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminResetUserPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminResetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminResetUserPassword{}, middleware.After)
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
	if err = addAdminResetUserPasswordResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAdminResetUserPasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminResetUserPassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminResetUserPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminResetUserPassword",
	}
}

type opAdminResetUserPasswordResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opAdminResetUserPasswordResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opAdminResetUserPasswordResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "cognito-idp"
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
				signingName = "cognito-idp"
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
				v4aScheme.SigningName = aws.String("cognito-idp")
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

func addAdminResetUserPasswordResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opAdminResetUserPasswordResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}