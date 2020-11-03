package cpu

import (
	"n64emu/pkg/bus"
	"n64emu/pkg/core/r4300i/inst"
	"n64emu/pkg/core/r4300i/reg"
	"n64emu/pkg/types"
	"n64emu/pkg/util"
)

// CPU is cpu registers and bus accessor
type CPU struct {
	gpr   reg.GPR          // 32 64-bit general purpose registers, GPRs
	fpr   reg.FPR          // 32 64-bit floating-point operation registers, FPRs
	pc    types.DoubleWord // Program Counter, the PC register
	hi    types.DoubleWord // HI register, containing the integer multiply and divide highorder doubleword result
	lo    types.DoubleWord // LO register, containing the integer multiply and divide loworder doubleword result
	llBit bool             //Load/Link LLBit register
	fcr0  types.Word       // 32-bit floating-point Implementation/Revision register, FCR0
	fcr31 types.Word       // 32-bit floating-point Control/Status register, FCR31
	bus   bus.Bus          // Bus accessor
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

func (c *CPU) Step() {
	// TODO: We need to consider about `pipline`.
	//       Implement later here.

	opcode := c.fetch()
	op := inst.GetOp(opcode)

	switch op {
	// R type instructions
	case 0x00:
		instR := inst.DecodeR(opcode)
		switch instR.Funct {
		case 0x00:
			util.TODO("SLL")
		case 0x02:
			util.TODO("SRL")
		case 0x03:
			util.TODO("SRA")
		case 0x04:
			util.TODO("SLLV")
		case 0x06:
			util.TODO("SRLL")
		case 0x07:
			util.TODO("SRAV")
		case 0x08:
			util.TODO("JR")
		case 0x09:
			util.TODO("JALR")
		case 0x0D:
			util.TODO("BREAK")
		case 0x0F:
			util.TODO("SYNC")
		case 0x10:
			util.TODO("MFHI")
		case 0x11:
			util.TODO("MTHI")
		case 0x12:
			util.TODO("MFLO")
		case 0x13:
			util.TODO("MTLO")
		case 0x14:
			util.TODO("DSLLV")
		case 0x16:
			util.TODO("DSRLV")
		case 0x17:
			util.TODO("DSRAV")
		case 0x18:
			util.TODO("MULT")
		case 0x19:
			util.TODO("MULTU")
		case 0x1A:
			util.TODO("DIV")
		case 0x1B:
			util.TODO("DIVU")
		case 0x1C:
			util.TODO("DMULT")
		case 0x1D:
			util.TODO("DMULTU")
		case 0x1E:
			util.TODO("DDIV")
		case 0x1F:
			util.TODO("DDIVU")
		case 0x20:
			util.TODO("ADD")
		case 0x21:
			util.TODO("ADDU")
		case 0x22:
			util.TODO("SUB")
		case 0x23:
			util.TODO("SUBU")
		case 0x24:
			util.TODO("AND")
		case 0x25:
			c.gpr.Write(instR.Rd, c.gpr.Read(instR.Rs)|c.gpr.Read(instR.Rt))
		case 0x26:
			util.TODO("XOR")
		case 0x27:
			util.TODO("NOR")
		case 0x2A:
			util.TODO("SLT")
		case 0x2B:
			util.TODO("SLTU")
		case 0x2C:
			util.TODO("DADD")
		case 0x2D:
			util.TODO("DADDU")
		case 0x2E:
			util.TODO("DSUB")
		case 0x2F:
			util.TODO("DSUBU")
		case 0x34:
			util.TODO("TEQ")
		case 0x38:
			util.TODO("DSLL")
		case 0x3A:
			util.TODO("DSRL")
		case 0x3B:
			util.TODO("DSRA")
		case 0x3C:
			util.TODO("DSLL32")
		case 0x3E:
			util.TODO("DSRL32")
		case 0x3F:
			util.TODO("DSRA32")
		}
		// TODO: map other instructions
	case 0x01:
		util.TODO("Other inst")
	}
}
