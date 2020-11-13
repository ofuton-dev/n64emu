package cpu

import (
	"n64emu/pkg/core/mips/r4300i/reg"
	"n64emu/pkg/types"
)

type aluOutput struct {
	dest   types.Byte
	result types.DoubleWord
}

// SLL rd, rt, sa
// The contents of general purpose register rt are shifted left by sa bits, inserting zeros into the low-order bits.
func sll(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord((int32(gpr.Read(inst.Rt)) << inst.Sa)),
	}
}

// SRL rd, rt, sa
// The contents of general purpose register rt are shifted right by sa bits, inserting zeros into the high-order bits.
func srl(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord((gpr.Read(inst.Rt)) >> inst.Sa),
	}
}

// SRA rd, rt, sa
// Shifts the contents of register rt sa bits to the right, and sign-extends the high- order bits.
func sra(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord(int32(gpr.Read(inst.Rt)) >> inst.Sa),
	}
}

// SLLV rd, rt, rs
// Shifts the contents of register rt to the left and inserts 0 to the low-order bits.
func sllv(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord(int32(gpr.Read(inst.Rt)) << (inst.Rs & 0x1F)),
	}
}

// SRLV rd, rt, rs
// Shifts the contents of register rt to the right, and inserts 0 to the high-order bits.
func srlv(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord(int32(gpr.Read(inst.Rt)) >> (inst.Rs & 0x1F)),
	}
}

// SRAV rd, rt, rs
// Shifts the contents of register rt to the right and sign-extends the high-order bits.
func srav(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord(int32(gpr.Read(inst.Rt)) >> (inst.Rs & 0x1F)),
	}
}

func or(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord(gpr.Read(inst.Rs) | gpr.Read(inst.Rt)),
	}
}
