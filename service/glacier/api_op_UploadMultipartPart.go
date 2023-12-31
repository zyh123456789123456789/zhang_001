// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// This operation uploads a part of an archive. You can upload archive parts in
// any order. You can also upload them in parallel. You can upload up to 10,000
// parts for a multipart upload. Amazon Glacier rejects your upload part request if
// any of the following conditions is true:
//   - SHA256 tree hash does not matchTo ensure that part data is not corrupted in
//     transmission, you compute a SHA256 tree hash of the part and include it in your
//     request. Upon receiving the part data, Amazon S3 Glacier also computes a SHA256
//     tree hash. If these hash values don't match, the operation fails. For
//     information about computing a SHA256 tree hash, see Computing Checksums (https://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html)
//     .
//   - Part size does not matchThe size of each part except the last must match
//     the size specified in the corresponding InitiateMultipartUpload request. The
//     size of the last part must be the same size as, or smaller than, the specified
//     size. If you upload a part whose size is smaller than the part size you
//     specified in your initiate multipart upload request and that part is not the
//     last part, then the upload part request will succeed. However, the subsequent
//     Complete Multipart Upload request will fail.
//   - Range does not alignThe byte range value in the request does not align with
//     the part size specified in the corresponding initiate request. For example, if
//     you specify a part size of 4194304 bytes (4 MB), then 0 to 4194303 bytes (4 MB -
//     1) and 4194304 (4 MB) to 8388607 (8 MB - 1) are valid part ranges. However, if
//     you set a range value of 2 MB to 6 MB, the range does not align with the part
//     size and the upload will fail.
//
// This operation is idempotent. If you upload the same part multiple times, the
// data included in the most recent request overwrites the previously uploaded
// data. An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see Access Control Using AWS Identity
// and Access Management (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html)
// . For conceptual information and underlying REST API, see Uploading Large
// Archives in Parts (Multipart Upload) (https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-archive-mpu.html)
// and Upload Part  (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-upload-part.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) UploadMultipartPart(ctx context.Context, params *UploadMultipartPartInput, optFns ...func(*Options)) (*UploadMultipartPartOutput, error) {
	if params == nil {
		params = &UploadMultipartPartInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadMultipartPart", params, optFns, c.addOperationUploadMultipartPartMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadMultipartPartOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to upload a part of an archive in a multipart upload operation.
type UploadMultipartPartInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single ' - ' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The upload ID of the multipart upload.
	//
	// This member is required.
	UploadId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The data to upload.
	Body io.Reader

	// The SHA256 tree hash of the data being uploaded.
	Checksum *string

	// Identifies the range of bytes in the assembled archive that will be uploaded in
	// this part. Amazon S3 Glacier uses this information to assemble the archive in
	// the proper sequence. The format of this header follows RFC 2616. An example
	// header is Content-Range:bytes 0-4194303/*.
	Range *string

	noSmithyDocumentSerde
}

// Contains the Amazon S3 Glacier response to your request.
type UploadMultipartPartOutput struct {

	// The SHA256 tree hash that Amazon S3 Glacier computed for the uploaded part.
	Checksum *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUploadMultipartPartMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUploadMultipartPart{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUploadMultipartPart{}, middleware.After)
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
	if err = addUploadMultipartPartResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUploadMultipartPartValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUploadMultipartPart(options.Region), middleware.Before); err != nil {
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
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
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

func newServiceMetadataMiddleware_opUploadMultipartPart(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "UploadMultipartPart",
	}
}

type opUploadMultipartPartResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUploadMultipartPartResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUploadMultipartPartResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "glacier"
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
				signingName = "glacier"
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
				v4aScheme.SigningName = aws.String("glacier")
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

func addUploadMultipartPartResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUploadMultipartPartResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
