package yext

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	HoursClosedAllWeek = "1:closed,2:closed,3:closed,4:closed,5:closed,6:closed,7:closed"
	HoursOpen24Hours   = "00:00:00:00"
	HoursClosed        = "closed"
)

type Weekday int

// Following the documentation of the Yext API,
// the indexing of days begins at 1 (with Sunday)
// and ends with 7 (Saturday)
// http://developer.yext.com/docs/api-reference/
const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekday) ToString() string {
	switch w {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	}
	return "Unknown"
}

type HoursHelper struct {
	Sunday    []string
	Monday    []string
	Tuesday   []string
	Wednesday []string
	Thursday  []string
	Friday    []string
	Saturday  []string
}

func HoursHelperFromString(str string) (*HoursHelper, error) {
	var (
		hoursHelper  = &HoursHelper{}
		hoursForDays = strings.Split(str, ",")
	)
	if len(str) == 0 {
		return hoursHelper, nil
	}
	for _, hoursForDay := range hoursForDays {
		weekday, hours, err := parseWeekdayAndHoursFromString(hoursForDay)
		if err != nil {
			return nil, err
		}
		hoursHelper.AppendHours(weekday, hours)
	}
	for weekday, hours := range hoursHelper.ToMap() {
		if hours == nil {
			hoursHelper.SetClosed(weekday)
		}
	}
	return hoursHelper, nil
}

func (h *HoursHelper) SetHours(weekday Weekday, hours []string) {
	switch weekday {
	case Sunday:
		h.Sunday = hours
	case Monday:
		h.Monday = hours
	case Tuesday:
		h.Tuesday = hours
	case Wednesday:
		h.Wednesday = hours
	case Thursday:
		h.Thursday = hours
	case Friday:
		h.Friday = hours
	case Saturday:
		h.Saturday = hours
	}
}

func (h *HoursHelper) AppendHours(weekday Weekday, hours string) {
	switch weekday {
	case Sunday:
		h.Sunday = append(h.Sunday, hours)
	case Monday:
		h.Monday = append(h.Monday, hours)
	case Tuesday:
		h.Tuesday = append(h.Tuesday, hours)
	case Wednesday:
		h.Wednesday = append(h.Wednesday, hours)
	case Thursday:
		h.Thursday = append(h.Thursday, hours)
	case Friday:
		h.Friday = append(h.Friday, hours)
	case Saturday:
		h.Saturday = append(h.Saturday, hours)
	}
}

func (h *HoursHelper) SetClosed(weekday Weekday) {
	h.SetHours(weekday, []string{HoursClosed})
}

func (h *HoursHelper) SetUnspecified(weekday Weekday) {
	h.SetHours(weekday, nil)
}

func (h *HoursHelper) GetHours(weekday Weekday) []string {
	switch weekday {
	case Sunday:
		return h.Sunday
	case Monday:
		return h.Monday
	case Tuesday:
		return h.Tuesday
	case Wednesday:
		return h.Wednesday
	case Thursday:
		return h.Thursday
	case Friday:
		return h.Friday
	case Saturday:
		return h.Saturday
	}
	return nil
}

func (h *HoursHelper) Serialize() string {
	if h.HoursAreAllUnspecified() {
		return ""
	}
	var days = []string{
		h.SerializeDay(Sunday),
		h.SerializeDay(Monday),
		h.SerializeDay(Tuesday),
		h.SerializeDay(Wednesday),
		h.SerializeDay(Thursday),
		h.SerializeDay(Friday),
		h.SerializeDay(Saturday),
	}
	return strings.Join(days, ",")
}

func (h *HoursHelper) SerializeDay(weekday Weekday) string {
	if h.HoursAreAllUnspecified() {
		return ""
	}
	var hoursStrings = []string{}
	if h.GetHours(weekday) == nil || len(h.GetHours(weekday)) == 0 {
		return fmt.Sprintf("%d:%s", weekday, HoursClosed)
	}
	for _, hours := range h.GetHours(weekday) {
		hoursStrings = append(hoursStrings, fmt.Sprintf("%d:%s", weekday, hours))
	}
	return strings.Join(hoursStrings, ",")
}

func parseWeekdayAndHoursFromString(str string) (Weekday, string, error) {
	if len(str) == 0 {
		return -1, "", fmt.Errorf("Error parsing weekday and hours from string: string has 0 length")
	}
	hoursParts := strings.Split(str, ":")
	if len(hoursParts) == 0 {
		return -1, "", fmt.Errorf("Error parsing weekday and hours from string: string in unexpectd format")
	}
	weekdayInt, err := strconv.Atoi(hoursParts[0])
	if err != nil {
		return -1, "", fmt.Errorf("Error parsing weekday hours from string; unable to convert index to num: %s", err)
	}
	return Weekday(weekdayInt), strings.Join(hoursParts[1:], ":"), nil
}

func ParseOpenAndCloseHoursFromString(hours string) (string, string, error) {
	if strings.Contains(hours, ":") {
		parts := strings.Split(hours, ":")
		if len(parts) == 4 {
			return fmt.Sprintf("%s:%s", parts[0], parts[1]), fmt.Sprintf("%s:%s", parts[2], parts[3]), nil
		} else if len(parts) == 5 {
			return fmt.Sprintf("%s:%s", parts[1], parts[2]), fmt.Sprintf("%s:%s", parts[3], parts[4]), nil
		}
	}
	return "", "", fmt.Errorf("Error parsing open and close hours from string %s: Unexpected format", hours)
}

func (h HoursHelper) ToMap() map[Weekday][]string {
	return map[Weekday][]string{
		Sunday:    h.Sunday,
		Monday:    h.Monday,
		Tuesday:   h.Tuesday,
		Wednesday: h.Wednesday,
		Thursday:  h.Thursday,
		Friday:    h.Friday,
		Saturday:  h.Saturday,
	}
}

func (h *HoursHelper) HoursAreAllUnspecified() bool {
	for _, hours := range h.ToMap() {
		if hours != nil {
			return false
		}
	}
	return true
}

func (h *HoursHelper) HoursAreUnspecified(weekday Weekday) bool {
	return h.GetHours(weekday) == nil
}

func (h *HoursHelper) HoursAreClosed(weekday Weekday) bool {
	var hours = h.GetHours(weekday)
	return hours != nil && len(hours) == 1 && hours[0] == HoursClosed
}

func (h *HoursHelper) HoursAreOpen24Hours(weekday Weekday) bool {
	var hours = h.GetHours(weekday)
	return hours != nil && len(hours) == 1 && hours[0] == HoursOpen24Hours
}

func ParseAndFormatHours(tFormat string, openHours string, closeHours string) (string, error) {
	openTime, err := time.Parse(tFormat, openHours)
	if err != nil {
		return "", fmt.Errorf("Error parsing hours %s with format %s: %s", openHours, tFormat, err)
	}
	closeTime, err := time.Parse(tFormat, closeHours)
	if err != nil {
		return "", fmt.Errorf("Error parsing hours %s with format %s: %s", closeHours, tFormat, err)
	}
	return fmt.Sprintf("%s:%s", openTime.Format("15:04"), closeTime.Format("15:04")), nil
}

func ConvertBetweenFormats(hours string, convertFromFormat string, convertToFormat string) (string, error) {
	t, err := time.Parse(convertFromFormat, hours)
	if err != nil {
		return "", fmt.Errorf("Hours %s was not in expected format %s: %s", hours, convertFromFormat, err)
	}
	return t.Format(convertToFormat), nil
}
