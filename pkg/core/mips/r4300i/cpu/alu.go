package cpu

import (
	"n64emu/pkg/core/mips/r4300i/reg"
	"n64emu/pkg/types"
)

type destType byte

const (
	destTypeGPR destType = iota
	destTypeHi
	destTypeLo
)

type aluOutput struct {
	destType     destType
	destGPRIndex types.Byte
	result       types.DoubleWord
}

// SLL rd, rt, sa
// The contents of general purpose register rt are shifted left by sa bits, inserting zeros into the low-order bits.
func sll(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       types.DoubleWord((int32(gpr.Read(inst.Rt)) << inst.Sa)),
	}
}

// SRL rd, rt, sa
// The contents of general purpose register rt are shifted right by sa bits, inserting zeros into the high-order bits.
func srl(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       types.DoubleWord((gpr.Read(inst.Rt)) >> inst.Sa),
	}
}

// SRA rd, rt, sa
// Shifts the contents of register rt sa bits to the right, and sign-extends the high- order bits.
func sra(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       types.DoubleWord(int32(gpr.Read(inst.Rt)) >> inst.Sa),
	}
}

// SLLV rd, rt, rs
// Shifts the contents of register rt to the left and inserts 0 to the low-order bits.
func sllv(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       types.DoubleWord(int32(gpr.Read(inst.Rt)) << (inst.Rs & 0x1F)),
	}
}

// SRLV rd, rt, rs
// Shifts the contents of register rt to the right, and inserts 0 to the high-order bits.
func srlv(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       types.DoubleWord(int32(gpr.Read(inst.Rt)) >> (inst.Rs & 0x1F)),
	}
}

// SRAV rd, rt, rs
// Shifts the contents of register rt to the right and sign-extends the high-order bits.
func srav(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       types.DoubleWord(int32(gpr.Read(inst.Rt)) >> (inst.Rs & 0x1F)),
	}
}

// MFHI rd
// Transfers the contents of special register HI to register rd.
func mfhi(hi types.DoubleWord, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       hi,
	}
}

// MFLO rd
// Transfers the contents of special register LO to register rd.
func mflo(lo types.DoubleWord, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       lo,
	}
}

// MTHI rs
// Transfers the contents of register rs to special register HI.
func mthi(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType: destTypeHi,
		result:   types.DoubleWord(gpr.Read(inst.Rs)),
	}
}

// MTLO rs
// Transfers the contents of register rs to special register LO.
func mtlo(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType: destTypeLo,
		result:   types.DoubleWord(gpr.Read(inst.Rs)),
	}
}

// DSLLV rd, rt, rs
// Shifts the contents of register rt to the left, and inserts 0 to the low-order bits.
func dsllv(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       types.DoubleWord((gpr.Read(inst.Rt)) << (inst.Rs & 0x3F)),
	}
}

// DSRLV rd, rt, rs
// Shifts the contents of register rt to the right, and inserts 0 to the higher bits.
func dsrlv(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       types.DoubleWord((gpr.Read(inst.Rt)) >> (inst.Rs & 0x3F)),
	}
}

// DSRAV rd, rt, rs
// Shifts the contents of register rt to the right, and sign-extends the high-order bits.
func dsrav(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       types.DoubleWord(int64(gpr.Read(inst.Rt)) >> (inst.Rs & 0x3F)),
	}
}

func or(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		destType:     destTypeGPR,
		destGPRIndex: inst.Rd,
		result:       types.DoubleWord(gpr.Read(inst.Rs) | gpr.Read(inst.Rt)),
	}
}
