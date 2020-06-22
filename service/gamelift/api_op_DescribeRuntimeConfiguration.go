// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type DescribeRuntimeConfigurationInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a fleet to get the runtime configuration for. You
	// can use either the fleet ID or ARN value.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRuntimeConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRuntimeConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRuntimeConfigurationInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type DescribeRuntimeConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Instructions describing how server processes should be launched and maintained
	// on each instance in the fleet.
	RuntimeConfiguration *RuntimeConfiguration `type:"structure"`
}

// String returns the string representation
func (s DescribeRuntimeConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRuntimeConfiguration = "DescribeRuntimeConfiguration"

// DescribeRuntimeConfigurationRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves a fleet's runtime configuration settings. The runtime configuration
// tells Amazon GameLift which server processes to run (and how) on each instance
// in the fleet.
//
// To get a runtime configuration, specify the fleet's unique identifier. If
// successful, a RuntimeConfiguration object is returned for the requested fleet.
// If the requested fleet has been deleted, the result set is empty.
//
// Learn more
//
// Setting up GameLift Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
//
// Running Multiple Processes on a Fleet (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-multiprocess.html)
//
// Related operations
//
//    * CreateFleet
//
//    * ListFleets
//
//    * DeleteFleet
//
//    * Describe fleets: DescribeFleetAttributes DescribeFleetCapacity DescribeFleetPortSettings
//    DescribeFleetUtilization DescribeRuntimeConfiguration DescribeEC2InstanceLimits
//    DescribeFleetEvents
//
//    * UpdateFleetAttributes
//
//    * StartFleetActions or StopFleetActions
//
//    // Example sending a request using DescribeRuntimeConfigurationRequest.
//    req := client.DescribeRuntimeConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeRuntimeConfiguration
func (c *Client) DescribeRuntimeConfigurationRequest(input *DescribeRuntimeConfigurationInput) DescribeRuntimeConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeRuntimeConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRuntimeConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeRuntimeConfigurationOutput{})
	return DescribeRuntimeConfigurationRequest{Request: req, Input: input, Copy: c.DescribeRuntimeConfigurationRequest}
}

// DescribeRuntimeConfigurationRequest is the request type for the
// DescribeRuntimeConfiguration API operation.
type DescribeRuntimeConfigurationRequest struct {
	*aws.Request
	Input *DescribeRuntimeConfigurationInput
	Copy  func(*DescribeRuntimeConfigurationInput) DescribeRuntimeConfigurationRequest
}

// Send marshals and sends the DescribeRuntimeConfiguration API request.
func (r DescribeRuntimeConfigurationRequest) Send(ctx context.Context) (*DescribeRuntimeConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRuntimeConfigurationResponse{
		DescribeRuntimeConfigurationOutput: r.Request.Data.(*DescribeRuntimeConfigurationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRuntimeConfigurationResponse is the response type for the
// DescribeRuntimeConfiguration API operation.
type DescribeRuntimeConfigurationResponse struct {
	*DescribeRuntimeConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRuntimeConfiguration request.
func (r *DescribeRuntimeConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
