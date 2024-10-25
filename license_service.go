package yext

import "fmt"

const (
	packsPath = "license-packs"
)

type LicenseService struct {
	client *Client
}

func (u *LicenseService) ListPacks(opts *ListOptions) (*LicensePacksListResponse, *Response, error) {
	if opts != nil {
		opts.usePageSize = true
	}

	requrl, err := addListOptions(packsPath, opts)
	if err != nil {
		return nil, nil, err
	}

	v := &LicensePacksListResponse{}
	r, err := u.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (u *LicenseService) ListAllPacks() ([]*LicensePack, error) {
	var licensePacks []*LicensePack

	var al tokenListRetriever = func(opts *ListOptions) (string, error) {
		lpr, _, err := u.ListPacks(opts)
		if err != nil {
			return "", err
		}
		licensePacks = append(licensePacks, lpr.LicensePacks...)
		return lpr.NextPageToken, err
	}

	opts := ListOptions{}
	if err := tokenListHelper(al, &opts); err != nil {
		return nil, err
	} else {
		return licensePacks, nil
	}
}

func (u *LicenseService) ListAssignments(licensePackId string, opts *ListOptions) (*LicenseAssignmentListResponse, *Response, error) {
	if opts != nil {
		opts.usePageSize = true
	}
	requrl, err := addListOptions(fmt.Sprintf("%s/%s/assigned-entities", packsPath, licensePackId), opts)
	if err != nil {
		return nil, nil, err
	}

	v := &LicenseAssignmentListResponse{}
	r, err := u.client.DoRequest("GET", requrl, v)
	if err != nil {
		return nil, r, err
	}
	return v, r, nil
}

func (u *LicenseService) ListAllAssignments(licensePackId string) (*LicenseAssignment, error) {
	var licenseAssignments = &LicenseAssignment{}

	var al tokenListRetriever = func(opts *ListOptions) (string, error) {
		lar, _, err := u.ListAssignments(licensePackId, opts)
		if err != nil {
			return "", err
		}

		if len(lar.LicenseAssignments) == 0 {
			return "", fmt.Errorf("error: no license assignments found for id %s", licensePackId)
		}

		licenseAssignments.LicensePacks = lar.LicenseAssignments[0].LicensePacks
		licenseAssignments.AssignedEntities = append(licenseAssignments.AssignedEntities, lar.LicenseAssignments[0].AssignedEntities...)
		
		return lar.NextPageToken, err
	}

	opts := ListOptions{}
	if err := tokenListHelper(al, &opts); err != nil {
		return nil, err
	} else {
		return licenseAssignments, nil
	}
}

func (u *LicenseService) EditAssignment(l *AssignedEntity) (*Response, error) {
	return u.client.DoRequestJSON("PUT", l.pathToAssignment(), l, nil)
}

func (u *LicenseService) CreateAssignment(l *AssignedEntity) (*Response, error) {
	return u.client.DoRequestJSON("POST", l.pathToAssignment(), nil, nil)
}

func (u *LicenseService) DeleteAssignment(l *AssignedEntity) (*Response, error) {
	return u.client.DoRequest("DELETE", l.pathToAssignment(), nil)
}
