package main

import yext "gopkg.in/yext/yext-go.v2"

func main() {
	locA := &yext.Location{}
	locB := &yext.Location{}

	locA.Diff(locB)
}
