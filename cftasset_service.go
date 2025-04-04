package yext

import (
	"encoding/json"
	"fmt"
)

const assetsPath = "assets"

var AssetListMaxLimit = 1000

type CFTAssetService struct {
	client   *Client
	registry Registry
}

type CFTAssetListResponse struct {
	Count  int         `json:"count"`
	Assets []*CFTAsset `json:"assets"`
}

func (a *CFTAssetService) RegisterDefaultAssetValues() {
	a.registry = make(Registry)
	// ASSETTYPE_TEXT doesn't need to register because it's just a string
	a.RegisterAssetValue(ASSETTYPE_IMAGE, &ImageValue{})
	a.RegisterAssetValue(ASSETTYPE_VIDEO, &VideoValue{})
	a.RegisterAssetValue(ASSETTYPE_COMPLEXIMAGE, &ComplexImageValue{})
	a.RegisterAssetValue(ASSETTYPE_COMPLEXVIDEO, &ComplexVideoValue{})
}

func (a *CFTAssetService) RegisterAssetValue(t AssetType, assetValue interface{}) {
	a.registry.Register(string(t), assetValue)
}

func (a *CFTAssetService) CreateAssetValue(t AssetType) (interface{}, error) {
	return a.registry.Initialize(string(t))
}

func (a *CFTAssetService) toAssetsWithValues(assets []*CFTAsset) error {
	for _, asset := range assets {
		if err := a.toAssetWithValue(asset); err != nil {
			return err
		}
	}
	return nil
}

func (a *CFTAssetService) toAssetWithValue(asset *CFTAsset) error {
	if asset.Type == ASSETTYPE_TEXT {
		asset.Value = TextValue(asset.Value.(string))
		return nil
	}
	var assetValueValsByKey = asset.Value.(map[string]interface{})
	assetValueObj, err := a.CreateAssetValue(asset.Type)
	if err != nil {
		return err
	}

	// Convert into struct of Asset Value Type
	assetValueJSON, err := json.Marshal(assetValueValsByKey)
	if err != nil {
		return fmt.Errorf("Marshaling asset value to JSON: %s", err)
	}

	err = json.Unmarshal(assetValueJSON, &assetValueObj)
	if err != nil {
		return fmt.Errorf("Unmarshaling asset value to JSON: %s", err)
	}
	asset.Value = assetValueObj
	return nil
}

func (a *CFTAssetService) ListAll() ([]*CFTAsset, error) {
	var assets []*CFTAsset
	var lr listRetriever = func(opts *ListOptions) (int, int, error) {
		alr, _, err := a.List(opts)
		if err != nil {
			return 0, 0, err
		}

		if err = a.toAssetsWithValues(alr.Assets); err != nil {
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

func (a *CFTAssetService) List(opts *ListOptions) (*CFTAssetListResponse, *Response, error) {
	requrl, err := addListOptions(assetsPath, opts)
	if err != nil {
		return nil, nil, err
	}
	var v CFTAssetListResponse
	r, err := a.client.DoRequest("GET", requrl, &v)
	if err != nil {
		return nil, r, err
	}

	return &v, r, nil
}

func (a *CFTAssetService) Create(asset *CFTAsset) (*Response, error) {
	r, err := a.client.DoRequestJSON("POST", fmt.Sprintf("%s", assetsPath), asset, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (a *CFTAssetService) Get(assetId string) (*CFTAsset, *Response, error) {
	var v *CFTAsset
	r, err := a.client.DoRequest("GET", fmt.Sprintf("%s/%s", assetsPath, assetId), &v)
	if err != nil {
		return nil, r, err
	}

	if err := a.toAssetWithValue(v); err != nil {
		return nil, r, err
	}

	return v, r, nil
}

func (a *CFTAssetService) Update(assetId string, asset *CFTAsset) (*Response, error) {
	r, err := a.client.DoRequestJSON("PUT", fmt.Sprintf("%s/%s", assetsPath, assetId), asset, nil)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (a *CFTAssetService) Delete(assetId string) (*Response, error) {
	r, err := a.client.DoRequest("DELETE", fmt.Sprintf("%s/%s", assetsPath, assetId), nil)
	if err != nil {
		return r, err
	}

	return r, nil
}
