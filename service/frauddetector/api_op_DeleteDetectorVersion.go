// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDetectorVersionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the parent detector for the detector version to delete.
	//
	// DetectorId is a required field
	DetectorId *string `locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The ID of the detector version to delete.
	//
	// DetectorVersionId is a required field
	DetectorVersionId *string `locationName:"detectorVersionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDetectorVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDetectorVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDetectorVersionInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.DetectorVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorVersionId"))
	}
	if s.DetectorVersionId != nil && len(*s.DetectorVersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorVersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDetectorVersionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDetectorVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDetectorVersion = "DeleteDetectorVersion"

// DeleteDetectorVersionRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Deletes the detector version. You cannot delete detector versions that are
// in ACTIVE status.
//
//    // Example sending a request using DeleteDetectorVersionRequest.
//    req := client.DeleteDetectorVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/DeleteDetectorVersion
func (c *Client) DeleteDetectorVersionRequest(input *DeleteDetectorVersionInput) DeleteDetectorVersionRequest {
	op := &aws.Operation{
		Name:       opDeleteDetectorVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDetectorVersionInput{}
	}

	req := c.newRequest(op, input, &DeleteDetectorVersionOutput{})
	return DeleteDetectorVersionRequest{Request: req, Input: input, Copy: c.DeleteDetectorVersionRequest}
}

// DeleteDetectorVersionRequest is the request type for the
// DeleteDetectorVersion API operation.
type DeleteDetectorVersionRequest struct {
	*aws.Request
	Input *DeleteDetectorVersionInput
	Copy  func(*DeleteDetectorVersionInput) DeleteDetectorVersionRequest
}

// Send marshals and sends the DeleteDetectorVersion API request.
func (r DeleteDetectorVersionRequest) Send(ctx context.Context) (*DeleteDetectorVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDetectorVersionResponse{
		DeleteDetectorVersionOutput: r.Request.Data.(*DeleteDetectorVersionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDetectorVersionResponse is the response type for the
// DeleteDetectorVersion API operation.
type DeleteDetectorVersionResponse struct {
	*DeleteDetectorVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDetectorVersion request.
func (r *DeleteDetectorVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
