package yext

func (a *CFTAsset) Diff(b *CFTAsset) (*CFTAsset, bool) {
	if a.GetAssetType() != b.GetAssetType() {
		return nil, true
	}

	delta, isDiff := diff(a, b, true, true)
	if !isDiff {
		return nil, isDiff
	}
	return delta.(*CFTAsset), isDiff
}
