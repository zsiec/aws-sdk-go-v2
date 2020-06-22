// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateApplicationInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether Application Insights can listen to CloudWatch events for
	// the application resources, such as instance terminated, failed deployment,
	// and others.
	CWEMonitorEnabled *bool `type:"boolean"`

	// When set to true, creates opsItems for any problems detected on an application.
	OpsCenterEnabled *bool `type:"boolean"`

	// The SNS topic provided to Application Insights that is associated to the
	// created opsItem. Allows you to receive notifications for updates to the opsItem.
	OpsItemSNSTopicArn *string `min:"20" type:"string"`

	// The name of the resource group.
	//
	// ResourceGroupName is a required field
	ResourceGroupName *string `min:"1" type:"string" required:"true"`

	// List of tags to add to the application. tag key (Key) and an associated tag
	// value (Value). The maximum length of a tag key is 128 characters. The maximum
	// length of a tag value is 256 characters.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateApplicationInput"}
	if s.OpsItemSNSTopicArn != nil && len(*s.OpsItemSNSTopicArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("OpsItemSNSTopicArn", 20))
	}

	if s.ResourceGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceGroupName"))
	}
	if s.ResourceGroupName != nil && len(*s.ResourceGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceGroupName", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateApplicationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the application.
	ApplicationInfo *ApplicationInfo `type:"structure"`
}

// String returns the string representation
func (s CreateApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateApplication = "CreateApplication"

// CreateApplicationRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Adds an application that is created from a resource group.
//
//    // Example sending a request using CreateApplicationRequest.
//    req := client.CreateApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/CreateApplication
func (c *Client) CreateApplicationRequest(input *CreateApplicationInput) CreateApplicationRequest {
	op := &aws.Operation{
		Name:       opCreateApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateApplicationInput{}
	}

	req := c.newRequest(op, input, &CreateApplicationOutput{})
	return CreateApplicationRequest{Request: req, Input: input, Copy: c.CreateApplicationRequest}
}

// CreateApplicationRequest is the request type for the
// CreateApplication API operation.
type CreateApplicationRequest struct {
	*aws.Request
	Input *CreateApplicationInput
	Copy  func(*CreateApplicationInput) CreateApplicationRequest
}

// Send marshals and sends the CreateApplication API request.
func (r CreateApplicationRequest) Send(ctx context.Context) (*CreateApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApplicationResponse{
		CreateApplicationOutput: r.Request.Data.(*CreateApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApplicationResponse is the response type for the
// CreateApplication API operation.
type CreateApplicationResponse struct {
	*CreateApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApplication request.
func (r *CreateApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
