// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows a topic owner to set an attribute of the topic to a new value. To remove
// the ability to change topic permissions, you must deny permissions to the
// AddPermission , RemovePermission , and SetTopicAttributes actions in your IAM
// policy.
func (c *Client) SetTopicAttributes(ctx context.Context, params *SetTopicAttributesInput, optFns ...func(*Options)) (*SetTopicAttributesOutput, error) {
	if params == nil {
		params = &SetTopicAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTopicAttributes", params, optFns, c.addOperationSetTopicAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTopicAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for SetTopicAttributes action.
type SetTopicAttributesInput struct {

	// A map of attributes with their corresponding values. The following lists the
	// names, descriptions, and values of the special request parameters that the
	// SetTopicAttributes action uses:
	//   - ApplicationSuccessFeedbackRoleArn – Indicates failed message delivery status
	//   for an Amazon SNS topic that is subscribed to a platform application endpoint.
	//   - DeliveryPolicy – The policy that defines how Amazon SNS retries failed
	//   deliveries to HTTP/S endpoints.
	//   - DisplayName – The display name to use for a topic with SMS subscriptions.
	//   - Policy – The policy that defines who can access your topic. By default, only
	//   the topic owner can publish or subscribe to the topic.
	//   - TracingConfig – Tracing mode of an Amazon SNS topic. By default
	//   TracingConfig is set to PassThrough , and the topic passes through the tracing
	//   header it receives from an Amazon SNS publisher to its subscriptions. If set to
	//   Active , Amazon SNS will vend X-Ray segment data to topic owner account if the
	//   sampled flag in the tracing header is true. This is only supported on standard
	//   topics.
	//   - HTTP
	//   - HTTPSuccessFeedbackRoleArn – Indicates successful message delivery status
	//   for an Amazon SNS topic that is subscribed to an HTTP endpoint.
	//   - HTTPSuccessFeedbackSampleRate – Indicates percentage of successful messages
	//   to sample for an Amazon SNS topic that is subscribed to an HTTP endpoint.
	//   - HTTPFailureFeedbackRoleArn – Indicates failed message delivery status for an
	//   Amazon SNS topic that is subscribed to an HTTP endpoint.
	//   - Amazon Kinesis Data Firehose
	//   - FirehoseSuccessFeedbackRoleArn – Indicates successful message delivery
	//   status for an Amazon SNS topic that is subscribed to an Amazon Kinesis Data
	//   Firehose endpoint.
	//   - FirehoseSuccessFeedbackSampleRate – Indicates percentage of successful
	//   messages to sample for an Amazon SNS topic that is subscribed to an Amazon
	//   Kinesis Data Firehose endpoint.
	//   - FirehoseFailureFeedbackRoleArn – Indicates failed message delivery status
	//   for an Amazon SNS topic that is subscribed to an Amazon Kinesis Data Firehose
	//   endpoint.
	//   - Lambda
	//   - LambdaSuccessFeedbackRoleArn – Indicates successful message delivery status
	//   for an Amazon SNS topic that is subscribed to an Lambda endpoint.
	//   - LambdaSuccessFeedbackSampleRate – Indicates percentage of successful
	//   messages to sample for an Amazon SNS topic that is subscribed to an Lambda
	//   endpoint.
	//   - LambdaFailureFeedbackRoleArn – Indicates failed message delivery status for
	//   an Amazon SNS topic that is subscribed to an Lambda endpoint.
	//   - Platform application endpoint
	//   - ApplicationSuccessFeedbackRoleArn – Indicates successful message delivery
	//   status for an Amazon SNS topic that is subscribed to an Amazon Web Services
	//   application endpoint.
	//   - ApplicationSuccessFeedbackSampleRate – Indicates percentage of successful
	//   messages to sample for an Amazon SNS topic that is subscribed to an Amazon Web
	//   Services application endpoint.
	//   - ApplicationFailureFeedbackRoleArn – Indicates failed message delivery status
	//   for an Amazon SNS topic that is subscribed to an Amazon Web Services application
	//   endpoint. In addition to being able to configure topic attributes for message
	//   delivery status of notification messages sent to Amazon SNS application
	//   endpoints, you can also configure application attributes for the delivery status
	//   of push notification messages sent to push notification services. For example,
	//   For more information, see Using Amazon SNS Application Attributes for Message
	//   Delivery Status (https://docs.aws.amazon.com/sns/latest/dg/sns-msg-status.html)
	//   .
	//   - Amazon SQS
	//   - SQSSuccessFeedbackRoleArn – Indicates successful message delivery status for
	//   an Amazon SNS topic that is subscribed to an Amazon SQS endpoint.
	//   - SQSSuccessFeedbackSampleRate – Indicates percentage of successful messages
	//   to sample for an Amazon SNS topic that is subscribed to an Amazon SQS endpoint.
	//   - SQSFailureFeedbackRoleArn – Indicates failed message delivery status for an
	//   Amazon SNS topic that is subscribed to an Amazon SQS endpoint.
	// The SuccessFeedbackRoleArn and FailureFeedbackRoleArn attributes are used to
	// give Amazon SNS write access to use CloudWatch Logs on your behalf. The
	// SuccessFeedbackSampleRate attribute is for specifying the sample rate percentage
	// (0-100) of successfully delivered messages. After you configure the
	// FailureFeedbackRoleArn attribute, then all failed message deliveries generate
	// CloudWatch Logs. The following attribute applies only to server-side-encryption (https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html)
	// :
	//   - KmsMasterKeyId – The ID of an Amazon Web Services managed customer master
	//   key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms (https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
	//   . For more examples, see KeyId (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)
	//   in the Key Management Service API Reference.
	//   - SignatureVersion – The signature version corresponds to the hashing
	//   algorithm used while creating the signature of the notifications, subscription
	//   confirmations, or unsubscribe confirmation messages sent by Amazon SNS. By
	//   default, SignatureVersion is set to 1 .
	// The following attribute applies only to FIFO topics (https://docs.aws.amazon.com/sns/latest/dg/sns-fifo-topics.html)
	// :
	//   - ContentBasedDeduplication – Enables content-based deduplication for FIFO
	//   topics.
	//   - By default, ContentBasedDeduplication is set to false . If you create a FIFO
	//   topic and this attribute is false , you must specify a value for the
	//   MessageDeduplicationId parameter for the Publish (https://docs.aws.amazon.com/sns/latest/api/API_Publish.html)
	//   action.
	//   - When you set ContentBasedDeduplication to true , Amazon SNS uses a SHA-256
	//   hash to generate the MessageDeduplicationId using the body of the message (but
	//   not the attributes of the message). (Optional) To override the generated value,
	//   you can specify a value for the MessageDeduplicationId parameter for the
	//   Publish action.
	//
	// This member is required.
	AttributeName *string

	// The ARN of the topic to modify.
	//
	// This member is required.
	TopicArn *string

	// The new value for the attribute.
	AttributeValue *string

	noSmithyDocumentSerde
}

type SetTopicAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetTopicAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetTopicAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetTopicAttributes{}, middleware.After)
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
	if err = addSetTopicAttributesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSetTopicAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetTopicAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetTopicAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "SetTopicAttributes",
	}
}

type opSetTopicAttributesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opSetTopicAttributesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSetTopicAttributesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "sns"
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
				signingName = "sns"
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
				v4aScheme.SigningName = aws.String("sns")
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

func addSetTopicAttributesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opSetTopicAttributesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
