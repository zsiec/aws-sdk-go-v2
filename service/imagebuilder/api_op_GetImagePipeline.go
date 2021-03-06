// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetImagePipelineInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the image pipeline that you want to retrieve.
	//
	// ImagePipelineArn is a required field
	ImagePipelineArn *string `location:"querystring" locationName:"imagePipelineArn" type:"string" required:"true"`
}

// String returns the string representation
func (s GetImagePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetImagePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetImagePipelineInput"}

	if s.ImagePipelineArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImagePipelineArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImagePipelineInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ImagePipelineArn != nil {
		v := *s.ImagePipelineArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "imagePipelineArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetImagePipelineOutput struct {
	_ struct{} `type:"structure"`

	// The image pipeline object.
	ImagePipeline *ImagePipeline `locationName:"imagePipeline" type:"structure"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s GetImagePipelineOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetImagePipelineOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ImagePipeline != nil {
		v := s.ImagePipeline

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "imagePipeline", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetImagePipeline = "GetImagePipeline"

// GetImagePipelineRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Gets an image pipeline.
//
//    // Example sending a request using GetImagePipelineRequest.
//    req := client.GetImagePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/GetImagePipeline
func (c *Client) GetImagePipelineRequest(input *GetImagePipelineInput) GetImagePipelineRequest {
	op := &aws.Operation{
		Name:       opGetImagePipeline,
		HTTPMethod: "GET",
		HTTPPath:   "/GetImagePipeline",
	}

	if input == nil {
		input = &GetImagePipelineInput{}
	}

	req := c.newRequest(op, input, &GetImagePipelineOutput{})
	return GetImagePipelineRequest{Request: req, Input: input, Copy: c.GetImagePipelineRequest}
}

// GetImagePipelineRequest is the request type for the
// GetImagePipeline API operation.
type GetImagePipelineRequest struct {
	*aws.Request
	Input *GetImagePipelineInput
	Copy  func(*GetImagePipelineInput) GetImagePipelineRequest
}

// Send marshals and sends the GetImagePipeline API request.
func (r GetImagePipelineRequest) Send(ctx context.Context) (*GetImagePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetImagePipelineResponse{
		GetImagePipelineOutput: r.Request.Data.(*GetImagePipelineOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetImagePipelineResponse is the response type for the
// GetImagePipeline API operation.
type GetImagePipelineResponse struct {
	*GetImagePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetImagePipeline request.
func (r *GetImagePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
