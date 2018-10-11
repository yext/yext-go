package yext

import (
	"reflect"
	"testing"
)

type diffTest struct {
	name       string
	property   string
	isDiff     bool
	baseValue  interface{}
	newValue   interface{}
	deltaValue interface{}
}

func SetValOnProperty(val interface{}, property string, entity Entity) {
	reflect.ValueOf(entity).Elem().FieldByName(property).Set(reflect.ValueOf(val))

}

func (c *CustomLocationEntity) SetValOnProperty(val interface{}, property string) *CustomLocationEntity {
	if val != nil {
		reflect.ValueOf(c).Elem().FieldByName(property).Set(reflect.ValueOf(val))
	}
	return c
}

func (c *CustomLocationEntity) SetName(name string) *CustomLocationEntity {
	SetValOnProperty(name, "Name", c)
	return c
}

func TestEntityDiff(t *testing.T) {
	/*
		ORDER MATTERS
		locA := &Location{
			Name: String("CTG"),
		}
		locB := &Location{}
		delta, isDiff := locA.Diff(locB)
		log.Println(isDiff)
		log.Println(delta)
		delta, isDiff = locB.Diff(locA)
		log.Println(isDiff)
		log.Println(delta)
	*/
	tests := []diffTest{
		diffTest{
			name:     "Base Entity (Entity Meta) is equal",
			property: "BaseEntity",
			baseValue: BaseEntity{
				Meta: &EntityMeta{
					CategoryIds: Strings([]string{"123", "456"}),
				},
			},
			newValue: BaseEntity{
				Meta: &EntityMeta{
					CategoryIds: Strings([]string{"123", "456"}),
				},
			},
			isDiff: false,
		},
		diffTest{
			name:       "String not equal",
			property:   "Name",
			baseValue:  String("Hupman's Hotdogs"),
			newValue:   String("Bryan's Bagels"),
			isDiff:     true,
			deltaValue: String("Bryan's Bagels"),
		},
		diffTest{
			name:      "*String equal",
			property:  "Name",
			baseValue: String("Hupman's Hotdogs"),
			newValue:  String("Hupman's Hotdogs"),
			isDiff:    false,
		},
		diffTest{
			name:       "*Float not equal",
			property:   "YearEstablished",
			baseValue:  Float(2018),
			newValue:   Float(1993),
			isDiff:     true,
			deltaValue: Float(1993),
		},
		diffTest{
			name:      "*Float equal",
			property:  "YearEstablished",
			baseValue: Float(2018),
			newValue:  Float(2018),
			isDiff:    false,
		},
		diffTest{
			name:       "*Bool not equal",
			property:   "SuppressAddress",
			baseValue:  Bool(true),
			newValue:   Bool(false),
			isDiff:     true,
			deltaValue: Bool(false),
		},
		diffTest{
			name:      "*Bool equal",
			property:  "SuppressAddress",
			baseValue: Bool(true),
			newValue:  Bool(true),
			isDiff:    false,
		},
		diffTest{
			name:     "Address Equal",
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
			name:     "Address Not Equal",
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
			name:     "Address Not Equal (New Value Non-Empty String)",
			property: "Address",
			baseValue: &Address{
				Line1: String(""),
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
			name:     "Address Non Equal (New Value Empty String)",
			property: "Address",
			baseValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
			newValue: &Address{
				Line1: String(""),
			},
			isDiff: true,
			deltaValue: &Address{
				Line1: String(""),
			},
		},
		diffTest{
			name:     "Address Equal (New Value nil)",
			property: "Address",
			baseValue: &Address{
				Line1: String("7900 Westpark Dr"),
			},
			newValue: nil,
			isDiff:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//a := new(CustomLocationEntity).SetName("Name A")
			var (
				a = new(CustomLocationEntity)
				b = new(CustomLocationEntity)
			)
			if test.property != "" {
				a.SetValOnProperty(test.baseValue, test.property)
				b.SetValOnProperty(test.newValue, test.property)
			}
			delta, isDiff := Diff(a, b)
			if isDiff != test.isDiff {
				t.Errorf("Expected isDiff: %t. Got: %t", test.isDiff, isDiff)
			} else if test.isDiff == false && delta != nil {
				t.Errorf("Expected isDiff: %t. Got delta: %v", test.isDiff, delta)
			} else if isDiff {
				expectedDelta := new(CustomLocationEntity)
				expectedDelta.SetValOnProperty(&EntityMeta{Id: String("")}, "Meta")
				if test.property != "" && test.deltaValue != nil {
					expectedDelta.SetValOnProperty(test.deltaValue, test.property)
				}
				if !reflect.DeepEqual(delta, expectedDelta) {
					t.Errorf("Expected delta: %v. Got: %v", expectedDelta, delta)
				}
			}
		})
	}

	// tests := []struct {
	// 	EntityA Entity
	// 	EntityB Entity
	// 	isDiff  bool
	// }{
	// 	{
	// 		EntityA: &LocationEntity{
	// 			BaseEntity: BaseEntity{
	// 				Meta: &EntityMeta{EntityType: ENTITYTYPE_LOCATION},
	// 			},
	// 		},
	// 		EntityB: &LocationEntity{
	// 			BaseEntity: BaseEntity{
	// 				Meta: &EntityMeta{EntityType: ENTITYTYPE_LOCATION},
	// 			},
	// 		},
	// 		isDiff: false,
	// 	},
	// 	{
	// 		EntityA: &LocationEntity{
	// 			BaseEntity: BaseEntity{
	// 				Meta: &EntityMeta{EntityType: ENTITYTYPE_LOCATION},
	// 			},
	// 		},
	// 		EntityB: &LocationEntity{
	// 			BaseEntity: BaseEntity{
	// 				Meta: &EntityMeta{EntityType: ENTITYTYPE_EVENT},
	// 			},
	// 		},
	// 		isDiff: true,
	// 	},
	// 	{
	// 		EntityA: &LocationEntity{
	// 			Name: String("Consulting"),
	// 		},
	// 		EntityB: &LocationEntity{
	// 			Name: String("Yext Consulting"),
	// 		},
	// 		isDiff: true,
	// 	},
	// 	{
	// 		EntityA: &LocationEntity{
	// 			Address: &Address{
	// 				Line1: String("7900 Westpark Dr"),
	// 			},
	// 		},
	// 		EntityB: &LocationEntity{
	// 			Address: &Address{
	// 				Line1: String("7900 Westpark St"),
	// 			},
	// 		},
	// 		isDiff: true,
	// 	},
	// 	{
	// 		EntityA: &LocationEntity{
	// 			Address: &Address{
	// 				Line1: String("7900 Westpark Dr"),
	// 			},
	// 		},
	// 		EntityB: &LocationEntity{
	// 			Address: &Address{
	// 				Line1: String("7900 Westpark Dr"),
	// 			},
	// 		},
	// 		isDiff: false,
	// 	},
	// 	{
	// 		EntityA: &LocationEntity{
	// 			Address: &Address{
	// 				Line1: String("7900 Westpark Dr"),
	// 			},
	// 		},
	// 		EntityB: &LocationEntity{},
	// 		isDiff:  false,
	// 	},
	// 	{
	// 		EntityA: &LocationEntity{},
	// 		EntityB: &LocationEntity{
	// 			Address: &Address{
	// 				Line1: String("7900 Westpark Dr"),
	// 			},
	// 		},
	// 		isDiff: true,
	// 	},
	// }
	// for _, test := range tests {
	// 	_, isDiff := Diff(test.EntityA, test.EntityB)
	// 	if isDiff != test.isDiff {
	// 		t.Errorf("Expected diff to be %t was %t", test.isDiff, isDiff)
	// 	}
	// }
}

//SetValOnProperty("Name", "")
//https://blog.golang.org/subtests
// subtests
