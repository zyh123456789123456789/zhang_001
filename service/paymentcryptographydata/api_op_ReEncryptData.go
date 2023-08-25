// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptographydata

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Re-encrypt ciphertext using DUKPT, Symmetric and Asymmetric Data Encryption
// Keys. You can either generate an encryption key within Amazon Web Services
// Payment Cryptography by calling CreateKey (https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html)
// or import your own encryption key by calling ImportKey (https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html)
// . The KeyArn for use with this operation must be in a compatible key state with
// KeyModesOfUse set to Encrypt . In asymmetric encryption, ciphertext is encrypted
// using public component (imported by calling ImportKey (https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html)
// ) of the asymmetric key pair created outside of Amazon Web Services Payment
// Cryptography. For symmetric and DUKPT encryption, Amazon Web Services Payment
// Cryptography supports TDES and AES algorithms. For asymmetric encryption,
// Amazon Web Services Payment Cryptography supports RSA . To encrypt using DUKPT,
// a DUKPT key must already exist within your account with KeyModesOfUse set to
// DeriveKey or a new DUKPT can be generated by calling CreateKey (https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html)
// . For information about valid keys for this operation, see Understanding key
// attributes (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html)
// and Key types for specific data operations (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html)
// in the Amazon Web Services Payment Cryptography User Guide. Cross-account use:
// This operation can't be used across different Amazon Web Services accounts.
// Related operations:
//   - DecryptData
//   - EncryptData
//   - GetPublicCertificate (https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetPublicKeyCertificate.html)
//   - ImportKey (https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html)
func (c *Client) ReEncryptData(ctx context.Context, params *ReEncryptDataInput, optFns ...func(*Options)) (*ReEncryptDataOutput, error) {
	if params == nil {
		params = &ReEncryptDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReEncryptData", params, optFns, c.addOperationReEncryptDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReEncryptDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReEncryptDataInput struct {

	// Ciphertext to be encrypted. The minimum allowed length is 16 bytes and maximum
	// allowed length is 4096 bytes.
	//
	// This member is required.
	CipherText *string

	// The attributes and values for incoming ciphertext.
	//
	// This member is required.
	IncomingEncryptionAttributes types.ReEncryptionAttributes

	// The keyARN of the encryption key of incoming ciphertext data.
	//
	// This member is required.
	IncomingKeyIdentifier *string

	// The attributes and values for outgoing ciphertext data after encryption by
	// Amazon Web Services Payment Cryptography.
	//
	// This member is required.
	OutgoingEncryptionAttributes types.ReEncryptionAttributes

	// The keyARN of the encryption key of outgoing ciphertext data after encryption
	// by Amazon Web Services Payment Cryptography.
	//
	// This member is required.
	OutgoingKeyIdentifier *string

	noSmithyDocumentSerde
}

type ReEncryptDataOutput struct {

	// The encrypted ciphertext.
	//
	// This member is required.
	CipherText *string

	// The keyARN (Amazon Resource Name) of the encryption key that Amazon Web
	// Services Payment Cryptography uses for plaintext encryption.
	//
	// This member is required.
	KeyArn *string

	// The key check value (KCV) of the encryption key. The KCV is used to check if
	// all parties holding a given key have the same key or to detect that a key has
	// changed. Amazon Web Services Payment Cryptography calculates the KCV by using
	// standard algorithms, typically by encrypting 8 or 16 bytes or "00" or "01" and
	// then truncating the result to the first 3 bytes, or 6 hex digits, of the
	// resulting cryptogram.
	//
	// This member is required.
	KeyCheckValue *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReEncryptDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpReEncryptData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpReEncryptData{}, middleware.After)
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
	if err = addReEncryptDataResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpReEncryptDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReEncryptData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReEncryptData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "payment-cryptography",
		OperationName: "ReEncryptData",
	}
}

type opReEncryptDataResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opReEncryptDataResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opReEncryptDataResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "payment-cryptography"
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
				signingName = "payment-cryptography"
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
				v4aScheme.SigningName = aws.String("payment-cryptography")
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

func addReEncryptDataResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opReEncryptDataResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}