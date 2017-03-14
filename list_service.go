package yext

import "fmt"

const (
	menusPath    = "menus"
	biosPath     = "bios"
	productsPath = "products"
	eventsPath   = "events"
)

var (
	ListListMaxLimit = 50
)

type ListService struct {
	client *Client
}

type ProductListsResponse struct {
	Count        int             `json:"count"`
	ProductLists []*ProductsList `json:"products"`
}

type BioListsResponse struct {
	Count    int         `json:"count"`
	BioLists []*BiosList `json:"bios"`
}

type EventListsResponse struct {
	Count      int           `json:"count"`
	EventLists []*EventsList `json:"events"`
}

type MenuListsResponse struct {
	Count     int          `json:"count"`
	MenuLists []*MenusList `json:"menus"`
}

// TODO: Genericize the List[type]Lists/ListAll[type]Lists endpoints in this
// service, they are basically clones of each other

func (e *ListService) ListAllProductLists() ([]*ProductsList, error) {
	var productLists []*ProductsList
	var lr listRetriever = func(opts *ListOptions) (int, int, error) {
		plr, _, err := e.ListProductLists(opts)
		if err != nil {
			return 0, 0, err
		}
		productLists = append(productLists, plr.ProductLists...)
		return len(plr.ProductLists), plr.Count, err
	}

	if err := listHelper(lr, ListListMaxLimit); err != nil {
		return nil, err
	} else {
		return productLists, nil
	}
}

func (e *ListService) ListProductLists(opts *ListOptions) (*ProductListsResponse, *Response, error) {
	requrl, err := addListOptions(productsPath, opts)
	if err != nil {
		return nil, nil, err
	}

	v := &ProductListsResponse{}
	r, err := e.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (e *ListService) ListAllBioLists() ([]*BiosList, error) {
	var bioLists []*BiosList
	var lr listRetriever = func(opts *ListOptions) (int, int, error) {
		plr, _, err := e.ListBioLists(opts)
		if err != nil {
			return 0, 0, err
		}
		bioLists = append(bioLists, plr.BioLists...)
		return len(plr.BioLists), plr.Count, err
	}

	if err := listHelper(lr, ListListMaxLimit); err != nil {
		return nil, err
	} else {
		return bioLists, nil
	}
}

func (e *ListService) ListBioLists(opts *ListOptions) (*BioListsResponse, *Response, error) {
	requrl, err := addListOptions(biosPath, opts)
	if err != nil {
		return nil, nil, err
	}

	v := &BioListsResponse{}
	r, err := e.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (e *ListService) ListAllEventLists() ([]*EventsList, error) {
	var eventLists []*EventsList
	var lr listRetriever = func(opts *ListOptions) (int, int, error) {
		plr, _, err := e.ListEventLists(opts)
		if err != nil {
			return 0, 0, err
		}
		eventLists = append(eventLists, plr.EventLists...)
		return len(plr.EventLists), plr.Count, err
	}

	if err := listHelper(lr, ListListMaxLimit); err != nil {
		return nil, err
	} else {
		return eventLists, nil
	}
}

func (e *ListService) ListEventLists(opts *ListOptions) (*EventListsResponse, *Response, error) {
	requrl, err := addListOptions(eventsPath, opts)
	if err != nil {
		return nil, nil, err
	}

	v := &EventListsResponse{}
	r, err := e.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (e *ListService) ListAllMenuLists() ([]*MenusList, error) {
	var menuLists []*MenusList
	var lr listRetriever = func(opts *ListOptions) (int, int, error) {
		plr, _, err := e.ListMenuLists(opts)
		if err != nil {
			return 0, 0, err
		}
		menuLists = append(menuLists, plr.MenuLists...)
		return len(plr.MenuLists), plr.Count, err
	}

	if err := listHelper(lr, ListListMaxLimit); err != nil {
		return nil, err
	} else {
		return menuLists, nil
	}
}

func (e *ListService) ListMenuLists(opts *ListOptions) (*MenuListsResponse, *Response, error) {
	requrl, err := addListOptions(menusPath, opts)
	if err != nil {
		return nil, nil, err
	}

	v := &MenuListsResponse{}
	r, err := e.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (e *ListService) CreateProductList(y *ProductsList) (*Response, error) {
	return e.client.DoRequestJSON("POST", fmt.Sprintf("%s", productsPath), y, nil)
}

func (e *ListService) CreateBioList(y *BiosList) (*Response, error) {
	return e.client.DoRequestJSON("POST", fmt.Sprintf("%s", biosPath), y, nil)
}

func (e *ListService) CreateEventList(y *EventsList) (*Response, error) {
	return e.client.DoRequestJSON("POST", fmt.Sprintf("%s", eventsPath), y, nil)
}

func (e *ListService) CreateMenuList(y *MenusList) (*Response, error) {
	return e.client.DoRequestJSON("POST", fmt.Sprintf("%s", menusPath), y, nil)
}

func (e *ListService) EditProductList(y *ProductsList) (*ProductsList, *Response, error) {
	var v ProductsList
	r, err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", productsPath, y.GetId()), y, &v)
	if err != nil {
		return nil, r, err
	}
	return &v, r, nil
}

func (e *ListService) EditBioList(y *BiosList) (*BiosList, *Response, error) {
	var v BiosList
	r, err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", biosPath, y.GetId()), y, &v)
	if err != nil {
		return nil, r, err
	}
	return &v, r, nil
}

func (e *ListService) EditEventList(y *EventsList) (*EventsList, *Response, error) {
	var v EventsList
	r, err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", eventsPath, y.GetId()), y, &v)
	if err != nil {
		return nil, r, err
	}
	return &v, r, nil
}

func (e *ListService) EditMenuList(y *MenusList) (*MenusList, *Response, error) {
	var v MenusList
	r, err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", menusPath, y.GetId()), y, &v)
	if err != nil {
		return nil, r, err
	}
	return &v, r, nil
}

func (e *ListService) GetProductList(id string) (*ProductsList, *Response, error) {
	var v ProductsList
	r, err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", productsPath, id), &v)
	if err != nil {
		return nil, r, err
	}
	return &v, r, nil
}

func (e *ListService) GetEventList(id string) (*EventsList, *Response, error) {
	var v EventsList
	r, err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", eventsPath, id), &v)
	if err != nil {
		return nil, r, err
	}
	return &v, r, nil
}

func (e *ListService) GetBioList(id string) (*BiosList, *Response, error) {
	var v BiosList
	r, err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", biosPath, id), &v)
	if err != nil {
		return nil, r, err
	}
	return &v, r, nil
}

func (e *ListService) GetMenuList(id string) (*MenusList, *Response, error) {
	var v MenusList
	r, err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", menusPath, id), &v)
	if err != nil {
		return nil, r, err
	}
	return &v, r, nil
}
