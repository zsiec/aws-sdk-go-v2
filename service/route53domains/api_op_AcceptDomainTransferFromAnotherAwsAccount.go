// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The AcceptDomainTransferFromAnotherAwsAccount request includes the following
// elements.
type AcceptDomainTransferFromAnotherAwsAccountInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that was specified when another AWS account submitted
	// a TransferDomainToAnotherAwsAccount (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html)
	// request.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// The password that was returned by the TransferDomainToAnotherAwsAccount (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html)
	// request.
	//
	// Password is a required field
	Password *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptDomainTransferFromAnotherAwsAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptDomainTransferFromAnotherAwsAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptDomainTransferFromAnotherAwsAccountInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The AcceptDomainTransferFromAnotherAwsAccount response includes the following
// element.
type AcceptDomainTransferFromAnotherAwsAccountOutput struct {
	_ struct{} `type:"structure"`

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	OperationId *string `type:"string"`
}

// String returns the string representation
func (s AcceptDomainTransferFromAnotherAwsAccountOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptDomainTransferFromAnotherAwsAccount = "AcceptDomainTransferFromAnotherAwsAccount"

// AcceptDomainTransferFromAnotherAwsAccountRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// Accepts the transfer of a domain from another AWS account to the current
// AWS account. You initiate a transfer between AWS accounts using TransferDomainToAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html).
//
// Use either ListOperations (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html)
// or GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// to determine whether the operation succeeded. GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
// provides additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled.
//
//    // Example sending a request using AcceptDomainTransferFromAnotherAwsAccountRequest.
//    req := client.AcceptDomainTransferFromAnotherAwsAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/AcceptDomainTransferFromAnotherAwsAccount
func (c *Client) AcceptDomainTransferFromAnotherAwsAccountRequest(input *AcceptDomainTransferFromAnotherAwsAccountInput) AcceptDomainTransferFromAnotherAwsAccountRequest {
	op := &aws.Operation{
		Name:       opAcceptDomainTransferFromAnotherAwsAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptDomainTransferFromAnotherAwsAccountInput{}
	}

	req := c.newRequest(op, input, &AcceptDomainTransferFromAnotherAwsAccountOutput{})
	return AcceptDomainTransferFromAnotherAwsAccountRequest{Request: req, Input: input, Copy: c.AcceptDomainTransferFromAnotherAwsAccountRequest}
}

// AcceptDomainTransferFromAnotherAwsAccountRequest is the request type for the
// AcceptDomainTransferFromAnotherAwsAccount API operation.
type AcceptDomainTransferFromAnotherAwsAccountRequest struct {
	*aws.Request
	Input *AcceptDomainTransferFromAnotherAwsAccountInput
	Copy  func(*AcceptDomainTransferFromAnotherAwsAccountInput) AcceptDomainTransferFromAnotherAwsAccountRequest
}

// Send marshals and sends the AcceptDomainTransferFromAnotherAwsAccount API request.
func (r AcceptDomainTransferFromAnotherAwsAccountRequest) Send(ctx context.Context) (*AcceptDomainTransferFromAnotherAwsAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptDomainTransferFromAnotherAwsAccountResponse{
		AcceptDomainTransferFromAnotherAwsAccountOutput: r.Request.Data.(*AcceptDomainTransferFromAnotherAwsAccountOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptDomainTransferFromAnotherAwsAccountResponse is the response type for the
// AcceptDomainTransferFromAnotherAwsAccount API operation.
type AcceptDomainTransferFromAnotherAwsAccountResponse struct {
	*AcceptDomainTransferFromAnotherAwsAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptDomainTransferFromAnotherAwsAccount request.
func (r *AcceptDomainTransferFromAnotherAwsAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
