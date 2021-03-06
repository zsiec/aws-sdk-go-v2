// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type CreateMatchmakingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// A flag that determines whether a match that was created with this configuration
	// must be accepted by the matched players. To require acceptance, set to TRUE.
	//
	// AcceptanceRequired is a required field
	AcceptanceRequired *bool `type:"boolean" required:"true"`

	// The length of time (in seconds) to wait for players to accept a proposed
	// match. If any player rejects the match or fails to accept before the timeout,
	// the ticket continues to look for an acceptable match.
	AcceptanceTimeoutSeconds *int64 `min:"1" type:"integer"`

	// The number of player slots in a match to keep open for future players. For
	// example, assume that the configuration's rule set specifies a match for a
	// single 12-person team. If the additional player count is set to 2, only 10
	// players are initially selected for the match.
	AdditionalPlayerCount *int64 `type:"integer"`

	// The method used to backfill game sessions that are created with this matchmaking
	// configuration. Specify MANUAL when your game manages backfill requests manually
	// or does not use the match backfill feature. Specify AUTOMATIC to have GameLift
	// create a StartMatchBackfill request whenever a game session has one or more
	// open slots. Learn more about manual and automatic backfill in Backfill Existing
	// Games with FlexMatch (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-backfill.html).
	BackfillMode BackfillMode `type:"string" enum:"true"`

	// Information to be added to all events related to this matchmaking configuration.
	CustomEventData *string `type:"string"`

	// A human-readable description of the matchmaking configuration.
	Description *string `min:"1" type:"string"`

	// A set of custom properties for a game session, formatted as key-value pairs.
	// These properties are passed to a game server process in the GameSession object
	// with a request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	// This information is added to the new GameSession object that is created for
	// a successful match.
	GameProperties []GameProperty `type:"list"`

	// A set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process in the GameSession object with
	// a request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	// This information is added to the new GameSession object that is created for
	// a successful match.
	GameSessionData *string `min:"1" type:"string"`

	// Amazon Resource Name (ARN (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html))
	// that is assigned to a GameLift game session queue resource and uniquely identifies
	// it. ARNs are unique across all Regions. These queues are used when placing
	// game sessions for matches that are created with this matchmaking configuration.
	// Queues can be located in any Region.
	//
	// GameSessionQueueArns is a required field
	GameSessionQueueArns []string `type:"list" required:"true"`

	// A unique identifier for a matchmaking configuration. This name is used to
	// identify the configuration associated with a matchmaking request or ticket.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// An SNS topic ARN that is set up to receive matchmaking notifications.
	NotificationTarget *string `type:"string"`

	// The maximum duration, in seconds, that a matchmaking ticket can remain in
	// process before timing out. Requests that fail due to timing out can be resubmitted
	// as needed.
	//
	// RequestTimeoutSeconds is a required field
	RequestTimeoutSeconds *int64 `min:"1" type:"integer" required:"true"`

	// A unique identifier for a matchmaking rule set to use with this configuration.
	// You can use either the rule set name or ARN value. A matchmaking configuration
	// can only use rule sets that are defined in the same Region.
	//
	// RuleSetName is a required field
	RuleSetName *string `min:"1" type:"string" required:"true"`

	// A list of labels to assign to the new matchmaking configuration resource.
	// Tags are developer-defined key-value pairs. Tagging AWS resources are useful
	// for resource management, access management and cost allocation. For more
	// information, see Tagging AWS Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the AWS General Reference. Once the resource is created, you can use TagResource,
	// UntagResource, and ListTagsForResource to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference
	// for actual tagging limits.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateMatchmakingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMatchmakingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMatchmakingConfigurationInput"}

	if s.AcceptanceRequired == nil {
		invalidParams.Add(aws.NewErrParamRequired("AcceptanceRequired"))
	}
	if s.AcceptanceTimeoutSeconds != nil && *s.AcceptanceTimeoutSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("AcceptanceTimeoutSeconds", 1))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.GameSessionData != nil && len(*s.GameSessionData) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionData", 1))
	}

	if s.GameSessionQueueArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameSessionQueueArns"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.RequestTimeoutSeconds == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestTimeoutSeconds"))
	}
	if s.RequestTimeoutSeconds != nil && *s.RequestTimeoutSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("RequestTimeoutSeconds", 1))
	}

	if s.RuleSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleSetName"))
	}
	if s.RuleSetName != nil && len(*s.RuleSetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleSetName", 1))
	}
	if s.GameProperties != nil {
		for i, v := range s.GameProperties {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GameProperties", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type CreateMatchmakingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Object that describes the newly created matchmaking configuration.
	Configuration *MatchmakingConfiguration `type:"structure"`
}

// String returns the string representation
func (s CreateMatchmakingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateMatchmakingConfiguration = "CreateMatchmakingConfiguration"

// CreateMatchmakingConfigurationRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Defines a new matchmaking configuration for use with FlexMatch. A matchmaking
// configuration sets out guidelines for matching players and getting the matches
// into games. You can set up multiple matchmaking configurations to handle
// the scenarios needed for your game. Each matchmaking ticket (StartMatchmaking
// or StartMatchBackfill) specifies a configuration for the match and provides
// player attributes to support the configuration being used.
//
// To create a matchmaking configuration, at a minimum you must specify the
// following: configuration name; a rule set that governs how to evaluate players
// and find acceptable matches; a game session queue to use when placing a new
// game session for the match; and the maximum time allowed for a matchmaking
// attempt.
//
// There are two ways to track the progress of matchmaking tickets: (1) polling
// ticket status with DescribeMatchmaking; or (2) receiving notifications with
// Amazon Simple Notification Service (SNS). To use notifications, you first
// need to set up an SNS topic to receive the notifications, and provide the
// topic ARN in the matchmaking configuration. Since notifications promise only
// "best effort" delivery, we recommend calling DescribeMatchmaking if no notifications
// are received within 30 seconds.
//
// Learn more
//
//  Design a FlexMatch Matchmaker (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-configuration.html)
//
//  Setting up Notifications for Matchmaking (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-notification.html)
//
// Related operations
//
//    * CreateMatchmakingConfiguration
//
//    * DescribeMatchmakingConfigurations
//
//    * UpdateMatchmakingConfiguration
//
//    * DeleteMatchmakingConfiguration
//
//    * CreateMatchmakingRuleSet
//
//    * DescribeMatchmakingRuleSets
//
//    * ValidateMatchmakingRuleSet
//
//    * DeleteMatchmakingRuleSet
//
//    // Example sending a request using CreateMatchmakingConfigurationRequest.
//    req := client.CreateMatchmakingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateMatchmakingConfiguration
func (c *Client) CreateMatchmakingConfigurationRequest(input *CreateMatchmakingConfigurationInput) CreateMatchmakingConfigurationRequest {
	op := &aws.Operation{
		Name:       opCreateMatchmakingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateMatchmakingConfigurationInput{}
	}

	req := c.newRequest(op, input, &CreateMatchmakingConfigurationOutput{})
	return CreateMatchmakingConfigurationRequest{Request: req, Input: input, Copy: c.CreateMatchmakingConfigurationRequest}
}

// CreateMatchmakingConfigurationRequest is the request type for the
// CreateMatchmakingConfiguration API operation.
type CreateMatchmakingConfigurationRequest struct {
	*aws.Request
	Input *CreateMatchmakingConfigurationInput
	Copy  func(*CreateMatchmakingConfigurationInput) CreateMatchmakingConfigurationRequest
}

// Send marshals and sends the CreateMatchmakingConfiguration API request.
func (r CreateMatchmakingConfigurationRequest) Send(ctx context.Context) (*CreateMatchmakingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMatchmakingConfigurationResponse{
		CreateMatchmakingConfigurationOutput: r.Request.Data.(*CreateMatchmakingConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMatchmakingConfigurationResponse is the response type for the
// CreateMatchmakingConfiguration API operation.
type CreateMatchmakingConfigurationResponse struct {
	*CreateMatchmakingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMatchmakingConfiguration request.
func (r *CreateMatchmakingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
