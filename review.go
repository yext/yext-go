package yext

import "encoding/json"

type Reviewer struct {
	LocationId *string       `json:"locationId,omitempty"`
	Entity     *ReviewEntity `json:"entity,omitempty"` // Must have v param >= 20210728
	FirstName  *string       `json:"firstName,omitempty"`
	LastName   *string       `json:"lastName,omitempty"`
	Contact    *string       `json:"contact,omitempty"`
	Image      *bool         `json:"image,omitempty"`
	TemplateId *string       `json:"templateId,omitempty"`
	LabelIds   []*string     `json:"labelIds,omitempty"`
}

type ReviewEntity struct {
	Id string `json:"id"`
}

func NullableReviewEntity(v ReviewEntity) **ReviewEntity {
	p := &v
	return &p
}

type Review struct {
	Id                 *int           `json:"id"`
	LocationId         *string        `json:"locationId"`
	AccountId          *string        `json:"accountId"`
	PublisherId        *string        `json:"publisherId"`
	Rating             *float64       `json:"rating"`
	Title              *string        `json:"title"`
	Content            *string        `json:"content"`
	AuthorName         *string        `json:"authorName"`
	AuthorEmail        *string        `json:"authorEmail"`
	URL                *string        `json:"url"`
	PublisherDate      *int           `json:"publisherDate"`
	LastYextUpdateTime *int           `json:"lastYextUpdateTime"`
	Status             *string        `json:"status"`
	FlagStatus         *string        `json:"flagStatus"`
	ReviewLanguage     *string        `json:"reviewLanguage"`
	Comments           *[]Comment     `json:"comments"`
	LabelIds           *[]int         `json:"labelIds"`
	ExternalId         *string        `json:"externalId"` // Must have v param >= 20220120
	ReviewLabels       *[]ReviewLabel `json:"reviewLabels"`
	ReviewType         *string        `json:"reviewType"`
	Recommendation     *string        `json:"recommendation"`
	TransactionId      *string        `json:"transactionId"`
	InvitationId       *string        `json:"invitationId"`
}

type ReviewCreate struct {
	LocationId     *string  `json:"locationId,omitempty"`
	ExternalId     *string  `json:"externalId,omitempty"` // Must have v param >= 20220120
	AccountId      *string  `json:"accountId,omitempty"`
	Rating         *float64 `json:"rating,omitempty"`
	Content        *string  `json:"content,omitempty"`
	AuthorName     *string  `json:"authorName,omitempty"`
	AuthorEmail    *string  `json:"authorEmail,omitempty"`
	Status         *string  `json:"status,omitempty"`
	FlagStatus     *string  `json:"flagStatus,omitempty"`
	ReviewLanguage *string  `json:"reviewLanguage,omitempty"`
	TransactionId  *string  `json:"transactionId,omitempty"`
	Date           *string  `json:"date,omitempty"`

	//LiveAPI Specific Fields
	ReviewEntity     **ReviewEntity `json:"entity,omitempty"`
	Title            *string        `json:"title,omitempty"`
	ReviewLabelNames *[]string      `json:"reviewLabelNames,omitempty"`
	InvitationUID    *string        `json:"invitationUid,omitempty"`
	ReviewDate       *string        `json:"reviewDate,omitempty"`
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

func (y Review) GetAccountId() string {
	if y.AccountId != nil {
		return *y.AccountId
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

func (y Review) GetLastYextUpdateTime() int {
	if y.LastYextUpdateTime != nil {
		return *y.LastYextUpdateTime
	}
	return 0
}

func (y Review) GetStatus() string {
	if y.Status != nil {
		return *y.Status
	}
	return ""
}

func (y Review) GetFlagStatus() string {
	if y.FlagStatus != nil {
		return *y.FlagStatus
	}
	return ""
}

func (y Review) GetReviewLanguage() string {
	if y.ReviewLanguage != nil {
		return *y.ReviewLanguage
	}
	return ""
}

func (y Review) GetComments() (v []Comment) {
	if y.Comments != nil {
		v = *y.Comments
	}
	return v
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

func (y Review) GetReviewType() string {
	if y.ReviewType != nil {
		return *y.ReviewType
	}
	return ""
}

func (y Review) GetRecommendation() string {
	if y.Recommendation != nil {
		return *y.Recommendation
	}
	return ""
}

func (y Review) GetTransactionId() string {
	if y.TransactionId != nil {
		return *y.TransactionId
	}
	return ""
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
