package cpu

import (
	"n64emu/pkg/core/r4300i/inst"
	"n64emu/pkg/core/r4300i/reg"
	"n64emu/pkg/types"
)

// SLL rd, rt, sa
// The contents of general purpose register rt are shifted left by sa bits, inserting zeros into the low-order bits.
func sll(gpr *reg.GPR, inst *inst.InstR) {
	gpr.Write(inst.Rd, types.DoubleWord((int32(gpr.Read(inst.Rt)) << inst.Sa)))
}

// SRL rd, rt, sa
// The contents of general purpose register rt are shifted right by sa bits, inserting zeros into the high-order bits.
func srl(gpr *reg.GPR, inst *inst.InstR) {
	gpr.Write(inst.Rd, types.DoubleWord((gpr.Read(inst.Rt))>>inst.Sa))
}

// SRA rd, rt, sa
// Shifts the contents of register rt sa bits to the right, and sign-extends the high- order bits.
func sra(gpr *reg.GPR, inst *inst.InstR) {
	gpr.Write(inst.Rd, types.DoubleWord(int32(gpr.Read(inst.Rt))>>inst.Sa))
}
