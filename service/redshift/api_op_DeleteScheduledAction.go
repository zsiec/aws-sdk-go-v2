// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteScheduledActionInput struct {
	_ struct{} `type:"structure"`

	// The name of the scheduled action to delete.
	//
	// ScheduledActionName is a required field
	ScheduledActionName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteScheduledActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteScheduledActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteScheduledActionInput"}

	if s.ScheduledActionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledActionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteScheduledActionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteScheduledActionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteScheduledAction = "DeleteScheduledAction"

// DeleteScheduledActionRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Deletes a scheduled action.
//
//    // Example sending a request using DeleteScheduledActionRequest.
//    req := client.DeleteScheduledActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DeleteScheduledAction
func (c *Client) DeleteScheduledActionRequest(input *DeleteScheduledActionInput) DeleteScheduledActionRequest {
	op := &aws.Operation{
		Name:       opDeleteScheduledAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteScheduledActionInput{}
	}

	req := c.newRequest(op, input, &DeleteScheduledActionOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteScheduledActionRequest{Request: req, Input: input, Copy: c.DeleteScheduledActionRequest}
}

// DeleteScheduledActionRequest is the request type for the
// DeleteScheduledAction API operation.
type DeleteScheduledActionRequest struct {
	*aws.Request
	Input *DeleteScheduledActionInput
	Copy  func(*DeleteScheduledActionInput) DeleteScheduledActionRequest
}

// Send marshals and sends the DeleteScheduledAction API request.
func (r DeleteScheduledActionRequest) Send(ctx context.Context) (*DeleteScheduledActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteScheduledActionResponse{
		DeleteScheduledActionOutput: r.Request.Data.(*DeleteScheduledActionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteScheduledActionResponse is the response type for the
// DeleteScheduledAction API operation.
type DeleteScheduledActionResponse struct {
	*DeleteScheduledActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteScheduledAction request.
func (r *DeleteScheduledActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
