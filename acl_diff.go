package yext

// Diff calculates the differences between a base ACL (a) and a second ACL (b).
// The result returned is an ACL object with only parameters that are strictly different.
func (a ACL) Diff(b ACL) (delta *ACL, diff bool) {
	roleDelta, roleDiff := a.Role.Diff(b.Role)
	if !roleDiff && a.On == b.On && a.AccessOn == b.AccessOn {
		return nil, false
	}

	delta = &ACL{}

	if roleDiff {
		delta.Role = roleDelta
	}

	if a.On != b.On {
		delta.On = b.On
	}

	if a.AccessOn != b.AccessOn {
		delta.AccessOn = b.AccessOn
	}

	return delta, true
}

// Diff for ACLList calculates if ACLList a is identical to ACLList b (order is ignored).
// If any ACL in a is missing in b, return the entire list b.
func (a ACLList) Diff(b ACLList) (delta ACLList, diff bool) {
	if len(a) != len(b) {
		return b, true
	}

	for _, aclA := range a {
		var found bool
		for _, aclB := range b {
			if _, hasDelta := aclA.Diff(aclB); !hasDelta {
				found = true
			}
		}
		if !found {
			return b, true
		}
	}

	// Handle vice-versa for true equality
	for _, aclB := range b {
		var found bool
		for _, aclA := range a {
			if _, hasDelta := aclB.Diff(aclA); !hasDelta {
				found = true
			}
		}
		if !found {
			return b, true
		}
	}

	return nil, false
}
