package yext

type Reviewer struct {
	LocationId *string `json:"locationId"`
	FirstName  *string `json:"firstName"`
	LastName   *string `json:"lastName"`
	Contact    *string `json:"contact"`
	Image      *bool   `json:"image"`
	TemplateId *string `json:"templateId"`
}

type Review struct {
	Id                 *int       `json:"id"`
	LocationId         *string    `json:"locationId"`
	PublisherId        *string    `json:"publisherId"`
	Rating             *float64   `json:"rating"`
	Title              *string    `json:"title"`
	Content            *string    `json:"content"`
	AuthorName         *string    `json:"authorName"`
	AuthorEmail        *string    `json:"authorEmail"`
	URL                *string    `json:"url"`
	PublisherDate      *int       `json:"publisherDate"`
	LastYextUpdateDate *int       `json:"lastYextUpdateDate"`
	Status             *string    `json:"status"`
	Comments           *[]Comment `json:"comments"`
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

func (y Comment) GetId() int {
	if y.Id != nil {
		return *y.Id
	}
	return 0
}

func (y Comment) GetParentId() int {
	if y.Id != nil {
		return *y.ParentId
	}
	return 0
}

func (y Comment) GetPublisherDate() int {
	if y.Id != nil {
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
