/*

Floating Point General Purpose Register(FGR) are reserved for floating point operations

Register map:
	F0 :  register reserved for floating point operations
	F1 :  register reserved for floating point operations
	F2 :  register reserved for floating point operations
	F3 :  register reserved for floating point operations
	F4 :  register reserved for floating point operations
	F5 :  register reserved for floating point operations
	F6 :  register reserved for floating point operations
	F7 :  register reserved for floating point operations
	F8 :  register reserved for floating point operations
	F9 :  register reserved for floating point operations
	F10:  register reserved for floating point operations
	F11:  register reserved for floating point operations
	F12:  register reserved for floating point operations
	F13:  register reserved for floating point operations
	F14:  register reserved for floating point operations
	F15:  register reserved for floating point operations
	F16:  register reserved for floating point operations
	F17:  register reserved for floating point operations
	F18:  register reserved for floating point operations
	F19:  register reserved for floating point operations
	F20:  register reserved for floating point operations
	F21:  register reserved for floating point operations
	F22:  register reserved for floating point operations
	F23:  register reserved for floating point operations
	F24:  register reserved for floating point operations
	F25:  register reserved for floating point operations
	F26:  register reserved for floating point operations
	F27:  register reserved for floating point operations
	F28:  register reserved for floating point operations
	F29:  register reserved for floating point operations
	F30:  register reserved for floating point operations
	F31:  register reserved for floating point operations
*/

package reg

const (
	NumOfRegsInFgr = 32
)

// Floating Point General Purpose Register
type FGR struct {
	// registers
	f [NumOfRegsInFgr]float64
}

// Read value of the register.
func (fgr *FGR) Read(index int) float64 {
	return fgr.f[index]
}

// Write value in register
func (fgr *FGR) Write(index int, value float64) {
	fgr.f[index] = value
}
