// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SearchDevicesRequest
type SearchDevicesInput struct {
	_ struct{} `type:"structure"`

	// The filters to use to list a specified set of devices. Supported filter keys
	// are DeviceName, DeviceStatus, DeviceStatusDetailCode, RoomName, DeviceType,
	// DeviceSerialNumber, UnassociatedOnly, and ConnectionStatus (ONLINE and OFFLINE).
	Filters []Filter `type:"list"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	MaxResults *int64 `min:"1" type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `min:"1" type:"string"`

	// The sort order to use in listing the specified set of devices. Supported
	// sort keys are DeviceName, DeviceStatus, RoomName, DeviceType, DeviceSerialNumber,
	// and ConnectionStatus.
	SortCriteria []Sort `type:"list"`
}

// String returns the string representation
func (s SearchDevicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchDevicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchDevicesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SortCriteria != nil {
		for i, v := range s.SortCriteria {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SortCriteria", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SearchDevicesResponse
type SearchDevicesOutput struct {
	_ struct{} `type:"structure"`

	// The devices that meet the specified set of filter criteria, in sort order.
	Devices []DeviceData `type:"list"`

	// The token returned to indicate that there is more data available.
	NextToken *string `min:"1" type:"string"`

	// The total number of devices returned.
	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s SearchDevicesOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchDevices = "SearchDevices"

// SearchDevicesRequest returns a request value for making API operation for
// Alexa For Business.
//
// Searches devices and lists the ones that meet a set of filter criteria.
//
//    // Example sending a request using SearchDevicesRequest.
//    req := client.SearchDevicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SearchDevices
func (c *Client) SearchDevicesRequest(input *SearchDevicesInput) SearchDevicesRequest {
	op := &aws.Operation{
		Name:       opSearchDevices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &SearchDevicesInput{}
	}

	req := c.newRequest(op, input, &SearchDevicesOutput{})
	return SearchDevicesRequest{Request: req, Input: input, Copy: c.SearchDevicesRequest}
}

// SearchDevicesRequest is the request type for the
// SearchDevices API operation.
type SearchDevicesRequest struct {
	*aws.Request
	Input *SearchDevicesInput
	Copy  func(*SearchDevicesInput) SearchDevicesRequest
}

// Send marshals and sends the SearchDevices API request.
func (r SearchDevicesRequest) Send(ctx context.Context) (*SearchDevicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchDevicesResponse{
		SearchDevicesOutput: r.Request.Data.(*SearchDevicesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchDevicesRequestPaginator returns a paginator for SearchDevices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchDevicesRequest(input)
//   p := alexaforbusiness.NewSearchDevicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchDevicesPaginator(req SearchDevicesRequest) SearchDevicesPaginator {
	return SearchDevicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *SearchDevicesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// SearchDevicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchDevicesPaginator struct {
	aws.Pager
}

func (p *SearchDevicesPaginator) CurrentPage() *SearchDevicesOutput {
	return p.Pager.CurrentPage().(*SearchDevicesOutput)
}

// SearchDevicesResponse is the response type for the
// SearchDevices API operation.
type SearchDevicesResponse struct {
	*SearchDevicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchDevices request.
func (r *SearchDevicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}