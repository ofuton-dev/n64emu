/*

General Purpose Register(GPR) are reserved for integer operations while the other thirty-two register

Register map:
	R0 : always zero. Any attempts to modify this register silently fail.
	R1 (AT): assembler temporary. Free use.
	R2 (V0): arithmetic values, function return values.
	R3 (V1): arithmetic values, function return values.
	R4 (A0): parameter passing to subroutines. Formal but not rigid.
	R5 (A1): parameter passing to subroutines. Formal but not rigid.
	R6 (A2): parameter passing to subroutines. Formal but not rigid.
	R7 (A3): parameter passing to subroutines. Formal but not rigid.
	R8 (T0): scratch registers. CPU RAM.
	R9 (T1): scratch registers. CPU RAM.
	R10(T2): scratch registers. CPU RAM.
	R11(T3): scratch registers. CPU RAM.
	R12(T4): scratch registers. CPU RAM.
	R13(T5): scratch registers. CPU RAM.
	R14(T6): scratch registers. CPU RAM.
	R15(T7): scratch registers. CPU RAM.
	R16(S0): registers saved upon function protocol. Trash at will if you know how.
	R17(S1): registers saved upon function protocol. Trash at will if you know how.
	R18(S2): registers saved upon function protocol. Trash at will if you know how.
	R19(S3): registers saved upon function protocol. Trash at will if you know how.
	R20(S4): registers saved upon function protocol. Trash at will if you know how.
	R21(S5): registers saved upon function protocol. Trash at will if you know how.
	R22(S6): registers saved upon function protocol. Trash at will if you know how.
	R23(S7): registers saved upon function protocol. Trash at will if you know how.
	R24(T8): scratch registers. CPU RAM.
	R25(T9): scratch registers. CPU RAM.
	R26(K0):
	R27(K1):
	R28(GP):
	R29(SP): stack pointer
	R30(S8):
	R31(RA): return address from subroutine. Not pulled from 'stack'. Change at convenience.
*/

package reg

import "n64emu/pkg/types"

const (
	NumOfRegsInGpr = 32
)

// General Purpose Register
type GPR struct {
	// registers (r[0] is not used because it is always zero.)
	r [NumOfRegsInGpr]types.DoubleWord
}

// NewGPR is GPR constructor
func NewGPR() GPR {
	return GPR{
		r: [NumOfRegsInGpr]types.DoubleWord{},
	}
}

// Read value of the register.
func (gpr *GPR) Read(index types.Byte) types.DoubleWord {
	if index == 0 {
		return 0
	} else {
		return gpr.r[index]
	}
}

// Write value in register
func (gpr *GPR) Write(index types.Byte, value types.DoubleWord) {
	if index != 0 {
		gpr.r[index] = value
	}
}
