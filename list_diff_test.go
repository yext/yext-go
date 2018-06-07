package yext

import (
	"testing"
)

func TestListEqual(t *testing.T) {
	type test struct {
		A, B   *BioList
		IsDiff bool
	}

	tests := []test{
		test{
			A:      &BioList{},
			B:      &BioList{},
			IsDiff: false,
		},
		test{
			A: &BioList{
				List: List{Name: String("List A")},
			},
			B: &BioList{
				List: List{Name: String("List A")},
			},
			IsDiff: false,
		},
		test{
			A: &BioList{
				List: List{Name: String("List A")},
			},
			B: &BioList{
				List: List{Name: String("List B")},
			},
			IsDiff: true,
		},
		test{
			A: &BioList{
				List: List{Name: String("List A")},
				Sections: []*BioListSection{
					&BioListSection{
						ListSection: ListSection{Name: String("Section")},
						Items: []*Bio{
							&Bio{
								Title:       "Doctor",
								PhoneNumber: "8888888888",
							},
						},
					},
				},
			},
			B: &BioList{
				List: List{Name: String("List A")},
			},
			IsDiff: true,
		},
		test{
			A: &BioList{
				List: List{Name: String("List A")},
				Sections: []*BioListSection{
					&BioListSection{
						ListSection: ListSection{Name: String("Section")},
						Items: []*Bio{
							&Bio{
								Title:       "Doctor",
								PhoneNumber: "8888888888",
							},
						},
					},
				},
			},
			B: &BioList{
				List: List{Name: String("List A")},
				Sections: []*BioListSection{
					&BioListSection{
						ListSection: ListSection{Name: String("Section")},
						Items: []*Bio{
							&Bio{
								Title:       "Physician",
								PhoneNumber: "8888888888",
							},
						},
					},
				},
			},
			IsDiff: true,
		},
		test{
			A: &BioList{
				List: List{Name: String("List A")},
				Sections: []*BioListSection{
					&BioListSection{
						ListSection: ListSection{Name: String("Section")},
						Items: []*Bio{
							&Bio{
								Title:       "Doctor",
								PhoneNumber: "8888888888",
							},
						},
					},
				},
			},
			B: &BioList{
				List: List{Name: String("List A")},
				Sections: []*BioListSection{
					&BioListSection{
						ListSection: ListSection{Name: String("Section")},
						Items: []*Bio{
							&Bio{
								Title:       "Doctor",
								PhoneNumber: "8888888888",
							},
						},
					},
				},
			},
			IsDiff: false,
		},
	}

	for _, test := range tests {
		if isDiff := test.A.Equal(test.B); isDiff != test.IsDiff {
			t.Errorf("\nA: %s\nB: %s\nWanted:%t\nGot:   %t", test.A, test.B, test.IsDiff, isDiff)
		}
	}
}
