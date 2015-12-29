package yext

// ListECLResponse is a calculated value that provides typed access to an underlying 'ListECL' response
type ListECLResponse struct {
	AllLists     []ECL
	ProductLists []ProductsECL
	BioLists     []BiosECL
	EventsLists  []EventsECL
}

type ECL struct {
	Id       *string      `json:"id"`
	Name     *string      `json:"name,omitempty"`  // max length 100
	Title    *string      `json:"title,omitempty"` // max length 100
	Type     *string      `json:"type,omitempty"`  // one of MENU, BIOS, PRODUCTS, EVENTS
	Size     *int         `json:"size,omitempty"`  // read only
	Publish  *bool        `json:"publish"`
	Currency *string      `json:"currency,omitempty"` // ISO Code for currency
	Sections []ECLSection `json:"sections,omitempty"`
}

func (e ECL) GetId() string {
	if e.Id != nil {
		return *e.Id
	}
	return ""
}

func (e ECL) GetName() string {
	if e.Name != nil {
		return *e.Name
	}
	return ""
}

func (e ECL) GetTitle() string {
	if e.Title != nil {
		return *e.Title
	}
	return ""
}

func (e ECL) GetType() string {
	if e.Type != nil {
		return *e.Type
	}
	return ""
}

func (e ECL) GetSize() int {
	if e.Size != nil {
		return *e.Size
	}
	return -1
}

func (e ECL) GetPublish() bool {
	if e.Publish != nil {
		return *e.Publish
	}
	return false
}

func (e ECL) GetCurrency() string {
	if e.Currency != nil {
		return *e.Currency
	}
	return ""
}

type ECLSection struct {
	Id          *string `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (e ECLSection) GetId() string {
	if e.Id != nil {
		return *e.Id
	}
	return ""
}

func (e ECLSection) GetName() string {
	if e.Name != nil {
		return *e.Name
	}
	return ""
}

func (e ECLSection) GetDescription() string {
	if e.Description != nil {
		return *e.Description
	}
	return ""
}

type ECLItem struct {
	Id          *string `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
}

func (e ECLItem) GetId() string {
	if e.Id != nil {
		return *e.Id
	}
	return ""
}

func (e ECLItem) GetName() string {
	if e.Name != nil {
		return *e.Name
	}
	return ""
}

func (e ECLItem) GetDescription() string {
	if e.Description != nil {
		return *e.Description
	}
	return ""
}

type Photo struct {
	Url    string `json:"url"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type Cost struct {
	Type    string        `json:"type,omitempty"`
	Price   string        `json:"price,omitempty"`
	Unit    string        `json:"unit,omitempty"`
	RangeTo string        `json:"rangeTo,omitempty"`
	Other   string        `json:"other,omitempty"`
	Options []CostOptions `json:"options,omitempty"`
}

type CostOptions struct {
	Name    string `json:"name,omitempty"`
	Price   string `json:"price,omitempty"`
	Calorie int    `json:"calorie,omitempty"`
}

type EventsECL struct {
	ECL
	Sections []EventsECLSection `json:"sections,omitempty"`
}

type EventsECLSection struct {
	ECLSection
	Items []Event `json:"items,omitempty"` // max 100 items
}

type Event struct {
	ECLItem
	Type   string `json:"type,omitempty"`
	Starts string `json:"starts,omitempty"`
	Ends   string `json:"ends,omitempty"`
	Url    string `json:"url,omitempty"`
}

type ProductsECL struct {
	ECL
	Sections []ProductsECLSection `json:"sections,omitempty"`
}

type ProductsECLSection struct {
	ECLSection
	Items []Product `json:"items,omitempty"` // max 100 items
}

type Product struct {
	ECLItem
	Type   string  `json:"idcode,omitempty"`
	Cost   *Cost   `json:"cost,omitempty"`
	Photos []Photo `json:"photos,omitempty"`
	Video  string  `json:"video,omitempty"`
	Url    string  `json:"url,omitempty"`
}

type BiosECL struct {
	ECL
	Sections []BiosECLSection `json:"sections,omitempty"`
}

type BiosECLSection struct {
	ECLSection
	Items []Bio `json:"items,omitempty"` // max 100 items
}

type Bio struct {
	ECLItem
	Title          string   `json:"idcode,omitempty"`
	Photo          *Photo   `json:"photo,omitempty"`
	Education      []string `json:"education"`
	Certifications []string `json:"certifications"`
	Services       []string `json:"services,omitempty"`
	Url            string   `json:"url,omitempty"`
}
