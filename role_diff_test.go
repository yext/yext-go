package yext_test

import (
	"testing"

	"github.com/yext/yext-go"
)

var (
	blankRole = yext.Role{}

	exampleRole = yext.Role{
		Id:   yext.String("3"),
		Name: yext.String("Example Role"),
	}

	identicalRole = yext.Role{
		Id:   yext.String("3"),
		Name: yext.String("Example Role"),
	}

	differentIdRole = yext.Role{
		Id:   yext.String("4"),
		Name: yext.String("Example Role"),
	}

	differentNameRole = yext.Role{
		Id:   yext.String("3"),
		Name: yext.String("Example Role Two"),
	}
)

func TestDiffIdenticalRole(t *testing.T) {
	d, isDiff := exampleRole.Diff(identicalRole)
	if isDiff {
		t.Errorf("Expected diff to be false but was true, diff result: %v", d)
	}

	if d != blankRole {
		t.Errorf("Expected empty diff, but got %v", d)
	}
}

func TestDiffIdRole(t *testing.T) {
	d, isDiff := exampleRole.Diff(differentIdRole)
	if !isDiff {
		t.Errorf("Expected diff to be true but was false, diff result: %v", d)
	}

	expectedDiffRole := yext.Role{
		Id: differentIdRole.Id,
	}

	if d != expectedDiffRole {
		t.Errorf("Expected %v, but got %v", expectedDiffRole, d)
	}
}

func TestDiffNameRole(t *testing.T) {
	d, isDiff := exampleRole.Diff(differentNameRole)
	if !isDiff {
		t.Errorf("Expected diff to be true but was false, diff result: %v", d)
	}

	expectedDiffRole := yext.Role{
		Name: differentNameRole.Name,
	}

	if d != expectedDiffRole {
		t.Errorf("Expected a diff in name, but got %v", d)
	}
}
