package yext

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const reviewsPath = "reviews"

var (
	ReviewListMaxLimit = 50
)

type ReviewService struct {
	client *Client
}

type ReviewListOptions struct {
	ListOptions
	LocationIds           []string
	FolderId              string
	Countries             []string
	LocationLabels        []string
	PublisherIds          []string
	ReviewContent         string
	MinRating             float64
	MaxRating             float64
	MinPublisherDate      string
	MaxPublisherDate      string
	MinLastYextUpdateDate string
	MaxLastYextUpdateDate string
	AwaitingResponse      string
	MinNonOwnerComments   int
	ReviewerName          string
	ReviewerEmail         string
}

type ReviewListResponse struct {
	Count         int       `json:"count"`
	AverageRating float64   `json:"averageRating"`
	Reviews       []*Review `json:"reviews"`
}

func (l *ReviewService) ListAll() ([]*Review, error) {
	var reviews []*Review
	var lg listRetriever = func(opts *ListOptions) (int, int, error) {
		llr, _, err := l.List(&ReviewListOptions{ListOptions: *opts})
		if err != nil {
			return 0, 0, err
		}
		reviews = append(reviews, llr.Reviews...)
		return len(llr.Reviews), llr.Count, err
	}

	if err := listHelper(lg, ReviewListMaxLimit); err != nil {
		return nil, err
	} else {
		return reviews, nil
	}
}

func addReviewListOptions(requrl string, opts *ReviewListOptions) (string, error) {
	u, err := url.Parse(requrl)
	if err != nil {
		return "", nil
	}

	if opts == nil {
		return requrl, nil
	}

	q := u.Query()
	if opts.LocationIds != nil {
		q.Add("locationIds", strings.Join(opts.LocationIds, ","))
	}
	if opts.FolderId != "" {
		q.Add("folderId", opts.FolderId)
	}
	if opts.Countries != nil {
		q.Add("countries", strings.Join(opts.Countries, ","))
	}
	if opts.LocationLabels != nil {
		q.Add("locationLabels", strings.Join(opts.LocationLabels, ","))
	}
	if opts.PublisherIds != nil {
		q.Add("publisherIds", strings.Join(opts.PublisherIds, ","))
	}
	if opts.ReviewContent != "" {
		q.Add("reviewContent", opts.ReviewContent)
	}
	if opts.MinRating != 0 {
		q.Add("minRating", strconv.FormatFloat(opts.MinRating, 'f', -1, 64))
	}
	if opts.MaxRating != 0 {
		q.Add("maxRating", strconv.FormatFloat(opts.MaxRating, 'f', -1, 64))
	}
	if opts.MinPublisherDate != "" {
		q.Add("minPublisherDate", opts.MinPublisherDate)
	}
	if opts.MaxPublisherDate != "" {
		q.Add("maxPublisherDate", opts.MaxPublisherDate)
	}
	if opts.MinLastYextUpdateDate != "" {
		q.Add("minLastYextUpdateDate", opts.MinLastYextUpdateDate)
	}
	if opts.MaxLastYextUpdateDate != "" {
		q.Add("maxLastYextUpdateDate", opts.MaxLastYextUpdateDate)
	}
	if opts.AwaitingResponse != "" {
		q.Add("awaitingResponse", opts.AwaitingResponse)
	}
	if opts.MinNonOwnerComments != 0 {
		q.Add("minNonOwnerComments", strconv.Itoa(opts.MinNonOwnerComments))
	}
	if opts.ReviewerName != "" {
		q.Add("reviewerName", opts.ReviewerName)
	}
	if opts.ReviewerEmail != "" {
		q.Add("reviewerEmail", opts.ReviewerEmail)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (l *ReviewService) List(opts *ReviewListOptions) (*ReviewListResponse, *Response, error) {

	requrl, err := addReviewListOptions(reviewsPath, opts)
	if err != nil {
		return nil, nil, err
	}

	if opts != nil {
		requrl, err = addListOptions(requrl, &opts.ListOptions)
		if err != nil {
			return nil, nil, err
		}
	}

	v := &ReviewListResponse{}
	r, err := l.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}

	return v, r, nil
}

func (l *ReviewService) Get(id int) (*Review, *Response, error) {
	var v Review
	r, err := l.client.DoRequest("GET", fmt.Sprintf("%s/%d", reviewsPath, id), &v)
	if err != nil {
		return nil, r, err
	}

	return &v, r, nil
}
