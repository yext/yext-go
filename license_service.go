package yext

import "fmt"

const licensePacksPath = "licensePacks"

type LicenseService struct {
	client *Client
}

func (l *LicenseService) Get(id string) (*LicensePack, error) {
	var v LicensePack
	err := l.client.DoRequest("GET", fmt.Sprintf("%s/%s", licensePacksPath, id), &v)
	return &v, err
}

func (l *LicenseService) AddLocationToLicense(licenseId string, locationId string) (*LicensePack, error) {
	var v LicensePack
	err := l.client.DoRequest("PUT", fmt.Sprintf("%s/%s/locationIds/%s", licensePacksPath, licenseId, locationId), &v)
	return &v, err
}

func (l *LicenseService) RemoveLocationFromLicense(licenseId string, locationId string) (*LicensePack, error) {
	var v LicensePack
	err := l.client.DoRequest("DELETE", fmt.Sprintf("%s/%s/locationIds/%s", licensePacksPath, licenseId, locationId), &v)
	return &v, err
}
