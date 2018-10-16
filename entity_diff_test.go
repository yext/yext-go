package yext

import (
	"log"
	"reflect"
	"testing"
)

type diffTest struct {
	name     string
	property string
	isDiff   bool

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
		// Meta
		// diffTest{
		// 	name:     "Base Entity: not equal",
		// 	property: "BaseEntity",
		// 	baseValue: BaseEntity{
		// 		Meta: &EntityMeta{
		// 			Id:          String("CTG"),
		// 			CategoryIds: Strings([]string{"123"}),
		// 		},
		// 	},
		// 	newValue: BaseEntity{
		// 		Meta: &EntityMeta{
		// 			Id:          String("CTG"),
		// 			CategoryIds: Strings([]string{"123", "456"}),
		// 		},
		// 	},
		// 	isDiff: true,
		// 	deltaValue: BaseEntity{
		// 		Meta: &EntityMeta{
		// 			Id:          String("CTG"),
		// 			CategoryIds: Strings([]string{"123", "456"}),
		// 		},
		// 	},
		// },
		// diffTest{
		// 	name:     "Base Entity: equal",
		// 	property: "BaseEntity",
		// 	baseValue: BaseEntity{
		// 		Meta: &EntityMeta{
		// 			CategoryIds: Strings([]string{"123", "456"}),
		// 		},
		// 	},
		// 	newValue: BaseEntity{
		// 		Meta: &EntityMeta{
		// 			CategoryIds: Strings([]string{"123", "456"}),
		// 		},
		// 	},
		// 	isDiff: false,
		// },
		// // String tests
		// diffTest{
		// 	name:      "*String: equal (A)",
		// 	property:  "Name",
		// 	baseValue: String("Hupman's Hotdogs"),
		// 	newValue:  String("Hupman's Hotdogs"),
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:       "*String: not equal (B)",
		// 	property:   "Name",
		// 	baseValue:  String("Hupman's Hotdogs"),
		// 	newValue:   String("Bryan's Bagels"),
		// 	isDiff:     true,
		// 	deltaValue: String("Bryan's Bagels"),
		// },
		// diffTest{
		// 	name:       "*String: base is empty string, new is not (C)",
		// 	property:   "Name",
		// 	baseValue:  String(""),
		// 	newValue:   String("Bryan's Bagels"),
		// 	isDiff:     true,
		// 	deltaValue: String("Bryan's Bagels"),
		// },
		// diffTest{
		// 	name:       "*String: new is empty string, base is not (D)",
		// 	property:   "Name",
		// 	baseValue:  String("Hupman's Hotdogs"),
		// 	newValue:   String(""),
		// 	isDiff:     true,
		// 	deltaValue: String(""),
		// },
		// diffTest{
		// 	name:      "*String: both are empty (E)",
		// 	property:  "Name",
		// 	baseValue: String(""),
		// 	newValue:  String(""),
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:      "*String: both are nil (F)",
		// 	property:  "Name",
		// 	baseValue: nil,
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:      "*String: base is not nil, new is nil (G)",
		// 	property:  "Name",
		// 	baseValue: String("Bryan's Bagels"),
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:       "*String: base is nil, new is not nil (H)",
		// 	property:   "Name",
		// 	baseValue:  nil,
		// 	newValue:   String("Bryan's Bagels"),
		// 	isDiff:     true,
		// 	deltaValue: String("Bryan's Bagels"),
		// },
		// diffTest{
		// 	name:       "*String: base is nil, new is empty string (I)",
		// 	property:   "Name",
		// 	baseValue:  nil,
		// 	newValue:   String(""),
		// 	isDiff:     true,
		// 	deltaValue: String(""),
		// },
		// diffTest{
		// 	name:      "*String: base is empty string, new is nil (J)",
		// 	property:  "Name",
		// 	baseValue: String(""),
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:           "*String: base is nil (with nil is empty), new is empty string (K)",
		// 	property:       "Name",
		// 	baseValue:      nil,
		// 	newValue:       String(""),
		// 	baseNilIsEmpty: true,
		// 	isDiff:         false,
		// },
		// diffTest{
		// 	name:           "*String: base is nil (with nil is empty), new is empty string (with nil is empty) (L)",
		// 	property:       "Name",
		// 	baseValue:      nil,
		// 	newValue:       String(""),
		// 	baseNilIsEmpty: true,
		// 	newNilIsEmpty:  true,
		// 	isDiff:         false,
		// },
		// diffTest{
		// 	name:          "*String: base is empty string, new is nil (with nil is empty) (M)",
		// 	property:      "Name",
		// 	baseValue:     String(""),
		// 	newValue:      nil,
		// 	newNilIsEmpty: true,
		// 	isDiff:        false,
		// },
		// diffTest{
		// 	name:           "*String: base is empty string (with nil is empty), new is nil (with nil is empty) (N)",
		// 	property:       "Name",
		// 	baseValue:      String(""),
		// 	baseNilIsEmpty: true,
		// 	newValue:       nil,
		// 	newNilIsEmpty:  true,
		// 	isDiff:         false,
		// },
		//
		// //Float tests
		// diffTest{
		// 	name:      "*Float: equal (A)",
		// 	property:  "YearEstablished",
		// 	baseValue: Float(2018),
		// 	newValue:  Float(2018),
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:       "*Float: not equal (B)",
		// 	property:   "YearEstablished",
		// 	baseValue:  Float(2018),
		// 	newValue:   Float(2006),
		// 	isDiff:     true,
		// 	deltaValue: Float(2006),
		// },
		// diffTest{
		// 	name:       "*Float: base is 0, new is not 0 (C)",
		// 	property:   "YearEstablished",
		// 	baseValue:  Float(0),
		// 	newValue:   Float(2006),
		// 	isDiff:     true,
		// 	deltaValue: Float(2006),
		// },
		// diffTest{
		// 	name:       "*Float: base is not 0, new is 0 (D)",
		// 	property:   "YearEstablished",
		// 	baseValue:  Float(2006),
		// 	newValue:   Float(0),
		// 	isDiff:     true,
		// 	deltaValue: Float(0),
		// },
		// diffTest{
		// 	name:      "*Float: both are 0 (E)",
		// 	property:  "YearEstablished",
		// 	baseValue: Float(2018),
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:      "*Float: both are nil (F)",
		// 	property:  "YearEstablished",
		// 	baseValue: nil,
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:      "*Float: base is not 0, new is nil (G)",
		// 	property:  "YearEstablished",
		// 	baseValue: Float(1993),
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:       "*Float: base is nil, new is not 0 (H)",
		// 	property:   "YearEstablished",
		// 	baseValue:  nil,
		// 	newValue:   Float(1993),
		// 	isDiff:     true,
		// 	deltaValue: Float(1993),
		// },
		// diffTest{
		// 	name:       "*Float: base is nil, new is 0 (I)",
		// 	property:   "YearEstablished",
		// 	baseValue:  nil,
		// 	newValue:   Float(0),
		// 	isDiff:     true,
		// 	deltaValue: Float(0),
		// },
		// diffTest{
		// 	name:      "*Float: base is 0, new is nil (J)",
		// 	property:  "YearEstablished",
		// 	baseValue: Float(0),
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:           "*Float: base is nil (nil is empty), new is 0 (K)",
		// 	property:       "YearEstablished",
		// 	baseValue:      nil,
		// 	newValue:       Float(0),
		// 	baseNilIsEmpty: true,
		// 	isDiff:         false,
		// },
		// diffTest{
		// 	name:           "*Float: base is nil (nil is empty), new is 0 (nil is empty) (L)",
		// 	property:       "YearEstablished",
		// 	baseValue:      nil,
		// 	newValue:       Float(0),
		// 	baseNilIsEmpty: true,
		// 	newNilIsEmpty:  true,
		// 	isDiff:         false,
		// },
		// diffTest{
		// 	name:          "*Float: base is 0, new is nil (nil is empty) (M)",
		// 	property:      "YearEstablished",
		// 	baseValue:     Float(0),
		// 	newValue:      nil,
		// 	newNilIsEmpty: true,
		// 	isDiff:        false,
		// },
		// diffTest{
		// 	name:           "*Float: base is 0 (nil is empty), new is nil (nil is empty) (N)",
		// 	property:       "YearEstablished",
		// 	baseValue:      Float(0),
		// 	baseNilIsEmpty: true,
		// 	newValue:       nil,
		// 	newNilIsEmpty:  true,
		// 	isDiff:         false,
		// },
		//
		// // Bool tests
		// diffTest{
		// 	name:      "*Bool: both true (A)",
		// 	property:  "SuppressAddress",
		// 	baseValue: Bool(true),
		// 	newValue:  Bool(true),
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:      "*Bool: both false (A/E)",
		// 	property:  "SuppressAddress",
		// 	baseValue: Bool(false),
		// 	newValue:  Bool(false),
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:       "*Bool: not equal, base true, new false (B/D)",
		// 	property:   "SuppressAddress",
		// 	baseValue:  Bool(true),
		// 	newValue:   Bool(false),
		// 	isDiff:     true,
		// 	deltaValue: Bool(false),
		// },
		// diffTest{
		// 	name:       "*Bool: not equal, base is false, new is true (B/C)",
		// 	property:   "SuppressAddress",
		// 	baseValue:  Bool(false),
		// 	newValue:   Bool(true),
		// 	isDiff:     true,
		// 	deltaValue: Bool(true),
		// },
		// diffTest{
		// 	name:      "*Bool: both are nil (F)",
		// 	property:  "SuppressAddress",
		// 	baseValue: nil,
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:      "*Bool: base is non-zero, new is nil (G)",
		// 	property:  "SuppressAddress",
		// 	baseValue: Bool(true),
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:       "*Bool: base is nil, new is non-zero value (H)",
		// 	property:   "SuppressAddress",
		// 	baseValue:  nil,
		// 	newValue:   Bool(true),
		// 	isDiff:     true,
		// 	deltaValue: Bool(true),
		// },
		// diffTest{
		// 	name:       "*Bool: base is nil, new is zero value (I)",
		// 	property:   "SuppressAddress",
		// 	baseValue:  nil,
		// 	newValue:   Bool(false),
		// 	isDiff:     true,
		// 	deltaValue: Bool(false),
		// },
		// diffTest{
		// 	name:      "*Bool: base is zero value, new is nil (J)",
		// 	property:  "SuppressAddress",
		// 	baseValue: Bool(false),
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:           "*Bool: base is nil (nil is empty), new is zero value (K)",
		// 	property:       "SuppressAddress",
		// 	baseValue:      nil,
		// 	newValue:       Bool(false),
		// 	baseNilIsEmpty: true,
		// 	isDiff:         false,
		// },
		// diffTest{
		// 	name:           "*Bool: base is nil (nil is empty), new is zero value (nil is empty) (L)",
		// 	property:       "SuppressAddress",
		// 	baseValue:      nil,
		// 	newValue:       Bool(false),
		// 	baseNilIsEmpty: true,
		// 	newNilIsEmpty:  true,
		// 	isDiff:         false,
		// },
		// diffTest{
		// 	name:          "*Bool: base is zero value, new is nil (nil is empty) (L)",
		// 	property:      "SuppressAddress",
		// 	baseValue:     Bool(false),
		// 	newValue:      nil,
		// 	newNilIsEmpty: true,
		// 	isDiff:        false,
		// },
		// diffTest{
		// 	name:           "*Bool: base is zero value (nil is empty), new is nil (nil is empty) (L)",
		// 	property:       "SuppressAddress",
		// 	baseValue:      Bool(false),
		// 	newValue:       nil,
		// 	baseNilIsEmpty: true,
		// 	newNilIsEmpty:  true,
		// 	isDiff:         false,
		// },
		//
		// Struct Diffs
		// diffTest{
		// 	name:     "Address: equal (A)",
		// 	property: "Address",
		// 	baseValue: &Address{
		// 		Line1: String("7900 Westpark"),
		// 	},
		// 	newValue: &Address{
		// 		Line1: String("7900 Westpark"),
		// 	},
		// 	isDiff: false,
		// },
		// diffTest{
		// 	name:     "Address: line 1 and line 2 equal (A)",
		// 	property: "Address",
		// 	baseValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 		Line2: String("Suite 500"),
		// 	},
		// 	newValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 		Line2: String("Suite 500"),
		// 	},
		// 	isDiff: false,
		// },
		// diffTest{
		// 	name:     "Address: not equal (B)",
		// 	property: "Address",
		// 	baseValue: &Address{
		// 		Line1: String("7900 Westpark"),
		// 	},
		// 	newValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// 	isDiff: true,
		// 	deltaValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// },
		// diffTest{
		// 	name:     "Address: line 1 equal, line 2 not equal (B)",
		// 	property: "Address",
		// 	baseValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 		Line2: String("Suite 500"),
		// 	},
		// 	newValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 		Line2: String("Suite 700"),
		// 	},
		// 	isDiff: true,
		// 	deltaValue: &Address{
		// 		Line2: String("Suite 700"),
		// 	},
		// },
		// diffTest{
		// 	name:      "Address: base is empty struct, new is not empty (C)",
		// 	property:  "Address",
		// 	baseValue: &Address{},
		// 	newValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// 	isDiff: true,
		// 	deltaValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// },
		// diffTest{
		// 	name:     "Address: base is not empty struct, new is empty struct (D)",
		// 	property: "Address",
		// 	baseValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// 	newValue:   &Address{},
		// 	isDiff:     true,
		// 	deltaValue: &Address{},
		// },
		// diffTest{
		// 	name:       "List: base is not empty struct, new is empty struct (D)",
		// 	property:   "CFTextList",
		// 	baseValue:  &[]string{"a", "b"},
		// 	newValue:   &[]string{},
		// 	isDiff:     true,
		// 	deltaValue: &[]string{},
		// },
		// diffTest{
		// 	name:      "Address: both are empty struct (E)",
		// 	property:  "Address",
		// 	baseValue: &Address{},
		// 	newValue:  &Address{},
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:      "Address: both are nil (F)",
		// 	property:  "Address",
		// 	baseValue: nil,
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:     "Address: base is non-empty struct, new is nil (G)",
		// 	property: "Address",
		// 	baseValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// 	newValue: nil,
		// 	isDiff:   false,
		// },
		// diffTest{
		// 	name:      "Address: base is nil, new is non-empty struct (H)",
		// 	property:  "Address",
		// 	baseValue: nil,
		// 	newValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// 	isDiff: true,
		// 	deltaValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// },
		// diffTest{
		// 	name:       "Address: base is nil, new is empty struct (I)",
		// 	property:   "Address",
		// 	baseValue:  nil,
		// 	newValue:   &Address{},
		// 	isDiff:     true,
		// 	deltaValue: &Address{},
		// },
		// diffTest{
		// 	name:      "Address: base is empty struct, new is nil (J)",
		// 	property:  "Address",
		// 	baseValue: &Address{},
		// 	newValue:  nil,
		// 	isDiff:    false,
		// },
		// diffTest{
		// 	name:           "Address: base is nil (nil is empty), new is empty struct (K)",
		// 	property:       "Address",
		// 	baseValue:      nil,
		// 	newValue:       &Address{Line1: String(""), Line2: String(""), City: String(""), Sublocality: String(""), Region: String(""), PostalCode: String("")}, // NOT &Address{}
		// 	baseNilIsEmpty: true,
		// 	isDiff:         false,
		// },
		diffTest{
			name:     "Address: base is empty struct, new is nil (nis is empty) (?)",
			property: "Address",

			baseValue:     &Address{Line1: String(""), Line2: String(""), City: String(""), Sublocality: String(""), Region: String(""), PostalCode: String("")},
			newValue:      nil,
			newNilIsEmpty: true,
			isDiff:        false,
		},
		diffTest{
			name:          "Address: base is empty struct, new is nil (nil is empty) (?)",
			property:      "Address",
			baseValue:     &Address{},
			newValue:      nil,
			newNilIsEmpty: true,
			isDiff:        false,
		},
		// diffTest{
		// 	name:           "Address: base is nil (nil is empty), new is empty struct (?)",
		// 	property:       "Address",
		// 	baseValue:      nil,
		// 	baseNilIsEmpty: true,
		// 	newValue:       &Address{},
		// 	isDiff:         false,
		// },
		diffTest{
			name:           "Address: base is nil (nil is empty), new is empty struct (?)",
			property:       "Address",
			baseValue:      nil,
			baseNilIsEmpty: true,
			newValue:       &Address{Line1: String(""), Line2: String(""), City: String(""), Sublocality: String(""), Region: String(""), PostalCode: String("")},
			isDiff:         false,
		},
		diffTest{
			name:          "Address: base is nil, new is nil (nil is empty) (?)",
			property:      "Address",
			baseValue:     nil,
			newValue:      nil,
			newNilIsEmpty: true,
			isDiff:        false,
		},
		// // TODO: add L, M, N
		//
		// diffTest{
		// 	name:     "Address: base line1 is empty",
		// 	property: "Address",
		// 	baseValue: &Address{
		// 		Line1: String(""),
		// 	},
		// 	newValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// 	isDiff: true,
		// 	deltaValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// },
		// diffTest{
		// 	name:     "Address: new line1 is empty",
		// 	property: "Address",
		// 	baseValue: &Address{
		// 		Line1: String("7900 Westpark Dr"),
		// 	},
		// 	newValue: &Address{
		// 		Line1: String(""),
		// 	},
		// 	isDiff: true,
		// 	deltaValue: &Address{
		// 		Line1: String(""),
		// 	},
		// },
	}

	log.Println("&Address{}, true", isZeroValue(reflect.ValueOf(&Address{}), true))
	log.Println("&Address{}, false", isZeroValue(reflect.ValueOf(&Address{}), false))
	log.Println("Address{}, true", isZeroValue(reflect.ValueOf(Address{}), true))
	log.Println("Address{}, false", isZeroValue(reflect.ValueOf(Address{}), false))

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
			delta, isDiff := Diff(baseEntity, newEntity)
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

// test comparable to make sure it gets used
// UnorderedStrings
// text list
// hours
