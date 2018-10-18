package yext

func (a *CFTAsset) Diff(b *CFTAsset) (*CFTAsset, bool) {
	if a == nil && b != nil {
		return b, true
	} else if b == nil && a != nil {
		return a, true
	} else if a == nil && b == nil {
		return nil, false
	}

	if a.GetAssetType() != b.GetAssetType() {
		return nil, true
	}

	delta, isDiff := diff(a, b, true, true)
	if !isDiff {
		return nil, isDiff
	}
	return delta.(*CFTAsset), isDiff
}
