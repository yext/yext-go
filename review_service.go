package yext

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const reviewsPath = "reviews"

const reviewInvitePath = "reviewinvites"

// Review update enums
const (
	ReviewStatusLive        = "LIVE"
	ReviewStatusQuarantined = "QUARANTINED"
	ReviewStatusRemoved     = "REMOVED"

	ReviewFlagStatusInappropriate = "INAPPROPRIATE"
	ReviewFlagStatusSpam          = "SPAM"
	ReviewFlagStatusIrrelevant    = "IRRELEVANT"
	ReviewFlagStatusSensitive     = "SENSITIVE"
	ReviewFlagStatusNotFlagged    = "NOT_FLAGGED"
)

var (
	ReviewListMaxLimit = 100
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
	LabelIds              []string
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
	Status                string
}

type ReviewListResponse struct {
	Count         int       `json:"count"`
	AverageRating float64   `json:"averageRating"`
	Reviews       []*Review `json:"reviews"`
	NextPageToken string    `json:"nextPageToken"`
}

type ReviewUpdateOptions struct {
	Rating      float64 `json:"rating,omitempty"`
	Content     string  `json:"content,omitempty"`
	AuthorName  string  `json:"authorName,omitempty"`
	AuthorEmail string  `json:"authorEmail,omitempty"`
	LocationId  string  `json:"locationId,omitempty"`
	Status      string  `json:"status,omitempty"`
	FlagStatus  string  `json:"flagStatus,omitempty"`
	ExternalId  *string `json:"externalId"` // Must have v param >= 20220120
}

type ReviewUpdateResponse struct {
	Id IdWrapper `json:"id"`
}

// Wrapper types are helpful because the Yext API versions have different data types that the ID is returned as.
// These wrappers allow unmarshalling and using as a string vs using an interface and type casting every time
type IdWrapper string

func (w *IdWrapper) UnmarshalJSON(data []byte) (err error) {
	if id, err := strconv.Atoi(string(data)); err == nil {
		str := strconv.Itoa(id)
		*w = IdWrapper(str)
		return nil
	}
	var str string
	err = json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	*w = IdWrapper(str)
	return nil
}

func (n IdWrapper) String() string {
	return string(n)
}

type ReviewCreateInvitationResponse struct {
	InvitationUID string         `json:"invitationUid"`
	Id            string         `json:"id"`
	LocationId    string         `json:"locationId"`
	Entity        **ReviewEntity `json:"entity,omitempty"` // Must have v param >= 20210728
	FirstName     string         `json:"firstName"`
	LastName      string         `json:"lastName"`
	Contact       string         `json:"contact"`
	Image         bool           `json:"image"`
	TemplateId    int            `json:"templateId"`
	Status        string         `json:"status"`
	Details       string         `json:"details"`
}

type ReviewCreateReviewResponse struct {
	Id string `json:"id"`
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
	if opts.LabelIds != nil {
		q.Add("labelIds", strings.Join(opts.LabelIds, ","))
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
	if opts.Status != "" {
		q.Add("status", opts.Status)
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

func (l *ReviewService) Update(id int, opts *ReviewUpdateOptions) (*ReviewUpdateResponse, *Response, error) {
	var v ReviewUpdateResponse
	r, err := l.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%d", reviewsPath, id), opts, &v)
	if err != nil {
		return nil, r, err
	}

	return &v, r, nil
}

func (l *ReviewService) CreateInvitation(jsonData []*Reviewer) ([]*ReviewCreateInvitationResponse, *Response, error) {
	var v []*ReviewCreateInvitationResponse
	r, err := l.client.DoRequestJSON("POST", reviewInvitePath, jsonData, &v)
	if err != nil {
		return nil, r, err
	}

	return v, r, nil
}

func (l *ReviewService) CreateReview(jsonData *ReviewCreate) (*ReviewCreateReviewResponse, *Response, error) {
	var v *ReviewCreateReviewResponse
	r, err := l.client.DoRequestJSON("POST", reviewsPath, jsonData, &v)
	if err != nil {
		return nil, r, err
	}

	return v, r, nil
}

// the new way to create reviews on the yext platform
// refer to https://yextops.slack.com/archives/C01269F1ZTL/p1634751884059700
func (l *ReviewService) CreateReviewLiveAPI(jsonData *ReviewCreate) (*ReviewCreateReviewResponse, *Response, error) {
	reviewCreateReviewResponse := &ReviewCreateReviewResponse{}
	baseURL := "https://liveapi.yext.com"
	if strings.Contains(l.client.Config.BaseUrl, "sandbox") {
		baseURL = "https://liveapi-sandbox.yext.com"
	}

	reviewSubmissionLiveAPIURL := fmt.Sprintf("%s/v2/accounts/%s/reviewSubmission?api_key=%s&v=%s", baseURL, l.client.Config.AccountId, l.client.Config.ApiKey, l.client.Config.Version)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, reviewSubmissionLiveAPIURL, strings.NewReader(jsonData.String()))
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	reviewResponse := &Response{}

	err = json.Unmarshal(body, reviewResponse)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(*reviewResponse.ResponseRaw, reviewCreateReviewResponse)
	if err != nil {
		return nil, nil, err
	}

	if len(reviewResponse.Meta.Errors) > 0 {
		return nil, reviewResponse, fmt.Errorf("error creating review: %s", reviewResponse.Meta.Errors[0].Message)
	}

	return reviewCreateReviewResponse, reviewResponse, nil
}

type ReviewUpdateLabelOptions struct {
	LabelIds *[]int `json:"labelIds,omitempty"`
}

func (l *ReviewService) UpdateLabels(id int, opts *ReviewUpdateLabelOptions) (*ReviewUpdateResponse, *Response, error) {
	var v ReviewUpdateResponse
	r, err := l.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%d/labels", reviewsPath, id), opts, &v)
	if err != nil {
		return nil, r, err
	}

	return &v, r, nil
}
