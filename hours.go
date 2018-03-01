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
	Sunday    string
	Monday    string
	Tuesday   string
	Wednesday string
	Thursday  string
	Friday    string
	Saturday  string
}

func HoursHelperFromString(str string) (*HoursHelper, error) {
	var (
		hoursHelper  = &HoursHelper{}
		hoursForDays = strings.Split(str, ",")
	)
	for _, hoursForDay := range hoursForDays {
		weekday, hours, err := parseWeekdayAndHoursFromString(hoursForDay)
		if err != nil {
			return nil, err
		}
		hoursHelper.SetHours(weekday, hours)
	}
	return hoursHelper, nil
}

func (h *HoursHelper) SetHours(weekday Weekday, hours string) {
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

func (h *HoursHelper) GetHours(weekday Weekday) string {
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
	return ""
}

func (h *HoursHelper) Serialize() string {
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
	return fmt.Sprintf("%d:%s", weekday, h.GetHours(weekday))
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

func (h HoursHelper) ToStringSlice() []string {
	return []string{
		h.Sunday,
		h.Monday,
		h.Tuesday,
		h.Wednesday,
		h.Thursday,
		h.Friday,
		h.Saturday,
	}
}

func HoursAreClosed(hours string) bool {
	return strings.ToLower(hours) == HoursClosed
}

func HoursAreOpen24Hours(hours string) bool {
	return hours == HoursOpen24Hours
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
