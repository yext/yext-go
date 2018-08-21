package yext

// Diff calculates the differences between a base ACL (a) and a second ACL (b).
// The result returned is an ACL object with only parameters that are strictly different.
func (a ACL) Diff(b ACL) (delta ACL, diff bool) {
	roleDelta, roleDiff := a.Role.Diff(b.Role)
	if !roleDiff && a.On == b.On && a.AccessOn == b.AccessOn {
		return ACL{}, false
	}

	delta = ACL{}

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
		for idxB, aclB := range b {
			_, hasDelta := aclA.Diff(aclB)

			if !hasDelta {
				break
			}

			if idxB == len(b)-1 {
				return b, true
			}
		}
	}

	return nil, false
}
