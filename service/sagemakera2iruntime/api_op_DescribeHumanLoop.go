// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakera2iruntime

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeHumanLoopInput struct {
	_ struct{} `type:"structure"`

	// The name of the human loop that you want information about.
	//
	// HumanLoopName is a required field
	HumanLoopName *string `location:"uri" locationName:"HumanLoopName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeHumanLoopInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeHumanLoopInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeHumanLoopInput"}

	if s.HumanLoopName == nil {
		invalidParams.Add(aws.NewErrParamRequired("HumanLoopName"))
	}
	if s.HumanLoopName != nil && len(*s.HumanLoopName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HumanLoopName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeHumanLoopInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.HumanLoopName != nil {
		v := *s.HumanLoopName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "HumanLoopName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeHumanLoopOutput struct {
	_ struct{} `type:"structure"`

	// The creation time when Amazon Augmented AI created the human loop.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" required:"true"`

	// A failure code that identifies the type of failure.
	FailureCode *string `type:"string"`

	// The reason why a human loop failed. The failure reason is returned when the
	// status of the human loop is Failed.
	FailureReason *string `type:"string"`

	// The Amazon Resource Name (ARN) of the flow definition.
	//
	// FlowDefinitionArn is a required field
	FlowDefinitionArn *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the human loop.
	//
	// HumanLoopArn is a required field
	HumanLoopArn *string `type:"string" required:"true"`

	// The name of the human loop. The name must be lowercase, unique within the
	// Region in your account, and can have up to 63 characters. Valid characters:
	// a-z, 0-9, and - (hyphen).
	//
	// HumanLoopName is a required field
	HumanLoopName *string `min:"1" type:"string" required:"true"`

	// An object that contains information about the output of the human loop.
	HumanLoopOutput *HumanLoopOutput `type:"structure"`

	// The status of the human loop.
	//
	// HumanLoopStatus is a required field
	HumanLoopStatus HumanLoopStatus `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DescribeHumanLoopOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeHumanLoopOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.FailureCode != nil {
		v := *s.FailureCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FailureCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FailureReason != nil {
		v := *s.FailureReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FailureReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FlowDefinitionArn != nil {
		v := *s.FlowDefinitionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FlowDefinitionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HumanLoopArn != nil {
		v := *s.HumanLoopArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HumanLoopArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HumanLoopName != nil {
		v := *s.HumanLoopName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HumanLoopName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HumanLoopOutput != nil {
		v := s.HumanLoopOutput

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HumanLoopOutput", v, metadata)
	}
	if len(s.HumanLoopStatus) > 0 {
		v := s.HumanLoopStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HumanLoopStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opDescribeHumanLoop = "DescribeHumanLoop"

// DescribeHumanLoopRequest returns a request value for making API operation for
// Amazon Augmented AI Runtime.
//
// Returns information about the specified human loop.
//
//    // Example sending a request using DescribeHumanLoopRequest.
//    req := client.DescribeHumanLoopRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-a2i-runtime-2019-11-07/DescribeHumanLoop
func (c *Client) DescribeHumanLoopRequest(input *DescribeHumanLoopInput) DescribeHumanLoopRequest {
	op := &aws.Operation{
		Name:       opDescribeHumanLoop,
		HTTPMethod: "GET",
		HTTPPath:   "/human-loops/{HumanLoopName}",
	}

	if input == nil {
		input = &DescribeHumanLoopInput{}
	}

	req := c.newRequest(op, input, &DescribeHumanLoopOutput{})
	return DescribeHumanLoopRequest{Request: req, Input: input, Copy: c.DescribeHumanLoopRequest}
}

// DescribeHumanLoopRequest is the request type for the
// DescribeHumanLoop API operation.
type DescribeHumanLoopRequest struct {
	*aws.Request
	Input *DescribeHumanLoopInput
	Copy  func(*DescribeHumanLoopInput) DescribeHumanLoopRequest
}

// Send marshals and sends the DescribeHumanLoop API request.
func (r DescribeHumanLoopRequest) Send(ctx context.Context) (*DescribeHumanLoopResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeHumanLoopResponse{
		DescribeHumanLoopOutput: r.Request.Data.(*DescribeHumanLoopOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeHumanLoopResponse is the response type for the
// DescribeHumanLoop API operation.
type DescribeHumanLoopResponse struct {
	*DescribeHumanLoopOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeHumanLoop request.
func (r *DescribeHumanLoopResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
