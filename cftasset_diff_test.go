package yext

import (
	"testing"
)

func TestAssetDiff(t *testing.T) {
	tests := []struct {
		A      *CFTAsset
		B      *CFTAsset
		IsDiff bool
	}{
		{
			A: &CFTAsset{
				Id: String("400122"),
				ForEntities: &ForEntities{
					MappingType: MappingTypeEntities,
					EntityIds: ToUnorderedStrings([]string{
						"DK738",
						"AG995",
						"BP579",
						"AP730",
						"57GLL",
						"0000646579",
						"AH350",
					}),
				},
			},
			B: &CFTAsset{
				Id: String("400122"),
				ForEntities: &ForEntities{
					MappingType: MappingTypeEntities,
					EntityIds: ToUnorderedStrings([]string{
						"599YM",
						"26FMJ",
						"0000679221",
					}),
				},
			},
			IsDiff: true,
		},
		{
			A: &CFTAsset{
				Id:     String("122400"),
				Name:   String("122400"),
				Type:   ASSETTYPE_TEXT,
				Locale: String("en"),
				ForEntities: &ForEntities{
					MappingType: MappingTypeEntities,
					EntityIds: ToUnorderedStrings([]string{
						"DK738",
						"AG995",
						"BP579",
						"AP730",
						"57GLL",
						"0000646579",
						"AH350",
					}),
				},
				Value: TextValue("122400"),
			},
			B: &CFTAsset{
				Id:     String("122400"),
				Name:   String("122400"),
				Type:   ASSETTYPE_TEXT,
				Locale: String("en"),
				ForEntities: &ForEntities{
					MappingType: MappingTypeEntities,
					EntityIds: ToUnorderedStrings([]string{
						"599YM",
						"26FMJ",
						"0000679221",
					}),
				},
				Value: TextValue("122400"),
			},
			IsDiff: true,
		},
		// Shouldn't be detecting diff but for some reason equal struct subfield Usage is throwing diff() off
		{
			A: &CFTAsset{
				Id:     String("122400"),
				Name:   String("122400"),
				Type:   ASSETTYPE_TEXT,
				Locale: String("en"),
				ForEntities: &ForEntities{
					MappingType: MappingTypeEntities,
					EntityIds: ToUnorderedStrings([]string{
						"DK738",
						"AG995",
						"BP579",
						"AP730",
						"57GLL",
						"0000646579",
						"AH350",
					}),
				},
				Usage: &AssetUsageList{
					{
						Type:       UsageTypeProfileFields,
						FieldNames: ToUnorderedStrings([]string{"c_contactFormJPN"}),
					},
				},
				Value: TextValue("122400"),
			},
			B: &CFTAsset{
				Id:     String("122400"),
				Name:   String("122400"),
				Type:   ASSETTYPE_TEXT,
				Locale: String("en"),
				ForEntities: &ForEntities{
					MappingType: MappingTypeEntities,
					EntityIds: ToUnorderedStrings([]string{
						"599YM",
						"26FMJ",
						"0000679221",
					}),
				},
				Usage: &AssetUsageList{
					{
						Type:       UsageTypeProfileFields,
						FieldNames: ToUnorderedStrings([]string{"c_contactFormJPN"}),
					},
				},
				Value: TextValue("122400"),
			},
			IsDiff: true,
		},
	}

	for _, test := range tests {
		diff, isDiff := test.A.Diff(test.B)
		if isDiff != test.IsDiff {
			t.Errorf("Expected isDiff: %t. Got: %t\nDelta: %+v", test.IsDiff, isDiff, diff)
		}
	}
}
