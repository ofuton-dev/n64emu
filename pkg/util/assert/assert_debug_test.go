// +build debug

package assert

import "testing"

// not doing anything
func TestAssert_Debug(t *testing.T) {
	t.Run("NoPanic", func(t *testing.T) {
		Assert(true, "")
	})
	t.Run("NoPanic", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != "assert message" {
				t.Errorf("got %v, want %v", err, "assert message")
			}
		}()
		Assert(false, "assert message")
	})
}

// not doing anything
func TestAssertEq_Debug(t *testing.T) {
	t.Run("NoPanic", func(t *testing.T) {
		AssertEq(0, 0, "")
	})
	t.Run("NoPanic", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != "assert message" {
				t.Errorf("got %v, want %v", err, "assert message")
			}
		}()
		AssertEq(0, 1, "assert message")
	})
}

// not doing anything
func TestAssertNe_Debug(t *testing.T) {
	t.Run("NoPanic", func(t *testing.T) {
		AssertNe(0, 1, "")
	})
	t.Run("NoPanic", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != "assert message" {
				t.Errorf("got %v, want %v", err, "assert message")
			}
		}()
		AssertNe(0, 0, "assert message")
	})
}
