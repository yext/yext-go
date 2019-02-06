package yext

import (
	"reflect"
	"testing"
)

type diffTest struct {
	name           string
	property       string
	isDiff         bool
	baseValue      interface{}
	newValue       interface{}
	baseNilIsEmpty bool
	newNilIsEmpty  bool
	deltaValue     interface{}
}

func setValOnProperty(val interface{}, property string, entity Entity) {
	if val != nil {
		reflect.ValueOf(entity).Elem().FieldByName(property).Set(reflect.ValueOf(val))
	}
}

func (c *CustomLocationEntity) SetValOnProperty(val interface{}, property string) *CustomLocationEntity {
	setValOnProperty(val, property, c)
	return c
}

func TestEntityDiff(t *testing.T) {
	// To ensure that we test all combinations, tests follow this pattern: (more or less):
	// BASIC EQUALITY:
	// A) equal
	// B) not equal

	// BASIC EQUALITY WITH ZERO VALUES:
	// C) base is zero value, new is non-zero value
	// D) base is non-zero value, new is zero value
	// E) both are zero value

	// EQUALITY WITH NIL (and no nil is empty)
	// F) both are nil
	// G) base is non-zero value, new is nil
	// H) base is nil, new is non-zero value
	// I) base is nil, new is zero value
	// J) base is zero value, new is nil

	// EQUALITY WITH NIL (and nil is empty)
	// K) base is nil (nil is empty), new is zero value
	// L) base is nil (nil is empty), new is zero value (nil is empty)
	// M) base is zero value, new is nil (nil is empty)
	// N) base is zero value (nil is empty), new is nil (nil is empty)

	tests := []diffTest{
		// **String tests
		diffTest{
			name:      "*String: equal (A)",
			property:  "Name",
			baseValue: String("Hupman's Hotdogs"),
			newValue:  String("Hupman's Hotdogs"),
			isDiff:    false,
		},
		diffTest{
			name:       "*String: not equal (B)",
			property:   "Name",
			baseValue:  String("Hupman's Hotdogs"),
			newValue:   String("Bryan's Bagels"),
			isDiff:     true,
			deltaValue: String("Bryan's Bagels"),
		},
		diffTest{
			name:       "*String: base is empty string, new is not (C)",
			property:   "Name",
			baseValue:  String(""),
			newValue:   String("Bryan's Bagels"),
			isDiff:     true,
			deltaValue: String("Bryan's Bagels"),
		},
		diffTest{
			name:       "*String: new is empty string, base is not (D)",
			property:   "Name",
			baseValue:  String("Hupman's Hotdogs"),
			newValue:   String(""),
			isDiff:     true,
			deltaValue: String(""),
		},
		diffTest{
			name:      "*String: both are empty (E)",
			property:  "Name",
			baseValue: String(""),
			newValue:  String(""),
			isDiff:    false,
		},
		diffTest{
			name:      "*String: both are nil (F)",
			property:  "Name",
			baseValue: nil,
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:      "*String: base is not nil, new is nil (G)",
			property:  "Name",
			baseValue: String("Bryan's Bagels"),
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:       "*String: base is nil, new is not nil (H)",
			property:   "Name",
			baseValue:  nil,
			newValue:   String("Bryan's Bagels"),
			isDiff:     true,
			deltaValue: String("Bryan's Bagels"),
		},
		diffTest{
			name:       "*String: base is nil, new is empty string (I)",
			property:   "Name",
			baseValue:  nil,
			newValue:   String(""),
			isDiff:     true,
			deltaValue: String(""),
		},
		diffTest{
			name:      "*String: base is empty string, new is nil (J)",
			property:  "Name",
			baseValue: String(""),
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:           "*String: base is nil (with nil is empty), new is empty string (K)",
			property:       "Name",
			baseValue:      nil,
			newValue:       String(""),
			baseNilIsEmpty: true,
			isDiff:         false,
		},
		diffTest{
			name:           "*String: base is nil (with nil is empty), new is empty string (with nil is empty) (L)",
			property:       "Name",
			baseValue:      nil,
			newValue:       String(""),
			baseNilIsEmpty: true,
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		diffTest{
			name:          "*String: base is empty string, new is nil (with nil is empty) (M)",
			property:      "Name",
			baseValue:     String(""),
			newValue:      nil,
			newNilIsEmpty: true,
			isDiff:        false,
		},
		diffTest{
			name:           "*String: base is empty string (with nil is empty), new is nil (with nil is empty) (N)",
			property:       "Name",
			baseValue:      String(""),
			baseNilIsEmpty: true,
			newValue:       nil,
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		// **Float tests
		diffTest{
			name:      "**Float: equal (A)",
			property:  "YearEstablished",
			baseValue: NullableFloat(2018),
			newValue:  NullableFloat(2018),
			isDiff:    false,
		},
		diffTest{
			name:       "**Float: not equal (B)",
			property:   "YearEstablished",
			baseValue:  NullableFloat(2018),
			newValue:   NullableFloat(2006),
			isDiff:     true,
			deltaValue: NullableFloat(2006),
		},
		diffTest{
			name:       "**Float: base is 0, new is not 0 (C)",
			property:   "YearEstablished",
			baseValue:  NullableFloat(0),
			newValue:   NullableFloat(2006),
			isDiff:     true,
			deltaValue: NullableFloat(2006),
		},
		diffTest{
			name:       "**Float: base is not 0, new is 0 (D)",
			property:   "YearEstablished",
			baseValue:  NullableFloat(2006),
			newValue:   NullableFloat(0),
			isDiff:     true,
			deltaValue: NullableFloat(0),
		},
		diffTest{
			name:      "**Float: both are 0 (E)",
			property:  "YearEstablished",
			baseValue: NullableFloat(2018),
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:      "**Float: both are nil (F)",
			property:  "YearEstablished",
			baseValue: nil,
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:      "**Float: base is not 0, new is nil (G)",
			property:  "YearEstablished",
			baseValue: NullableFloat(1993),
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:       "**Float: base is nil, new is not 0 (H)",
			property:   "YearEstablished",
			baseValue:  nil,
			newValue:   NullableFloat(1993),
			isDiff:     true,
			deltaValue: NullableFloat(1993),
		},
		diffTest{
			name:       "**Float: base is nil, new is 0 (I)",
			property:   "YearEstablished",
			baseValue:  nil,
			newValue:   NullableFloat(0),
			isDiff:     true,
			deltaValue: NullableFloat(0),
		},
		diffTest{
			name:      "**Float: base is 0, new is nil (J)",
			property:  "YearEstablished",
			baseValue: NullableFloat(0),
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:           "**Float: base is nil (nil is empty), new is 0 (K)",
			property:       "YearEstablished",
			baseValue:      nil,
			newValue:       NullableFloat(0),
			baseNilIsEmpty: true,
			isDiff:         false,
		},
		diffTest{
			name:           "**Float: base is nil (nil is empty), new is 0 (nil is empty) (L)",
			property:       "YearEstablished",
			baseValue:      nil,
			newValue:       NullableFloat(0),
			baseNilIsEmpty: true,
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		diffTest{
			name:          "**Float: base is 0, new is nil (nil is empty) (M)",
			property:      "YearEstablished",
			baseValue:     NullableFloat(0),
			newValue:      nil,
			newNilIsEmpty: true,
			isDiff:        false,
		},
		diffTest{
			name:           "**Float: base is 0 (nil is empty), new is nil (nil is empty) (N)",
			property:       "YearEstablished",
			baseValue:      NullableFloat(0),
			baseNilIsEmpty: true,
			newValue:       nil,
			newNilIsEmpty:  true,
			isDiff:         false,
		},

		// **Bool tests
		diffTest{
			name:      "**Bool: both true (A)",
			property:  "Closed",
			baseValue: NullableBool(true),
			newValue:  NullableBool(true),
			isDiff:    false,
		},
		diffTest{
			name:      "**Bool: both false (A/E)",
			property:  "Closed",
			baseValue: NullableBool(false),
			newValue:  NullableBool(false),
			isDiff:    false,
		},
		diffTest{
			name:       "**Bool: not equal, base true, new false (B/D)",
			property:   "Closed",
			baseValue:  NullableBool(true),
			newValue:   NullableBool(false),
			isDiff:     true,
			deltaValue: NullableBool(false),
		},
		diffTest{
			name:       "**Bool: not equal, base is false, new is true (B/C)",
			property:   "Closed",
			baseValue:  NullableBool(false),
			newValue:   NullableBool(true),
			isDiff:     true,
			deltaValue: NullableBool(true),
		},
		diffTest{
			name:      "**Bool: both are nil (F)",
			property:  "Closed",
			baseValue: nil,
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:      "**Bool: base is non-zero, new is nil (G)",
			property:  "Closed",
			baseValue: NullableBool(true),
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:       "**Bool: base is nil, new is non-zero value (H)",
			property:   "Closed",
			baseValue:  nil,
			newValue:   NullableBool(true),
			isDiff:     true,
			deltaValue: NullableBool(true),
		},
		diffTest{
			name:       "**Bool: base is nil, new is zero value (I)",
			property:   "Closed",
			baseValue:  nil,
			newValue:   NullableBool(false),
			isDiff:     true,
			deltaValue: NullableBool(false),
		},
		diffTest{
			name:      "**Bool: base is zero value, new is nil (J)",
			property:  "Closed",
			baseValue: NullableBool(false),
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:           "**Bool: base is nil (nil is empty), new is zero value (K)",
			property:       "Closed",
			baseValue:      nil,
			newValue:       NullableBool(false),
			baseNilIsEmpty: true,
			isDiff:         false,
		},
		diffTest{
			name:           "**Bool: base is nil (nil is empty), new is zero value (nil is empty) (L)",
			property:       "Closed",
			baseValue:      nil,
			newValue:       NullableBool(false),
			baseNilIsEmpty: true,
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		diffTest{
			name:          "**Bool: base is zero value, new is nil (nil is empty) (L)",
			property:      "Closed",
			baseValue:     NullableBool(false),
			newValue:      nil,
			newNilIsEmpty: true,
			isDiff:        false,
		},
		diffTest{
			name:           "**Bool: base is zero value (nil is empty), new is nil (nil is empty) (L)",
			property:       "Closed",
			baseValue:      NullableBool(false),
			newValue:       nil,
			baseNilIsEmpty: true,
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		// Struct tests (Address)
		diffTest{
			name:     "Address: equal (A)",
			property: "Address",
			baseValue: &Address{
				Line1: String("7900 Westpark"),
			},
			newValue: &Address{
				Line1: String("7900 Westpark"),
			},
			isDiff: false,
		},
		diffTest{
			name:     "Address: line 1 and line 2 equal (A)",
			property: "Address",
			baseValue: &Address{
				Line1: String("7900 Westpark Dr"),
				Line2: String("Suite 500"),
			},
			newValue: &Address{
				Line1: String("7900 Westpark Dr"),
				Line2: String("Suite 500"),
			},
			isDiff: false,
		},
		diffTest{
			name:     "Address: not equal (B)",
			property: "Address",
			baseValue: &Address{
				Line1: String("7900 Westpark"),
			},
			newValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
			isDiff: true,
			deltaValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
		},
		diffTest{
			name:     "Address: line 1 equal, line 2 not equal (B)",
			property: "Address",
			baseValue: &Address{
				Line1: String("7900 Westpark Dr"),
				Line2: String("Suite 500"),
			},
			newValue: &Address{
				Line1: String("7900 Westpark Dr"),
				Line2: String("Suite 700"),
			},
			isDiff: true,
			deltaValue: &Address{
				Line2: String("Suite 700"),
			},
		},
		diffTest{
			name:     "Address: base has line 1, new has line 2 (B)",
			property: "Address",
			baseValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
			newValue: &Address{
				Line2: String("Suite 700"),
			},
			isDiff: true,
			deltaValue: &Address{
				Line2: String("Suite 700"),
			},
		},
		diffTest{
			name:      "Address: base is empty struct, new is not empty (C)",
			property:  "Address",
			baseValue: &Address{},
			newValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
			isDiff: true,
			deltaValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
		},
		diffTest{
			name:     "Address: base is not empty struct, new is empty struct (D)",
			property: "Address",
			baseValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
			newValue:   &Address{},
			isDiff:     true,
			deltaValue: &Address{},
		},
		diffTest{
			name:      "Address: both are empty struct (E)",
			property:  "Address",
			baseValue: &Address{},
			newValue:  &Address{},
			isDiff:    false,
		},
		diffTest{
			name:      "Address: both are nil (F)",
			property:  "Address",
			baseValue: nil,
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:     "Address: base is non-empty struct, new is nil (G)",
			property: "Address",
			baseValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
			newValue: nil,
			isDiff:   false,
		},
		diffTest{
			name:      "Address: base is nil, new is non-empty struct (H)",
			property:  "Address",
			baseValue: nil,
			newValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
			isDiff: true,
			deltaValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
		},
		diffTest{
			name:       "Address: base is nil, new is empty struct (I)",
			property:   "Address",
			baseValue:  nil,
			newValue:   &Address{},
			isDiff:     true,
			deltaValue: &Address{},
		},
		diffTest{
			name:      "Address: base is empty struct, new is nil (J)",
			property:  "Address",
			baseValue: &Address{},
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:           "Address: base is nil (nil is empty), new is struct with zero values (K)",
			property:       "Address",
			baseValue:      nil,
			newValue:       &Address{Line1: String(""), Line2: String(""), City: String(""), Sublocality: String(""), Region: String(""), PostalCode: String("")},
			baseNilIsEmpty: true,
			isDiff:         false,
		},
		diffTest{
			name:           "Address: base is nil (nil is empty), new is struct with zero value (K-2)",
			property:       "Address",
			baseValue:      nil,
			newValue:       &Address{Line1: String("")},
			baseNilIsEmpty: true,
			isDiff:         false,
		},
		diffTest{
			name:           "Address: base is nil (nil is empty), new is empty struct (K-3)",
			property:       "Address",
			baseValue:      nil,
			newValue:       &Address{},
			baseNilIsEmpty: true,
			isDiff:         false,
		},

		diffTest{
			name:           "Address: base is nil (nil is empty), new is struct with zero values (nil is empty) (L)",
			property:       "Address",
			baseValue:      nil,
			baseNilIsEmpty: true,
			newValue:       &Address{Line1: String(""), Line2: String(""), City: String(""), Sublocality: String(""), Region: String(""), PostalCode: String("")},
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		diffTest{
			name:           "Address: base is nil (nil is empty), new is struct with zero value (nil is empty) (L-2)",
			property:       "Address",
			baseValue:      nil,
			baseNilIsEmpty: true,
			newValue:       &Address{Line1: String("")},
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		diffTest{
			name:           "Address: base is nil (nil is empty), new is empty struct (nil is empty) (L-3)",
			property:       "Address",
			baseValue:      nil,
			baseNilIsEmpty: true,
			newValue:       &Address{},
			newNilIsEmpty:  true,
			isDiff:         false,
		},

		diffTest{
			name:          "Address: base is struct with zero values, new is nil (nil is empty) (M)",
			property:      "Address",
			baseValue:     &Address{Line1: String(""), Line2: String(""), City: String(""), Sublocality: String(""), Region: String(""), PostalCode: String("")},
			newValue:      nil,
			newNilIsEmpty: true,
			isDiff:        false,
		},
		diffTest{
			name:          "Address: base is struct with zero value, new is nil (nil is empty) (M-2)",
			property:      "Address",
			baseValue:     &Address{Line1: String("")},
			newValue:      nil,
			newNilIsEmpty: true,
			isDiff:        false,
		},
		diffTest{
			name:          "Address: base empty struct, new nis nil (nil is empty) (M-3)",
			property:      "Address",
			baseValue:     &Address{},
			newValue:      nil,
			newNilIsEmpty: true,
			isDiff:        false,
		},

		diffTest{
			name:           "Address: base is struct with zero values (nil is empty), new is nil (nil is empty) (N)",
			property:       "Address",
			baseValue:      &Address{Line1: String(""), Line2: String(""), City: String(""), Sublocality: String(""), Region: String(""), PostalCode: String("")},
			baseNilIsEmpty: true,
			newValue:       nil,
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		diffTest{
			name:           "Address: base is struct with zero value (nil is empty), new is nil (nil is empty) (N-2)",
			property:       "Address",
			baseValue:      &Address{Line1: String("")},
			baseNilIsEmpty: true,
			newValue:       nil,
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		diffTest{
			name:           "Address: base empty struct (nil is empty), new nis nil (nil is empty) (N-3)",
			property:       "Address",
			baseValue:      &Address{},
			baseNilIsEmpty: true,
			newValue:       nil,
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		diffTest{
			name:       "List: base is not empty struct, new is empty struct (D)",
			property:   "CFTextList",
			baseValue:  &[]string{"a", "b"},
			newValue:   &[]string{},
			isDiff:     true,
			deltaValue: &[]string{},
		},
		diffTest{
			name:       "List: not equal (B)",
			property:   "CFTextList",
			baseValue:  &[]string{"a", "b"},
			newValue:   &[]string{"b", "c"},
			isDiff:     true,
			deltaValue: &[]string{"b", "c"},
		},
		// Comparable tests (Unordered Strings)
		diffTest{
			name:      "UnorderedStrings: equal (ordered) (A)",
			property:  "CFMultiOption",
			baseValue: ToUnorderedStrings([]string{"a", "b"}),
			newValue:  ToUnorderedStrings([]string{"a", "b"}),
			isDiff:    false,
		},
		diffTest{
			name:      "UnorderedStrings: equal (unordered) (A)",
			property:  "CFMultiOption",
			baseValue: ToUnorderedStrings([]string{"a", "b"}),
			newValue:  ToUnorderedStrings([]string{"b", "a"}),
			isDiff:    false,
		},
		diffTest{
			name:       "UnorderedStrings: not equal (B)",
			property:   "CFMultiOption",
			baseValue:  ToUnorderedStrings([]string{"a", "b"}),
			newValue:   ToUnorderedStrings([]string{"c", "b"}),
			isDiff:     true,
			deltaValue: ToUnorderedStrings([]string{"c", "b"}),
		},
		diffTest{
			name:       "UnorderedStrings: base is empty, new is non-empty (C)",
			property:   "CFMultiOption",
			baseValue:  ToUnorderedStrings([]string{}),
			newValue:   ToUnorderedStrings([]string{"a", "b"}),
			isDiff:     true,
			deltaValue: ToUnorderedStrings([]string{"a", "b"}),
		},
		diffTest{
			name:       "UnorderedStrings: base is non-empty, new is empty (D)",
			property:   "CFMultiOption",
			baseValue:  ToUnorderedStrings([]string{"a", "b"}),
			newValue:   ToUnorderedStrings([]string{}),
			isDiff:     true,
			deltaValue: ToUnorderedStrings([]string{}),
		},
		diffTest{
			name:      "UnorderedStrings: both are empty (E)",
			property:  "CFMultiOption",
			baseValue: ToUnorderedStrings([]string{}),
			newValue:  ToUnorderedStrings([]string{}),
			isDiff:    false,
		},
		diffTest{
			name:      "UnorderedStrings: both are nil (F)",
			property:  "CFMultiOption",
			baseValue: nil,
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:      "UnorderedStrings: base is non-zero, new is nil (G)",
			property:  "CFMultiOption",
			baseValue: ToUnorderedStrings([]string{"a", "b"}),
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:       "UnorderedStrings: base is nil, new is non-zero (H)",
			property:   "CFMultiOption",
			baseValue:  nil,
			newValue:   ToUnorderedStrings([]string{"a", "b"}),
			isDiff:     true,
			deltaValue: ToUnorderedStrings([]string{"a", "b"}),
		},
		diffTest{
			name:       "UnorderedStrings: base is nil, new is zero (I)",
			property:   "CFMultiOption",
			baseValue:  nil,
			newValue:   ToUnorderedStrings([]string{}),
			isDiff:     true,
			deltaValue: ToUnorderedStrings([]string{}),
		},
		diffTest{
			name:      "UnorderedStrings: base is zero, new is nil (J)",
			property:  "CFMultiOption",
			baseValue: ToUnorderedStrings([]string{}),
			newValue:  nil,
			isDiff:    false,
		},
		diffTest{
			name:           "UnorderedStrings: base is nil (nil is empty), new is zero (K)",
			property:       "CFMultiOption",
			baseValue:      nil,
			baseNilIsEmpty: true,
			newValue:       ToUnorderedStrings([]string{}),
			isDiff:         false,
		},
		diffTest{
			name:           "UnorderedStrings: base is nil (nil is empty), new is zero (nil is empty) (L)",
			property:       "CFMultiOption",
			baseValue:      nil,
			baseNilIsEmpty: true,
			newValue:       ToUnorderedStrings([]string{}),
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		diffTest{
			name:          "UnorderedStrings: base is zero value, new is nil (nil is empty) (L)",
			property:      "CFMultiOption",
			baseValue:     ToUnorderedStrings([]string{}),
			newValue:      nil,
			newNilIsEmpty: true,
			isDiff:        false,
		},
		diffTest{
			name:           "UnorderedStrings: base is zero value (nil is empty), new is nil (nil is empty) (L)",
			property:       "CFMultiOption",
			baseValue:      ToUnorderedStrings([]string{}),
			baseNilIsEmpty: true,
			newValue:       nil,
			newNilIsEmpty:  true,
			isDiff:         false,
		},
		// EmbeddedStruct tests
		diffTest{
			name:     "Base Entity: different Ids",
			property: "BaseEntity",
			baseValue: BaseEntity{
				Meta: &EntityMeta{
					Id: String("CTG"),
				},
			},
			newValue: BaseEntity{
				Meta: &EntityMeta{
					Id: String("CTG2"),
				},
			},
			isDiff: true,
			deltaValue: BaseEntity{
				Meta: &EntityMeta{
					Id: String("CTG2"),
				},
			},
		},
		diffTest{
			name:     "Base Entity: not equal",
			property: "BaseEntity",
			baseValue: BaseEntity{
				Meta: &EntityMeta{
					Id:       String("CTG"),
					FolderId: String("123"),
				},
			},
			newValue: BaseEntity{
				Meta: &EntityMeta{
					Id:       String("CTG"),
					FolderId: String("456"),
				},
			},
			isDiff: true,
			deltaValue: BaseEntity{
				Meta: &EntityMeta{
					FolderId: String("456"),
				},
			},
		},
		diffTest{
			name:     "Base Entity: equal",
			property: "BaseEntity",
			baseValue: BaseEntity{
				Meta: &EntityMeta{
					FolderId: String("123"),
				},
			},
			newValue: BaseEntity{
				Meta: &EntityMeta{
					FolderId: String("123"),
				},
			},
			isDiff: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var (
				baseEntity = new(CustomLocationEntity)
				newEntity  = new(CustomLocationEntity)
			)
			if test.property != "" {
				baseEntity.SetValOnProperty(test.baseValue, test.property)
				newEntity.SetValOnProperty(test.newValue, test.property)
			}
			if test.baseNilIsEmpty {
				setNilIsEmpty(baseEntity)
			}
			if test.newNilIsEmpty {
				setNilIsEmpty(newEntity)
			}
			delta, isDiff, _ := Diff(baseEntity, newEntity)
			if isDiff != test.isDiff {
				t.Errorf("Expected isDiff: %t. Got: %t", test.isDiff, isDiff)
			} else if test.isDiff == false && delta != nil {
				t.Errorf("Expected isDiff: %t. Got delta: %v", test.isDiff, delta)
			} else if isDiff {
				expectedDelta := new(CustomLocationEntity)
				if test.property != "" && test.deltaValue != nil {
					expectedDelta.SetValOnProperty(test.deltaValue, test.property)
				}
				if !reflect.DeepEqual(delta, expectedDelta) {
					t.Errorf("Expected delta: %v. Got: %v", expectedDelta, delta)
				}
			}
		})
	}
}

func TestEntityDiffComplex(t *testing.T) {
	custom1 := &CustomLocationEntity{
		LocationEntity: LocationEntity{
			Address: &Address{
				Line1: String("7900 Westpark"),
			},
		},
	}
	custom2 := &CustomLocationEntity{
		LocationEntity: LocationEntity{
			Address: &Address{
				Line1: String("7900 Westpark"),
				Line2: String("Suite T200"),
			},
		},
	}
	custom3 := &CustomLocationEntity{
		LocationEntity: LocationEntity{
			Address: &Address{
				Line2: String("Suite T200"),
			},
		},
	}
	custom4 := &CustomLocationEntity{}

	tests := []struct {
		name   string
		base   *CustomLocationEntity
		new    *CustomLocationEntity
		isDiff bool
		delta  *CustomLocationEntity
	}{
		{
			name:   "equal",
			base:   custom1,
			new:    custom1,
			isDiff: false,
		},
		{
			name:   "not equal",
			base:   custom1,
			new:    custom2,
			isDiff: true,
			delta: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Address: &Address{
						Line2: String("Suite T200"),
					},
				},
			},
		},
		{
			name:   "not equal (2)",
			base:   custom1,
			new:    custom3,
			isDiff: true,
			delta: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Address: &Address{
						Line2: String("Suite T200"),
					},
				},
			},
		},
		{
			name:   "empty struct",
			base:   custom4,
			new:    custom3,
			isDiff: true,
			delta:  custom3,
		},
		// Though the test below might look incorrect this is how the location.Diff() works
		{
			name:   "empty struct (new is empty struct)",
			base:   custom3,
			new:    custom4,
			isDiff: false,
		},
		{
			name: "Address, partial diff",
			base: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Address: &Address{
						Line1: String("7900 Westpark"),
						City:  String("McLean"),
					},
				},
			},
			new: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Address: &Address{
						Line1: String("7900 Westpark"),
						City:  String(""),
					},
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Address: &Address{
						City: String(""),
					},
				},
			},
		},
		{
			name: "CFGallery",
			base: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFGallery: &[]Photo{
						Photo{
							Description: String("Description 1"),
						},
						Photo{
							Description: String("Description 2"),
						},
					},
				},
			},
			new: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFGallery: &[]Photo{
						Photo{
							Description: String("New Description 1"),
						},
						Photo{
							Description: String("Description 2"),
						},
					},
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFGallery: &[]Photo{
						Photo{
							Description: String("New Description 1"),
						},
						Photo{
							Description: String("Description 2"),
						},
					},
				},
			},
		},
		{
			name: "CFHolidayHours", // TODO this is not real!!
			base: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date:     String("2019-01-21"),
							IsClosed: NullableBool(true),
						},
					},
				},
			},
			new: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date:     String("2019-01-22"),
							IsClosed: NullableBool(true),
						},
					},
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				CustomEntity: CustomEntity{CFHolidayHours: &[]HolidayHours{
					HolidayHours{
						Date:     String("2019-01-22"),
						IsClosed: NullableBool(true),
					},
				},
				}},
		},
		{
			name: "CFHolidayHours",
			base: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date:     String("2019-01-21"),
							IsClosed: NullableBool(false),
						},
					},
				},
			},
			new: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date:     String("2019-01-21"),
							IsClosed: NullableBool(true),
						},
					},
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date:     String("2019-01-21"),
							IsClosed: NullableBool(true),
						},
					},
				},
			},
		},
		{
			name: "CFHolidayHours",
			base: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date: String("2019-01-21"),
							OpenIntervals: &[]Interval{
								Interval{
									Start: "09:00",
									End:   "16:30",
								},
							},
						},
					},
				},
			},
			new: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date:     String("2019-01-21"),
							IsClosed: NullableBool(true),
						},
					},
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date:     String("2019-01-21"),
							IsClosed: NullableBool(true),
						},
					},
				},
			},
		},
		{
			name: "CFHolidayHours",
			base: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date: String("2019-01-21"),
							OpenIntervals: &[]Interval{
								Interval{
									Start: "09:00",
									End:   "16:30",
								},
							},
						},
					},
				},
			},
			new: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date: String("2019-01-21"),
							OpenIntervals: &[]Interval{
								Interval{
									Start: "09:00",
									End:   "17:00",
								},
							},
						},
					},
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date: String("2019-01-21"),
							OpenIntervals: &[]Interval{
								Interval{
									Start: "09:00",
									End:   "17:00",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "CFHolidayHours",
			base: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date:     String("2019-01-21"),
							IsClosed: NullableBool(true),
						},
						HolidayHours{
							Date:     String("2019-01-22"),
							IsClosed: NullableBool(true),
						},
					},
				},
			},
			new: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date:     String("2019-01-23"),
							IsClosed: NullableBool(true),
						},
						HolidayHours{
							Date:     String("2019-01-24"),
							IsClosed: NullableBool(true),
						},
					},
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				CustomEntity: CustomEntity{
					CFHolidayHours: &[]HolidayHours{
						HolidayHours{
							Date:     String("2019-01-23"),
							IsClosed: NullableBool(true),
						},
						HolidayHours{
							Date:     String("2019-01-24"),
							IsClosed: NullableBool(true),
						},
					},
				},
			},
		},
		{
			name: "Hours Closed Change",
			base: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
					}),
				},
			},
			new: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(false),
						}),
					}),
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(false),
						}),
					}),
				},
			},
		},
		{
			name: "Holiday Hours Date Change",
			base: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-21-2019"),
								IsClosed: NullableBool(true),
							},
							HolidayHours{
								Date:     String("01-23-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
			new: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-22-2019"),
								IsClosed: NullableBool(true),
							},
							HolidayHours{
								Date:     String("01-23-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-22-2019"),
								IsClosed: NullableBool(true),
							},
							HolidayHours{
								Date:     String("01-23-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
		},
		{
			name: "Hours Change",
			base: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Tuesday: NullableDayHours(&DayHours{
							OpenIntervals: &[]Interval{
								Interval{
									Start: "08:00",
									End:   "19:00",
								},
							},
						}),
						Wednesday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Thursday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Friday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Saturday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Sunday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-21-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
			new: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Tuesday: NullableDayHours(&DayHours{
							OpenIntervals: &[]Interval{
								Interval{
									Start: "08:00",
									End:   "20:00",
								},
							},
						}),
						Wednesday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Thursday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Friday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Saturday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Sunday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-21-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Tuesday: NullableDayHours(&DayHours{
							OpenIntervals: &[]Interval{
								Interval{
									Start: "08:00",
									End:   "20:00",
								},
							},
						}),
					}),
				},
			},
		},
		{
			name: "Hours Change (**DayHours) is nil -> null",
			base: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Tuesday: NullableDayHours(&DayHours{
							OpenIntervals: &[]Interval{
								Interval{
									Start: "08:00",
									End:   "19:00",
								},
							},
						}),
						Wednesday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Thursday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Friday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Saturday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Sunday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-21-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
			new: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday:    NullDayHours(),
						Tuesday:   NullDayHours(),
						Wednesday: NullDayHours(),
						Thursday:  NullDayHours(),
						Friday:    NullDayHours(),
						Saturday:  NullDayHours(),
						Sunday:    NullDayHours(),
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-21-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday:    NullDayHours(),
						Tuesday:   NullDayHours(),
						Wednesday: NullDayHours(),
						Thursday:  NullDayHours(),
						Friday:    NullDayHours(),
						Saturday:  NullDayHours(),
						Sunday:    NullDayHours(),
					}),
				},
			},
		},
		{
			name: "Hours Change null -> value",
			base: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday:    NullDayHours(),
						Tuesday:   NullDayHours(),
						Wednesday: NullDayHours(),
						Thursday:  NullDayHours(),
						Friday:    NullDayHours(),
						Saturday:  NullDayHours(),
						Sunday:    NullDayHours(),
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-21-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
			new: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Tuesday: NullableDayHours(&DayHours{
							OpenIntervals: &[]Interval{
								Interval{
									Start: "08:00",
									End:   "19:00",
								},
							},
						}),
						Wednesday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Thursday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Friday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Saturday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Sunday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-21-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Tuesday: NullableDayHours(&DayHours{
							OpenIntervals: &[]Interval{
								Interval{
									Start: "08:00",
									End:   "19:00",
								},
							},
						}),
						Wednesday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Thursday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Friday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Saturday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Sunday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
					}),
				},
			},
		},
		{
			name: "Hours No Change (showing nil is not the same as **DayHours(nil) )",
			base: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Tuesday: NullableDayHours(&DayHours{
							OpenIntervals: &[]Interval{
								Interval{
									Start: "08:00",
									End:   "19:00",
								},
							},
						}),
						Wednesday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Thursday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Friday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Saturday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						Sunday: NullableDayHours(&DayHours{
							IsClosed: NullableBool(true),
						}),
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-21-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
			new: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Hours: NullableHours(&Hours{
						Monday:    nil,
						Tuesday:   nil,
						Wednesday: nil,
						Thursday:  nil,
						Friday:    nil,
						Saturday:  nil,
						Sunday:    nil,
						HolidayHours: &[]HolidayHours{
							HolidayHours{
								Date:     String("01-21-2019"),
								IsClosed: NullableBool(true),
							},
						},
					}),
				},
			},
			isDiff: false,
		},

		{
			name: "Name",
			base: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Name: String("CTG"),
				},
			},
			new: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Name: String("CTG2"),
				},
			},
			isDiff: true,
			delta: &CustomLocationEntity{
				LocationEntity: LocationEntity{
					Name: String("CTG2"),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			delta, isDiff, _ := Diff(test.base, test.new)
			if isDiff != test.isDiff {
				t.Log(delta)
				t.Errorf("Expected isDiff: %t.\nGot: %t", test.isDiff, isDiff)
			} else if test.isDiff == false && delta != nil {
				t.Errorf("Expected isDiff: %t.\nGot delta: %v", test.isDiff, delta)
			} else if isDiff {
				if !reflect.DeepEqual(delta, test.delta) {
					t.Errorf("Expected delta: %v.\nGot: %v", test.delta, delta)
				}
			}
		})
	}
}

func TestInstanceOf(t *testing.T) {
	var (
		b = String("apple")
		i = instanceOf(b)
	)
	if _, ok := i.(*string); !ok {
		t.Error("Expected *string")
	}

	var (
		h = &[]HolidayHours{
			HolidayHours{
				IsClosed: NullableBool(true),
			},
		}
		j = instanceOf(h)
	)
	if _, ok := j.(*[]HolidayHours); !ok {
		t.Error("Expected *[]HolidayHours")
	}

	var (
		hours = NullableHours(&Hours{
			Monday: NullableDayHours(&DayHours{
				IsClosed: NullableBool(true),
			}),
		})
		k = instanceOf(hours)
	)
	if iHours, ok := k.(**Hours); !ok {
		t.Error("Expected **Hours")
	} else if *iHours == nil {
		t.Errorf("*Hours is nil")
	} else if GetHours(iHours) == nil {
		t.Error("**Hours instance is nil")
	}

	var (
		address = &Address{
			Line1: String("7900 Westpark"),
		}
		l = instanceOf(address)
	)
	if iAddress, ok := l.(*Address); !ok {
		t.Error("Expected *Address")
	} else if iAddress == nil {
		t.Error("*Address instance is nil")
	}
}
