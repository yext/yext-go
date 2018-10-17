package yext

import (
	"fmt"
	"reflect"
	"testing"
)

func TestACL_Diff(t *testing.T) {
	tests := []struct {
		name      string
		aclA      ACL
		aclB      ACL
		wantDelta *ACL
		wantDiff  bool
	}{
		{
			name: "Identical ACLs",
			aclA: ACL{
				Role: Role{
					Id:   String("3"),
					Name: String("Example Role"),
				},
				On:       "12345",
				AccessOn: ACCESS_FOLDER,
			},
			aclB: ACL{
				Role: Role{
					Id:   String("3"),
					Name: String("Example Role"),
				},
				On:       "12345",
				AccessOn: ACCESS_FOLDER,
			},
			wantDelta: nil,
			wantDiff:  false,
		},
		{
			name: "Different Roles in ACL",
			aclA: ACL{
				Role: Role{
					Id:   String("3"),
					Name: String("Example Role"),
				},
				On:       "12345",
				AccessOn: ACCESS_FOLDER,
			},
			aclB: ACL{
				Role: Role{
					Id:   String("4"),
					Name: String("Example Role Two"),
				},
				On:       "12345",
				AccessOn: ACCESS_FOLDER,
			},
			wantDelta: &ACL{
				Role: Role{
					Id:   String("4"),
					Name: String("Example Role Two"),
				},
			},
			wantDiff: true,
		},
		{
			name: "Different 'On' params in ACL",
			aclA: ACL{
				Role: Role{
					Id:   String("3"),
					Name: String("Example Role"),
				},
				On:       "12345",
				AccessOn: ACCESS_FOLDER,
			},
			aclB: ACL{
				Role: Role{
					Id:   String("3"),
					Name: String("Example Role"),
				},
				On:       "123456",
				AccessOn: ACCESS_FOLDER,
			},
			wantDelta: &ACL{
				On: "123456",
			},
			wantDiff: true,
		},
		{
			name: "Different 'AccessOn' params in ACL",
			aclA: ACL{
				Role: Role{
					Id:   String("3"),
					Name: String("Example Role"),
				},
				On:       "12345",
				AccessOn: ACCESS_FOLDER,
			},
			aclB: ACL{
				Role: Role{
					Id:   String("3"),
					Name: String("Example Role"),
				},
				On:       "12345",
				AccessOn: ACCESS_LOCATION,
			},
			wantDelta: &ACL{
				AccessOn: ACCESS_LOCATION,
			},
			wantDiff: true,
		},
	}

	for _, test := range tests {
		if gotDelta, gotDiff := test.aclA.Diff(test.aclB); !reflect.DeepEqual(test.wantDelta, gotDelta) || test.wantDiff != gotDiff {
			t.Error(fmt.Sprintf("test '%s' failed, got diff: %t, wanted diff: %t, got delta: %+v, wanted delta: %+v", test.name, test.wantDiff, gotDiff, test.wantDelta, gotDelta))
		}
	}
}

func TestACLList_Diff(t *testing.T) {
	tests := []struct {
		name      string
		aclListA  ACLList
		aclListB  ACLList
		wantDelta ACLList
		wantDiff  bool
	}{
		{
			name: "Identical ACLLists",
			aclListA: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("4"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
			},
			aclListB: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("4"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
			},
			wantDelta: nil,
			wantDiff:  false,
		},
		{
			name: "Identical ACLs in ACLLists",
			aclListA: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
			},
			aclListB: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
			},
			wantDelta: nil,
			wantDiff:  false,
		},
		{
			name: "Different Length in ACLLists",
			aclListA: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("4"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
			},
			aclListB: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("4"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
				ACL{
					Role: Role{
						Id:   String("5"),
						Name: String("Example Role Three"),
					},
					On:       "1234567",
					AccessOn: ACCESS_LOCATION,
				},
			},
			wantDelta: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("4"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
				ACL{
					Role: Role{
						Id:   String("5"),
						Name: String("Example Role Three"),
					},
					On:       "1234567",
					AccessOn: ACCESS_LOCATION,
				},
			},
			wantDiff: true,
		},
		{
			name: "Different Items in ACLLists",
			aclListA: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("4"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
			},
			aclListB: ACLList{
				ACL{
					Role: Role{
						Id:   String("33"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("44"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
			},
			wantDelta: ACLList{
				ACL{
					Role: Role{
						Id:   String("33"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("44"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
			},
			wantDiff: true,
		},
		{
			name: "Some Identical and Some Different Items in ACLLists",
			aclListA: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("4"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
			},
			aclListB: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("4"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
				ACL{
					Role: Role{
						Id:   String("5"),
						Name: String("Example Role Three"),
					},
					On:       "1234567",
					AccessOn: ACCESS_LOCATION,
				},
			},
			wantDelta: ACLList{
				ACL{
					Role: Role{
						Id:   String("3"),
						Name: String("Example Role"),
					},
					On:       "12345",
					AccessOn: ACCESS_FOLDER,
				},
				ACL{
					Role: Role{
						Id:   String("4"),
						Name: String("Example Role Two"),
					},
					On:       "123456",
					AccessOn: ACCESS_LOCATION,
				},
				ACL{
					Role: Role{
						Id:   String("5"),
						Name: String("Example Role Three"),
					},
					On:       "1234567",
					AccessOn: ACCESS_LOCATION,
				},
			},
			wantDiff: true,
		},
	}

	for _, test := range tests {
		if gotDelta, gotDiff := test.aclListA.Diff(test.aclListB); !reflect.DeepEqual(test.wantDelta, gotDelta) || test.wantDiff != gotDiff {
			t.Error(fmt.Sprintf("test '%s' failed, got diff: %t, wanted diff: %t, got delta: %+v, wanted delta: %+v", test.name, test.wantDiff, gotDiff, test.wantDelta, gotDelta))
		}
	}
}
