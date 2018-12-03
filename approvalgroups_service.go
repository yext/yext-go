package yext

import (
	"fmt"
)

const approvalGroupsPath = "approvalgroups"

type ApprovalGroupsService struct {
	client *Client
}

type ApprovalGroupsListReponse struct {
	Count          int              `json:"count"`
	ApprovalGroups []*ApprovalGroup `json:"approvalGroups"`
	NextPageToken  string           `json:"nextPageToken"`
}

func (a *ApprovalGroupsService) ListAll() ([]*ApprovalGroup, error) {
	var approvalGroups []*ApprovalGroup

	var al tokenListRetriever = func(opts *ListOptions) (string, error) {
		alr, _, err := a.List(opts)
		if err != nil {
			return "", err
		}
		approvalGroups = append(approvalGroups, alr.ApprovalGroups...)
		return alr.NextPageToken, err
	}

	if err := tokenListHelper(al, nil); err != nil {
		return nil, err
	} else {
		return approvalGroups, nil
	}
}

func (a *ApprovalGroupsService) List(opts *ListOptions) (*ApprovalGroupsListReponse, *Response, error) {
	requrl, err := addListOptions(approvalGroupsPath, opts)
	if err != nil {
		return nil, nil, err
	}

	v := &ApprovalGroupsListReponse{}
	r, err := a.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (a *ApprovalGroup) pathToApprovalGroup() string {
	return pathToApprovalGroupId(a.GetId())
}

func pathToApprovalGroupId(id string) string {
	return fmt.Sprintf("%s/%s", approvalGroupsPath, id)
}

func (a *ApprovalGroupsService) Get(id string) (*ApprovalGroup, *Response, error) {
	var v = &ApprovalGroup{}
	r, err := a.client.DoRequest("GET", pathToApprovalGroupId(id), v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (a *ApprovalGroupsService) Edit(y *ApprovalGroup) (*Response, error) {
	return a.client.DoRequestJSON("PUT", y.pathToApprovalGroup(), y, nil)
}

func (a *ApprovalGroupsService) Create(y *ApprovalGroup) (*Response, error) {
	return a.client.DoRequestJSON("POST", y.pathToApprovalGroup(), y, nil)
}

func (a *ApprovalGroupsService) Delete(y *ApprovalGroup) (*Response, error) {
	return a.client.DoRequest("DELETE", y.pathToApprovalGroup(), nil)
}
