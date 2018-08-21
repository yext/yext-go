package yext

// Diff performs a diff of two roles and their parameters.
func (a Role) Diff(b Role) (Role, bool) {
	if a.GetId() == b.GetId() && a.GetName() == b.GetName() {
		return Role{}, false
	}

	delta := Role{}
	if a.GetId() != b.GetId() {
		delta.Id = b.Id
	}

	if a.GetName() != b.GetName() {
		delta.Name = b.Name
	}

	return delta, true
}
