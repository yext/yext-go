package yext

import "encoding/json"

type Reviewer struct {
	LocationId *string   `json:"locationId,omitempty"`
	FirstName  *string   `json:"firstName,omitempty"`
	LastName   *string   `json:"lastName,omitempty"`
	Contact    *string   `json:"contact,omitempty"`
	Image      *bool     `json:"image,omitempty"`
	TemplateId *string   `json:"templateId,omitempty"`
	LabelIds   []*string `json:"labelIds,omitempty"`
}

type Review struct {
	Id                 *int           `json:"id"`
	LocationId         *string        `json:"locationId"`
	PublisherId        *string        `json:"publisherId"`
	Rating             *float64       `json:"rating"`
	Title              *string        `json:"title"`
	Content            *string        `json:"content"`
	AuthorName         *string        `json:"authorName"`
	AuthorEmail        *string        `json:"authorEmail"`
	URL                *string        `json:"url"`
	PublisherDate      *int           `json:"publisherDate"`
	LastYextUpdateDate *int           `json:"lastYextUpdateDate"`
	Status             *string        `json:"status"`
	Comments           *[]Comment     `json:"comments"`
	LabelIds           *[]int         `json:"labelIds"`
	ExternalId         *string        `json:"externalId"`
	ReviewLabels       *[]ReviewLabel `json:"reviewLabels"`
	InvitationId       *string        `json:"invitationId"`
}

type ReviewCreate struct {
	LocationId     *string  `json:"locationId"`
	AccountId      *string  `json:"accountId"`
	Rating         *float64 `json:"rating"`
	Content        *string  `json:"content"`
	AuthorName     *string  `json:"authorName"`
	AuthorEmail    *string  `json:"authorEmail,omitempty"`
	Status         *string  `json:"status,omitempty"`
	FlagStatus     *string  `json:"flagStatus,omitempty"`
	ReviewLanguage *string  `json:"reviewLanguage,omitempty"`
	TransactionId  *string  `json:"transactionId,omitempty"`
	Date           *string  `json:"date,omitempty"`
}

type Comment struct {
	Id            *int    `json:"id"`
	ParentId      *int    `json:"parentId"`
	PublisherDate *int    `json:"publisherDate"`
	AuthorName    *string `json:"authorName"`
	AuthorEmail   *string `json:"authorEmail"`
	AuthorRole    *string `json:"authorRole"`
	Content       *string `json:"content"`
	Visibility    *string `json:"visibility"`
}

type ReviewLabel struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
}

func (y Review) GetId() int {
	if y.Id != nil {
		return *y.Id
	}
	return 0
}

func (y Review) GetLocationId() string {
	if y.LocationId != nil {
		return *y.LocationId
	}
	return ""
}

func (y Review) GetPublisherId() string {
	if y.PublisherId != nil {
		return *y.PublisherId
	}
	return ""
}

func (y Review) GetRating() float64 {
	if y.Rating != nil {
		return *y.Rating
	}
	return -1
}

func (y Review) GetTitle() string {
	if y.Title != nil {
		return *y.Title
	}
	return ""
}

func (y Review) GetContent() string {
	if y.Content != nil {
		return *y.Content
	}
	return ""
}

func (y Review) GetAuthorName() string {
	if y.AuthorName != nil {
		return *y.AuthorName
	}
	return ""
}

func (y Review) GetAuthorEmail() string {
	if y.AuthorEmail != nil {
		return *y.AuthorEmail
	}
	return ""
}

func (y Review) GetURL() string {
	if y.URL != nil {
		return *y.URL
	}
	return ""
}

func (y Review) GetPublisherDate() int {
	if y.PublisherDate != nil {
		return *y.PublisherDate
	}
	return 0
}

func (y Review) GetLastYextUpdateDate() int {
	if y.LastYextUpdateDate != nil {
		return *y.LastYextUpdateDate
	}
	return 0
}

func (y Review) GetStatus() string {
	if y.Status != nil {
		return *y.Status
	}
	return ""
}

func (y Review) GetLabelIds() (v []int) {
	if y.LabelIds != nil {
		v = *y.LabelIds
	}
	return v
}

func (y Review) GetExternalId() string {
	if y.ExternalId != nil {
		return *y.ExternalId
	}
	return ""
}

func (y Review) GetReviewLabels() (v []ReviewLabel) {
	if y.ReviewLabels != nil {
		v = *y.ReviewLabels
	}
	return v
}

func (y Review) GetComments() (v []Comment) {
	if y.Comments != nil {
		v = *y.Comments
	}
	return v
}

func (y Comment) GetId() int {
	if y.Id != nil {
		return *y.Id
	}
	return 0
}

func (y Comment) GetParentId() int {
	if y.ParentId != nil {
		return *y.ParentId
	}
	return 0
}

func (y Comment) GetPublisherDate() int {
	if y.PublisherDate != nil {
		return *y.PublisherDate
	}
	return 0
}

func (y Comment) GetAuthorName() string {
	if y.AuthorName != nil {
		return *y.AuthorName
	}
	return ""
}

func (y Comment) GetAuthorEmail() string {
	if y.AuthorEmail != nil {
		return *y.AuthorEmail
	}
	return ""
}

func (y Comment) GetAuthorRole() string {
	if y.AuthorRole != nil {
		return *y.AuthorRole
	}
	return ""
}

func (y Comment) GetContent() string {
	if y.Content != nil {
		return *y.Content
	}
	return ""
}

func (y Comment) GetVisibility() string {
	if y.Visibility != nil {
		return *y.Visibility
	}
	return ""
}

func (y ReviewLabel) GetId() int {
	if y.Id != nil {
		return *y.Id
	}
	return 0
}

func (y ReviewLabel) GetName() string {
	if y.Name != nil {
		return *y.Name
	}
	return ""
}

func (y ReviewCreate) GetLocationId() string {
	if y.LocationId != nil {
		return *y.LocationId
	}
	return ""
}

func (y ReviewCreate) GetAccountId() string {
	if y.AccountId != nil {
		return *y.AccountId
	}
	return ""
}

func (y ReviewCreate) GetRating() float64 {
	if y.Rating != nil {
		return *y.Rating
	}
	return 0
}

func (y ReviewCreate) GetContent() string {
	if y.Content != nil {
		return *y.Content
	}
	return ""
}

func (y ReviewCreate) GetAuthorName() string {
	if y.AuthorName != nil {
		return *y.AuthorName
	}
	return ""
}

func (y ReviewCreate) GetAuthorEmail() string {
	if y.AuthorEmail != nil {
		return *y.AuthorEmail
	}
	return ""
}

func (y ReviewCreate) GetStatus() string {
	if y.Status != nil {
		return *y.Status
	}
	return ""
}

func (y ReviewCreate) GetFlagStatus() string {
	if y.FlagStatus != nil {
		return *y.FlagStatus
	}
	return ""
}

func (y ReviewCreate) GetReviewLanguage() string {
	if y.ReviewLanguage != nil {
		return *y.ReviewLanguage
	}
	return ""
}

func (y ReviewCreate) GetTransactionId() string {
	if y.TransactionId != nil {
		return *y.TransactionId
	}
	return ""
}

func (y ReviewCreate) GetDate() string {
	if y.Date != nil {
		return *y.Date
	}
	return ""
}

func (y ReviewCreate) String() string {
	b, _ := json.Marshal(y)
	return string(b)
}
