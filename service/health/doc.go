// Code generated by smithy-go-codegen DO NOT EDIT.

// Package health provides the API client, operations, and parameter types for AWS
// Health APIs and Notifications.
//
// Health The Health API provides access to the Health information that appears in
// the Health Dashboard (https://health.aws.amazon.com/health/home) . You can use
// the API operations to get information about events that might affect your Amazon
// Web Services and resources. You must have a Business, Enterprise On-Ramp, or
// Enterprise Support plan from Amazon Web Services Support (http://aws.amazon.com/premiumsupport/)
// to use the Health API. If you call the Health API from an Amazon Web Services
// account that doesn't have a Business, Enterprise On-Ramp, or Enterprise Support
// plan, you receive a SubscriptionRequiredException error. For API access, you
// need an access key ID and a secret access key. Use temporary credentials instead
// of long-term access keys when possible. Temporary credentials include an access
// key ID, a secret access key, and a security token that indicates when the
// credentials expire. For more information, see Best practices for managing
// Amazon Web Services access keys (https://docs.aws.amazon.com/general/latest/gr/aws-access-keys-best-practices.html)
// in the Amazon Web Services General Reference. You can use the Health endpoint
// health.us-east-1.amazonaws.com (HTTPS) to call the Health API operations. Health
// supports a multi-Region application architecture and has two regional endpoints
// in an active-passive configuration. You can use the high availability endpoint
// example to determine which Amazon Web Services Region is active, so that you can
// get the latest information from the API. For more information, see Accessing
// the Health API (https://docs.aws.amazon.com/health/latest/ug/health-api.html) in
// the Health User Guide. For authentication of requests, Health uses the
// Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
// . If your Amazon Web Services account is part of Organizations, you can use the
// Health organizational view feature. This feature provides a centralized view of
// Health events across all accounts in your organization. You can aggregate Health
// events in real time to identify accounts in your organization that are affected
// by an operational event or get notified of security vulnerabilities. Use the
// organizational view API operations to enable this feature and return event
// information. For more information, see Aggregating Health events (https://docs.aws.amazon.com/health/latest/ug/aggregate-events.html)
// in the Health User Guide. When you use the Health API operations to return
// Health events, see the following recommendations:
//   - Use the eventScopeCode (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html#AWSHealth-Type-Event-eventScopeCode)
//     parameter to specify whether to return Health events that are public or
//     account-specific.
//   - Use pagination to view all events from the response. For example, if you
//     call the DescribeEventsForOrganization operation to get all events in your
//     organization, you might receive several page results. Specify the nextToken in
//     the next request to return more results.
package health
