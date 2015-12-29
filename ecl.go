package yext

type ListECLResponse struct {
	AllLists     []ECL
	ProductLists []ProductsECL
	BioLists     []BiosECL
	EventsLists  []EventsECL
}

type ECL struct {
	Id       string       `json:"id"`
	Name     string       `json:"name,omitempty"`  // max length 100
	Title    string       `json:"title,omitempty"` // max length 100
	EclType  string       `json:"type,omitempty"`  // one of MENU, BIOS, PRODUCTS, EVENTS
	Size     int          `json:"size,omitempty"`  // read only
	Publish  bool         `json:"publish"`
	Currency string       `json:"currency,omitempty"` //ISO Code for currency
	Sections []ECLSection `json:"sections,omitempty"`
}

type ECLSection struct {
	Id          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ECLItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
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

// ***************** Events ***************************
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

// *****************************************************

// ***************** Product ***************************
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

// *****************************************************

// ***************** Bio ***************************
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

// *****************************************************
