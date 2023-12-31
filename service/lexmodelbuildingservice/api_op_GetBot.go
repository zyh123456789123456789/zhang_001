// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata information for a specific bot. You must provide the bot name
// and the bot version or alias. This operation requires permissions for the
// lex:GetBot action.
func (c *Client) GetBot(ctx context.Context, params *GetBotInput, optFns ...func(*Options)) (*GetBotOutput, error) {
	if params == nil {
		params = &GetBotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBot", params, optFns, c.addOperationGetBotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotInput struct {

	// The name of the bot. The name is case sensitive.
	//
	// This member is required.
	Name *string

	// The version or alias of the bot.
	//
	// This member is required.
	VersionOrAlias *string

	noSmithyDocumentSerde
}

type GetBotOutput struct {

	// The message that Amazon Lex returns when the user elects to end the
	// conversation without completing it. For more information, see PutBot .
	AbortStatement *types.Statement

	// Checksum of the bot used to identify a specific revision of the bot's $LATEST
	// version.
	Checksum *string

	// For each Amazon Lex bot created with the Amazon Lex Model Building Service, you
	// must specify whether your use of Amazon Lex is related to a website, program, or
	// other application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to the Children's Online Privacy Protection Act (COPPA)
	// by specifying true or false in the childDirected field. By specifying true in
	// the childDirected field, you confirm that your use of Amazon Lex is related to
	// a website, program, or other application that is directed or targeted, in whole
	// or in part, to children under age 13 and subject to COPPA. By specifying false
	// in the childDirected field, you confirm that your use of Amazon Lex is not
	// related to a website, program, or other application that is directed or
	// targeted, in whole or in part, to children under age 13 and subject to COPPA.
	// You may not specify a default value for the childDirected field that does not
	// accurately reflect whether your use of Amazon Lex is related to a website,
	// program, or other application that is directed or targeted, in whole or in part,
	// to children under age 13 and subject to COPPA. If your use of Amazon Lex relates
	// to a website, program, or other application that is directed in whole or in
	// part, to children under age 13, you must obtain any required verifiable parental
	// consent under COPPA. For information regarding the use of Amazon Lex in
	// connection with websites, programs, or other applications that are directed or
	// targeted, in whole or in part, to children under age 13, see the Amazon Lex FAQ. (https://aws.amazon.com/lex/faqs#data-security)
	ChildDirected *bool

	// The message Amazon Lex uses when it doesn't understand the user's request. For
	// more information, see PutBot .
	ClarificationPrompt *types.Prompt

	// The date that the bot was created.
	CreatedDate *time.Time

	// A description of the bot.
	Description *string

	// Indicates whether user utterances should be sent to Amazon Comprehend for
	// sentiment analysis.
	DetectSentiment *bool

	// Indicates whether the bot uses accuracy improvements. true indicates that the
	// bot is using the improvements, otherwise, false .
	EnableModelImprovements *bool

	// If status is FAILED , Amazon Lex explains why it failed to build the bot.
	FailureReason *string

	// The maximum time in seconds that Amazon Lex retains the data gathered in a
	// conversation. For more information, see PutBot .
	IdleSessionTTLInSeconds *int32

	// An array of intent objects. For more information, see PutBot .
	Intents []types.Intent

	// The date that the bot was updated. When you create a resource, the creation
	// date and last updated date are the same.
	LastUpdatedDate *time.Time

	// The target locale for the bot.
	Locale types.Locale

	// The name of the bot.
	Name *string

	// The score that determines where Amazon Lex inserts the AMAZON.FallbackIntent ,
	// AMAZON.KendraSearchIntent , or both when returning alternative intents in a
	// PostContent (https://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostContent.html)
	// or PostText (https://docs.aws.amazon.com/lex/latest/dg/API_runtime_PostText.html)
	// response. AMAZON.FallbackIntent is inserted if the confidence score for all
	// intents is below this value. AMAZON.KendraSearchIntent is only inserted if it
	// is configured for the bot.
	NluIntentConfidenceThreshold *float64

	// The status of the bot. When the status is BUILDING Amazon Lex is building the
	// bot for testing and use. If the status of the bot is READY_BASIC_TESTING , you
	// can test the bot using the exact utterances specified in the bot's intents. When
	// the bot is ready for full testing or to run, the status is READY . If there was
	// a problem with building the bot, the status is FAILED and the failureReason
	// field explains why the bot did not build. If the bot was saved but not built,
	// the status is NOT_BUILT .
	Status types.Status

	// The version of the bot. For a new bot, the version is always $LATEST .
	Version *string

	// The Amazon Polly voice ID that Amazon Lex uses for voice interaction with the
	// user. For more information, see PutBot .
	VoiceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBot{}, middleware.After)
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
	if err = addGetBotResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetBotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBot",
	}
}

type opGetBotResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetBotResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetBotResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "lex"
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
				signingName = "lex"
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
				v4aScheme.SigningName = aws.String("lex")
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

func addGetBotResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetBotResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
