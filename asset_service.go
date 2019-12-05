package yext

import "fmt"

type AssetService struct {
	client *Client
}

type AssetListResponse struct {
	Count  int      `json:"count"`
	Assets []*Asset `json:"assets"`
}

func (a *AssetService) ListAll() ([]*Asset, error) {
	var assets []*Asset
	var lr listRetriever = func(opts *ListOptions) (int, int, error) {
		alr, _, err := a.List(opts)
		if err != nil {
			return 0, 0, err
		}
		assets = append(assets, alr.Assets...)
		return len(alr.Assets), alr.Count, err
	}

	if err := listHelper(lr, &ListOptions{Limit: AssetListMaxLimit}); err != nil {
		return nil, err
	} else {
		return assets, nil
	}
}

func (a *AssetService) List(opts *ListOptions) (*AssetListResponse, *Response, error) {
	requrl, err := addListOptions(assetsPath, opts)
	if err != nil {
		return nil, nil, err
	}
	var v AssetListResponse
	r, err := a.client.DoRequest("GET", requrl, &v)
	if err != nil {
		return nil, r, err
	}

	return &v, r, nil
}

func (a *AssetService) Create(asset *Asset) (*Response, error) {
	r, err := a.client.DoRequestJSON("POST", fmt.Sprintf("%s", assetsPath), asset, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (a *AssetService) Get(assetId string) (*Asset, *Response, error) {
	var v Asset
	r, err := a.client.DoRequest("GET", fmt.Sprintf("%s/%s", assetsPath, assetId), &v)
	if err != nil {
		return nil, r, err
	}

	return &v, r, nil
}

func (a *AssetService) Update(assetId string, asset *Asset) (*Response, error) {
	r, err := a.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", assetsPath, assetId), asset, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (a *AssetService) Delete(assetId string) (*Response, error) {
	r, err := a.client.DoRequest("DELETE", fmt.Sprintf("%s/%s", assetsPath, assetId), nil)
	if err != nil {
		return r, err
	}

	return r, nil
}
