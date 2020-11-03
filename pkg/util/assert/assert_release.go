// +build !debug

package assert

// panic if 'cond' is false
func Assert(cond bool, msg interface{}) {
	// nothing
}

// panic if 'a' is not equal to 'b'
func AssertEq(a interface{}, b interface{}, msg interface{}) {
	// nothing
}

// panic if 'a' is equal to 'b'
func AssertNe(a interface{}, b interface{}, msg interface{}) {
	// nothing
}
