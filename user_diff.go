package yext

import "reflect"

// Diff calculates the differences betwee a base User (a) and a proposed set of
// changes represented by a second User (b). The diffing logic will ignore fields in
// (b) location that aren't set (nil).  This characteristic makes the function ideal
// for calculating the minimal PUT required to up date an object via the API.
func (a *User) Diff(b *User) (d *User, diff bool) {
	diff, d = false, &User{Id: String(a.GetId())}

	var (
		aValue, bValue = reflect.ValueOf(a).Elem(), reflect.ValueOf(b).Elem()
		aType          = reflect.TypeOf(*a)
		numA           = aValue.NumField()
	)

	for i := 0; i < numA; i++ {
		var (
			nameA = aType.Field(i).Name
			valA  = aValue.Field(i)
			valB  = bValue.Field(i)
		)

		if valB.IsNil() || nameA == "Id" {
			continue
		}

		if nameA == "ACLs" {
			if len(a.ACLs) != len(b.ACLs) {
				diff = true
				d.ACLs = b.ACLs
			} else {
				aACLHash, bACLHash := make(map[string]bool), make(map[string]bool)
				for j := 0; j < len(a.ACLs); j++ {
					aACLHash[a.ACLs[j].Hash()] = true
					bACLHash[b.ACLs[j].Hash()] = true
				}

				for bHash, _ := range bACLHash {
					if _, ok := aACLHash[bHash]; !ok {
						diff = true
						d.ACLs = b.ACLs
						break
					}
				}

				for aHash, _ := range aACLHash {
					if _, ok := bACLHash[aHash]; !ok {
						diff = true
						d.ACLs = b.ACLs
						break
					}
				}
			}
		} else if !reflect.DeepEqual(valA.Interface(), valB.Interface()) {
			diff = true
			reflect.ValueOf(d).Elem().FieldByName(nameA).Set(valB)
		}
	}
	if !diff {
		// ensure that d remains nil if nothing has changed
		d = nil
	}
	return d, diff
}

func (u *User) Copy() (n *User) {
	n = &User{
		Id:           String(u.GetId()),
		FirstName:    String(u.GetFirstName()),
		LastName:     String(u.GetLastName()),
		PhoneNumber:  String(u.GetPhoneNumber()),
		EmailAddress: String(u.GetEmailAddress()),
		UserName:     String(u.GetUserName()),
		Password:     String(u.GetPassword()),
		ACLs:         make([]ACL, len(u.ACLs)),
	}

	for i := 0; i < len(u.ACLs); i++ {
		uACL := u.ACLs[i]
		n.ACLs[i] = ACL{
			Role: Role{
				Id:   uACL.Id,
				Name: String(uACL.GetName()),
			},
			On:       uACL.On,
			AccessOn: uACL.AccessOn,
		}
	}

	return n
}
