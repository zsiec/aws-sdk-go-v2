// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package frauddetectoriface provides an interface to enable mocking the Amazon Fraud Detector service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package frauddetectoriface

import (
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
)

// ClientAPI provides an interface to enable mocking the
// frauddetector.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Fraud Detector.
//    func myFunc(svc frauddetectoriface.ClientAPI) bool {
//        // Make svc.BatchCreateVariable request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := frauddetector.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        frauddetectoriface.ClientPI
//    }
//    func (m *mockClientClient) BatchCreateVariable(input *frauddetector.BatchCreateVariableInput) (*frauddetector.BatchCreateVariableOutput, error) {
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
	BatchCreateVariableRequest(*frauddetector.BatchCreateVariableInput) frauddetector.BatchCreateVariableRequest

	BatchGetVariableRequest(*frauddetector.BatchGetVariableInput) frauddetector.BatchGetVariableRequest

	CreateDetectorVersionRequest(*frauddetector.CreateDetectorVersionInput) frauddetector.CreateDetectorVersionRequest

	CreateModelVersionRequest(*frauddetector.CreateModelVersionInput) frauddetector.CreateModelVersionRequest

	CreateRuleRequest(*frauddetector.CreateRuleInput) frauddetector.CreateRuleRequest

	CreateVariableRequest(*frauddetector.CreateVariableInput) frauddetector.CreateVariableRequest

	DeleteDetectorRequest(*frauddetector.DeleteDetectorInput) frauddetector.DeleteDetectorRequest

	DeleteDetectorVersionRequest(*frauddetector.DeleteDetectorVersionInput) frauddetector.DeleteDetectorVersionRequest

	DeleteEventRequest(*frauddetector.DeleteEventInput) frauddetector.DeleteEventRequest

	DeleteRuleVersionRequest(*frauddetector.DeleteRuleVersionInput) frauddetector.DeleteRuleVersionRequest

	DescribeDetectorRequest(*frauddetector.DescribeDetectorInput) frauddetector.DescribeDetectorRequest

	DescribeModelVersionsRequest(*frauddetector.DescribeModelVersionsInput) frauddetector.DescribeModelVersionsRequest

	GetDetectorVersionRequest(*frauddetector.GetDetectorVersionInput) frauddetector.GetDetectorVersionRequest

	GetDetectorsRequest(*frauddetector.GetDetectorsInput) frauddetector.GetDetectorsRequest

	GetExternalModelsRequest(*frauddetector.GetExternalModelsInput) frauddetector.GetExternalModelsRequest

	GetModelVersionRequest(*frauddetector.GetModelVersionInput) frauddetector.GetModelVersionRequest

	GetModelsRequest(*frauddetector.GetModelsInput) frauddetector.GetModelsRequest

	GetOutcomesRequest(*frauddetector.GetOutcomesInput) frauddetector.GetOutcomesRequest

	GetPredictionRequest(*frauddetector.GetPredictionInput) frauddetector.GetPredictionRequest

	GetRulesRequest(*frauddetector.GetRulesInput) frauddetector.GetRulesRequest

	GetVariablesRequest(*frauddetector.GetVariablesInput) frauddetector.GetVariablesRequest

	PutDetectorRequest(*frauddetector.PutDetectorInput) frauddetector.PutDetectorRequest

	PutExternalModelRequest(*frauddetector.PutExternalModelInput) frauddetector.PutExternalModelRequest

	PutModelRequest(*frauddetector.PutModelInput) frauddetector.PutModelRequest

	PutOutcomeRequest(*frauddetector.PutOutcomeInput) frauddetector.PutOutcomeRequest

	UpdateDetectorVersionRequest(*frauddetector.UpdateDetectorVersionInput) frauddetector.UpdateDetectorVersionRequest

	UpdateDetectorVersionMetadataRequest(*frauddetector.UpdateDetectorVersionMetadataInput) frauddetector.UpdateDetectorVersionMetadataRequest

	UpdateDetectorVersionStatusRequest(*frauddetector.UpdateDetectorVersionStatusInput) frauddetector.UpdateDetectorVersionStatusRequest

	UpdateModelVersionRequest(*frauddetector.UpdateModelVersionInput) frauddetector.UpdateModelVersionRequest

	UpdateRuleMetadataRequest(*frauddetector.UpdateRuleMetadataInput) frauddetector.UpdateRuleMetadataRequest

	UpdateRuleVersionRequest(*frauddetector.UpdateRuleVersionInput) frauddetector.UpdateRuleVersionRequest

	UpdateVariableRequest(*frauddetector.UpdateVariableInput) frauddetector.UpdateVariableRequest
}

var _ ClientAPI = (*frauddetector.Client)(nil)
