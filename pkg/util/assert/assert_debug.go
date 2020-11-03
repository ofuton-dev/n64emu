// +build debug

package assert

import "reflect"

// panic if 'cond' is false
func Assert(cond bool, msg interface{}) {
	if !cond {
		panic(msg)
	}
}

// panic if 'a' is not equal to 'b'
func AssertEq(a interface{}, b interface{}, msg interface{}) {
	if !reflect.DeepEqual(a, b) {
		panic(msg)
	}
}

// panic if 'a' is equal to 'b'
func AssertNe(a interface{}, b interface{}, msg interface{}) {
	if reflect.DeepEqual(a, b) {
		panic(msg)
	}
}
