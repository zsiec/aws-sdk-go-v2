// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TagResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html))
	// that is assigned to and uniquely identifies the GameLift resource that you
	// want to assign tags to. GameLift resource ARNs are included in the data object
	// for the resource, which can be retrieved by calling a List or Describe action
	// for the resource type.
	//
	// ResourceARN is a required field
	ResourceARN *string `min:"1" type:"string" required:"true"`

	// A list of one or more tags to assign to the specified GameLift resource.
	// Tags are developer-defined and structured as key-value pairs. The maximum
	// tag limit may be lower than stated. See Tagging AWS Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for actual tagging limits.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s TagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagResourceInput"}

	if s.ResourceARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceARN"))
	}
	if s.ResourceARN != nil && len(*s.ResourceARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceARN", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
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

type TagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagResource = "TagResource"

// TagResourceRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Assigns a tag to a GameLift resource. AWS resource tags provide an additional
// management tool set. You can use tags to organize resources, create IAM permissions
// policies to manage access to groups of resources, customize AWS cost breakdowns,
// etc. This action handles the permissions necessary to manage tags for the
// following GameLift resource types:
//
//    * Build
//
//    * Script
//
//    * Fleet
//
//    * Alias
//
//    * GameSessionQueue
//
//    * MatchmakingConfiguration
//
//    * MatchmakingRuleSet
//
// To add a tag to a resource, specify the unique ARN value for the resource
// and provide a tag list containing one or more tags. The operation succeeds
// even if the list includes tags that are already assigned to the specified
// resource.
//
// Learn more
//
// Tagging AWS Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
// in the AWS General Reference
//
//  AWS Tagging Strategies (http://aws.amazon.com/answers/account-management/aws-tagging-strategies/)
//
// Related operations
//
//    * TagResource
//
//    * UntagResource
//
//    * ListTagsForResource
//
//    // Example sending a request using TagResourceRequest.
//    req := client.TagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/TagResource
func (c *Client) TagResourceRequest(input *TagResourceInput) TagResourceRequest {
	op := &aws.Operation{
		Name:       opTagResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagResourceInput{}
	}

	req := c.newRequest(op, input, &TagResourceOutput{})
	return TagResourceRequest{Request: req, Input: input, Copy: c.TagResourceRequest}
}

// TagResourceRequest is the request type for the
// TagResource API operation.
type TagResourceRequest struct {
	*aws.Request
	Input *TagResourceInput
	Copy  func(*TagResourceInput) TagResourceRequest
}

// Send marshals and sends the TagResource API request.
func (r TagResourceRequest) Send(ctx context.Context) (*TagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagResourceResponse{
		TagResourceOutput: r.Request.Data.(*TagResourceOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagResourceResponse is the response type for the
// TagResource API operation.
type TagResourceResponse struct {
	*TagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagResource request.
func (r *TagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
