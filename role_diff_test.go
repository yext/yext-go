package yext

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRole_Diff(t *testing.T) {
	tests := []struct {
		name      string
		roleA     Role
		roleB     Role
		wantDelta Role
		wantDiff  bool
	}{
		{
			name: "Identical Roles",
			roleA: Role{
				Id:   String("3"),
				Name: String("Example Role"),
			},
			roleB: Role{
				Id:   String("3"),
				Name: String("Example Role"),
			},
			wantDelta: Role{},
			wantDiff:  false,
		},
		{
			name: "Different 'Id' params in Roles",
			roleA: Role{
				Id:   String("3"),
				Name: String("Example Role"),
			},
			roleB: Role{
				Id:   String("4"),
				Name: String("Example Role"),
			},
			wantDelta: Role{
				Id: String("4"),
			},
			wantDiff: true,
		},
		{
			name: "Different 'Name' params in Roles",
			roleA: Role{
				Id:   String("3"),
				Name: String("Example Role"),
			},
			roleB: Role{
				Id:   String("3"),
				Name: String("Example Role Two"),
			},
			wantDelta: Role{
				Name: String("Example Role Two"),
			},
			wantDiff: true,
		},
	}

	for _, test := range tests {
		if gotDelta, gotDiff := test.roleA.Diff(test.roleB); !reflect.DeepEqual(test.wantDelta, gotDelta) || test.wantDiff != gotDiff {
			t.Error(fmt.Sprintf("test '%s' failed, got diff: %t, wanted diff: %t, got delta: %+v, wanted delta: %+v", test.name, test.wantDiff, gotDiff, test.wantDelta, gotDelta))
		}
	}
}
