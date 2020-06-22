// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type UpdateRuntimeConfigurationInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a fleet to update runtime configuration for. You
	// can use either the fleet ID or ARN value.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// Instructions for launching server processes on each instance in the fleet.
	// Server processes run either a custom game build executable or a Realtime
	// Servers script. The runtime configuration lists the types of server processes
	// to run on an instance and includes the following configuration settings:
	// the server executable or launch script file, launch parameters, and the number
	// of processes to run concurrently on each instance. A CreateFleet request
	// must include a runtime configuration with at least one server process configuration.
	//
	// RuntimeConfiguration is a required field
	RuntimeConfiguration *RuntimeConfiguration `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateRuntimeConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRuntimeConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRuntimeConfigurationInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if s.RuntimeConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuntimeConfiguration"))
	}
	if s.RuntimeConfiguration != nil {
		if err := s.RuntimeConfiguration.Validate(); err != nil {
			invalidParams.AddNested("RuntimeConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type UpdateRuntimeConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The runtime configuration currently in force. If the update was successful,
	// this object matches the one in the request.
	RuntimeConfiguration *RuntimeConfiguration `type:"structure"`
}

// String returns the string representation
func (s UpdateRuntimeConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRuntimeConfiguration = "UpdateRuntimeConfiguration"

// UpdateRuntimeConfigurationRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Updates the current runtime configuration for the specified fleet, which
// tells Amazon GameLift how to launch server processes on instances in the
// fleet. You can update a fleet's runtime configuration at any time after the
// fleet is created; it does not need to be in an ACTIVE status.
//
// To update runtime configuration, specify the fleet ID and provide a RuntimeConfiguration
// object with an updated set of server process configurations.
//
// Each instance in a Amazon GameLift fleet checks regularly for an updated
// runtime configuration and changes how it launches server processes to comply
// with the latest version. Existing server processes are not affected by the
// update; runtime configuration changes are applied gradually as existing processes
// shut down and new processes are launched during Amazon GameLift's normal
// process recycling activity.
//
// Learn more
//
// Setting up GameLift Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
//
// Related operations
//
//    * CreateFleet
//
//    * ListFleets
//
//    * DeleteFleet
//
//    * DescribeFleetAttributes
//
//    * Update fleets: UpdateFleetAttributes UpdateFleetCapacity UpdateFleetPortSettings
//    UpdateRuntimeConfiguration
//
//    * StartFleetActions or StopFleetActions
//
//    // Example sending a request using UpdateRuntimeConfigurationRequest.
//    req := client.UpdateRuntimeConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/UpdateRuntimeConfiguration
func (c *Client) UpdateRuntimeConfigurationRequest(input *UpdateRuntimeConfigurationInput) UpdateRuntimeConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateRuntimeConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRuntimeConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateRuntimeConfigurationOutput{})
	return UpdateRuntimeConfigurationRequest{Request: req, Input: input, Copy: c.UpdateRuntimeConfigurationRequest}
}

// UpdateRuntimeConfigurationRequest is the request type for the
// UpdateRuntimeConfiguration API operation.
type UpdateRuntimeConfigurationRequest struct {
	*aws.Request
	Input *UpdateRuntimeConfigurationInput
	Copy  func(*UpdateRuntimeConfigurationInput) UpdateRuntimeConfigurationRequest
}

// Send marshals and sends the UpdateRuntimeConfiguration API request.
func (r UpdateRuntimeConfigurationRequest) Send(ctx context.Context) (*UpdateRuntimeConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRuntimeConfigurationResponse{
		UpdateRuntimeConfigurationOutput: r.Request.Data.(*UpdateRuntimeConfigurationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRuntimeConfigurationResponse is the response type for the
// UpdateRuntimeConfiguration API operation.
type UpdateRuntimeConfigurationResponse struct {
	*UpdateRuntimeConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRuntimeConfiguration request.
func (r *UpdateRuntimeConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
