// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/UpdateServiceActionInput
type UpdateServiceActionInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// A map that defines the self-service action.
	Definition map[string]string `min:"1" type:"map"`

	// The self-service action description.
	Description *string `type:"string"`

	// The self-service action identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The self-service action name.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateServiceActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServiceActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateServiceActionInput"}
	if s.Definition != nil && len(s.Definition) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Definition", 1))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/UpdateServiceActionOutput
type UpdateServiceActionOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about the self-service action.
	ServiceActionDetail *ServiceActionDetail `type:"structure"`
}

// String returns the string representation
func (s UpdateServiceActionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateServiceAction = "UpdateServiceAction"

// UpdateServiceActionRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Updates a self-service action.
//
//    // Example sending a request using UpdateServiceActionRequest.
//    req := client.UpdateServiceActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/UpdateServiceAction
func (c *Client) UpdateServiceActionRequest(input *UpdateServiceActionInput) UpdateServiceActionRequest {
	op := &aws.Operation{
		Name:       opUpdateServiceAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateServiceActionInput{}
	}

	req := c.newRequest(op, input, &UpdateServiceActionOutput{})
	return UpdateServiceActionRequest{Request: req, Input: input, Copy: c.UpdateServiceActionRequest}
}

// UpdateServiceActionRequest is the request type for the
// UpdateServiceAction API operation.
type UpdateServiceActionRequest struct {
	*aws.Request
	Input *UpdateServiceActionInput
	Copy  func(*UpdateServiceActionInput) UpdateServiceActionRequest
}

// Send marshals and sends the UpdateServiceAction API request.
func (r UpdateServiceActionRequest) Send(ctx context.Context) (*UpdateServiceActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateServiceActionResponse{
		UpdateServiceActionOutput: r.Request.Data.(*UpdateServiceActionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateServiceActionResponse is the response type for the
// UpdateServiceAction API operation.
type UpdateServiceActionResponse struct {
	*UpdateServiceActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateServiceAction request.
func (r *UpdateServiceActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}