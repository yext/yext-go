package yext

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const reviewsPath = "reviews"

const reviewInvitePath = "reviewinvites"

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
	NextPageToken string    `json:"nextPageToken"`
}

type ReviewCreateInvitationResponse struct {
	Id         string `json:"id"`
	LocationId string `json:"locationId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Contact    string `json:"contact"`
	Image      bool   `json:"image"`
	TemplateId string `json:"templateId"`
	Status     string `json:"status"`
	Details    string `json:"details"`
}

func (l *ReviewService) ListAllWithOptions(rlOpts *ReviewListOptions) ([]*Review, error) {
	var (
		reviews []*Review
		lo      = rlOpts
	)
	if lo == nil {
		lo = &ReviewListOptions{}
	}

	lo.ListOptions = ListOptions{Limit: ReviewListMaxLimit, DisableCountValidation: true}

	var lg tokenListRetriever = func(opts *ListOptions) (string, error) {
		lo.ListOptions = *opts
		rlr, _, err := l.List(lo)
		if err != nil {
			return "", err
		}
		reviews = append(reviews, rlr.Reviews...)
		return rlr.NextPageToken, err
	}

	if err := tokenListHelper(lg, &lo.ListOptions); err != nil {
		return nil, err
	} else {
		return reviews, nil
	}
}

func (l *ReviewService) ListAll() ([]*Review, error) {
	return l.ListAllWithOptions(nil)
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

func (l *ReviewService) CreateInvitation(jsonData []Reviewer) (*[]ReviewCreateInvitationResponse, *Response, error) {
	v := &[]ReviewCreateInvitationResponse{}
	r, err := l.client.DoRequestJSON("POST", reviewInvitePath, jsonData, v)
	if err != nil {
		return nil, r, err
	}

	return v, r, nil
}
