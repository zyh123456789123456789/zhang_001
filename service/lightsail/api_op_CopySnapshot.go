// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Copies a manual snapshot of an instance or disk as another manual snapshot, or
// copies an automatic snapshot of an instance or disk as a manual snapshot. This
// operation can also be used to copy a manual or automatic snapshot of an instance
// or a disk from one Amazon Web Services Region to another in Amazon Lightsail.
// When copying a manual snapshot, be sure to define the source region , source
// snapshot name , and target snapshot name parameters. When copying an automatic
// snapshot, be sure to define the source region , source resource name , target
// snapshot name , and either the restore date or the use latest restorable auto
// snapshot parameters.
func (c *Client) CopySnapshot(ctx context.Context, params *CopySnapshotInput, optFns ...func(*Options)) (*CopySnapshotOutput, error) {
	if params == nil {
		params = &CopySnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopySnapshot", params, optFns, c.addOperationCopySnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopySnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopySnapshotInput struct {

	// The Amazon Web Services Region where the source manual or automatic snapshot is
	// located.
	//
	// This member is required.
	SourceRegion types.RegionName

	// The name of the new manual snapshot to be created as a copy.
	//
	// This member is required.
	TargetSnapshotName *string

	// The date of the source automatic snapshot to copy. Use the get auto snapshots
	// operation to identify the dates of the available automatic snapshots.
	// Constraints:
	//   - Must be specified in YYYY-MM-DD format.
	//   - This parameter cannot be defined together with the use latest restorable
	//   auto snapshot parameter. The restore date and use latest restorable auto
	//   snapshot parameters are mutually exclusive.
	//   - Define this parameter only when copying an automatic snapshot as a manual
	//   snapshot. For more information, see the Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-keeping-automatic-snapshots)
	//   .
	RestoreDate *string

	// The name of the source instance or disk from which the source automatic
	// snapshot was created. Constraint:
	//   - Define this parameter only when copying an automatic snapshot as a manual
	//   snapshot. For more information, see the Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-keeping-automatic-snapshots)
	//   .
	SourceResourceName *string

	// The name of the source manual snapshot to copy. Constraint:
	//   - Define this parameter only when copying a manual snapshot as another manual
	//   snapshot.
	SourceSnapshotName *string

	// A Boolean value to indicate whether to use the latest available automatic
	// snapshot of the specified source instance or disk. Constraints:
	//   - This parameter cannot be defined together with the restore date parameter.
	//   The use latest restorable auto snapshot and restore date parameters are
	//   mutually exclusive.
	//   - Define this parameter only when copying an automatic snapshot as a manual
	//   snapshot. For more information, see the Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-keeping-automatic-snapshots)
	//   .
	UseLatestRestorableAutoSnapshot *bool

	noSmithyDocumentSerde
}

type CopySnapshotOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopySnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCopySnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCopySnapshot{}, middleware.After)
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
	if err = addCopySnapshotResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCopySnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopySnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopySnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CopySnapshot",
	}
}

type opCopySnapshotResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCopySnapshotResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCopySnapshotResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "lightsail"
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
				signingName = "lightsail"
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
				v4aScheme.SigningName = aws.String("lightsail")
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

func addCopySnapshotResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCopySnapshotResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
