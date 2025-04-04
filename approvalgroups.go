package yext

type ApprovalGroup struct {
	Id        *string           `json:"id,omitempty"`
	Name      *string           `json:"name,omitempty"`
	IsDefault *bool             `json:"isDefault,omitempty"`
	Users     *UnorderedStrings `json:"users,omitempty"`
}

func (a *ApprovalGroup) GetId() string {
	if a.Id == nil {
		return ""
	}
	return *a.Id
}

func (a *ApprovalGroup) GetName() string {
	if a.Name == nil {
		return ""
	}
	return *a.Name
}

func (a *ApprovalGroup) GetIsDefault() bool {
	if a.IsDefault == nil {
		return false
	}
	return *a.IsDefault
}

func (a *ApprovalGroup) GetUsers() (v UnorderedStrings) {
	if a.Users != nil {
		v = *a.Users
	}
	return v
}
