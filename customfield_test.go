package yext

import (
	"fmt"
	"testing"
)

func TestSetMultiOptionId(t *testing.T) {
	a := MultiOption([]string{"a", "b", "c"})
	a.SetOptionId("d")
	b := MultiOption([]string{"a", "b", "c", "d"})
	stringA, stringB := fmt.Sprint(a), fmt.Sprint(b)
	if stringA != stringB {
		t.Errorf("expected %v was %v", b, a)
	}
}

func TestUnsetMultiOptionId(t *testing.T) {
	a := MultiOption([]string{"a", "b", "c"})
	b := MultiOption([]string{"a", "c"})
	a.UnsetOptionId("b")
	stringA, stringB := fmt.Sprint(a), fmt.Sprint(b)
	if stringA != stringB {
		t.Errorf("expected %v was %v", b, a)
	}
}

func TestIsMultiOptionIdSet(t *testing.T) {
	a := MultiOption([]string{"a", "b", "c"})
	if !a.IsOptionIdSet("a") {
		t.Errorf("expected true but was false for 'a' in '%s' ", a)
	}
	if a.IsOptionIdSet("d") {
		t.Errorf("expected false but was true for 'd' in '%s'", a)
	}
}

func TestSetSingleOptionId(t *testing.T) {
	a := SingleOption("a")
	b := SingleOption("b")
	a.SetOptionId("b")
	stringA, stringB := fmt.Sprint(a), fmt.Sprint(b)
	if stringA != stringB {
		t.Errorf("expected '%v' was '%v'", b, a)
	}
}

func TestUnsetSingleOptionId(t *testing.T) {
	a := SingleOption("a")
	b := SingleOption("a")
	a.UnsetOptionId("d") // not currently set
	stringA, stringB := fmt.Sprint(a), fmt.Sprint(b)
	if stringA != stringB {
		t.Errorf("unset not the same: expected '%v' was '%v'", b, a)
	}
	a.UnsetOptionId("a")
	stringA, stringB = fmt.Sprint(a), ""
	if stringA != stringB {
		t.Errorf("unset not the same: expected '%v' was '%v'", b, a)
	}
}

func TestIsSingleOptionIdSet(t *testing.T) {
	a := SingleOption("a")
	if !a.IsOptionIdSet("a") {
		t.Errorf("expected true but was false for 'a' being set on '%s'", a)
	}
	if a.IsOptionIdSet("d") {
		t.Errorf("expected false but was true for 'd' being set on '%s'", a)
	}
}
