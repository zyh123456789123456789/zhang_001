// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Displays the attributes associated with a single Amazon Web Services account.
type AccountAttribute struct {

	// The name of the account attribute.
	//
	// This member is required.
	Name AccountAttributeName

	// The value associated with the account attribute name.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The current resource quotas associated with an Amazon Web Services account.
type AccountLimit struct {

	// The Amazon Web Services set limit for that resource type, in US dollars.
	//
	// This member is required.
	Max int64

	// The name of the attribute to apply the account limit to.
	//
	// This member is required.
	Name AccountLimitName

	// The current amount that has been spent, in US dollars.
	//
	// This member is required.
	Used int64

	noSmithyDocumentSerde
}

// Contains the destination configuration to use when publishing message sending
// events.
type CloudWatchLogsDestination struct {

	// The Amazon Resource Name (ARN) of an Amazon Identity and Access Management
	// (IAM) role that is able to write event data to an Amazon CloudWatch destination.
	//
	// This member is required.
	IamRoleArn *string

	// The name of the Amazon CloudWatch log group that you want to record events in.
	//
	// This member is required.
	LogGroupArn *string

	noSmithyDocumentSerde
}

// The information for configuration sets that meet a specified criteria.
type ConfigurationSetFilter struct {

	// The name of the attribute to filter on.
	//
	// This member is required.
	Name ConfigurationSetFilterName

	// An array values to filter for.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// Information related to a given configuration set in your Amazon Web Services
// account.
type ConfigurationSetInformation struct {

	// The Resource Name (ARN) of the ConfigurationSet.
	//
	// This member is required.
	ConfigurationSetArn *string

	// The name of the ConfigurationSet.
	//
	// This member is required.
	ConfigurationSetName *string

	// The time when the ConfigurationSet was created, in UNIX epoch time (https://www.epochconverter.com/)
	// format.
	//
	// This member is required.
	CreatedTimestamp *time.Time

	// An array of EventDestination objects that describe any events to log and where
	// to log them.
	//
	// This member is required.
	EventDestinations []EventDestination

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	DefaultMessageType MessageType

	// The default sender ID used by the ConfigurationSet.
	DefaultSenderId *string

	noSmithyDocumentSerde
}

// Contains information about an event destination. Event destinations are
// associated with configuration sets, which enable you to publish message sending
// events to Amazon CloudWatch, Amazon Kinesis Data Firehose, or Amazon SNS.
type EventDestination struct {

	// When set to true events will be logged.
	//
	// This member is required.
	Enabled *bool

	// The name of the EventDestination.
	//
	// This member is required.
	EventDestinationName *string

	// An array of event types that determine which events to log.
	//
	// This member is required.
	MatchingEventTypes []EventType

	// An object that contains information about an event destination that sends
	// logging events to Amazon CloudWatch logs.
	CloudWatchLogsDestination *CloudWatchLogsDestination

	// An object that contains information about an event destination for logging to
	// Amazon Kinesis Data Firehose.
	KinesisFirehoseDestination *KinesisFirehoseDestination

	// An object that contains information about an event destination that sends
	// logging events to Amazon SNS.
	SnsDestination *SnsDestination

	noSmithyDocumentSerde
}

// The information for keywords that meet a specified criteria.
type KeywordFilter struct {

	// The name of the attribute to filter on.
	//
	// This member is required.
	Name KeywordFilterName

	// An array values to filter for.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// The information for all keywords in a pool.
type KeywordInformation struct {

	// The keyword as a string.
	//
	// This member is required.
	Keyword *string

	// The action to perform for the keyword.
	//
	// This member is required.
	KeywordAction KeywordAction

	// A custom message that can be used with the keyword.
	//
	// This member is required.
	KeywordMessage *string

	noSmithyDocumentSerde
}

// Contains the delivery stream Amazon Resource Name (ARN), and the ARN of the
// Identity and Access Management (IAM) role associated with an Kinesis Data
// Firehose event destination. Event destinations, such as Kinesis Data Firehose,
// are associated with configuration sets, which enable you to publish message
// sending events.
type KinesisFirehoseDestination struct {

	// The Amazon Resource Name (ARN) of the delivery stream.
	//
	// This member is required.
	DeliveryStreamArn *string

	// The ARN of an Amazon Identity and Access Management (IAM) role that is able to
	// write event data to an Amazon Firehose destination.
	//
	// This member is required.
	IamRoleArn *string

	noSmithyDocumentSerde
}

// The information for opted out numbers that meet a specified criteria.
type OptedOutFilter struct {

	// The name of the attribute to filter on.
	//
	// This member is required.
	Name OptedOutFilterName

	// An array of values to filter for.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// The information for an opted out number in an Amazon Web Services account.
type OptedOutNumberInformation struct {

	// This is set to true if it was the end recipient that opted out.
	//
	// This member is required.
	EndUserOptedOut bool

	// The phone number that is opted out.
	//
	// This member is required.
	OptedOutNumber *string

	// The time that the op tout occurred, in UNIX epoch time (https://www.epochconverter.com/)
	// format.
	//
	// This member is required.
	OptedOutTimestamp *time.Time

	noSmithyDocumentSerde
}

// The information for all OptOutList in an Amazon Web Services account.
type OptOutListInformation struct {

	// The time when the OutOutList was created, in UNIX epoch time (https://www.epochconverter.com/)
	// format.
	//
	// This member is required.
	CreatedTimestamp *time.Time

	// The Amazon Resource Name (ARN) of the OptOutList.
	//
	// This member is required.
	OptOutListArn *string

	// The name of the OptOutList.
	//
	// This member is required.
	OptOutListName *string

	noSmithyDocumentSerde
}

// The metadata for an origination identity associated with a pool.
type OriginationIdentityMetadata struct {

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	//
	// This member is required.
	IsoCountryCode *string

	// Describes if the origination identity can be used for text messages, voice
	// calls or both.
	//
	// This member is required.
	NumberCapabilities []NumberCapability

	// The unique identifier of the origination identity.
	//
	// This member is required.
	OriginationIdentity *string

	// The Amazon Resource Name (ARN) associated with the origination identity.
	//
	// This member is required.
	OriginationIdentityArn *string

	noSmithyDocumentSerde
}

// The information for a phone number that meets a specified criteria.
type PhoneNumberFilter struct {

	// The name of the attribute to filter on.
	//
	// This member is required.
	Name PhoneNumberFilterName

	// An array values to filter for.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// The information for a phone number in an Amazon Web Services account.
type PhoneNumberInformation struct {

	// The time when the phone number was created, in UNIX epoch time (https://www.epochconverter.com/)
	// format.
	//
	// This member is required.
	CreatedTimestamp *time.Time

	// When set to true the phone number can't be deleted.
	//
	// This member is required.
	DeletionProtectionEnabled bool

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	//
	// This member is required.
	IsoCountryCode *string

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	//
	// This member is required.
	MessageType MessageType

	// The price, in US dollars, to lease the phone number.
	//
	// This member is required.
	MonthlyLeasingPrice *string

	// Describes if the origination identity can be used for text messages, voice
	// calls or both.
	//
	// This member is required.
	NumberCapabilities []NumberCapability

	// The type of phone number.
	//
	// This member is required.
	NumberType NumberType

	// The name of the OptOutList associated with the phone number.
	//
	// This member is required.
	OptOutListName *string

	// The phone number in E.164 format.
	//
	// This member is required.
	PhoneNumber *string

	// The Amazon Resource Name (ARN) associated with the phone number.
	//
	// This member is required.
	PhoneNumberArn *string

	// When set to false an end recipient sends a message that begins with HELP or
	// STOP to one of your dedicated numbers, Amazon Pinpoint automatically replies
	// with a customizable message and adds the end recipient to the OptOutList. When
	// set to true you're responsible for responding to HELP and STOP requests. You're
	// also responsible for tracking and honoring opt-out request. For more information
	// see Self-managed opt-outs (https://docs.aws.amazon.com/pinpoint/latest/userguide/settings-sms-managing.html#settings-account-sms-self-managed-opt-out)
	//
	// This member is required.
	SelfManagedOptOutsEnabled bool

	// The current status of the phone number.
	//
	// This member is required.
	Status NumberStatus

	// By default this is set to false. When set to true you can receive incoming text
	// messages from your end recipients using the TwoWayChannelArn.
	//
	// This member is required.
	TwoWayEnabled bool

	// The unique identifier for the phone number.
	PhoneNumberId *string

	// The unique identifier of the pool associated with the phone number.
	PoolId *string

	// The Amazon Resource Name (ARN) of the two way channel.
	TwoWayChannelArn *string

	noSmithyDocumentSerde
}

// The information for a pool that meets a specified criteria.
type PoolFilter struct {

	// The name of the attribute to filter on.
	//
	// This member is required.
	Name PoolFilterName

	// An array values to filter for.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// The information for a pool in an Amazon Web Services account.
type PoolInformation struct {

	// The time when the pool was created, in UNIX epoch time (https://www.epochconverter.com/)
	// format.
	//
	// This member is required.
	CreatedTimestamp *time.Time

	// When set to true the pool can't be deleted.
	//
	// This member is required.
	DeletionProtectionEnabled bool

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	//
	// This member is required.
	MessageType MessageType

	// The name of the OptOutList associated with the pool.
	//
	// This member is required.
	OptOutListName *string

	// The Amazon Resource Name (ARN) for the pool.
	//
	// This member is required.
	PoolArn *string

	// The unique identifier for the pool.
	//
	// This member is required.
	PoolId *string

	// When set to false, an end recipient sends a message that begins with HELP or
	// STOP to one of your dedicated numbers, Amazon Pinpoint automatically replies
	// with a customizable message and adds the end recipient to the OptOutList. When
	// set to true you're responsible for responding to HELP and STOP requests. You're
	// also responsible for tracking and honoring opt-out requests. For more
	// information see Self-managed opt-outs (https://docs.aws.amazon.com/pinpoint/latest/userguide/settings-sms-managing.html#settings-account-sms-self-managed-opt-out)
	//
	// This member is required.
	SelfManagedOptOutsEnabled bool

	// Allows you to enable shared routes on your pool. By default, this is set to
	// False . If you set this value to True , your messages are sent using phone
	// numbers or sender IDs (depending on the country) that are shared with other
	// Amazon Pinpoint users. In some countries, such as the United States, senders
	// aren't allowed to use shared routes and must use a dedicated phone number or
	// short code.
	//
	// This member is required.
	SharedRoutesEnabled bool

	// The current status of the pool.
	//
	// This member is required.
	Status PoolStatus

	// When set to true you can receive incoming text messages from your end
	// recipients using the TwoWayChannelArn.
	//
	// This member is required.
	TwoWayEnabled bool

	// The Amazon Resource Name (ARN) of the two way channel.
	TwoWayChannelArn *string

	noSmithyDocumentSerde
}

// Information about origination identities associated with a pool that meets a
// specified criteria.
type PoolOriginationIdentitiesFilter struct {

	// The name of the attribute to filter on.
	//
	// This member is required.
	Name PoolOriginationIdentitiesFilterName

	// An array values to filter for.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// The alphanumeric sender ID in a specific country that you want to describe. For
// more information on sender IDs see Requesting sender IDs for SMS messaging with
// Amazon Pinpoint  (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-awssupport-sender-id.html)
// in the Amazon Pinpoint User Guide.
type SenderIdAndCountry struct {

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	//
	// This member is required.
	IsoCountryCode *string

	// The unique identifier of the sender.
	//
	// This member is required.
	SenderId *string

	noSmithyDocumentSerde
}

// The information for a sender ID that meets a specified criteria.
type SenderIdFilter struct {

	// The name of the attribute to filter on.
	//
	// This member is required.
	Name SenderIdFilterName

	// An array of values to filter for.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// The information for all SenderIds in an Amazon Web Services account.
type SenderIdInformation struct {

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	//
	// This member is required.
	IsoCountryCode *string

	// The type of message. Valid values are TRANSACTIONAL for messages that are
	// critical or time-sensitive and PROMOTIONAL for messages that aren't critical or
	// time-sensitive.
	//
	// This member is required.
	MessageTypes []MessageType

	// The monthly leasing price, in US dollars.
	//
	// This member is required.
	MonthlyLeasingPrice *string

	// The alphanumeric sender ID in a specific country that you'd like to describe.
	//
	// This member is required.
	SenderId *string

	// The Amazon Resource Name (ARN) associated with the SenderId.
	//
	// This member is required.
	SenderIdArn *string

	noSmithyDocumentSerde
}

// An object that defines an Amazon SNS destination for events. You can use Amazon
// SNS to send notification when certain events occur.
type SnsDestination struct {

	// The Amazon Resource Name (ARN) of the Amazon SNS topic that you want to publish
	// events to.
	//
	// This member is required.
	TopicArn *string

	noSmithyDocumentSerde
}

// Describes the current Amazon Pinpoint monthly spend limits for sending voice
// and text messages. For more information on increasing your monthly spend limit,
// see Requesting increases to your monthly SMS spending quota for Amazon Pinpoint  (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-awssupport-spend-threshold.html)
// in the Amazon Pinpoint User Guide.
type SpendLimit struct {

	// The maximum amount of money, in US dollars, that you want to be able to spend
	// sending messages each month. This value has to be less than or equal to the
	// amount in MaxLimit . To use this custom limit, Overridden must be set to true.
	//
	// This member is required.
	EnforcedLimit int64

	// The maximum amount of money that you are able to spend to send messages each
	// month, in US dollars.
	//
	// This member is required.
	MaxLimit int64

	// The name for the SpendLimit.
	//
	// This member is required.
	Name SpendLimitName

	// When set to True , the value that has been specified in the EnforcedLimit is
	// used to determine the maximum amount in US dollars that can be spent to send
	// messages each month, in US dollars.
	//
	// This member is required.
	Overridden bool

	noSmithyDocumentSerde
}

// The list of tags to be added to the specified topic.
type Tag struct {

	// The key identifier, or name, of the tag.
	//
	// This member is required.
	Key *string

	// The string value associated with the key of the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The field associated with the validation exception.
type ValidationExceptionField struct {

	// The message associated with the validation exception with information to help
	// determine its cause.
	//
	// This member is required.
	Message *string

	// The name of the field.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
