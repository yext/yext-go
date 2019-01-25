package yext_test

// import (
// 	"fmt"
// 	"reflect"
// 	"testing"
//
// 	"github.com/yext/yext-go"
// )
//
// func TestRole_Diff(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		roleA     yext.Role
// 		roleB     yext.Role
// 		wantDelta yext.Role
// 		wantDiff  bool
// 	}{
// 		{
// 			name: "Identical Roles",
// 			roleA: yext.Role{
// 				Id:   yext.SingleString("3"),
// 				Name: yext.SingleString("Example Role"),
// 			},
// 			roleB: yext.Role{
// 				Id:   yext.SingleString("3"),
// 				Name: yext.SingleString("Example Role"),
// 			},
// 			wantDelta: yext.Role{},
// 			wantDiff:  false,
// 		},
// 		{
// 			name: "Different 'Id' params in Roles",
// 			roleA: yext.Role{
// 				Id:   yext.SingleString("3"),
// 				Name: yext.SingleString("Example Role"),
// 			},
// 			roleB: yext.Role{
// 				Id:   yext.SingleString("4"),
// 				Name: yext.SingleString("Example Role"),
// 			},
// 			wantDelta: yext.Role{
// 				Id: yext.SingleString("4"),
// 			},
// 			wantDiff: true,
// 		},
// 		{
// 			name: "Different 'Name' params in Roles",
// 			roleA: yext.Role{
// 				Id:   yext.SingleString("3"),
// 				Name: yext.SingleString("Example Role"),
// 			},
// 			roleB: yext.Role{
// 				Id:   yext.SingleString("3"),
// 				Name: yext.SingleString("Example Role Two"),
// 			},
// 			wantDelta: yext.Role{
// 				Name: yext.SingleString("Example Role Two"),
// 			},
// 			wantDiff: true,
// 		},
// 	}
//
// 	for _, test := range tests {
// 		if gotDelta, gotDiff := test.roleA.Diff(test.roleB); !reflect.DeepEqual(test.wantDelta, gotDelta) || test.wantDiff != gotDiff {
// 			t.Error(fmt.Sprintf("test '%s' failed, got diff: %t, wanted diff: %t, got delta: %+v, wanted delta: %+v", test.name, test.wantDiff, gotDiff, test.wantDelta, gotDelta))
// 		}
// 	}
// }
