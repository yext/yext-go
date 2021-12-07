package yext

import "fmt"

const createSubAccountPath = "createsubaccount"

type AccountService struct {
	client *Client
}

type CreateSubAccountRequest struct {
	Id          string `json:"newSubAccountId"`
	Name        string `json:"newSubAccountName"`
	CountryCode string `json:"countryCode"`
}

type SubAccount struct {
	AccountId        *string `json:"accountId"`
	LocationCount    *int    `json:"locationCount"`
	SubAccountCount  *int    `json:"subAccountCount"`
	AccountName      *string `json:"accountName"`
	ContactFirstName *string `json:"contactFirstName"`
	ContactLastName  *string `json:"contactLastName"`
	ContactPhone     *string `json:"contactPhone"`
	ContactEmail     *string `json:"contactEmail"`
}

type ListResponse struct {
	Count    *int          `json:"count"`
	Accounts []*SubAccount `json:"accounts"`
}

func (a *AccountService) CreateSubAccount(createSubAccountRequest *CreateSubAccountRequest) (*Response, error) {
	r, err := a.client.DoRequestJSON("POST", fmt.Sprintf("%s", createSubAccountPath), createSubAccountRequest, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (a *AccountService) List(opts *ListOptions) (*ListResponse, *Response, error) {
	arr := &ListResponse{}
	requrl, err := addListOptions(ACCOUNTS_PATH, opts)
	if err != nil {
		return nil, nil, err
	}

	r, err := a.client.DoRootRequest("GET", requrl, arr)
	if err != nil {
		return nil, r, err
	}

	return arr, r, nil
}

func (a *AccountService) Get(accountId string) (*SubAccount, *Response, error) {
	v := &SubAccount{}

	r, err := a.client.DoRootRequest("GET", fmt.Sprintf("%s/%s", ACCOUNTS_PATH, accountId), v)
	if err != nil {
		return nil, r, err
	}

	return v, r, nil
}
