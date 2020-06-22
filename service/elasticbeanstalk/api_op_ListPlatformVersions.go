// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListPlatformVersionsInput struct {
	_ struct{} `type:"structure"`

	// Criteria for restricting the resulting list of platform versions. The filter
	// is interpreted as a logical conjunction (AND) of the separate PlatformFilter
	// terms.
	Filters []PlatformFilter `type:"list"`

	// The maximum number of platform version values returned in one call.
	MaxRecords *int64 `min:"1" type:"integer"`

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical
	// to the ones specified in the initial request.
	//
	// If no NextToken is specified, the first page is retrieved.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListPlatformVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPlatformVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPlatformVersionsInput"}
	if s.MaxRecords != nil && *s.MaxRecords < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxRecords", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListPlatformVersionsOutput struct {
	_ struct{} `type:"structure"`

	// In a paginated request, if this value isn't null, it's the token that you
	// can pass in a subsequent request to get the next response page.
	NextToken *string `type:"string"`

	// Summary information about the platform versions.
	PlatformSummaryList []PlatformSummary `type:"list"`
}

// String returns the string representation
func (s ListPlatformVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPlatformVersions = "ListPlatformVersions"

// ListPlatformVersionsRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Lists the platform versions available for your account in an AWS Region.
// Provides summary information about each platform version. Compare to DescribePlatformVersion,
// which provides full details about a single platform version.
//
// For definitions of platform version and other platform-related terms, see
// AWS Elastic Beanstalk Platforms Glossary (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-glossary.html).
//
//    // Example sending a request using ListPlatformVersionsRequest.
//    req := client.ListPlatformVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/ListPlatformVersions
func (c *Client) ListPlatformVersionsRequest(input *ListPlatformVersionsInput) ListPlatformVersionsRequest {
	op := &aws.Operation{
		Name:       opListPlatformVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListPlatformVersionsInput{}
	}

	req := c.newRequest(op, input, &ListPlatformVersionsOutput{})
	return ListPlatformVersionsRequest{Request: req, Input: input, Copy: c.ListPlatformVersionsRequest}
}

// ListPlatformVersionsRequest is the request type for the
// ListPlatformVersions API operation.
type ListPlatformVersionsRequest struct {
	*aws.Request
	Input *ListPlatformVersionsInput
	Copy  func(*ListPlatformVersionsInput) ListPlatformVersionsRequest
}

// Send marshals and sends the ListPlatformVersions API request.
func (r ListPlatformVersionsRequest) Send(ctx context.Context) (*ListPlatformVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPlatformVersionsResponse{
		ListPlatformVersionsOutput: r.Request.Data.(*ListPlatformVersionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListPlatformVersionsResponse is the response type for the
// ListPlatformVersions API operation.
type ListPlatformVersionsResponse struct {
	*ListPlatformVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPlatformVersions request.
func (r *ListPlatformVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
