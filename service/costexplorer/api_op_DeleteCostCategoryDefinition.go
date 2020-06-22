// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteCostCategoryDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for your Cost Category.
	//
	// CostCategoryArn is a required field
	CostCategoryArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCostCategoryDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCostCategoryDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCostCategoryDefinitionInput"}

	if s.CostCategoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CostCategoryArn"))
	}
	if s.CostCategoryArn != nil && len(*s.CostCategoryArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("CostCategoryArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteCostCategoryDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for your Cost Category.
	CostCategoryArn *string `min:"20" type:"string"`

	// The effective end date of the Cost Category as a result of deleting it. No
	// costs after this date will be categorized by the deleted Cost Category.
	EffectiveEnd *string `min:"20" type:"string"`
}

// String returns the string representation
func (s DeleteCostCategoryDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCostCategoryDefinition = "DeleteCostCategoryDefinition"

// DeleteCostCategoryDefinitionRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Deletes a Cost Category. Expenses from this month going forward will no longer
// be categorized with this Cost Category.
//
//    // Example sending a request using DeleteCostCategoryDefinitionRequest.
//    req := client.DeleteCostCategoryDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/DeleteCostCategoryDefinition
func (c *Client) DeleteCostCategoryDefinitionRequest(input *DeleteCostCategoryDefinitionInput) DeleteCostCategoryDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeleteCostCategoryDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCostCategoryDefinitionInput{}
	}

	req := c.newRequest(op, input, &DeleteCostCategoryDefinitionOutput{})
	return DeleteCostCategoryDefinitionRequest{Request: req, Input: input, Copy: c.DeleteCostCategoryDefinitionRequest}
}

// DeleteCostCategoryDefinitionRequest is the request type for the
// DeleteCostCategoryDefinition API operation.
type DeleteCostCategoryDefinitionRequest struct {
	*aws.Request
	Input *DeleteCostCategoryDefinitionInput
	Copy  func(*DeleteCostCategoryDefinitionInput) DeleteCostCategoryDefinitionRequest
}

// Send marshals and sends the DeleteCostCategoryDefinition API request.
func (r DeleteCostCategoryDefinitionRequest) Send(ctx context.Context) (*DeleteCostCategoryDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCostCategoryDefinitionResponse{
		DeleteCostCategoryDefinitionOutput: r.Request.Data.(*DeleteCostCategoryDefinitionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCostCategoryDefinitionResponse is the response type for the
// DeleteCostCategoryDefinition API operation.
type DeleteCostCategoryDefinitionResponse struct {
	*DeleteCostCategoryDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCostCategoryDefinition request.
func (r *DeleteCostCategoryDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
