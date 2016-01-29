package yext_test

import (
	"testing"

	"github.com/yext/yext-go"
)

var (
	exampleUser = &yext.User{
		Id:           yext.String("ding@yext.com"),
		FirstName:    yext.String("dang"),
		LastName:     yext.String("dangle"),
		PhoneNumber:  yext.String("2025562637"),
		EmailAddress: yext.String("ding@yext.com"),
		UserName:     yext.String("ding@yext.com"),
		Password:     yext.String("terriblepassword"),
		ACLs: []yext.ACL{
			yext.ACL{
				Role: yext.Role{
					Id:   3,
					Name: yext.String("Example Role"),
				},
				On:       "12345",
				AccessOn: yext.ACCESS_FOLDER,
			},
		},
	}
	secondUser *yext.User
)

func TestDiffIdentical(t *testing.T) {
	secondUser = exampleUser.Copy()
	d, isDiff := exampleUser.Diff(secondUser)
	if isDiff {
		t.Errorf("Expected diff to be false but was true, diff result", d)
	}
	if d != nil {
		t.Errorf("Expected empty diff, but got %v", d)
	}
}

func TestStringChanges(t *testing.T) {
	secondUser = exampleUser.Copy()
	secondUser.UserName = yext.String("someotherdang")
	secondUser.FirstName = yext.String("john")
	secondUser.LastName = yext.String("dang")
	secondUser.PhoneNumber = yext.String("5555555555")
	secondUser.EmailAddress = yext.String("ding@ding.com")
	secondUser.Password = yext.String("someotherpassword")
	d, isDiff := exampleUser.Diff(secondUser)
	if !isDiff {
		t.Errorf("Expected diff true, was false, diff result: %v", d)
	} else {
		if d.GetPassword() != secondUser.GetPassword() {
			t.Errorf("Expected Password to be '%s' but was '%v%', diff result", d.GetPassword(), secondUser.GetPassword(), d)
		}
		if d.GetUserName() != secondUser.GetUserName() {
			t.Errorf("Expected UserName to be '%s' but was '%v%', diff result", d.GetUserName(), secondUser.GetUserName(), d)
		}

		if d.GetFirstName() != secondUser.GetFirstName() {
			t.Errorf("Expected FirstName to be '%s' but was '%v%', diff result", d.GetFirstName(), secondUser.GetFirstName(), d)
		}

		if d.GetLastName() != secondUser.GetLastName() {
			t.Errorf("Expected LastName to be '%s' but was '%v%', diff result", d.GetLastName(), secondUser.GetLastName(), d)
		}

		if d.GetPhoneNumber() != secondUser.GetPhoneNumber() {
			t.Errorf("Expected PhoneNumber to be '%s' but was '%v%', diff result", d.GetPhoneNumber(), secondUser.GetPhoneNumber(), d)
		}

		if d.GetEmailAddress() != secondUser.GetEmailAddress() {
			t.Errorf("Expected EmailAddress to be '%s' but was '%v%', diff result", d.GetEmailAddress(), secondUser.GetEmailAddress(), d)
		}
	}
}

func TestAppendACL(t *testing.T) {
	expectACL := yext.ACL{
		Role: yext.Role{
			Id:   123,
			Name: yext.String("Crazy Role"),
		},
		On:       "123456",
		AccessOn: yext.ACCESS_CUSTOMER,
	}
	secondUser = exampleUser.Copy()
	secondUser.ACLs = append(secondUser.ACLs, expectACL)

	d, isDiff := exampleUser.Diff(secondUser)
	if !isDiff {
		t.Errorf("Expected diff true, was false, diff result: %v", d)
	} else {
		if len(d.ACLs) != len(secondUser.ACLs) {
			t.Errorf("Expected ACL to be added, got %v, diff result:", d.ACLs, d)
		}

		hasCorrectACL := false
		for _, acl := range d.ACLs {
			if acl.On == expectACL.On {
				hasCorrectACL = true
			}
		}

		if !hasCorrectACL {
			t.Errorf("Expected ACLS to contain %v but didn't, diff result: %v", expectACL, d)
		}
	}
}

func TestDeleteACL(t *testing.T) {
	secondUser = exampleUser.Copy()
	secondUser.ACLs = []yext.ACL{}

	d, isDiff := exampleUser.Diff(secondUser)
	if !isDiff {
		t.Errorf("Expected diff true, was false, diff result: %v", d)
	} else {
		if len(d.ACLs) != len(secondUser.ACLs) {
			t.Errorf("Expected ACL to be deleted, got %v, diff result:", d.ACLs, d)
		}
	}
}

func TestModifyACL(t *testing.T) {
	secondUser = exampleUser.Copy()
	acl := secondUser.ACLs[0]
	acl.Role.Name = yext.String("Dingle Role")
	acl.On = "987456"
	secondUser.ACLs[0] = acl

	d, isDiff := exampleUser.Diff(secondUser)

	if !isDiff {
		t.Errorf("Expected diff true, was false, diff result: %v", d)
	} else if d == nil {
		t.Errorf("Expected non nill diff: '%v' res: '%v'", d, isDiff)
	} else {
		if d.ACLs == nil {
			t.Errorf("Expected ACLS!%v", d)
		}

		if secondUser.ACLs == nil {
			t.Errorf("Should have had ACLS!%v", secondUser)
		}

		if len(d.ACLs) != len(secondUser.ACLs) {
			t.Errorf("Expected ACL to be the same length, got %v, diff result:", d.ACLs, d)
		}

		if d.ACLs[0].GetName() != acl.GetName() {
			t.Errorf("Expected ACL Name to be modified, got %v, diff result:", d.ACLs, d)
		}

		if d.ACLs[0].On != acl.On {
			t.Errorf("Expected ACL On to be modified, got %v, diff result:", d.ACLs, d)
		}
	}
}
