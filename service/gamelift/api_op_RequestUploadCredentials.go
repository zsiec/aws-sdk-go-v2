// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type RequestUploadCredentialsInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for a build to get credentials for. You can use either
	// the build ID or ARN value.
	//
	// BuildId is a required field
	BuildId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RequestUploadCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestUploadCredentialsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestUploadCredentialsInput"}

	if s.BuildId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BuildId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type RequestUploadCredentialsOutput struct {
	_ struct{} `type:"structure"`

	// Amazon S3 path and key, identifying where the game build files are stored.
	StorageLocation *S3Location `type:"structure"`

	// AWS credentials required when uploading a game build to the storage location.
	// These credentials have a limited lifespan and are valid only for the build
	// they were issued for.
	UploadCredentials *AwsCredentials `type:"structure" sensitive:"true"`
}

// String returns the string representation
func (s RequestUploadCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}

const opRequestUploadCredentials = "RequestUploadCredentials"

// RequestUploadCredentialsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves a fresh set of credentials for use when uploading a new set of
// game build files to Amazon GameLift's Amazon S3. This is done as part of
// the build creation process; see CreateBuild.
//
// To request new credentials, specify the build ID as returned with an initial
// CreateBuild request. If successful, a new set of credentials are returned,
// along with the S3 storage location associated with the build ID.
//
// Learn more
//
//  Create a Build with Files in S3 (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-cli-uploading.html#gamelift-build-cli-uploading-create-build)
//
// Related operations
//
//    * CreateBuild
//
//    * ListBuilds
//
//    * DescribeBuild
//
//    * UpdateBuild
//
//    * DeleteBuild
//
//    // Example sending a request using RequestUploadCredentialsRequest.
//    req := client.RequestUploadCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/RequestUploadCredentials
func (c *Client) RequestUploadCredentialsRequest(input *RequestUploadCredentialsInput) RequestUploadCredentialsRequest {
	op := &aws.Operation{
		Name:       opRequestUploadCredentials,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RequestUploadCredentialsInput{}
	}

	req := c.newRequest(op, input, &RequestUploadCredentialsOutput{})
	return RequestUploadCredentialsRequest{Request: req, Input: input, Copy: c.RequestUploadCredentialsRequest}
}

// RequestUploadCredentialsRequest is the request type for the
// RequestUploadCredentials API operation.
type RequestUploadCredentialsRequest struct {
	*aws.Request
	Input *RequestUploadCredentialsInput
	Copy  func(*RequestUploadCredentialsInput) RequestUploadCredentialsRequest
}

// Send marshals and sends the RequestUploadCredentials API request.
func (r RequestUploadCredentialsRequest) Send(ctx context.Context) (*RequestUploadCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RequestUploadCredentialsResponse{
		RequestUploadCredentialsOutput: r.Request.Data.(*RequestUploadCredentialsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RequestUploadCredentialsResponse is the response type for the
// RequestUploadCredentials API operation.
type RequestUploadCredentialsResponse struct {
	*RequestUploadCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RequestUploadCredentials request.
func (r *RequestUploadCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
