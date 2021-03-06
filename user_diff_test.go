package yext

import (
	"reflect"
	"testing"
)

var (
	exampleUser = &User{
		Id:           String("ding@yext.com"),
		FirstName:    String("dang"),
		LastName:     String("dangle"),
		PhoneNumber:  String("2025562637"),
		EmailAddress: String("ding@yext.com"),
		UserName:     String("ding@yext.com"),
		Password:     String("terriblepassword"),
		ACLs: []ACL{
			ACL{
				Role: Role{
					Id:   String("3"),
					Name: String("Example Role"),
				},
				On:       "12345",
				AccessOn: ACCESS_FOLDER,
			},
		},
	}
	secondUser *User
)

func TestStringChanges(t *testing.T) {
	secondUser = exampleUser.Copy()
	secondUser.UserName = String("someotherdang")
	secondUser.FirstName = String("john")
	secondUser.LastName = String("dang")
	secondUser.PhoneNumber = String("5555555555")
	secondUser.EmailAddress = String("ding@ding.com")
	secondUser.Password = String("someotherpassword")
	d, isDiff := exampleUser.Diff(secondUser)
	if !isDiff {
		t.Errorf("Expected diff true, was false, diff result: %v", d)
	} else {
		if d.GetPassword() != secondUser.GetPassword() {
			t.Errorf("Expected Password to be '%s' but was '%s', diff result: %v", d.GetPassword(), secondUser.GetPassword(), d)
		}
		if d.GetUserName() != secondUser.GetUserName() {
			t.Errorf("Expected UserName to be '%s' but was '%s', diff result: %v", d.GetUserName(), secondUser.GetUserName(), d)
		}

		if d.GetFirstName() != secondUser.GetFirstName() {
			t.Errorf("Expected FirstName to be '%s' but was '%s', diff result: %v", d.GetFirstName(), secondUser.GetFirstName(), d)
		}

		if d.GetLastName() != secondUser.GetLastName() {
			t.Errorf("Expected LastName to be '%s' but was '%s', diff result: %v", d.GetLastName(), secondUser.GetLastName(), d)
		}

		if d.GetPhoneNumber() != secondUser.GetPhoneNumber() {
			t.Errorf("Expected PhoneNumber to be '%s' but was '%s', diff result: %v", d.GetPhoneNumber(), secondUser.GetPhoneNumber(), d)
		}

		if d.GetEmailAddress() != secondUser.GetEmailAddress() {
			t.Errorf("Expected EmailAddress to be '%s' but was '%s', diff result: %v", d.GetEmailAddress(), secondUser.GetEmailAddress(), d)
		}
	}
}

func TestAppendACL(t *testing.T) {
	expectACL := ACL{
		Role: Role{
			Id:   String("123"),
			Name: String("Crazy Role"),
		},
		On:       "123456",
		AccessOn: ACCESS_ACCOUNT,
	}
	secondUser = exampleUser.Copy()
	secondUser.ACLs = append(secondUser.ACLs, expectACL)

	d, isDiff := exampleUser.Diff(secondUser)
	if !isDiff {
		t.Errorf("Expected diff true, was false, diff result: %v", d)
	} else {
		if len(d.ACLs) != len(secondUser.ACLs) {
			t.Errorf("Expected ACL to be added, got %v, diff result: %v", d.ACLs, d)
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
	secondUser.ACLs = []ACL{}

	d, isDiff := exampleUser.Diff(secondUser)
	if !isDiff {
		t.Errorf("Expected diff true, was false, diff result: %v", d)
	} else {
		if len(d.ACLs) != len(secondUser.ACLs) {
			t.Errorf("Expected ACL to be deleted, got %v, diff result: %v", d.ACLs, d)
		}
	}
}

func TestModifyACL(t *testing.T) {
	secondUser = exampleUser.Copy()
	acl := secondUser.ACLs[0]
	acl.Role.Name = String("Dingle Role")
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
			t.Errorf("Expected ACL to be the same length, got %v, diff result: %v", d.ACLs, d)
		}

		if d.ACLs[0].GetName() != acl.GetName() {
			t.Errorf("Expected ACL Name to be modified, got %v, diff result: %v", d.ACLs, d)
		}

		if d.ACLs[0].On != acl.On {
			t.Errorf("Expected ACL On to be modified, got %v, diff result: %v", d.ACLs, d)
		}
	}
}

func TestDiff(t *testing.T) {
	type test struct {
		A, B   *User
		IsDiff bool
		Diff   *User
	}

	tests := []test{
		test{
			A:      &User{},
			B:      &User{},
			IsDiff: false,
			Diff:   nil,
		},
		test{
			A:      &User{},
			B:      &User{},
			IsDiff: false,
			Diff:   nil,
		},
		test{
			A:      &User{ACLs: []ACL{}},
			B:      &User{ACLs: []ACL{}},
			IsDiff: false,
			Diff:   nil,
		},
		test{
			A:      &User{ACLs: []ACL{ACL{On: "foo", AccessOn: ACCESS_LOCATION, AccountId: "123"}}},
			B:      &User{ACLs: []ACL{ACL{On: "foo", AccessOn: "LOCATION", AccountId: "123"}}},
			IsDiff: false,
			Diff:   nil,
		},
		test{
			A:      &User{ACLs: []ACL{ACL{Role: Role{Id: String("228"), Name: String("Farmers Agent Basic Users")}}}},
			B:      &User{ACLs: []ACL{ACL{Role: Role{Id: String("228"), Name: String("Farmers Agent Basic Users")}}}},
			IsDiff: false,
			Diff:   nil,
		},
		test{
			A: &User{
				Id:           String("jdoe@yext.com"),
				FirstName:    String("jane"),
				LastName:     String("doe"),
				PhoneNumber:  String("2025562637"),
				EmailAddress: String("jdoe@yext.com"),
				UserName:     String("jdoe@yext.com"),
				Password:     String("sekret"),
				ACLs: []ACL{
					ACL{
						Role: Role{
							Id:   String("3"),
							Name: String("Example Role"),
						},
						On:       "12345",
						AccessOn: ACCESS_FOLDER,
					},
				},
			},
			B: &User{
				Id:           String("jdoe@yext.com"),
				FirstName:    String("jane"),
				LastName:     String("doe"),
				PhoneNumber:  String("2025562637"),
				EmailAddress: String("jdoe@yext.com"),
				UserName:     String("jdoe@yext.com"),
				Password:     String("sekret"),
				ACLs: []ACL{
					ACL{
						Role: Role{
							Id:   String("3"),
							Name: String("Example Role"),
						},
						On:       "12345",
						AccessOn: ACCESS_FOLDER,
					},
				},
			},
			IsDiff: false,
			Diff:   nil,
		},
	}

	for _, tst := range tests {
		diff, isDiff := tst.A.Diff(tst.B)
		if isDiff != tst.IsDiff || !reflect.DeepEqual(diff, tst.Diff) {
			t.Errorf("\nA: %s\nB: %s\nWanted:%s\nGot:   %s", tst.A, tst.B, tst.Diff, diff)
		}
	}
}
