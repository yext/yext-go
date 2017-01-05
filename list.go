package yext

import "reflect"

// Generic base of a list
type List struct {
	Id       *string        `json:"id"`
	Name     *string        `json:"name,omitempty"`  // max length 100
	Title    *string        `json:"title,omitempty"` // max length 100
	Type     *string        `json:"type,omitempty"`  // one of MENU, BIOS, PRODUCTS, EVENTS
	Size     *int           `json:"size,omitempty"`  // read only
	Publish  *bool          `json:"publish"`
	Currency *string        `json:"currency,omitempty"` // ISO Code for currency
	Sections []*ListSection `json:"sections,omitempty"`
}

func (e List) GetId() string {
	if e.Id != nil {
		return *e.Id
	}
	return ""
}

func (e List) GetName() string {
	if e.Name != nil {
		return *e.Name
	}
	return ""
}

func (e List) GetTitle() string {
	if e.Title != nil {
		return *e.Title
	}
	return ""
}

func (e List) GetType() string {
	if e.Type != nil {
		return *e.Type
	}
	return ""
}

func (e List) GetSize() int {
	if e.Size != nil {
		return *e.Size
	}
	return -1
}

func (e List) GetPublish() bool {
	if e.Publish != nil {
		return *e.Publish
	}
	return false
}

func (e List) GetCurrency() string {
	if e.Currency != nil {
		return *e.Currency
	}
	return ""
}

type ListSection struct {
	Id          *string `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (e ListSection) GetId() string {
	if e.Id != nil {
		return *e.Id
	}
	return ""
}

func (e ListSection) GetName() string {
	if e.Name != nil {
		return *e.Name
	}
	return ""
}

func (e ListSection) GetDescription() string {
	if e.Description != nil {
		return *e.Description
	}
	return ""
}

type ListItem struct {
	Id          *string `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (e ListItem) GetId() string {
	if e.Id != nil {
		return *e.Id
	}
	return ""
}

func (e ListItem) GetName() string {
	if e.Name != nil {
		return *e.Name
	}
	return ""
}

func (e ListItem) GetDescription() string {
	if e.Description != nil {
		return *e.Description
	}
	return ""
}

type ListPhoto struct {
	Url    string `json:"url"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type Cost struct {
	Type    string         `json:"type,omitempty"`
	Price   string         `json:"price,omitempty"`
	Unit    string         `json:"unit,omitempty"`
	RangeTo string         `json:"rangeTo,omitempty"`
	Other   string         `json:"other,omitempty"`
	Options []*CostOptions `json:"options,omitempty"`
}

type CostOptions struct {
	Name    string `json:"name,omitempty"`
	Price   string `json:"price,omitempty"`
	Calorie int    `json:"calorie,omitempty"`
}

type EventsList struct {
	List
	Sections []*EventsListSection `json:"sections,omitempty"`
}

type EventsListSection struct {
	ListSection
	Items []*Event `json:"items,omitempty"` // max 100 items
}

type Event struct {
	ListItem
	Type   string       `json:"type,omitempty"`
	Starts string       `json:"starts,omitempty"`
	Ends   string       `json:"ends,omitempty"`
	Photos []*ListPhoto `json:"photos,omitempty"`
	Url    string       `json:"url,omitempty"`
}

type ProductsList struct {
	List
	Sections []*ProductsListSection `json:"sections,omitempty"`
}

type ProductsListSection struct {
	ListSection
	Items []*Product `json:"items,omitempty"` // max 100 items
}

type Product struct {
	ListItem
	Type   string       `json:"idcode,omitempty"`
	Cost   *Cost        `json:"cost,omitempty"`
	Photos []*ListPhoto `json:"photos,omitempty"`
	Video  string       `json:"video,omitempty"`
	Url    string       `json:"url,omitempty"`
}

type BiosList struct {
	List
	Sections []*BiosListSection `json:"sections,omitempty"`
}

type BiosListSection struct {
	ListSection
	Items []*Bio `json:"items,omitempty"` // max 100 items
}

type Bio struct {
	ListItem
	Title          string     `json:"title,omitempty"`
	Photo          *ListPhoto `json:"photo,omitempty"`
	PhoneNumber    string     `json:"phone,omitempty"`
	EmailAddress   string     `json:"email,omitempty"`
	Education      []string   `json:"education"`
	Certifications []string   `json:"certifications"`
	Services       []string   `json:"services,omitempty"`
	Url            string     `json:"url,omitempty"`
}

func BioItemCompare(itemA Bio, itemB Bio) bool {
	return reflect.DeepEqual(itemA, itemB)
}
