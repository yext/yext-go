package yext

func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

func String(v string) *string {
	p := new(string)
	*p = v
	return p
}

func Float(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}

func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}
