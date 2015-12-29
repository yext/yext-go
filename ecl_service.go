package yext

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const listsPath = "lists"

type ECLService struct {
	client *Client
}

func (e *ECLService) Create(y *ECL) (*ECL, error) {
	var v ECL
	err := e.client.DoRequestJSON("POST", fmt.Sprintf("%s", listsPath), y, &v)
	return &v, err
}

func (e *ECLService) CreateProductList(y *ProductsECL) (*ProductsECL, error) {
	var v ProductsECL
	err := e.client.DoRequestJSON("POST", fmt.Sprintf("%s", listsPath), y, &v)
	return &v, err
}

func (e *ECLService) CreateBioList(y *BiosECL) (*BiosECL, error) {
	var v BiosECL
	err := e.client.DoRequestJSON("POST", fmt.Sprintf("%s", listsPath), y, &v)
	return &v, err
}

func (e *ECLService) CreateEventList(y *EventsECL) (*EventsECL, error) {
	var v EventsECL
	err := e.client.DoRequestJSON("POST", fmt.Sprintf("%s", listsPath), y, &v)
	return &v, err
}

func (e *ECLService) Edit(y *ECL) (*ECL, error) {
	var v ECL
	err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", listsPath, y.Id), y, &v)
	return &v, err
}

func (e *ECLService) EditProductList(y *ProductsECL) (*ProductsECL, error) {
	var v ProductsECL
	err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", listsPath, y.Id), y, &v)
	return &v, err
}

func (e *ECLService) EditBioList(y *BiosECL) (*BiosECL, error) {
	var v BiosECL
	err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", listsPath, y.Id), y, &v)
	return &v, err
}

func (e *ECLService) EditEventList(y *EventsECL) (*EventsECL, error) {
	var v EventsECL
	err := e.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", listsPath, y.Id), y, &v)
	return &v, err
}

type genericECLResponse struct {
	Lists []ECL `json:"lists"`
}

func (e *ECLService) List() (*ListECLResponse, error) {
	var buf bytes.Buffer
	err := e.client.DoRequest("GET", listsPath, &buf)
	if err != nil {
		return nil, err
	}
	return parseECLResponse(buf.Bytes())
}

func parseECLResponse(buf []byte) (*ListECLResponse, error) {
	var (
		gen = &genericECLResponse{}
		res = &ListECLResponse{}
		raw map[string][]map[string]interface{}
	)

	err := json.Unmarshal(buf, gen)
	if err != nil {
		return nil, err
	}
	res.AllLists = gen.Lists

	err = json.Unmarshal(buf, &raw)
	if err != nil {
		return nil, err
	}

	for _, ecl := range raw["lists"] {
		reMarshalled, err := json.Marshal(ecl)
		if err != nil {
			return nil, err
		}

		switch ecl["type"] {
		case "PRODUCTS":
			var prod ProductsECL
			err = json.Unmarshal(reMarshalled, &prod)
			if err != nil {
				return nil, err
			}
			res.ProductLists = append(res.ProductLists, prod)
		case "BIOS":
			var bio BiosECL
			err = json.Unmarshal(reMarshalled, &bio)
			if err != nil {
				return nil, err
			}
			res.BioLists = append(res.BioLists, bio)
		case "EVENTS":
			var event EventsECL
			err = json.Unmarshal(reMarshalled, &event)
			if err != nil {
				return nil, err
			}
			res.EventsLists = append(res.EventsLists, event)
		}
	}

	return res, nil
}

func (e *ECLService) GetProductList(id string) (*ProductsECL, error) {
	var v ProductsECL
	err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", listsPath, id), &v)
	return &v, err
}

func (e *ECLService) GetEventList(id string) (*EventsECL, error) {
	var v EventsECL
	err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", listsPath, id), &v)
	return &v, err
}

func (e *ECLService) GetBioList(id string) (*BiosECL, error) {
	var v BiosECL
	err := e.client.DoRequest("GET", fmt.Sprintf("%s/%s", listsPath, id), &v)
	return &v, err
}
