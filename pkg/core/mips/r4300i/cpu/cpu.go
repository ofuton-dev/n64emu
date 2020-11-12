package cpu

import (
	"n64emu/pkg/core/bus"
	"n64emu/pkg/core/mips/r4300i/reg"
	"n64emu/pkg/types"
)

// CPU is cpu registers and bus accessor
type CPU struct {
	gpr      reg.GPR          // 32 64-bit general purpose registers, GPRs
	fpr      reg.FPR          // 32 64-bit floating-point operation registers, FPRs
	pc       types.DoubleWord // Program Counter, the PC register
	hi       types.DoubleWord // HI register, containing the integer multiply and divide highorder doubleword result
	lo       types.DoubleWord // LO register, containing the integer multiply and divide loworder doubleword result
	llBit    bool             //Load/Link LLBit register
	fcr0     types.Word       // 32-bit floating-point Implementation/Revision register, FCR0
	fcr31    types.Word       // 32-bit floating-point Control/Status register, FCR31
	bus      bus.Bus          // Bus accessor
	pipeline Pipeline
}

// NewCPU is CPU constructor
func NewCPU(bus bus.Bus) *CPU {
	// TODO: Please check default value after power up.
	cpu := &CPU{
		gpr:   reg.NewGPR(),
		fpr:   reg.NewFGR(),
		pc:    0,
		hi:    0,
		lo:    0,
		llBit: false,
		fcr0:  0,
		fcr31: 0,
		bus:   bus,
	}
	return cpu
}

func (c *CPU) endian() types.Endianness {
	// TODO: For now, return only `BIG`.
	return types.Big
}

func (c *CPU) fetch() types.Word {
	data := c.bus.ReadWord(c.endian(), types.Word(c.pc))
	c.pc += 4
	return data
}

// Step runs 1 pclk cycle CPU
func (c *CPU) Step() {
	// TODO: We need to consider about `pipline`.
	//       Implement later here.
	c.pipeline.step(&c.gpr, c.fetch)
}

// RunUntil runs CPU until specified cycles
func (c *CPU) RunUntil(cycle types.Word) {
	for cycle > 0 {
		c.Step()
		cycle--
	}
}
