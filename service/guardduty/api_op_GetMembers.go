// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetMembersInput struct {
	_ struct{} `type:"structure"`

	// A list of account IDs of the GuardDuty member accounts that you want to describe.
	//
	// AccountIds is a required field
	AccountIds []string `locationName:"accountIds" min:"1" type:"list" required:"true"`

	// The unique ID of the detector of the GuardDuty account whose members you
	// want to retrieve.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMembersInput"}

	if s.AccountIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountIds"))
	}
	if s.AccountIds != nil && len(s.AccountIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountIds", 1))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountIds != nil {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "accountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetMembersOutput struct {
	_ struct{} `type:"structure"`

	// A list of members.
	//
	// Members is a required field
	Members []Member `locationName:"members" type:"list" required:"true"`

	// A list of objects that contain the unprocessed account and a result string
	// that explains why it was unprocessed.
	//
	// UnprocessedAccounts is a required field
	UnprocessedAccounts []UnprocessedAccount `locationName:"unprocessedAccounts" type:"list" required:"true"`
}

// String returns the string representation
func (s GetMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Members != nil {
		v := s.Members

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "members", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "unprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetMembers = "GetMembers"

// GetMembersRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Retrieves GuardDuty member accounts (to the current GuardDuty master account)
// specified by the account IDs.
//
//    // Example sending a request using GetMembersRequest.
//    req := client.GetMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetMembers
func (c *Client) GetMembersRequest(input *GetMembersInput) GetMembersRequest {
	op := &aws.Operation{
		Name:       opGetMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/member/get",
	}

	if input == nil {
		input = &GetMembersInput{}
	}

	req := c.newRequest(op, input, &GetMembersOutput{})
	return GetMembersRequest{Request: req, Input: input, Copy: c.GetMembersRequest}
}

// GetMembersRequest is the request type for the
// GetMembers API operation.
type GetMembersRequest struct {
	*aws.Request
	Input *GetMembersInput
	Copy  func(*GetMembersInput) GetMembersRequest
}

// Send marshals and sends the GetMembers API request.
func (r GetMembersRequest) Send(ctx context.Context) (*GetMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMembersResponse{
		GetMembersOutput: r.Request.Data.(*GetMembersOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMembersResponse is the response type for the
// GetMembers API operation.
type GetMembersResponse struct {
	*GetMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMembers request.
func (r *GetMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
