// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListVocabulariesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of vocabularies to return in the response. If there are
	// fewer results in the list, this response contains only the actual results.
	MaxResults *int64 `min:"1" type:"integer"`

	// When specified, the vocabularies returned in the list are limited to vocabularies
	// whose name contains the specified string. The search is case-insensitive,
	// ListVocabularies returns both "vocabularyname" and "VocabularyName" in the
	// response list.
	NameContains *string `min:"1" type:"string"`

	// If the result of the previous request to ListVocabularies was truncated,
	// include the NextToken to fetch the next set of jobs.
	NextToken *string `type:"string"`

	// When specified, only returns vocabularies with the VocabularyState field
	// equal to the specified state.
	StateEquals VocabularyState `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListVocabulariesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVocabulariesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVocabulariesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListVocabulariesOutput struct {
	_ struct{} `type:"structure"`

	// The ListVocabularies operation returns a page of vocabularies at a time.
	// The maximum size of the page is set by the MaxResults parameter. If there
	// are more jobs in the list than the page size, Amazon Transcribe returns the
	// NextPage token. Include the token in the next request to the ListVocabularies
	// operation to return in the next page of jobs.
	NextToken *string `type:"string"`

	// The requested vocabulary state.
	Status TranscriptionJobStatus `type:"string" enum:"true"`

	// A list of objects that describe the vocabularies that match the search criteria
	// in the request.
	Vocabularies []VocabularyInfo `type:"list"`
}

// String returns the string representation
func (s ListVocabulariesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListVocabularies = "ListVocabularies"

// ListVocabulariesRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Returns a list of vocabularies that match the specified criteria. If no criteria
// are specified, returns the entire list of vocabularies.
//
//    // Example sending a request using ListVocabulariesRequest.
//    req := client.ListVocabulariesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/ListVocabularies
func (c *Client) ListVocabulariesRequest(input *ListVocabulariesInput) ListVocabulariesRequest {
	op := &aws.Operation{
		Name:       opListVocabularies,
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
		input = &ListVocabulariesInput{}
	}

	req := c.newRequest(op, input, &ListVocabulariesOutput{})
	return ListVocabulariesRequest{Request: req, Input: input, Copy: c.ListVocabulariesRequest}
}

// ListVocabulariesRequest is the request type for the
// ListVocabularies API operation.
type ListVocabulariesRequest struct {
	*aws.Request
	Input *ListVocabulariesInput
	Copy  func(*ListVocabulariesInput) ListVocabulariesRequest
}

// Send marshals and sends the ListVocabularies API request.
func (r ListVocabulariesRequest) Send(ctx context.Context) (*ListVocabulariesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVocabulariesResponse{
		ListVocabulariesOutput: r.Request.Data.(*ListVocabulariesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListVocabulariesRequestPaginator returns a paginator for ListVocabularies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListVocabulariesRequest(input)
//   p := transcribe.NewListVocabulariesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListVocabulariesPaginator(req ListVocabulariesRequest) ListVocabulariesPaginator {
	return ListVocabulariesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListVocabulariesInput
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

// ListVocabulariesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListVocabulariesPaginator struct {
	aws.Pager
}

func (p *ListVocabulariesPaginator) CurrentPage() *ListVocabulariesOutput {
	return p.Pager.CurrentPage().(*ListVocabulariesOutput)
}

// ListVocabulariesResponse is the response type for the
// ListVocabularies API operation.
type ListVocabulariesResponse struct {
	*ListVocabulariesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVocabularies request.
func (r *ListVocabulariesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
