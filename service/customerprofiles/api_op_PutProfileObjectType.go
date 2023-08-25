// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Defines a ProfileObjectType. To add or remove tags on an existing ObjectType,
// see TagResource (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_TagResource.html)
// / UntagResource (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UntagResource.html)
// .
func (c *Client) PutProfileObjectType(ctx context.Context, params *PutProfileObjectTypeInput, optFns ...func(*Options)) (*PutProfileObjectTypeOutput, error) {
	if params == nil {
		params = &PutProfileObjectTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutProfileObjectType", params, optFns, c.addOperationPutProfileObjectTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutProfileObjectTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutProfileObjectTypeInput struct {

	// Description of the profile object type.
	//
	// This member is required.
	Description *string

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The name of the profile object type.
	//
	// This member is required.
	ObjectTypeName *string

	// Indicates whether a profile should be created when data is received if one
	// doesn’t exist for an object of this type. The default is FALSE . If the
	// AllowProfileCreation flag is set to FALSE , then the service tries to fetch a
	// standard profile and associate this object with the profile. If it is set to
	// TRUE , and if no match is found, then the service creates a new standard profile.
	AllowProfileCreation bool

	// The customer-provided key to encrypt the profile object that will be created in
	// this profile object type.
	EncryptionKey *string

	// The number of days until the data in the object expires.
	ExpirationDays *int32

	// A map of the name and ObjectType field.
	Fields map[string]types.ObjectTypeField

	// A list of unique keys that can be used to map data to the profile.
	Keys map[string][]types.ObjectTypeKey

	// The format of your sourceLastUpdatedTimestamp that was previously set up.
	SourceLastUpdatedTimestampFormat *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// A unique identifier for the object template. For some attributes in the
	// request, the service will use the default value from the object template when
	// TemplateId is present. If these attributes are present in the request, the
	// service may return a BadRequestException . These attributes include:
	// AllowProfileCreation, SourceLastUpdatedTimestampFormat, Fields, and Keys. For
	// example, if AllowProfileCreation is set to true when TemplateId is set, the
	// service may return a BadRequestException .
	TemplateId *string

	noSmithyDocumentSerde
}

type PutProfileObjectTypeOutput struct {

	// Description of the profile object type.
	//
	// This member is required.
	Description *string

	// The name of the profile object type.
	//
	// This member is required.
	ObjectTypeName *string

	// Indicates whether a profile should be created when data is received if one
	// doesn’t exist for an object of this type. The default is FALSE . If the
	// AllowProfileCreation flag is set to FALSE , then the service tries to fetch a
	// standard profile and associate this object with the profile. If it is set to
	// TRUE , and if no match is found, then the service creates a new standard profile.
	AllowProfileCreation bool

	// The timestamp of when the domain was created.
	CreatedAt *time.Time

	// The customer-provided key to encrypt the profile object that will be created in
	// this profile object type.
	EncryptionKey *string

	// The number of days until the data in the object expires.
	ExpirationDays *int32

	// A map of the name and ObjectType field.
	Fields map[string]types.ObjectTypeField

	// A list of unique keys that can be used to map data to the profile.
	Keys map[string][]types.ObjectTypeKey

	// The timestamp of when the domain was most recently edited.
	LastUpdatedAt *time.Time

	// The format of your sourceLastUpdatedTimestamp that was previously set up in
	// fields that were parsed using SimpleDateFormat (https://docs.oracle.com/javase/10/docs/api/java/text/SimpleDateFormat.html)
	// . If you have sourceLastUpdatedTimestamp in your field, you must set up
	// sourceLastUpdatedTimestampFormat .
	SourceLastUpdatedTimestampFormat *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// A unique identifier for the object template.
	TemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutProfileObjectTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutProfileObjectType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutProfileObjectType{}, middleware.After)
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
	if err = addPutProfileObjectTypeResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutProfileObjectTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutProfileObjectType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutProfileObjectType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "PutProfileObjectType",
	}
}

type opPutProfileObjectTypeResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPutProfileObjectTypeResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutProfileObjectTypeResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "profile"
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
				signingName = "profile"
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
				v4aScheme.SigningName = aws.String("profile")
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

func addPutProfileObjectTypeResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutProfileObjectTypeResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
