// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type UntagAttendeeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK attendee ID.
	//
	// AttendeeId is a required field
	AttendeeId *string `location:"uri" locationName:"attendeeId" type:"string" required:"true"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`

	// The tag keys.
	//
	// TagKeys is a required field
	TagKeys []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s UntagAttendeeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagAttendeeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagAttendeeInput"}

	if s.AttendeeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttendeeId"))
	}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}
	if s.TagKeys != nil && len(s.TagKeys) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TagKeys", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UntagAttendeeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TagKeys != nil {
		v := s.TagKeys

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TagKeys", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.AttendeeId != nil {
		v := *s.AttendeeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "attendeeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeetingId != nil {
		v := *s.MeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UntagAttendeeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UntagAttendeeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UntagAttendeeOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUntagAttendee = "UntagAttendee"

// UntagAttendeeRequest returns a request value for making API operation for
// Amazon Chime.
//
// Untags the specified tags from the specified Amazon Chime SDK attendee.
//
//    // Example sending a request using UntagAttendeeRequest.
//    req := client.UntagAttendeeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UntagAttendee
func (c *Client) UntagAttendeeRequest(input *UntagAttendeeInput) UntagAttendeeRequest {
	op := &aws.Operation{
		Name:       opUntagAttendee,
		HTTPMethod: "POST",
		HTTPPath:   "/meetings/{meetingId}/attendees/{attendeeId}/tags?operation=delete",
	}

	if input == nil {
		input = &UntagAttendeeInput{}
	}

	req := c.newRequest(op, input, &UntagAttendeeOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UntagAttendeeRequest{Request: req, Input: input, Copy: c.UntagAttendeeRequest}
}

// UntagAttendeeRequest is the request type for the
// UntagAttendee API operation.
type UntagAttendeeRequest struct {
	*aws.Request
	Input *UntagAttendeeInput
	Copy  func(*UntagAttendeeInput) UntagAttendeeRequest
}

// Send marshals and sends the UntagAttendee API request.
func (r UntagAttendeeRequest) Send(ctx context.Context) (*UntagAttendeeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagAttendeeResponse{
		UntagAttendeeOutput: r.Request.Data.(*UntagAttendeeOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagAttendeeResponse is the response type for the
// UntagAttendee API operation.
type UntagAttendeeResponse struct {
	*UntagAttendeeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagAttendee request.
func (r *UntagAttendeeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
