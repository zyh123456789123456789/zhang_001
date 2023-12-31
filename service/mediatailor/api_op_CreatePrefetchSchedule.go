// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a prefetch schedule for a playback configuration. A prefetch schedule
// allows you to tell MediaTailor to fetch and prepare certain ads before an ad
// break happens. For more information about ad prefetching, see Using ad
// prefetching (https://docs.aws.amazon.com/mediatailor/latest/ug/prefetching-ads.html)
// in the MediaTailor User Guide.
func (c *Client) CreatePrefetchSchedule(ctx context.Context, params *CreatePrefetchScheduleInput, optFns ...func(*Options)) (*CreatePrefetchScheduleOutput, error) {
	if params == nil {
		params = &CreatePrefetchScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePrefetchSchedule", params, optFns, c.addOperationCreatePrefetchScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePrefetchScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePrefetchScheduleInput struct {

	// The configuration settings for MediaTailor's consumption of the prefetched ads
	// from the ad decision server. Each consumption configuration contains an end time
	// and an optional start time that define the consumption window. Prefetch
	// schedules automatically expire no earlier than seven days after the end time.
	//
	// This member is required.
	Consumption *types.PrefetchConsumption

	// The name to assign to the schedule request.
	//
	// This member is required.
	Name *string

	// The name to assign to the playback configuration.
	//
	// This member is required.
	PlaybackConfigurationName *string

	// The configuration settings for retrieval of prefetched ads from the ad decision
	// server. Only one set of prefetched ads will be retrieved and subsequently
	// consumed for each ad break.
	//
	// This member is required.
	Retrieval *types.PrefetchRetrieval

	// An optional stream identifier that MediaTailor uses to prefetch ads for
	// multiple streams that use the same playback configuration. If StreamId is
	// specified, MediaTailor returns all of the prefetch schedules with an exact match
	// on StreamId . If not specified, MediaTailor returns all of the prefetch
	// schedules for the playback configuration, regardless of StreamId .
	StreamId *string

	noSmithyDocumentSerde
}

type CreatePrefetchScheduleOutput struct {

	// The ARN to assign to the prefetch schedule.
	Arn *string

	// The configuration settings for MediaTailor's consumption of the prefetched ads
	// from the ad decision server. Each consumption configuration contains an end time
	// and an optional start time that define the consumption window. Prefetch
	// schedules automatically expire no earlier than seven days after the end time.
	Consumption *types.PrefetchConsumption

	// The name to assign to the prefetch schedule.
	Name *string

	// The name to assign to the playback configuration.
	PlaybackConfigurationName *string

	// The configuration settings for retrieval of prefetched ads from the ad decision
	// server. Only one set of prefetched ads will be retrieved and subsequently
	// consumed for each ad break.
	Retrieval *types.PrefetchRetrieval

	// An optional stream identifier that MediaTailor uses to prefetch ads for
	// multiple streams that use the same playback configuration. If StreamId is
	// specified, MediaTailor returns all of the prefetch schedules with an exact match
	// on StreamId . If not specified, MediaTailor returns all of the prefetch
	// schedules for the playback configuration, regardless of StreamId .
	StreamId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePrefetchScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePrefetchSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePrefetchSchedule{}, middleware.After)
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
	if err = addCreatePrefetchScheduleResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePrefetchScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePrefetchSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePrefetchSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "CreatePrefetchSchedule",
	}
}

type opCreatePrefetchScheduleResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreatePrefetchScheduleResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreatePrefetchScheduleResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "mediatailor"
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
				signingName = "mediatailor"
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
				v4aScheme.SigningName = aws.String("mediatailor")
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

func addCreatePrefetchScheduleResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreatePrefetchScheduleResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
