// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation returns detailed information about a specified domain that is
// associated with the current Amazon Web Services account. Contact information for
// the domain is also returned as part of the output.
func (c *Client) GetDomainDetail(ctx context.Context, params *GetDomainDetailInput, optFns ...func(*Options)) (*GetDomainDetailOutput, error) {
	if params == nil {
		params = &GetDomainDetailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainDetail", params, optFns, c.addOperationGetDomainDetailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainDetailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GetDomainDetail request includes the following element.
type GetDomainDetailInput struct {

	// The name of the domain that you want to get detailed information about.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

// The GetDomainDetail response includes the following elements.
type GetDomainDetailOutput struct {

	// Email address to contact to report incorrect contact information for a domain,
	// to report that the domain is being used to send spam, to report that someone is
	// cybersquatting on a domain name, or report some other type of abuse.
	AbuseContactEmail *string

	// Phone number for reporting abuse.
	AbuseContactPhone *string

	// Provides details about the domain administrative contact.
	AdminContact *types.ContactDetail

	// Specifies whether contact information is concealed from WHOIS queries. If the
	// value is true , WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If the value is false , WHOIS queries
	// return the information that you entered for the admin contact.
	AdminPrivacy *bool

	// Specifies whether the domain registration is set to renew automatically.
	AutoRenew *bool

	// The date when the domain was created as found in the response to a WHOIS query.
	// The date and time is in Unix time format and Coordinated Universal time (UTC).
	CreationDate *time.Time

	// Deprecated.
	DnsSec *string

	// A complex type that contains information about the DNSSEC configuration.
	DnssecKeys []types.DnssecKey

	// The name of a domain.
	DomainName *string

	// The date when the registration for the domain is set to expire. The date and
	// time is in Unix time format and Coordinated Universal time (UTC).
	ExpirationDate *time.Time

	// The name servers of the domain.
	Nameservers []types.Nameserver

	// Provides details about the domain registrant.
	RegistrantContact *types.ContactDetail

	// Specifies whether contact information is concealed from WHOIS queries. If the
	// value is true , WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If the value is false , WHOIS queries
	// return the information that you entered for the registrant contact (domain
	// owner).
	RegistrantPrivacy *bool

	// Name of the registrar of the domain as identified in the registry. Domains with
	// a .com, .net, or .org TLD are registered by Amazon Registrar. All other domains
	// are registered by our registrar associate, Gandi. The value for domains that are
	// registered by Gandi is "GANDI SAS" .
	RegistrarName *string

	// Web address of the registrar.
	RegistrarUrl *string

	// Reserved for future use.
	RegistryDomainId *string

	// Reseller of the domain. Domains registered or transferred using Route 53
	// domains will have "Amazon" as the reseller.
	Reseller *string

	// An array of domain name status codes, also known as Extensible Provisioning
	// Protocol (EPP) status codes. ICANN, the organization that maintains a central
	// database of domain names, has developed a set of domain name status codes that
	// tell you the status of a variety of operations on a domain name, for example,
	// registering a domain name, transferring a domain name to another registrar,
	// renewing the registration for a domain name, and so on. All registrars use this
	// same set of status codes. For a current list of domain name status codes and an
	// explanation of what each code means, go to the ICANN website (https://www.icann.org/)
	// and search for epp status codes . (Search on the ICANN website; web searches
	// sometimes return an old version of the document.)
	StatusList []string

	// Provides details about the domain technical contact.
	TechContact *types.ContactDetail

	// Specifies whether contact information is concealed from WHOIS queries. If the
	// value is true , WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If the value is false , WHOIS queries
	// return the information that you entered for the technical contact.
	TechPrivacy *bool

	// The last updated date of the domain as found in the response to a WHOIS query.
	// The date and time is in Unix time format and Coordinated Universal time (UTC).
	UpdatedDate *time.Time

	// The fully qualified name of the WHOIS server that can answer the WHOIS query
	// for the domain.
	WhoIsServer *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDomainDetailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDomainDetail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDomainDetail{}, middleware.After)
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
	if err = addGetDomainDetailResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetDomainDetailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainDetail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDomainDetail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "GetDomainDetail",
	}
}

type opGetDomainDetailResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetDomainDetailResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetDomainDetailResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "route53domains"
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
				signingName = "route53domains"
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
				v4aScheme.SigningName = aws.String("route53domains")
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

func addGetDomainDetailResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetDomainDetailResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}