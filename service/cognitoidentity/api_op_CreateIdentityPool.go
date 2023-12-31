// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new identity pool. The identity pool is a store of user identity
// information that is specific to your AWS account. The keys for
// SupportedLoginProviders are as follows:
//   - Facebook: graph.facebook.com
//   - Google: accounts.google.com
//   - Amazon: www.amazon.com
//   - Twitter: api.twitter.com
//   - Digits: www.digits.com
//
// You must use AWS Developer credentials to call this API.
func (c *Client) CreateIdentityPool(ctx context.Context, params *CreateIdentityPoolInput, optFns ...func(*Options)) (*CreateIdentityPoolOutput, error) {
	if params == nil {
		params = &CreateIdentityPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIdentityPool", params, optFns, c.addOperationCreateIdentityPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIdentityPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the CreateIdentityPool action.
type CreateIdentityPoolInput struct {

	// TRUE if the identity pool supports unauthenticated logins.
	//
	// This member is required.
	AllowUnauthenticatedIdentities bool

	// A string that you provide.
	//
	// This member is required.
	IdentityPoolName *string

	// Enables or disables the Basic (Classic) authentication flow. For more
	// information, see Identity Pools (Federated Identities) Authentication Flow (https://docs.aws.amazon.com/cognito/latest/developerguide/authentication-flow.html)
	// in the Amazon Cognito Developer Guide.
	AllowClassicFlow *bool

	// An array of Amazon Cognito user pools and their client IDs.
	CognitoIdentityProviders []types.CognitoIdentityProvider

	// The "domain" by which Cognito will refer to your users. This name acts as a
	// placeholder that allows your backend and the Cognito service to communicate
	// about the developer provider. For the DeveloperProviderName , you can use
	// letters as well as period ( . ), underscore ( _ ), and dash ( - ). Once you have
	// set a developer provider name, you cannot change it. Please take care in setting
	// this parameter.
	DeveloperProviderName *string

	// Tags to assign to the identity pool. A tag is a label that you can apply to
	// identity pools to categorize and manage them in different ways, such as by
	// purpose, owner, environment, or other criteria.
	IdentityPoolTags map[string]string

	// The Amazon Resource Names (ARN) of the OpenID Connect providers.
	OpenIdConnectProviderARNs []string

	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity
	// pool.
	SamlProviderARNs []string

	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]string

	noSmithyDocumentSerde
}

// An object representing an Amazon Cognito identity pool.
type CreateIdentityPoolOutput struct {

	// TRUE if the identity pool supports unauthenticated logins.
	//
	// This member is required.
	AllowUnauthenticatedIdentities bool

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	// A string that you provide.
	//
	// This member is required.
	IdentityPoolName *string

	// Enables or disables the Basic (Classic) authentication flow. For more
	// information, see Identity Pools (Federated Identities) Authentication Flow (https://docs.aws.amazon.com/cognito/latest/developerguide/authentication-flow.html)
	// in the Amazon Cognito Developer Guide.
	AllowClassicFlow *bool

	// A list representing an Amazon Cognito user pool and its client ID.
	CognitoIdentityProviders []types.CognitoIdentityProvider

	// The "domain" by which Cognito will refer to your users.
	DeveloperProviderName *string

	// The tags that are assigned to the identity pool. A tag is a label that you can
	// apply to identity pools to categorize and manage them in different ways, such as
	// by purpose, owner, environment, or other criteria.
	IdentityPoolTags map[string]string

	// The ARNs of the OpenID Connect providers.
	OpenIdConnectProviderARNs []string

	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity
	// pool.
	SamlProviderARNs []string

	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIdentityPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateIdentityPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateIdentityPool{}, middleware.After)
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
	if err = addCreateIdentityPoolResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateIdentityPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIdentityPool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIdentityPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "CreateIdentityPool",
	}
}

type opCreateIdentityPoolResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateIdentityPoolResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateIdentityPoolResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "cognito-identity"
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
				signingName = "cognito-identity"
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
				v4aScheme.SigningName = aws.String("cognito-identity")
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

func addCreateIdentityPoolResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateIdentityPoolResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
