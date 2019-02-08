package yext

import "fmt"

func NullableBool(v bool) **bool {
	p := &v
	return &p
}

func GetNullableBool(v **bool) bool {
	if v == nil || *v == nil {
		return false
	}
	return **v
}

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

func NullBool() **bool {
	var v *bool
	return &v
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

func NullableString(v string) **string {
	y := &v
	return &y
}

func GetNullableString(v **string) string {
	if v == nil || *v == nil {
		return ""
	}
	return **v
}

func NullString() **string {
	var v *string
	return &v
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

func NullableFloat(v float64) **float64 {
	p := &v
	return &p
}

func GetNullableFloat(v **float64) float64 {
	if v == nil || *v == nil {
		return 0
	}
	return **v
}

func NullFloat() **float64 {
	var v *float64
	return &v
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

func NullableInt(v int) **int {
	p := &v
	return &p
}

func GetNullableInt(v **int) int {
	if v == nil || *v == nil {
		return 0
	}
	return **v
}

func NullInt() **int {
	var v *int
	return &v
}

func NullableVideo(v *Video) **Video {
	return &v
}

func GetVideo(v **Video) *Video {
	if v == nil {
		return nil
	}
	return *v
}

func NullVideo() **Video {
	var v *Video
	return &v
}

func NullablePhoto(v *Photo) **Photo {
	return &v
}

func GetPhoto(v **Photo) *Photo {
	if v == nil {
		return nil
	}
	return *v
}

func NullPhoto() **Photo {
	var v *Photo
	return &v
}

func NullableDailyTimes(v *DailyTimes) **DailyTimes {
	return &v
}

func GetDailyTimes(v **DailyTimes) *DailyTimes {
	if v == nil {
		return nil
	}
	return *v
}

func NullDailyTimes() **DailyTimes {
	var v *DailyTimes
	return &v
}

func NullableHours(v *Hours) **Hours {
	return &v
}

func GetHours(v **Hours) *Hours {
	if v == nil {
		return nil
	}
	return *v
}

func NullHours() **Hours {
	var v *Hours
	return &v
}

// UnorderedStrings masks []string properties for which Order doesn't matter, such as LabelIds
type UnorderedStrings []string

// Equal compares UnorderedStrings
func (a *UnorderedStrings) Equal(b Comparable) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Value of A: %+v, Value of B:%+v, Type Of A: %T, Type Of B: %T\n", a, b, a, b)
			panic(r)
		}
	}()

	if a == nil || b == nil {
		return false
	}

	var (
		u = []string(*a)
		s = []string(*b.(*UnorderedStrings))
	)
	if len(u) != len(s) {
		return false
	}

	for i := 0; i < len(u); i++ {
		var found bool
		for j := 0; j < len(s); j++ {
			if u[i] == s[j] {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func ToUnorderedStrings(v []string) *UnorderedStrings {
	u := UnorderedStrings(v)
	return &u
}
