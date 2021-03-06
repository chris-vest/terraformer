// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListObjectsV2Input struct {
	_ struct{} `type:"structure"`

	// Bucket name to list.
	//
	// When using this API with an access point, you must direct requests to the
	// access point hostname. The access point hostname takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this operation using an access point through the AWS SDKs, you
	// provide the access point ARN in place of the bucket name. For more information
	// about access point ARNs, see Using Access Points (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// ContinuationToken indicates Amazon S3 that the list is being continued on
	// this bucket with a token. ContinuationToken is obfuscated and is not a real
	// key.
	ContinuationToken *string `location:"querystring" locationName:"continuation-token" type:"string"`

	// A delimiter is a character you use to group keys.
	Delimiter *string `location:"querystring" locationName:"delimiter" type:"string"`

	// Encoding type used by Amazon S3 to encode object keys in the response.
	EncodingType EncodingType `location:"querystring" locationName:"encoding-type" type:"string" enum:"true"`

	// The owner field is not present in listV2 by default, if you want to return
	// owner field with each key in the result then set the fetch owner field to
	// true.
	FetchOwner *bool `location:"querystring" locationName:"fetch-owner" type:"boolean"`

	// Sets the maximum number of keys returned in the response. The response might
	// contain fewer keys but will never contain more.
	MaxKeys *int64 `location:"querystring" locationName:"max-keys" type:"integer"`

	// Limits the response to keys that begin with the specified prefix.
	Prefix *string `location:"querystring" locationName:"prefix" type:"string"`

	// Confirms that the requester knows that she or he will be charged for the
	// list objects request in V2 style. Bucket owners need not specify this parameter
	// in their requests.
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// StartAfter is where you want Amazon S3 to start listing from. Amazon S3 starts
	// listing after this specified key. StartAfter can be any key in the bucket.
	StartAfter *string `location:"querystring" locationName:"start-after" type:"string"`
}

// String returns the string representation
func (s ListObjectsV2Input) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListObjectsV2Input) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListObjectsV2Input"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *ListObjectsV2Input) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectsV2Input) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.ContinuationToken != nil {
		v := *s.ContinuationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "continuation-token", protocol.StringValue(v), metadata)
	}
	if s.Delimiter != nil {
		v := *s.Delimiter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "delimiter", protocol.StringValue(v), metadata)
	}
	if len(s.EncodingType) > 0 {
		v := s.EncodingType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "encoding-type", v, metadata)
	}
	if s.FetchOwner != nil {
		v := *s.FetchOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "fetch-owner", protocol.BoolValue(v), metadata)
	}
	if s.MaxKeys != nil {
		v := *s.MaxKeys

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-keys", protocol.Int64Value(v), metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "prefix", protocol.StringValue(v), metadata)
	}
	if s.StartAfter != nil {
		v := *s.StartAfter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "start-after", protocol.StringValue(v), metadata)
	}
	return nil
}

type ListObjectsV2Output struct {
	_ struct{} `type:"structure"`

	// All of the keys rolled up into a common prefix count as a single return when
	// calculating the number of returns.
	//
	// A response can contain CommonPrefixes only if you specify a delimiter.
	//
	// CommonPrefixes contains all (if there are any) keys between Prefix and the
	// next occurrence of the string specified by a delimiter.
	//
	// CommonPrefixes lists keys that act like subdirectories in the directory specified
	// by Prefix.
	//
	// For example, if the prefix is notes/ and the delimiter is a slash (/) as
	// in notes/summer/july, the common prefix is notes/summer/. All of the keys
	// that roll up into a common prefix count as a single return when calculating
	// the number of returns.
	CommonPrefixes []CommonPrefix `type:"list" flattened:"true"`

	// Metadata about each object returned.
	Contents []Object `type:"list" flattened:"true"`

	// If ContinuationToken was sent with the request, it is included in the response.
	ContinuationToken *string `type:"string"`

	// Causes keys that contain the same string between the prefix and the first
	// occurrence of the delimiter to be rolled up into a single result element
	// in the CommonPrefixes collection. These rolled-up keys are not returned elsewhere
	// in the response. Each rolled-up result counts as only one return against
	// the MaxKeys value.
	Delimiter *string `type:"string"`

	// Encoding type used by Amazon S3 to encode object key names in the XML response.
	//
	// If you specify the encoding-type request parameter, Amazon S3 includes this
	// element in the response, and returns encoded key name values in the following
	// response elements:
	//
	// Delimiter, Prefix, Key, and StartAfter.
	EncodingType EncodingType `type:"string" enum:"true"`

	// Set to false if all of the results were returned. Set to true if more keys
	// are available to return. If the number of results exceeds that specified
	// by MaxKeys, all of the results might not be returned.
	IsTruncated *bool `type:"boolean"`

	// KeyCount is the number of keys returned with this request. KeyCount will
	// always be less than equals to MaxKeys field. Say you ask for 50 keys, your
	// result will include less than equals 50 keys
	KeyCount *int64 `type:"integer"`

	// Sets the maximum number of keys returned in the response. The response might
	// contain fewer keys but will never contain more.
	MaxKeys *int64 `type:"integer"`

	// Bucket name.
	//
	// When using this API with an access point, you must direct requests to the
	// access point hostname. The access point hostname takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this operation using an access point through the AWS SDKs, you
	// provide the access point ARN in place of the bucket name. For more information
	// about access point ARNs, see Using Access Points (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html)
	// in the Amazon Simple Storage Service Developer Guide.
	Name *string `type:"string"`

	// NextContinuationToken is sent when isTruncated is true, which means there
	// are more keys in the bucket that can be listed. The next list requests to
	// Amazon S3 can be continued with this NextContinuationToken. NextContinuationToken
	// is obfuscated and is not a real key
	NextContinuationToken *string `type:"string"`

	// Keys that begin with the indicated prefix.
	Prefix *string `type:"string"`

	// If StartAfter was sent with the request, it is included in the response.
	StartAfter *string `type:"string"`
}

