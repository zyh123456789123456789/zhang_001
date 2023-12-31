// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the process of verifying an email identity. An identity is an email
// address or domain that you use when you send email. Before you can use an
// identity to send email, you first have to verify it. By verifying an identity,
// you demonstrate that you're the owner of the identity, and that you've given
// Amazon SES API v2 permission to send email from the identity. When you verify an
// email address, Amazon SES sends an email to the address. Your email address is
// verified as soon as you follow the link in the verification email. When you
// verify a domain without specifying the DkimSigningAttributes object, this
// operation provides a set of DKIM tokens. You can convert these tokens into CNAME
// records, which you then add to the DNS configuration for your domain. Your
// domain is verified when Amazon SES detects these records in the DNS
// configuration for your domain. This verification method is known as Easy DKIM (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html)
// . Alternatively, you can perform the verification process by providing your own
// public-private key pair. This verification method is known as Bring Your Own
// DKIM (BYODKIM). To use BYODKIM, your call to the CreateEmailIdentity operation
// has to include the DkimSigningAttributes object. When you specify this object,
// you provide a selector (a component of the DNS record name that identifies the
// public key to use for DKIM authentication) and a private key. When you verify a
// domain, this operation provides a set of DKIM tokens, which you can convert into
// CNAME tokens. You add these CNAME tokens to the DNS configuration for your
// domain. Your domain is verified when Amazon SES detects these records in the DNS
// configuration for your domain. For some DNS providers, it can take 72 hours or
// more to complete the domain verification process. Additionally, you can
// associate an existing configuration set with the email identity that you're
// verifying.
func (c *Client) CreateEmailIdentity(ctx context.Context, params *CreateEmailIdentityInput, optFns ...func(*Options)) (*CreateEmailIdentityOutput, error) {
	if params == nil {
		params = &CreateEmailIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEmailIdentity", params, optFns, c.addOperationCreateEmailIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEmailIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to begin the verification process for an email identity (an email
// address or domain).
type CreateEmailIdentityInput struct {

	// The email address or domain to verify.
	//
	// This member is required.
	EmailIdentity *string

	// The configuration set to use by default when sending from this identity. Note
	// that any configuration set defined in the email sending request takes
	// precedence.
	ConfigurationSetName *string

	// If your request includes this object, Amazon SES configures the identity to use
	// Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures
	// the key length to be used for Easy DKIM (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html)
	// . You can only specify this object if the email identity is a domain, as opposed
	// to an address.
	DkimSigningAttributes *types.DkimSigningAttributes

	// An array of objects that define the tags (keys and values) to associate with
	// the email identity.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// If the email identity is a domain, this object contains information about the
// DKIM verification status for the domain. If the email identity is an email
// address, this object is empty.
type CreateEmailIdentityOutput struct {

	// An object that contains information about the DKIM attributes for the identity.
	DkimAttributes *types.DkimAttributes

	// The email identity type. Note: the MANAGED_DOMAIN identity type is not
	// supported.
	IdentityType types.IdentityType

	// Specifies whether or not the identity is verified. You can only send email from
	// verified email addresses or domains. For more information about verifying
	// identities, see the Amazon Pinpoint User Guide (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html)
	// .
	VerifiedForSendingStatus bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEmailIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEmailIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEmailIdentity{}, middleware.After)
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
	if err = addCreateEmailIdentityResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEmailIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEmailIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEmailIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateEmailIdentity",
	}
}

type opCreateEmailIdentityResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateEmailIdentityResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateEmailIdentityResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ses"
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
				signingName = "ses"
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
				v4aScheme.SigningName = aws.String("ses")
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

func addCreateEmailIdentityResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateEmailIdentityResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
