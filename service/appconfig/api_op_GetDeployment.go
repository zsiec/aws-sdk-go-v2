// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetDeploymentInput struct {
	_ struct{} `type:"structure"`

	// The ID of the application that includes the deployment you want to get.
	//
	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"ApplicationId" type:"string" required:"true"`

	// The sequence number of the deployment.
	//
	// DeploymentNumber is a required field
	DeploymentNumber *int64 `location:"uri" locationName:"DeploymentNumber" type:"integer" required:"true"`

	// The ID of the environment that includes the deployment you want to get.
	//
	// EnvironmentId is a required field
	EnvironmentId *string `location:"uri" locationName:"EnvironmentId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeploymentInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.DeploymentNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentNumber"))
	}

	if s.EnvironmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnvironmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeploymentNumber != nil {
		v := *s.DeploymentNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DeploymentNumber", protocol.Int64Value(v), metadata)
	}
	if s.EnvironmentId != nil {
		v := *s.EnvironmentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EnvironmentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the application that was deployed.
	ApplicationId *string `type:"string"`

	// The time the deployment completed.
	CompletedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// Information about the source location of the configuration.
	ConfigurationLocationUri *string `min:"1" type:"string"`

	// The name of the configuration.
	ConfigurationName *string `min:"1" type:"string"`

	// The ID of the configuration profile that was deployed.
	ConfigurationProfileId *string `type:"string"`

	// The configuration version that was deployed.
	ConfigurationVersion *string `min:"1" type:"string"`

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes *int64 `type:"integer"`

	// The sequence number of the deployment.
	DeploymentNumber *int64 `type:"integer"`

	// The ID of the deployment strategy that was deployed.
	DeploymentStrategyId *string `type:"string"`

	// The description of the deployment.
	Description *string `type:"string"`

	// The ID of the environment that was deployed.
	EnvironmentId *string `type:"string"`

	// A list containing all events related to a deployment. The most recent events
	// are displayed first.
	EventLog []DeploymentEvent `type:"list"`

	// The amount of time AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes *int64 `type:"integer"`

	// The percentage of targets to receive a deployed configuration during each
	// interval.
	GrowthFactor *float64 `min:"1" type:"float"`

	// The algorithm used to define how percentage grew over time.
	GrowthType GrowthType `type:"string" enum:"true"`

	// The percentage of targets for which the deployment is available.
	PercentageComplete *float64 `min:"1" type:"float"`

	// The time the deployment started.
	StartedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The state of the deployment.
	State DeploymentState `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CompletedAt != nil {
		v := *s.CompletedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CompletedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.ConfigurationLocationUri != nil {
		v := *s.ConfigurationLocationUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConfigurationLocationUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationName != nil {
		v := *s.ConfigurationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConfigurationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationProfileId != nil {
		v := *s.ConfigurationProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConfigurationProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationVersion != nil {
		v := *s.ConfigurationVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConfigurationVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeploymentDurationInMinutes != nil {
		v := *s.DeploymentDurationInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentDurationInMinutes", protocol.Int64Value(v), metadata)
	}
	if s.DeploymentNumber != nil {
		v := *s.DeploymentNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentNumber", protocol.Int64Value(v), metadata)
	}
	if s.DeploymentStrategyId != nil {
		v := *s.DeploymentStrategyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentStrategyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnvironmentId != nil {
		v := *s.EnvironmentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EnvironmentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EventLog != nil {
		v := s.EventLog

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "EventLog", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.FinalBakeTimeInMinutes != nil {
		v := *s.FinalBakeTimeInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FinalBakeTimeInMinutes", protocol.Int64Value(v), metadata)
	}
	if s.GrowthFactor != nil {
		v := *s.GrowthFactor

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GrowthFactor", protocol.Float64Value(v), metadata)
	}
	if len(s.GrowthType) > 0 {
		v := s.GrowthType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GrowthType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.PercentageComplete != nil {
		v := *s.PercentageComplete

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PercentageComplete", protocol.Float64Value(v), metadata)
	}
	if s.StartedAt != nil {
		v := *s.StartedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StartedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opGetDeployment = "GetDeployment"

// GetDeploymentRequest returns a request value for making API operation for
// Amazon AppConfig.
//
// Retrieve information about a configuration deployment.
//
//    // Example sending a request using GetDeploymentRequest.
//    req := client.GetDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appconfig-2019-10-09/GetDeployment
func (c *Client) GetDeploymentRequest(input *GetDeploymentInput) GetDeploymentRequest {
	op := &aws.Operation{
		Name:       opGetDeployment,
		HTTPMethod: "GET",
		HTTPPath:   "/applications/{ApplicationId}/environments/{EnvironmentId}/deployments/{DeploymentNumber}",
	}

	if input == nil {
		input = &GetDeploymentInput{}
	}

	req := c.newRequest(op, input, &GetDeploymentOutput{})
	return GetDeploymentRequest{Request: req, Input: input, Copy: c.GetDeploymentRequest}
}

// GetDeploymentRequest is the request type for the
// GetDeployment API operation.
type GetDeploymentRequest struct {
	*aws.Request
	Input *GetDeploymentInput
	Copy  func(*GetDeploymentInput) GetDeploymentRequest
}

// Send marshals and sends the GetDeployment API request.
func (r GetDeploymentRequest) Send(ctx context.Context) (*GetDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeploymentResponse{
		GetDeploymentOutput: r.Request.Data.(*GetDeploymentOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeploymentResponse is the response type for the
// GetDeployment API operation.
type GetDeploymentResponse struct {
	*GetDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeployment request.
func (r *GetDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