// String returns the string representation
func (s ListObjectsV2Output) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectsV2Output) MarshalFields(e protocol.FieldEncoder) error {
	if s.CommonPrefixes != nil {
		v := s.CommonPrefixes

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "CommonPrefixes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Contents != nil {
		v := s.Contents

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "Contents", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ContinuationToken != nil {
		v := *s.ContinuationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ContinuationToken", protocol.StringValue(v), metadata)
	}
	if s.Delimiter != nil {
		v := *s.Delimiter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Delimiter", protocol.StringValue(v), metadata)
	}
	if len(s.EncodingType) > 0 {
		v := s.EncodingType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EncodingType", v, metadata)
	}
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.KeyCount != nil {
		v := *s.KeyCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KeyCount", protocol.Int64Value(v), metadata)
	}
	if s.MaxKeys != nil {
		v := *s.MaxKeys

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxKeys", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.StringValue(v), metadata)
	}
	if s.NextContinuationToken != nil {
		v := *s.NextContinuationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextContinuationToken", protocol.StringValue(v), metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Prefix", protocol.StringValue(v), metadata)
	}
	if s.StartAfter != nil {
		v := *s.StartAfter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StartAfter", protocol.StringValue(v), metadata)
	}
	return nil
}

const opListObjectsV2 = "ListObjectsV2"

// ListObjectsV2Request returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns some or all (up to 1,000) of the objects in a bucket. You can use
// the request parameters as selection criteria to return a subset of the objects
// in a bucket. A 200 OK response can contain valid or invalid XML. Make sure
// to design your application to parse the contents of the response and handle
// it appropriately.
//
// To use this operation, you must have READ access to the bucket.
//
// To use this operation in an AWS Identity and Access Management (IAM) policy,
// you must have permissions to perform the s3:ListBucket action. The bucket
// owner has this permission by default and can grant this permission to others.
// For more information about permissions, see Permissions Related to Bucket
// Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// This section describes the latest revision of the API. We recommend that
// you use this revised API for application development. For backward compatibility,
// Amazon S3 continues to support the prior version of this API, ListObjects.
//
// To get a list of your buckets, see ListBuckets.
//
// The following operations are related to ListObjectsV2:
//
//    * GetObject
//
//    * PutObject
//
//    * CreateBucket
//
//    // Example sending a request using ListObjectsV2Request.
//    req := client.ListObjectsV2Request(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListObjectsV2
func (c *Client) ListObjectsV2Request(input *ListObjectsV2Input) ListObjectsV2Request {
	op := &aws.Operation{
		Name:       opListObjectsV2,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?list-type=2",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"ContinuationToken"},
			OutputTokens:    []string{"NextContinuationToken"},
			LimitToken:      "MaxKeys",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListObjectsV2Input{}
	}

	req := c.newRequest(op, input, &ListObjectsV2Output{})
	return ListObjectsV2Request{Request: req, Input: input, Copy: c.ListObjectsV2Request}
}

// ListObjectsV2Request is the request type for the
// ListObjectsV2 API operation.
type ListObjectsV2Request struct {
	*aws.Request
	Input *ListObjectsV2Input
	Copy  func(*ListObjectsV2Input) ListObjectsV2Request
}

// Send marshals and sends the ListObjectsV2 API request.
func (r ListObjectsV2Request) Send(ctx context.Context) (*ListObjectsV2Response, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListObjectsV2Response{
		ListObjectsV2Output: r.Request.Data.(*ListObjectsV2Output),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListObjectsV2RequestPaginator returns a paginator for ListObjectsV2.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListObjectsV2Request(input)
//   p := s3.NewListObjectsV2RequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListObjectsV2Paginator(req ListObjectsV2Request) ListObjectsV2Paginator {
	return ListObjectsV2Paginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListObjectsV2Input
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

// ListObjectsV2Paginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListObjectsV2Paginator struct {
	aws.Pager
}

func (p *ListObjectsV2Paginator) CurrentPage() *ListObjectsV2Output {
	return p.Pager.CurrentPage().(*ListObjectsV2Output)
}

// ListObjectsV2Response is the response type for the
// ListObjectsV2 API operation.
type ListObjectsV2Response struct {
	*ListObjectsV2Output

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListObjectsV2 request.
func (r *ListObjectsV2Response) SDKResponseMetdata() *aws.Response {
	return r.response
}
