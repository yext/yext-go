package yext

func Bool(v bool) **bool {
	p := &v
	return &p
}

func GetBool(v **bool) bool {
	if v == nil || *v == nil {
		return false
	}
	return **v
}

func SingleBool(v bool) *bool {
	return &v
}

func GetSingleBool(v *bool) bool {
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

func NullString() *string {
	var v string
	return &v
}

// Necessary for single option custom fields
// TODO: rename to single option?
func DoubleString(v string) **string {
	y := &v
	return &y
}

func GetDoubleString(v **string) string {
	if v == nil || *v == nil {
		return ""
	}
	return **v
}

func NullDoubleString() **string {
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

func SingleFloat(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}

func GetSingleFloat(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func Float(v float64) **float64 {
	p := &v
	return &p
}

func GetFloat(v **float64) float64 {
	if v == nil || *v == nil {
		return 0
	}
	return **v
}

func Int(v int) **int {
	p := &v
	return &p
}

func GetInt(v **int) int {
	if v == nil || *v == nil {
		return 0
	}
	return **v
}

func SingleInt(v int) *int {
	p := &v
	return p
}

func GetSingleInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func NullInt() **int {
	var v *int
	return &v
}

func ToDate(v *Date) **Date {
	return &v
}

func GetDate(v **Date) *Date {
	if v == nil {
		return nil
	}
	return *v
}

func NullDate() **Date {
	var v *Date
	return &v
}

func ToVideo(v *Video) **Video {
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

func ToPhoto(v *Photo) **Photo {
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

func ToDailyTimes(v *DailyTimes) **DailyTimes {
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

func ToHours(v *Hours) **Hours {
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
