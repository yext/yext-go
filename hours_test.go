package yext

import (
	"reflect"
	"testing"
)

func TestHoursHelperFromString(t *testing.T) {
	var tests = []struct {
		Have string
		Want *LocationHoursHelper
	}{
		{
			Have: "1:09:00:20:00,2:closed,3:00:00:00:00,4:07:00:22:00,5:09:00:19:00,6:09:00:21:00,7:08:00:20:00",
			Want: &LocationHoursHelper{
				Sunday:    []string{"09:00:20:00"},
				Monday:    []string{"closed"},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
		},
		{
			Have: "1:9:00:20:00,2:closed,3:0:00:00:00,4:7:00:22:00,5:9:00:19:00,6:9:00:21:00,7:8:00:20:00",
			Want: &LocationHoursHelper{
				Sunday:    []string{"09:00:20:00"},
				Monday:    []string{"closed"},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
		},
		{
			Have: "1:09:00:20:00,3:00:00:00:00,4:07:00:22:00,5:09:00:19:00,6:09:00:21:00,7:08:00:20:00",
			Want: &LocationHoursHelper{
				Sunday:    []string{"09:00:20:00"},
				Monday:    []string{"closed"},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
		},
		{
			Have: "1:09:00:11:00,1:13:00:16:00,2:closed,3:00:00:00:00,4:07:00:22:00,5:09:00:19:00,6:09:00:21:00,7:08:00:20:00",
			Want: &LocationHoursHelper{
				Sunday:    []string{"09:00:11:00", "13:00:16:00"},
				Monday:    []string{"closed"},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
		},
		{
			Have: "2:9:00:18:00,3:9:00:18:00,4:9:00:18:00,5:9:00:18:00,6:9:00:18:00,7:10:00:2:00",
			Want: &LocationHoursHelper{
				Sunday:    []string{"closed"},
				Monday:    []string{"09:00:18:00"},
				Tuesday:   []string{"09:00:18:00"},
				Wednesday: []string{"09:00:18:00"},
				Thursday:  []string{"09:00:18:00"},
				Friday:    []string{"09:00:18:00"},
				Saturday:  []string{"10:00:02:00"},
			},
		},
		{
			Have: "",
			Want: &LocationHoursHelper{
				Sunday:    nil,
				Monday:    nil,
				Tuesday:   nil,
				Wednesday: nil,
				Thursday:  nil,
				Friday:    nil,
				Saturday:  nil,
			},
		},
	}

	for i, test := range tests {
		if got, err := LocationHoursHelperFromString(test.Have); err != nil {
			t.Errorf("Test HoursHelperFromString %d\nGot error: %s", i+1, err)
		} else if !reflect.DeepEqual(got, test.Want) {
			t.Errorf("Test HoursHelperFromString %d\nWanted: %v\nGot: %v", i+1, *test.Want, *got)
		}
	}
}

func TestStringSerialize(t *testing.T) {
	var tests = []struct {
		Have *LocationHoursHelper
		Want string
	}{
		{
			Have: &LocationHoursHelper{
				Sunday:    []string{"09:00:20:00"},
				Monday:    []string{"closed"},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
			Want: "1:09:00:20:00,2:closed,3:00:00:00:00,4:07:00:22:00,5:09:00:19:00,6:09:00:21:00,7:08:00:20:00",
		},
		{
			Have: &LocationHoursHelper{
				Sunday:    []string{"9:00:20:00"},
				Monday:    []string{"closed"},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"7:00:22:00"},
				Thursday:  []string{"9:00:19:00"},
				Friday:    []string{"9:00:21:00"},
				Saturday:  []string{"8:00:20:00"},
			},
			Want: "1:09:00:20:00,2:closed,3:00:00:00:00,4:07:00:22:00,5:09:00:19:00,6:09:00:21:00,7:08:00:20:00",
		},
		{
			Have: &LocationHoursHelper{
				Sunday:    []string{"09:00:20:00"},
				Monday:    []string{},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
			Want: "1:09:00:20:00,2:closed,3:00:00:00:00,4:07:00:22:00,5:09:00:19:00,6:09:00:21:00,7:08:00:20:00",
		},
		{
			Have: &LocationHoursHelper{
				Sunday:    []string{"09:00:20:00"},
				Monday:    nil,
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
			Want: "1:09:00:20:00,2:closed,3:00:00:00:00,4:07:00:22:00,5:09:00:19:00,6:09:00:21:00,7:08:00:20:00",
		},
		{
			Have: &LocationHoursHelper{
				Sunday:    []string{"09:00:11:00", "13:00:16:00"},
				Monday:    []string{"closed"},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"10:00:00:00"},
			},
			Want: "1:09:00:11:00,1:13:00:16:00,2:closed,3:00:00:00:00,4:07:00:22:00,5:09:00:19:00,6:09:00:21:00,7:10:00:00:00",
		},
		{
			Have: &LocationHoursHelper{
				Sunday:    nil,
				Monday:    nil,
				Tuesday:   nil,
				Wednesday: nil,
				Thursday:  nil,
				Friday:    nil,
				Saturday:  nil,
			},
			Want: "",
		},
	}

	for i, test := range tests {
		if got := test.Have.StringSerialize(); got != test.Want {
			t.Errorf("Test Serialize %d.\nWanted: %s\nGot: %s", i+1, test.Want, got)
		}
	}
}

func TestStructSerialize(t *testing.T) {
	var tests = []struct {
		Have *LocationHoursHelper
		Want **Hours
	}{
		{
			Have: &LocationHoursHelper{
				Sunday:    []string{"09:00:20:00"},
				Monday:    []string{"closed"},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
			Want: NullableHours(&Hours{
				Sunday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "20:00",
						},
					},
				}),
				Monday: NullableDayHours(&DayHours{
					IsClosed: NullableBool(true),
				}),
				Tuesday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "00:00",
							End:   "00:00",
						},
					},
				}),
				Wednesday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "07:00",
							End:   "22:00",
						},
					},
				}),
				Thursday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "19:00",
						},
					},
				}),
				Friday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "21:00",
						},
					},
				}),
				Saturday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "08:00",
							End:   "20:00",
						},
					},
				}),
			}),
		},
		{
			Have: &LocationHoursHelper{
				Sunday:    []string{"09:00:20:00"},
				Monday:    []string{},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
			Want: NullableHours(&Hours{
				Sunday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "20:00",
						},
					},
				}),
				Monday: NullableDayHours(&DayHours{
					IsClosed: NullableBool(true),
				}),
				Tuesday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "00:00",
							End:   "00:00",
						},
					},
				}),
				Wednesday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "07:00",
							End:   "22:00",
						},
					},
				}),
				Thursday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "19:00",
						},
					},
				}),
				Friday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "21:00",
						},
					},
				}),
				Saturday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "08:00",
							End:   "20:00",
						},
					},
				}),
			}),
		},
		{
			Have: &LocationHoursHelper{
				Sunday:    []string{"09:00:20:00"},
				Monday:    nil,
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
			Want: NullableHours(&Hours{
				Sunday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "20:00",
						},
					},
				}),
				Monday: NullableDayHours(&DayHours{
					IsClosed: NullableBool(true),
				}),
				Tuesday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "00:00",
							End:   "00:00",
						},
					},
				}),
				Wednesday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "07:00",
							End:   "22:00",
						},
					},
				}),
				Thursday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "19:00",
						},
					},
				}),
				Friday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "21:00",
						},
					},
				}),
				Saturday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "08:00",
							End:   "20:00",
						},
					},
				}),
			}),
		},
		{
			Have: &LocationHoursHelper{
				Sunday:    []string{"09:00:11:00", "13:00:16:00"},
				Monday:    []string{"closed"},
				Tuesday:   []string{"00:00:00:00"},
				Wednesday: []string{"07:00:22:00"},
				Thursday:  []string{"09:00:19:00"},
				Friday:    []string{"09:00:21:00"},
				Saturday:  []string{"08:00:20:00"},
			},
			Want: NullableHours(&Hours{
				Sunday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "11:00",
						},
						Interval{
							Start: "13:00",
							End:   "16:00",
						},
					},
				}),
				Monday: NullableDayHours(&DayHours{
					IsClosed: NullableBool(true),
				}),
				Tuesday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "00:00",
							End:   "00:00",
						},
					},
				}),
				Wednesday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "07:00",
							End:   "22:00",
						},
					},
				}),
				Thursday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "19:00",
						},
					},
				}),
				Friday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "09:00",
							End:   "21:00",
						},
					},
				}),
				Saturday: NullableDayHours(&DayHours{
					OpenIntervals: &[]Interval{
						Interval{
							Start: "08:00",
							End:   "20:00",
						},
					},
				}),
			}),
		},
		{
			Have: &LocationHoursHelper{
				Sunday:    nil,
				Monday:    nil,
				Tuesday:   nil,
				Wednesday: nil,
				Thursday:  nil,
				Friday:    nil,
				Saturday:  nil,
			},
			Want: NullHours(),
		},
	}

	for i, test := range tests {
		if got := test.Have.StructSerialize(); !reflect.DeepEqual(GetHours(got), GetHours(test.Want)) {
			//if want != nil {
			t.Errorf("Test StructSerialize %d.\nWanted: %v\nGot: %v", i+1, GetHours(test.Want), GetHours(got))
			//}
		}
	}
}

