// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codeguruprofileriface provides an interface to enable mocking the Amazon CodeGuru Profiler service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codeguruprofileriface

import (
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler"
)

// ClientAPI provides an interface to enable mocking the
// codeguruprofiler.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CodeGuru Profiler.
//    func myFunc(svc codeguruprofileriface.ClientAPI) bool {
//        // Make svc.ConfigureAgent request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := codeguruprofiler.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        codeguruprofileriface.ClientPI
//    }
//    func (m *mockClientClient) ConfigureAgent(input *codeguruprofiler.ConfigureAgentInput) (*codeguruprofiler.ConfigureAgentOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	ConfigureAgentRequest(*codeguruprofiler.ConfigureAgentInput) codeguruprofiler.ConfigureAgentRequest

	CreateProfilingGroupRequest(*codeguruprofiler.CreateProfilingGroupInput) codeguruprofiler.CreateProfilingGroupRequest

	DeleteProfilingGroupRequest(*codeguruprofiler.DeleteProfilingGroupInput) codeguruprofiler.DeleteProfilingGroupRequest

	DescribeProfilingGroupRequest(*codeguruprofiler.DescribeProfilingGroupInput) codeguruprofiler.DescribeProfilingGroupRequest

	GetPolicyRequest(*codeguruprofiler.GetPolicyInput) codeguruprofiler.GetPolicyRequest

	GetProfileRequest(*codeguruprofiler.GetProfileInput) codeguruprofiler.GetProfileRequest

	ListProfileTimesRequest(*codeguruprofiler.ListProfileTimesInput) codeguruprofiler.ListProfileTimesRequest

	ListProfilingGroupsRequest(*codeguruprofiler.ListProfilingGroupsInput) codeguruprofiler.ListProfilingGroupsRequest

	PostAgentProfileRequest(*codeguruprofiler.PostAgentProfileInput) codeguruprofiler.PostAgentProfileRequest

	PutPermissionRequest(*codeguruprofiler.PutPermissionInput) codeguruprofiler.PutPermissionRequest

	RemovePermissionRequest(*codeguruprofiler.RemovePermissionInput) codeguruprofiler.RemovePermissionRequest

	UpdateProfilingGroupRequest(*codeguruprofiler.UpdateProfilingGroupInput) codeguruprofiler.UpdateProfilingGroupRequest
}

var _ ClientAPI = (*codeguruprofiler.Client)(nil)
