// +build !debug

package assert

import "testing"

// not doing anything
func TestAssert_Release(t *testing.T) {
	Assert(true, "")
	Assert(false, "illegal processing")
}

// not doing anything
func TestAssertEq_Release(t *testing.T) {
	AssertEq(0, 0, "")
	AssertEq(0, 1, "illegal processing")
}

// not doing anything
func TestAssertNe_Release(t *testing.T) {
	AssertNe(0, 1, "")
	AssertNe(0, 0, "illegal processing")
}
