// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteQualificationTypeInput struct {
	_ struct{} `type:"structure"`

	// The ID of the QualificationType to dispose.
	//
	// QualificationTypeId is a required field
	QualificationTypeId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteQualificationTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteQualificationTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteQualificationTypeInput"}

	if s.QualificationTypeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QualificationTypeId"))
	}
	if s.QualificationTypeId != nil && len(*s.QualificationTypeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QualificationTypeId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteQualificationTypeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteQualificationTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteQualificationType = "DeleteQualificationType"

// DeleteQualificationTypeRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The DeleteQualificationType deletes a Qualification type and deletes any
// HIT types that are associated with the Qualification type.
//
// This operation does not revoke Qualifications already assigned to Workers
// because the Qualifications might be needed for active HITs. If there are
// any pending requests for the Qualification type, Amazon Mechanical Turk rejects
// those requests. After you delete a Qualification type, you can no longer
// use it to create HITs or HIT types.
//
// DeleteQualificationType must wait for all the HITs that use the deleted Qualification
// type to be deleted before completing. It may take up to 48 hours before DeleteQualificationType
// completes and the unique name of the Qualification type is available for
// reuse with CreateQualificationType.
//
//    // Example sending a request using DeleteQualificationTypeRequest.
//    req := client.DeleteQualificationTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/DeleteQualificationType
func (c *Client) DeleteQualificationTypeRequest(input *DeleteQualificationTypeInput) DeleteQualificationTypeRequest {
	op := &aws.Operation{
		Name:       opDeleteQualificationType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteQualificationTypeInput{}
	}

	req := c.newRequest(op, input, &DeleteQualificationTypeOutput{})
	return DeleteQualificationTypeRequest{Request: req, Input: input, Copy: c.DeleteQualificationTypeRequest}
}

// DeleteQualificationTypeRequest is the request type for the
// DeleteQualificationType API operation.
type DeleteQualificationTypeRequest struct {
	*aws.Request
	Input *DeleteQualificationTypeInput
	Copy  func(*DeleteQualificationTypeInput) DeleteQualificationTypeRequest
}

// Send marshals and sends the DeleteQualificationType API request.
func (r DeleteQualificationTypeRequest) Send(ctx context.Context) (*DeleteQualificationTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteQualificationTypeResponse{
		DeleteQualificationTypeOutput: r.Request.Data.(*DeleteQualificationTypeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteQualificationTypeResponse is the response type for the
// DeleteQualificationType API operation.
type DeleteQualificationTypeResponse struct {
	*DeleteQualificationTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteQualificationType request.
func (r *DeleteQualificationTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
