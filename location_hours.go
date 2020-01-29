package yext

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Legacy hours format used for Yext v2 locations API

const (
	LocationHoursClosedAllWeek = "1:closed,2:closed,3:closed,4:closed,5:closed,6:closed,7:closed"
	LocationHoursOpen24Hours   = "00:00:00:00"
	LocationHoursClosed        = "closed"
	HoursFormat                = "15:04"
	hoursLen                   = 11 // XX:XX:XX:XX format
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

type LocationHoursHelper struct {
	Sunday    []string
	Monday    []string
	Tuesday   []string
	Wednesday []string
	Thursday  []string
	Friday    []string
	Saturday  []string
}

// String Format used from LocationService
func LocationHoursHelperFromString(str string) (*LocationHoursHelper, error) {
	var (
		hoursHelper  = &LocationHoursHelper{}
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

func MustLocationHoursHelperFromString(str string) *LocationHoursHelper {
	hoursHelper, err := LocationHoursHelperFromString(str)
	if err != nil {
		panic(err)
	}
	return hoursHelper
}

func (h *LocationHoursHelper) SetHours(weekday Weekday, hours []string) {
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

func (h *LocationHoursHelper) AppendHours(weekday Weekday, hours string) {
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

func (h *LocationHoursHelper) SetClosed(weekday Weekday) {
	h.SetHours(weekday, []string{LocationHoursClosed})
}

func (h *LocationHoursHelper) SetUnspecified(weekday Weekday) {
	h.SetHours(weekday, nil)
}

func (h *LocationHoursHelper) SetOpen24Hours(weekday Weekday) {
	h.SetHours(weekday, []string{LocationHoursOpen24Hours})
}

func (h *LocationHoursHelper) GetHours(weekday Weekday) []string {
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

func (h *LocationHoursHelper) StringSerialize() string {
	if h.HoursAreAllUnspecified() {
		return ""
	}
	var days = []string{
		h.StringSerializeDay(Sunday),
		h.StringSerializeDay(Monday),
		h.StringSerializeDay(Tuesday),
		h.StringSerializeDay(Wednesday),
		h.StringSerializeDay(Thursday),
		h.StringSerializeDay(Friday),
		h.StringSerializeDay(Saturday),
	}
	return strings.Join(days, ",")
}

func (h *LocationHoursHelper) StringSerializeDay(weekday Weekday) string {
	if h.HoursAreAllUnspecified() {
		return ""
	}
	var hoursStrings = []string{}
	if h.GetHours(weekday) == nil || len(h.GetHours(weekday)) == 0 || h.HoursAreClosed(weekday) {
		return fmt.Sprintf("%d:%s", weekday, LocationHoursClosed)
	}
	for _, hours := range h.GetHours(weekday) {
		if len(hours) != hoursLen {
			hours = "0" + hours
		}
		hoursStrings = append(hoursStrings, fmt.Sprintf("%d:%s", weekday, hours))
	}
	return strings.Join(hoursStrings, ",")
}

func (h *LocationHoursHelper) StructSerialize() **Hours {
	if h.HoursAreAllUnspecified() {
		return NullRegularHours()
	}
	hours := &Hours{}
	hours.Sunday = NullableDayHours(h.StructSerializeDay(Sunday))
	hours.Monday = NullableDayHours(h.StructSerializeDay(Monday))
	hours.Tuesday = NullableDayHours(h.StructSerializeDay(Tuesday))
	hours.Wednesday = NullableDayHours(h.StructSerializeDay(Wednesday))
	hours.Thursday = NullableDayHours(h.StructSerializeDay(Thursday))
	hours.Friday = NullableDayHours(h.StructSerializeDay(Friday))
	hours.Saturday = NullableDayHours(h.StructSerializeDay(Saturday))
	return NullableHours(hours)

}

func (h *LocationHoursHelper) StructSerializeDay(weekday Weekday) *DayHours {
	if h.HoursAreUnspecified(weekday) || h.HoursAreClosed(weekday) || len(h.GetHours(weekday)) == 0 {
		return &DayHours{
			IsClosed: NullableBool(true),
		}
	}
	var d = &DayHours{}
	intervals := []Interval{}
	for _, interval := range h.GetHours(weekday) {
		parts := strings.Split(interval, ":")
		intervals = append(intervals,
			Interval{
				Start: fmt.Sprintf("%s:%s", parts[0], parts[1]),
				End:   fmt.Sprintf("%s:%s", parts[2], parts[3]),
			},
		)
	}
	d.OpenIntervals = &intervals
	return d
}

func (h *LocationHoursHelper) ToStringSlice() ([]string, error) {
	var hoursStringSlice = make([][]string, 7)
	for i := range hoursStringSlice {
		weekday := Weekday(i + 1)
		if h.HoursAreClosed(weekday) {
			hoursStringSlice[i] = []string{"Closed"}
		} else if h.HoursAreOpen24Hours(weekday) {
			hoursStringSlice[i] = []string{"24hr"}
		} else {
			hours := h.GetHours(weekday)
			for _, h := range hours {
				open, close, err := ParseOpenAndCloseHoursFromString(h)
				if err != nil {
					return nil, err
				}
				if open, err = ConvertBetweenFormats(open, "15:04", "03:04pm"); err != nil {
					return nil, err
				}
				if close, err = ConvertBetweenFormats(close, "15:04", "03:04pm"); err != nil {
					return nil, err
				}
				hoursStringSlice[i] = append(hoursStringSlice[i], fmt.Sprintf("%s - %s", open, close))
			}
		}
	}
	var hoursSlice = make([]string, 7)
	for i, h := range hoursStringSlice {
		hoursSlice[i] = strings.Join(h, ",")
	}
	return hoursSlice, nil
}

func (h *LocationHoursHelper) MustToStringSlice() []string {
	hoursStringSlice, err := h.ToStringSlice()
	if err != nil {
		panic(err)
	}
	return hoursStringSlice
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
	if strings.ToLower(hoursParts[1]) == LocationHoursClosed {
		return Weekday(weekdayInt), LocationHoursClosed, nil
	}
	// Pad hours with leading 0s
	for i := 1; i < len(hoursParts); i += 2 {
		if len(hoursParts[i])+len(hoursParts[i+1]) != 4 {
			hoursParts[i] = "0" + hoursParts[i]
		}
	}
	hours := strings.Join(hoursParts[1:], ":")
	return Weekday(weekdayInt), hours, nil
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

func (h LocationHoursHelper) ToMap() map[Weekday][]string {
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

func (h *LocationHoursHelper) HoursAreAllUnspecified() bool {
	for _, hours := range h.ToMap() {
		if hours != nil {
			return false
		}
	}
	return true
}

func (h *LocationHoursHelper) HoursAreUnspecified(weekday Weekday) bool {
	return h.GetHours(weekday) == nil
}

func (h *LocationHoursHelper) HoursAreClosed(weekday Weekday) bool {
	var hours = h.GetHours(weekday)
	return hours != nil && len(hours) == 1 && hours[0] == LocationHoursClosed
}

func (h *LocationHoursHelper) HoursAreOpen24Hours(weekday Weekday) bool {
	var hours = h.GetHours(weekday)
	return hours != nil && len(hours) == 1 && hours[0] == LocationHoursOpen24Hours
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

func LocationHolidayHoursToHolidayHours(l *LocationHolidayHours) (*HolidayHours, error) {
	if l == nil {
		return nil, nil
	}
	var h = &HolidayHours{
		Date: String(l.Date),
	}
	if l.Hours == "" {
		h.IsClosed = NullableBool(true)
	} else {
		intervalsList := []Interval{}
		intervals := strings.Split(l.Hours, ",")
		for _, i := range intervals {
			open, close, err := ParseOpenAndCloseHoursFromString(i)
			if err != nil {
				return nil, err
			}
			if len(open) != 5 {
				open = "0" + open
			}
			if len(close) != 5 {
				close = "0" + close
			}
			intervalsList = append(intervalsList, Interval{
				Start: open,
				End:   close,
			})
		}
		h.OpenIntervals = &intervalsList
	}
	return h, nil
}
