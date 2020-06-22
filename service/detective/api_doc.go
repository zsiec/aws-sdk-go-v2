// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package detective provides the client and types for making API
// requests to Amazon Detective.
//
// Detective uses machine learning and purpose-built visualizations to help
// you analyze and investigate security issues across your Amazon Web Services
// (AWS) workloads. Detective automatically extracts time-based events such
// as login attempts, API calls, and network traffic from AWS CloudTrail and
// Amazon Virtual Private Cloud (Amazon VPC) flow logs. It also extracts findings
// detected by Amazon GuardDuty.
//
// The Detective API primarily supports the creation and management of behavior
// graphs. A behavior graph contains the extracted data from a set of member
// accounts, and is created and managed by a master account.
//
// Every behavior graph is specific to a Region. You can only use the API to
// manage graphs that belong to the Region that is associated with the currently
// selected endpoint.
//
// A Detective master account can use the Detective API to do the following:
//
//    * Enable and disable Detective. Enabling Detective creates a new behavior
//    graph.
//
//    * View the list of member accounts in a behavior graph.
//
//    * Add member accounts to a behavior graph.
//
//    * Remove member accounts from a behavior graph.
//
// A member account can use the Detective API to do the following:
//
//    * View the list of behavior graphs that they are invited to.
//
//    * Accept an invitation to contribute to a behavior graph.
//
//    * Decline an invitation to contribute to a behavior graph.
//
//    * Remove their account from a behavior graph.
//
// All API actions are logged as CloudTrail events. See Logging Detective API
// Calls with CloudTrail (https://docs.aws.amazon.com/detective/latest/adminguide/logging-using-cloudtrail.html).
//
// See https://docs.aws.amazon.com/goto/WebAPI/detective-2018-10-26 for more information on this service.
//
// See detective package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/detective/
//
// Using the Client
//
// To use Amazon Detective with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Detective client for more information on
// creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/detective/#New
package detective
