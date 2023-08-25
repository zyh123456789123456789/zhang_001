// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A context is a variable that contains information about the current state of
// the conversation between a user and Amazon Lex. Context can be set automatically
// by Amazon Lex when an intent is fulfilled, or it can be set at runtime using the
// PutContent , PutText , or PutSession operation.
type ActiveContext struct {

	// The name of the context.
	//
	// This member is required.
	Name *string

	// State variables for the current context. You can use these values as default
	// values for slots in subsequent events.
	//
	// This member is required.
	Parameters map[string]string

	// The length of time or number of turns that a context remains active.
	//
	// This member is required.
	TimeToLive *ActiveContextTimeToLive

	noSmithyDocumentSerde
}

// The length of time or number of turns that a context remains active.
type ActiveContextTimeToLive struct {

	// The number of seconds that the context should be active after it is first sent
	// in a PostContent or PostText response. You can set the value between 5 and
	// 86,400 seconds (24 hours).
	TimeToLiveInSeconds *int32

	// The number of conversation turns that the context should be active. A
	// conversation turn is one PostContent or PostText request and the corresponding
	// response from Amazon Lex.
	TurnsToLive *int32

	noSmithyDocumentSerde
}

// Represents an option to be shown on the client platform (Facebook, Slack, etc.)
type Button struct {

	// Text that is visible to the user on the button.
	//
	// This member is required.
	Text *string

	// The value sent to Amazon Lex when a user chooses the button. For example,
	// consider button text "NYC." When the user chooses the button, the value sent can
	// be "New York City."
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Describes the next action that the bot should take in its interaction with the
// user and provides information about the context in which the action takes place.
// Use the DialogAction data type to set the interaction to a specific state, or
// to return the interaction to a previous state.
type DialogAction struct {

	// The next action that the bot should take in its interaction with the user. The
	// possible values are:
	//   - ConfirmIntent - The next action is asking the user if the intent is complete
	//   and ready to be fulfilled. This is a yes/no question such as "Place the order?"
	//   - Close - Indicates that the there will not be a response from the user. For
	//   example, the statement "Your order has been placed" does not require a response.
	//
	//   - Delegate - The next action is determined by Amazon Lex.
	//   - ElicitIntent - The next action is to determine the intent that the user
	//   wants to fulfill.
	//   - ElicitSlot - The next action is to elicit a slot value from the user.
	//
	// This member is required.
	Type DialogActionType

	// The fulfillment state of the intent. The possible values are:
	//   - Failed - The Lambda function associated with the intent failed to fulfill
	//   the intent.
	//   - Fulfilled - The intent has fulfilled by the Lambda function associated with
	//   the intent.
	//   - ReadyForFulfillment - All of the information necessary for the intent is
	//   present and the intent ready to be fulfilled by the client application.
	FulfillmentState FulfillmentState

	// The name of the intent.
	IntentName *string

	// The message that should be shown to the user. If you don't specify a message,
	// Amazon Lex will use the message configured for the intent.
	Message *string

	//   - PlainText - The message contains plain UTF-8 text.
	//   - CustomPayload - The message is a custom format for the client.
	//   - SSML - The message contains text formatted for voice output.
	//   - Composite - The message contains an escaped JSON object containing one or
	//   more messages. For more information, see Message Groups (https://docs.aws.amazon.com/lex/latest/dg/howitworks-manage-prompts.html)
	//   .
	MessageFormat MessageFormatType

	// The name of the slot that should be elicited from the user.
	SlotToElicit *string

	// Map of the slots that have been gathered and their values.
	Slots map[string]string

	noSmithyDocumentSerde
}

// Represents an option rendered to the user when a prompt is shown. It could be
// an image, a button, a link, or text.
type GenericAttachment struct {

	// The URL of an attachment to the response card.
	AttachmentLinkUrl *string

	// The list of options to show to the user.
	Buttons []Button

	// The URL of an image that is displayed to the user.
	ImageUrl *string

	// The subtitle shown below the title.
	SubTitle *string

	// The title of the option.
	Title *string

	noSmithyDocumentSerde
}

// Provides a score that indicates the confidence that Amazon Lex has that an
// intent is the one that satisfies the user's intent.
type IntentConfidence struct {

	// A score that indicates how confident Amazon Lex is that an intent satisfies the
	// user's intent. Ranges between 0.00 and 1.00. Higher scores indicate higher
	// confidence.
	Score float64

	noSmithyDocumentSerde
}

// Provides information about the state of an intent. You can use this information
// to get the current state of an intent so that you can process the intent, or so
// that you can return the intent to its previous state.
type IntentSummary struct {

	// The next action that the bot should take in its interaction with the user. The
	// possible values are:
	//   - ConfirmIntent - The next action is asking the user if the intent is complete
	//   and ready to be fulfilled. This is a yes/no question such as "Place the order?"
	//   - Close - Indicates that the there will not be a response from the user. For
	//   example, the statement "Your order has been placed" does not require a response.
	//
	//   - ElicitIntent - The next action is to determine the intent that the user
	//   wants to fulfill.
	//   - ElicitSlot - The next action is to elicit a slot value from the user.
	//
	// This member is required.
	DialogActionType DialogActionType

	// A user-defined label that identifies a particular intent. You can use this
	// label to return to a previous intent. Use the checkpointLabelFilter parameter
	// of the GetSessionRequest operation to filter the intents returned by the
	// operation to those with only the specified label.
	CheckpointLabel *string

	// The status of the intent after the user responds to the confirmation prompt. If
	// the user confirms the intent, Amazon Lex sets this field to Confirmed . If the
	// user denies the intent, Amazon Lex sets this value to Denied . The possible
	// values are:
	//   - Confirmed - The user has responded "Yes" to the confirmation prompt,
	//   confirming that the intent is complete and that it is ready to be fulfilled.
	//   - Denied - The user has responded "No" to the confirmation prompt.
	//   - None - The user has never been prompted for confirmation; or, the user was
	//   prompted but did not confirm or deny the prompt.
	ConfirmationStatus ConfirmationStatus

	// The fulfillment state of the intent. The possible values are:
	//   - Failed - The Lambda function associated with the intent failed to fulfill
	//   the intent.
	//   - Fulfilled - The intent has fulfilled by the Lambda function associated with
	//   the intent.
	//   - ReadyForFulfillment - All of the information necessary for the intent is
	//   present and the intent ready to be fulfilled by the client application.
	FulfillmentState FulfillmentState

	// The name of the intent.
	IntentName *string

	// The next slot to elicit from the user. If there is not slot to elicit, the
	// field is blank.
	SlotToElicit *string

	// Map of the slots that have been gathered and their values.
	Slots map[string]string

	noSmithyDocumentSerde
}

// An intent that Amazon Lex suggests satisfies the user's intent. Includes the
// name of the intent, the confidence that Amazon Lex has that the user's intent is
// satisfied, and the slots defined for the intent.
type PredictedIntent struct {

	// The name of the intent that Amazon Lex suggests satisfies the user's intent.
	IntentName *string

	// Indicates how confident Amazon Lex is that an intent satisfies the user's
	// intent.
	NluIntentConfidence *IntentConfidence

	// The slot and slot values associated with the predicted intent.
	Slots map[string]string

	noSmithyDocumentSerde
}

// If you configure a response card when creating your bots, Amazon Lex
// substitutes the session attributes and slot values that are available, and then
// returns it. The response card can also come from a Lambda function (
// dialogCodeHook and fulfillmentActivity on an intent).
type ResponseCard struct {

	// The content type of the response.
	ContentType ContentType

	// An array of attachment objects representing options.
	GenericAttachments []GenericAttachment

	// The version of the response card format.
	Version *string

	noSmithyDocumentSerde
}

// The sentiment expressed in an utterance. When the bot is configured to send
// utterances to Amazon Comprehend for sentiment analysis, this field structure
// contains the result of the analysis.
type SentimentResponse struct {

	// The inferred sentiment that Amazon Comprehend has the highest confidence in.
	SentimentLabel *string

	// The likelihood that the sentiment was correctly inferred.
	SentimentScore *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde