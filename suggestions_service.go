package yext

import (
	"net/url"
	"strings"
)

const (
	suggestionsPath         = "listings"
	suggestionsListMaxLimit = 50
)

type SuggestionsResponse struct {
	Count         int            `json:"count"`
	Suggestions   *[]Suggestions `json:"suggestions,omitempty"`
	NextPageToken string         `json:"nextPageToken"`
}

type Source struct {
	UserID      string `json:"userId"`
	AppID       string `json:"appId"`
	PublisherID string `json:"publisherId"`
	YextSource  string `json:"yextSource"`
}

type EntityContent struct {
	ID       string   `json:"id"`
	UID      string   `json:"uid"`
	Type     string   `json:"type"`
	Language string   `json:"language"`
	FolderID string   `json:"folderId"`
	Labels   []string `json:"labels"`
}

type EntityFieldSuggestion struct {
	Entity           EntityContent          `json:"entity"`
	ExistingContent  map[string]interface{} `json:"existingContent,omitempty"`
	SuggestedContent map[string]interface{} `json:"suggestedContent,omitempty"`
}

type Assignee struct {
	UserID      string `json:"userId"`
	UserGroupID string `json:"userGroupId"`
}

type Commenter struct {
	UserID     string `json:"userId"`
	AppID      string `json:"appId"`
	YextSource string `json:"yextSource"`
}

type Attachments struct {
	Name        string `json:"name"`
	DownloadURL string `json:"downloadUrl"`
}

type Comments struct {
	Commenter   Commenter      `json:"commenter"`
	Text        string         `json:"text"`
	Createddate string         `json:"createdDate"`
	Attachments *[]Attachments `json:"attachments,omitempty"`
}

type Approver struct {
	UserID string `json:"userId"`
	AppID  string `json:"appId"`
}

type Suggestions struct {
	UID                   string                `json:"uid"`
	Accountid             string                `json:"accountId"`
	Createddate           string                `json:"createdDate"`
	Lastupdateddate       string                `json:"lastUpdatedDate"`
	Resolveddate          string                `json:"resolvedDate"`
	Source                Source                `json:"source"`
	EntityFieldSuggestion EntityFieldSuggestion `json:"entityFieldSuggestion"`
	Status                string                `json:"status"`
	Locked                bool                  `json:"locked"`
	Assignee              Assignee              `json:"assignee"`
	Comments              *[]Comments           `json:"comments,omitempty"`
	Approver              Approver              `json:"approver"`
}

type SuggestionsListOptions struct {
	ListOptions
	Format     RichTextFormat          `json:"format,omitempty"`
	EntityIds  []string                `json:"entity_ids,omitempty"`
	EntityUIDs []string                `json:"entity_ui_ds,omitempty"`
	Statuses   []SuggestionsStatusType `json:"statuses,omitempty"`
}

// APPROVED, REJECTED, PENDING, CANCELED
type SuggestionsStatusType int

const (
	Approved SuggestionsStatusType = iota
	Rejected
	Pending
	Canceled
)

func (l SuggestionsStatusType) ToString() string {
	switch l {
	case Approved:
		return "APPROVED"
	case Rejected:
		return "REJECTED"
	case Pending:
		return "PENDING"
	case Canceled:
		return "CANCELED"
	}
	return ""
}

type SuggestionsService struct {
	client *Client
}

// List performs the API call outlined here
// https://hitchhikers.yext.com/docs/knowledgeapis/knowledgegraph/suggestions/
func (l *SuggestionsService) List(opts *SuggestionsListOptions) (*SuggestionsResponse, *Response, error) {
	var (
		requrl = suggestionsPath
		err    error
	)

	if opts != nil {
		requrl, err = addSuggestionListOptions(requrl, opts)
		if err != nil {
			return nil, nil, err
		}
	}

	if opts != nil {
		requrl, err = addListOptions(requrl, &opts.ListOptions)
		if err != nil {
			return nil, nil, err
		}
	}

	v := &SuggestionsResponse{}
	r, err := l.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}

	return v, r, nil
}

// addSuggestionListOptions adds options to query that are specific to the listings API
func addSuggestionListOptions(requrl string, opts *SuggestionsListOptions) (string, error) {
	if opts == nil {
		return requrl, nil
	}

	u, err := url.Parse(requrl)
	if err != nil {
		return "", err
	}

	q := u.Query()
	if len(opts.EntityIds) > 0 {
		q.Add("entityIds", strings.Join(opts.EntityIds, ","))
	}

	if len(opts.Statuses) > 0 {
		var statuses []string
		for _, status := range opts.Statuses {
			statuses = append(statuses, status.ToString())
		}
		q.Add("statuses", strings.Join(statuses, ","))
	}

	if opts.Format != RichTextFormatDefault {
		q.Add("format", opts.Format.ToString())
	}

	if len(opts.EntityUIDs) > 0 {
		q.Add("entityUids", strings.Join(opts.EntityUIDs, ","))
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}
