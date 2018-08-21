package yext_test

import (
	"reflect"
	"testing"

	"github.com/yext/yext-go"
)

var (
	blankACL = yext.ACL{}

	exampleACL = yext.ACL{
		Role: yext.Role{
			Id:   yext.String("3"),
			Name: yext.String("Example Role"),
		},
		On:       "12345",
		AccessOn: yext.ACCESS_FOLDER,
	}

	identicalACL = yext.ACL{
		Role: yext.Role{
			Id:   yext.String("3"),
			Name: yext.String("Example Role"),
		},
		On:       "12345",
		AccessOn: yext.ACCESS_FOLDER,
	}

	differentRoleACL = yext.ACL{
		Role: yext.Role{
			Id:   yext.String("4"),
			Name: yext.String("Example Role Two"),
		},
		On:       "12345",
		AccessOn: yext.ACCESS_FOLDER,
	}

	differentOnACL = yext.ACL{
		Role: yext.Role{
			Id:   yext.String("3"),
			Name: yext.String("Example Role"),
		},
		On:       "123456",
		AccessOn: yext.ACCESS_FOLDER,
	}

	differentAccessOnACL = yext.ACL{
		Role: yext.Role{
			Id:   yext.String("3"),
			Name: yext.String("Example Role"),
		},
		On:       "12345",
		AccessOn: yext.ACCESS_LOCATION,
	}

	exampleACLList = yext.ACLList{
		exampleACL,
	}

	identicalACLList = yext.ACLList{
		exampleACL,
	}

	differentItemACLList = yext.ACLList{
		differentRoleACL,
	}

	differentLenACLList = yext.ACLList{
		exampleACL,
		identicalACL,
	}

	differentLenIdenticalACLList = yext.ACLList{
		exampleACL,
		identicalACL,
	}
)

func TestDiffIdenticalACL(t *testing.T) {
	d, isDiff := exampleACL.Diff(identicalACL)
	if isDiff {
		t.Errorf("Expected diff to be false but was true, diff result: %v", d)
	}

	if d != blankACL {
		t.Errorf("Expected empty diff, but got %v", d)
	}
}

func TestDiffRoleACL(t *testing.T) {
	d, isDiff := exampleACL.Diff(differentRoleACL)
	if !isDiff {
		t.Errorf("Expected diff to be true but was false, diff result: %v", d)
	}

	expectedDiffACL := yext.ACL{
		Role: yext.Role{
			Id:   differentRoleACL.Id,
			Name: differentRoleACL.Name,
		},
	}

	if d != expectedDiffACL {
		t.Errorf("Expected diff in role, but got %v", d)
	}
}

func TestDiffOnACL(t *testing.T) {
	d, isDiff := exampleACL.Diff(differentOnACL)
	if !isDiff {
		t.Errorf("Expected diff to be true but was false, diff result: %v", d)
	}

	expectedDiffACL := yext.ACL{
		On: "123456",
	}

	if d != expectedDiffACL {
		t.Errorf("Expected diff in values, but got %v", d)
	}
}

func TestDiffAccessOnACL(t *testing.T) {
	d, isDiff := exampleACL.Diff(differentAccessOnACL)
	if !isDiff {
		t.Errorf("Expected diff to be true but was false, diff result: %v", d)
	}

	expectedDiffACL := yext.ACL{
		AccessOn: yext.ACCESS_LOCATION,
	}

	if d != expectedDiffACL {
		t.Errorf("Expected diff in values, but got %v", d)
	}
}

func TestDiffIdenticalACLList(t *testing.T) {
	d, isDiff := exampleACLList.Diff(identicalACLList)

	if isDiff {
		t.Errorf("Expected diff to be false but was true, diff result: %v", d)
	}

	if d != nil {
		t.Errorf("Expected nil diff, but got %v", d)
	}
}

func TestDiffIdenticalWithMultipleItemsACLList(t *testing.T) {
	d, isDiff := differentLenACLList.Diff(differentLenIdenticalACLList)

	if isDiff {
		t.Errorf("Expected diff to be false but was true, diff result: %v", d)
	}

	if d != nil {
		t.Errorf("Expected nil diff, but got %v", d)
	}
}

func TestDiffItemACLList(t *testing.T) {
	d, isDiff := exampleACLList.Diff(differentItemACLList)

	if !isDiff {
		t.Errorf("Expected diff to be true but was false, diff result: %v", d)
	}

	if !reflect.DeepEqual(d, differentItemACLList) {
		t.Errorf("Expected diff to be second ACLList, but got %v", d)
	}
}

func TestDiffLenACLList(t *testing.T) {
	d, isDiff := exampleACLList.Diff(differentLenACLList)

	if !isDiff {
		t.Errorf("Expected diff to be true but was false, diff result: %v", d)
	}

	if !reflect.DeepEqual(d, differentLenACLList) {
		t.Errorf("Expected diff to be second ACLList, but got %v", d)
	}
}