func TestHolidayHoursConvert(t *testing.T) {

	var tests = []struct {
		Have *LocationHolidayHours
		Want *HolidayHours
	}{
		{
			Have: &LocationHolidayHours{
				Date:  "2018-12-25",
				Hours: "8:00:16:00",
			},
			Want: &HolidayHours{
				Date: String("2018-12-25"),
				OpenIntervals: &[]Interval{
					Interval{
						Start: "08:00",
						End:   "16:00",
					},
				},
			},
		},
		{
			Have: &LocationHolidayHours{
				Date:  "2018-12-25",
				Hours: "08:00:16:00",
			},
			Want: &HolidayHours{
				Date: String("2018-12-25"),
				OpenIntervals: &[]Interval{
					Interval{
						Start: "08:00",
						End:   "16:00",
					},
				},
			},
		},
		{
			Have: &LocationHolidayHours{
				Date:  "2018-12-25",
				Hours: "9:00:15:00,17:00:19:00",
			},
			Want: &HolidayHours{
				Date: String("2018-12-25"),
				OpenIntervals: &[]Interval{
					Interval{
						Start: "09:00",
						End:   "15:00",
					},
					Interval{
						Start: "17:00",
						End:   "19:00",
					},
				},
			},
		},
		{
			Have: &LocationHolidayHours{
				Date:  "2018-12-25",
				Hours: "",
			},
			Want: &HolidayHours{
				Date:     String("2018-12-25"),
				IsClosed: NullableBool(true),
			},
		},
	}

	for i, test := range tests {
		if got, _ := LocationHolidayHoursToHolidayHours(test.Have); !reflect.DeepEqual(got, test.Want) {
			t.Errorf("Test LocattionHolidayHoursToHolidayHours %d.\nWanted: %v\nGot: %v", i+1, *test.Want, *got)
		}
	}

}
