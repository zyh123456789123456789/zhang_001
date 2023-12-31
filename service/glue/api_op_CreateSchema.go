// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new schema set and registers the schema definition. Returns an error
// if the schema set already exists without actually registering the version. When
// the schema set is created, a version checkpoint will be set to the first
// version. Compatibility mode "DISABLED" restricts any additional schema versions
// from being added after the first schema version. For all other compatibility
// modes, validation of compatibility settings will be applied only from the second
// version onwards when the RegisterSchemaVersion API is used. When this API is
// called without a RegistryId , this will create an entry for a "default-registry"
// in the registry database tables, if it is not already present.
func (c *Client) CreateSchema(ctx context.Context, params *CreateSchemaInput, optFns ...func(*Options)) (*CreateSchemaOutput, error) {
	if params == nil {
		params = &CreateSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSchema", params, optFns, c.addOperationCreateSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSchemaInput struct {

	// The data format of the schema definition. Currently AVRO , JSON and PROTOBUF
	// are supported.
	//
	// This member is required.
	DataFormat types.DataFormat

	// Name of the schema to be created of max length of 255, and may only contain
	// letters, numbers, hyphen, underscore, dollar sign, or hash mark. No whitespace.
	//
	// This member is required.
	SchemaName *string

	// The compatibility mode of the schema. The possible values are:
	//   - NONE: No compatibility mode applies. You can use this choice in development
	//   scenarios or if you do not know the compatibility mode that you want to apply to
	//   schemas. Any new version added will be accepted without undergoing a
	//   compatibility check.
	//   - DISABLED: This compatibility choice prevents versioning for a particular
	//   schema. You can use this choice to prevent future versioning of a schema.
	//   - BACKWARD: This compatibility choice is recommended as it allows data
	//   receivers to read both the current and one previous schema version. This means
	//   that for instance, a new schema version cannot drop data fields or change the
	//   type of these fields, so they can't be read by readers using the previous
	//   version.
	//   - BACKWARD_ALL: This compatibility choice allows data receivers to read both
	//   the current and all previous schema versions. You can use this choice when you
	//   need to delete fields or add optional fields, and check compatibility against
	//   all previous schema versions.
	//   - FORWARD: This compatibility choice allows data receivers to read both the
	//   current and one next schema version, but not necessarily later versions. You can
	//   use this choice when you need to add fields or delete optional fields, but only
	//   check compatibility against the last schema version.
	//   - FORWARD_ALL: This compatibility choice allows data receivers to read
	//   written by producers of any new registered schema. You can use this choice when
	//   you need to add fields or delete optional fields, and check compatibility
	//   against all previous schema versions.
	//   - FULL: This compatibility choice allows data receivers to read data written
	//   by producers using the previous or next version of the schema, but not
	//   necessarily earlier or later versions. You can use this choice when you need to
	//   add or remove optional fields, but only check compatibility against the last
	//   schema version.
	//   - FULL_ALL: This compatibility choice allows data receivers to read data
	//   written by producers using all previous schema versions. You can use this choice
	//   when you need to add or remove optional fields, and check compatibility against
	//   all previous schema versions.
	Compatibility types.Compatibility

	// An optional description of the schema. If description is not provided, there
	// will not be any automatic default value for this.
	Description *string

	// This is a wrapper shape to contain the registry identity fields. If this is not
	// provided, the default registry will be used. The ARN format for the same will
	// be: arn:aws:glue:us-east-2::registry/default-registry:random-5-letter-id .
	RegistryId *types.RegistryId

	// The schema definition using the DataFormat setting for SchemaName .
	SchemaDefinition *string

	// Amazon Web Services tags that contain a key value pair and may be searched by
	// console, command line, or API. If specified, follows the Amazon Web Services
	// tags-on-create pattern.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSchemaOutput struct {

	// The schema compatibility mode.
	Compatibility types.Compatibility

	// The data format of the schema definition. Currently AVRO , JSON and PROTOBUF
	// are supported.
	DataFormat types.DataFormat

	// A description of the schema if specified when created.
	Description *string

	// The latest version of the schema associated with the returned schema definition.
	LatestSchemaVersion int64

	// The next version of the schema associated with the returned schema definition.
	NextSchemaVersion int64

	// The Amazon Resource Name (ARN) of the registry.
	RegistryArn *string

	// The name of the registry.
	RegistryName *string

	// The Amazon Resource Name (ARN) of the schema.
	SchemaArn *string

	// The version number of the checkpoint (the last time the compatibility mode was
	// changed).
	SchemaCheckpoint int64

	// The name of the schema.
	SchemaName *string

	// The status of the schema.
	SchemaStatus types.SchemaStatus

	// The unique identifier of the first schema version.
	SchemaVersionId *string

	// The status of the first schema version created.
	SchemaVersionStatus types.SchemaVersionStatus

	// The tags for the schema.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSchema{}, middleware.After)
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
	if err = addCreateSchemaResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateSchema",
	}
}

type opCreateSchemaResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateSchemaResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateSchemaResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "glue"
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
				signingName = "glue"
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
				v4aScheme.SigningName = aws.String("glue")
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

func addCreateSchemaResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateSchemaResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
