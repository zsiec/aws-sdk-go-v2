// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Requests API Gateway to create a new BasePathMapping resource.
type CreateBasePathMappingInput struct {
	_ struct{} `type:"structure"`

	// The base path name that callers of the API must provide as part of the URL
	// after the domain name. This value must be unique for all of the mappings
	// across a single API. Specify '(none)' if you do not want callers to specify
	// a base path name after the domain name.
	BasePath *string `locationName:"basePath" type:"string"`

	// [Required] The domain name of the BasePathMapping resource to create.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domain_name" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `locationName:"restApiId" type:"string" required:"true"`

	// The name of the API's stage that you want to use for this mapping. Specify
	// '(none)' if you want callers to explicitly specify the stage name after any
	// base path name.
	Stage *string `locationName:"stage" type:"string"`
}

// String returns the string representation
func (s CreateBasePathMappingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBasePathMappingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBasePathMappingInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBasePathMappingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BasePath != nil {
		v := *s.BasePath

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "basePath", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "restApiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Stage != nil {
		v := *s.Stage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domain_name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents the base path that callers of the API must provide as part of
// the URL after the domain name.
//
// A custom domain name plus a BasePathMapping specification identifies a deployed
// RestApi in a given stage of the owner Account.
//
// Use Custom Domain Names (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type CreateBasePathMappingOutput struct {
	_ struct{} `type:"structure"`

	// The base path name that callers of the API must provide as part of the URL
	// after the domain name.
	BasePath *string `locationName:"basePath" type:"string"`

	// The string identifier of the associated RestApi.
	RestApiId *string `locationName:"restApiId" type:"string"`

	// The name of the associated stage.
	Stage *string `locationName:"stage" type:"string"`
}

// String returns the string representation
func (s CreateBasePathMappingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateBasePathMappingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BasePath != nil {
		v := *s.BasePath

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "basePath", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "restApiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Stage != nil {
		v := *s.Stage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateBasePathMapping = "CreateBasePathMapping"

// CreateBasePathMappingRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Creates a new BasePathMapping resource.
//
//    // Example sending a request using CreateBasePathMappingRequest.
//    req := client.CreateBasePathMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateBasePathMappingRequest(input *CreateBasePathMappingInput) CreateBasePathMappingRequest {
	op := &aws.Operation{
		Name:       opCreateBasePathMapping,
		HTTPMethod: "POST",
		HTTPPath:   "/domainnames/{domain_name}/basepathmappings",
	}

	if input == nil {
		input = &CreateBasePathMappingInput{}
	}

	req := c.newRequest(op, input, &CreateBasePathMappingOutput{})
	return CreateBasePathMappingRequest{Request: req, Input: input, Copy: c.CreateBasePathMappingRequest}
}

// CreateBasePathMappingRequest is the request type for the
// CreateBasePathMapping API operation.
type CreateBasePathMappingRequest struct {
	*aws.Request
	Input *CreateBasePathMappingInput
	Copy  func(*CreateBasePathMappingInput) CreateBasePathMappingRequest
}

// Send marshals and sends the CreateBasePathMapping API request.
func (r CreateBasePathMappingRequest) Send(ctx context.Context) (*CreateBasePathMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBasePathMappingResponse{
		CreateBasePathMappingOutput: r.Request.Data.(*CreateBasePathMappingOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBasePathMappingResponse is the response type for the
// CreateBasePathMapping API operation.
type CreateBasePathMappingResponse struct {
	*CreateBasePathMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBasePathMapping request.
func (r *CreateBasePathMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
