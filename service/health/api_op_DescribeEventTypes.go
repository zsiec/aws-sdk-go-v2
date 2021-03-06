// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package health

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEventTypesInput struct {
	_ struct{} `type:"structure"`

	// Values to narrow the results returned.
	Filter *EventTypeFilter `locationName:"filter" type:"structure"`

	// The locale (language) to return information in. English (en) is the default
	// and the only supported value at this time.
	Locale *string `locationName:"locale" min:"2" type:"string"`

	// The maximum number of items to return in one batch, between 10 and 100, inclusive.
	MaxResults *int64 `locationName:"maxResults" min:"10" type:"integer"`

	// If the results of a search are large, only a portion of the results are returned,
	// and a nextToken pagination token is returned in the response. To retrieve
	// the next batch of results, reissue the search request and include the returned
	// token. When all results have been returned, the response does not contain
	// a pagination token value.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s DescribeEventTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventTypesInput"}
	if s.Locale != nil && len(*s.Locale) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Locale", 2))
	}
	if s.MaxResults != nil && *s.MaxResults < 10 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 10))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEventTypesOutput struct {
	_ struct{} `type:"structure"`

	// A list of event types that match the filter criteria. Event types have a
	// category (issue, accountNotification, or scheduledChange), a service (for
	// example, EC2, RDS, DATAPIPELINE, BILLING), and a code (in the format AWS_SERVICE_DESCRIPTION
	// ; for example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT).
	EventTypes []EventType `locationName:"eventTypes" type:"list"`

	// If the results of a search are large, only a portion of the results are returned,
	// and a nextToken pagination token is returned in the response. To retrieve
	// the next batch of results, reissue the search request and include the returned
	// token. When all results have been returned, the response does not contain
	// a pagination token value.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s DescribeEventTypesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEventTypes = "DescribeEventTypes"

// DescribeEventTypesRequest returns a request value for making API operation for
// AWS Health APIs and Notifications.
//
// Returns the event types that meet the specified filter criteria. If no filter
// criteria are specified, all event types are returned, in no particular order.
//
//    // Example sending a request using DescribeEventTypesRequest.
//    req := client.DescribeEventTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/health-2016-08-04/DescribeEventTypes
func (c *Client) DescribeEventTypesRequest(input *DescribeEventTypesInput) DescribeEventTypesRequest {
	op := &aws.Operation{
		Name:       opDescribeEventTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEventTypesInput{}
	}

	req := c.newRequest(op, input, &DescribeEventTypesOutput{})
	return DescribeEventTypesRequest{Request: req, Input: input, Copy: c.DescribeEventTypesRequest}
}

// DescribeEventTypesRequest is the request type for the
// DescribeEventTypes API operation.
type DescribeEventTypesRequest struct {
	*aws.Request
	Input *DescribeEventTypesInput
	Copy  func(*DescribeEventTypesInput) DescribeEventTypesRequest
}

// Send marshals and sends the DescribeEventTypes API request.
func (r DescribeEventTypesRequest) Send(ctx context.Context) (*DescribeEventTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventTypesResponse{
		DescribeEventTypesOutput: r.Request.Data.(*DescribeEventTypesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEventTypesRequestPaginator returns a paginator for DescribeEventTypes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEventTypesRequest(input)
//   p := health.NewDescribeEventTypesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEventTypesPaginator(req DescribeEventTypesRequest) DescribeEventTypesPaginator {
	return DescribeEventTypesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEventTypesInput
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

// DescribeEventTypesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEventTypesPaginator struct {
	aws.Pager
}

func (p *DescribeEventTypesPaginator) CurrentPage() *DescribeEventTypesOutput {
	return p.Pager.CurrentPage().(*DescribeEventTypesOutput)
}

// DescribeEventTypesResponse is the response type for the
// DescribeEventTypes API operation.
type DescribeEventTypesResponse struct {
	*DescribeEventTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEventTypes request.
func (r *DescribeEventTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
