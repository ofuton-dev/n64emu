package cpu

import (
	"fmt"
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
	fmt.Println("sra", inst.Rd)
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

// JR rs
// Jumps to the address of register rs, delayed by one instruction.
func jr(pc *types.DoubleWord, gpr *reg.GPR, inst *InstR) *aluOutput {
	*pc = gpr.Read(inst.Rs)
	return nil
}

// JALR rs, rd
// Jumps to the address of register rs, delayed by one instruction.
// Stores the address of the instruction following the delay slot to register rd.
// See also U10504EJ7V0UM00 p98
func jalr(pc *types.DoubleWord, gpr *reg.GPR, inst *InstR) *aluOutput {
	result := *pc + 8
	*pc = gpr.Read(inst.Rs)
	return &aluOutput{
		dest:   inst.Rd,
		result: result,
	}
}

// MFHI rd
// Transfers the contents of special register HI to register rd.
func mfhi(hi types.DoubleWord, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: hi,
	}
}

// MFLO rd
// Transfers the contents of special register LO to register rd.
func mflo(lo types.DoubleWord, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: lo,
	}
}

// MTHI rs
// Transfers the contents of register rs to special register HI.
func mthi(gpr *reg.GPR, hi *types.DoubleWord, inst *InstR) *aluOutput {
	// TODO: We need to do some investigation about write back timing
	*hi = types.DoubleWord(gpr.Read(inst.Rs))
	return nil
}

// MTLO rs
// Transfers the contents of register rs to special register LO.
func mtlo(gpr *reg.GPR, lo *types.DoubleWord, inst *InstR) *aluOutput {
	// TODO: We need to do some investigation about write back timing
	*lo = types.DoubleWord(gpr.Read(inst.Rs))
	return nil
}

// DSLLV rd, rt, rs
// Shifts the contents of register rt to the left, and inserts 0 to the low-order bits.
func dsllv(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord((gpr.Read(inst.Rt)) << (inst.Rs & 0x3F)),
	}
}

// DSRLV rd, rt, rs
// Shifts the contents of register rt to the right, and inserts 0 to the higher bits.
func dsrlv(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord((gpr.Read(inst.Rt)) >> (inst.Rs & 0x3F)),
	}
}

// DSRAV rd, rt, rs
// Shifts the contents of register rt to the right, and sign-extends the high-order bits.
func dsrav(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord(int64(gpr.Read(inst.Rt)) >> (inst.Rs & 0x3F)),
	}
}

// MULT rs, rt
// Multiplies the contents of register rs by the contents of register rt as a 32-bit signed integer.
// Number of required cycles 5
func mult(gpr *reg.GPR, hi *types.DoubleWord, lo *types.DoubleWord, inst *InstR) *aluOutput {
	result := types.DoubleWord(int64(gpr.Read(inst.Rt)) * int64(gpr.Read(inst.Rs)))
	// TODO: We need to do some investigation about write back timing
	// .     Should we add 20 cycle delay for 64bit mode?
	//       ref. https://en.wikipedia.org/wiki/R4000#Integer_execution
	// .     See also, https://github.com/ofuton-dev/n64emu/pull/18
	*hi = result >> 32
	*lo = result & 0xFFFFFFFF
	return nil
}

// MULTU rs, rt
// The contents of general purpose register rs and the contents of general purpose
// register rt are multiplied, treating both operands as 32-bit unsigned values.
// Number of required cycles 5
func multu(gpr *reg.GPR, hi *types.DoubleWord, lo *types.DoubleWord, inst *InstR) *aluOutput {
	result := types.DoubleWord(gpr.Read(inst.Rt) * gpr.Read(inst.Rs))
	// TODO: We need to do some investigation about write back timing
	// .     Should we add 20 cycle delay for 64bit mode?
	//       ref. https://en.wikipedia.org/wiki/R4000#Integer_execution
	// .     See also, https://github.com/ofuton-dev/n64emu/pull/18
	*hi = result >> 32
	*lo = result & 0xFFFFFFFF
	return nil
}

// DIV rs, rt
// Divides the contents of register rs by the contents of register rt. The operand
// is treated as a 32-bit signed integer.
// Number of required cycles 37
func div(gpr *reg.GPR, hi *types.DoubleWord, lo *types.DoubleWord, inst *InstR) *aluOutput {
	rs := int32(gpr.Read(inst.Rs))
	rt := int32(gpr.Read(inst.Rt))
	*hi = types.DoubleWord(rs / rt)
	*lo = types.DoubleWord(rs % rt)
	return nil
}

// DIVU rs, rt
// The contents of general purpose register rs are divided by the contents of general
// purpose register rt, treating both operands as unsigned integers.
// Number of required cycles 37
func divu(gpr *reg.GPR, hi *types.DoubleWord, lo *types.DoubleWord, inst *InstR) *aluOutput {
	rs := types.Word(gpr.Read(inst.Rs))
	rt := types.Word(gpr.Read(inst.Rt))
	*hi = types.DoubleWord(rs / rt)
	*lo = types.DoubleWord(rs % rt)
	return nil
}

func or(gpr *reg.GPR, inst *InstR) *aluOutput {
	return &aluOutput{
		dest:   inst.Rd,
		result: types.DoubleWord(gpr.Read(inst.Rs) | gpr.Read(inst.Rt)),
	}
}
