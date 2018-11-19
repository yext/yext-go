package yext

func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

func GetBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func String(v string) *string {
	p := new(string)
	*p = v
	return p
}

func GetString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func Float(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}

func GetFloat(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

func GetInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func Strings(v []string) *[]string {
	return &v
}

func GetStrings(v *[]string) []string {
	if v == nil {
		return []string{}
	}
	return *v
}

func ToUnorderedStrings(v []string) *UnorderedStrings {
	u := UnorderedStrings(v)
	return &u
}

func ToGoogleAttributes(v []*GoogleAttribute) *GoogleAttributes {
	u := GoogleAttributes(v)
	return &u
}

func ToHolidayHours(v []HolidayHours) *[]HolidayHours {
	return &v
}
