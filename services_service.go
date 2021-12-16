package yext

const createExistingSubAccountPath = "existingsubaccountaddrequest"
const createExistingLocationPath = "existinglocationaddrequest"

type ServicesService struct {
	client *Client
}

type ExistingSubAccountAddRequest struct {
	SubAccountId string        `json:"subAccountId"`
	SkuAdditions []SkuAddition `json:"skuAdditions"`
	AgreementId  string        `json:"agreementId"`
}

type ExistingLocationAddRequest struct {
	ExistingLocationId        string   `json:"existingLocationId"`
	ExistingLocationAccountId string   `json:"existingLocationAccountId"`
	Skus                      []string `json:"skus"`
	AgreementId               string   `json:"agreementId"`
	ForceReview               string   `json:"forceReview"`
}

type SkuAddition struct {
	Sku      string `json:"sku"`
	Quantity string `json:"quantity"`
}

type ExistingSubAccountAddResponse struct {
	Id            int           `json:"id"`
	SubAccountId  string        `json:"subAccountId"`
	SkuAdditions  []SkuAddition `json:"skuAdditions"`
	AgreementId   string        `json:"agreementId"`
	Status        string        `json:"status"`
	DateSubmitted string        `json:"dateSubmitted"`
	StatusDetail  string        `json:"statusDetail"`
}

type ExistingLocationAddResponse struct {
	Id                        int      `json:"id"`
	LocationMode              string   `json:"locationMode"`
	ExistingLocationId        string   `json:"existingLocationId"`
	NewLocationId             string   `json:"newLocationId"`
	NewLocationAccountId      string   `json:"newLocationAccountId"`
	NewLocationAccountName    string   `json:"newLocationAccountName"`
	NewAccountParentAccountId string   `json:"newAccountParentAccountId"`
	NewLocationData           string   `json:"newLocationData"`
	NewEntityData             string   `json:"newEntityData"`
	Skus                      []string `json:"skus"`
	AgreementId               string   `json:"agreementId"`
	Status                    string   `json:"status"`
	DateSubmitted             string   `json:"dateSubmitted"`
	DateCompleted             string   `json:"dateCompleted"`
	StatusDetail              string   `json:"statusDetail"`
}

func (a *ServicesService) CreateAddRequestExistingSubAccount(existingSubAccountAddRequest *ExistingSubAccountAddRequest) (*ExistingSubAccountAddResponse, *Response, error) {
	var v *ExistingSubAccountAddResponse
	r, err := a.client.DoRequest("POST", createExistingSubAccountPath, &v)
	if err != nil {
		return v, r, err
	}

	return v, r, nil
}

func (a *ServicesService) CreateAddRequestExistingLocation(existingLocationAddRequest *ExistingLocationAddRequest) (*ExistingLocationAddResponse, *Response, error) {
	var v *ExistingLocationAddResponse
	r, err := a.client.DoRequest("POST", createExistingLocationPath, &v)
	if err != nil {
		return v, r, err
	}

	return v, r, nil
}
