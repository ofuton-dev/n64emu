package cpu

import (
	"n64emu/pkg/core/bus"
	"n64emu/pkg/core/mips/r4300i/reg"
	"n64emu/pkg/types"
	"n64emu/pkg/util"
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
	pipeline *Pipeline
}

// NewCPU is CPU constructor
func NewCPU(bus bus.Bus) *CPU {
	// TODO: Please check default value after power up.
	cpu := &CPU{
		gpr:      reg.NewGPR(),
		fpr:      reg.NewFGR(),
		pc:       0,
		hi:       0,
		lo:       0,
		llBit:    false,
		fcr0:     0,
		fcr31:    0,
		bus:      bus,
		pipeline: NewPipeline(bus),
	}
	return cpu
}

func (c *CPU) endian() types.Endianness {
	// TODO: For now, return only `BIG`.
	return types.Big
}

func (c *CPU) fetch(addr types.DoubleWord) types.Word {
	data := c.bus.ReadWord(c.endian(), types.Word(addr))
	return data
}

// Step runs 1 pclk cycle CPU
func (c *CPU) Step() {
	// TODO: We need to consider about `pipline`.
	//       Implement later here.
	c.pipeline.step(c.endian(), &c.pc, &c.gpr, c.execute, c.fetch)
}

// RunUntil runs CPU until specified cycles
func (c *CPU) RunUntil(cycle types.Word) {
	for cycle > 0 {
		c.Step()
		cycle--
	}
}

func (c *CPU) trapIntegerOverflow() {
	util.TODO("Please implement integer overflow excrption.")
}

func (c *CPU) execute(opcode types.Word) *aluOutput {
	op := GetOp(opcode)

	instI := DecodeI(opcode)

	switch op {
	// R type instructions
	// SPECIAL
	case 0x00:
		instR := DecodeR(opcode)
		switch instR.Funct {
		case 0x00: // SLL
			return sll(&c.gpr, &instR)
		case 0x02: // SRL
			return srl(&c.gpr, &instR)
		case 0x03: // SRA
			return sra(&c.gpr, &instR)
		case 0x04: // SLLV
			return sllv(&c.gpr, &instR)
		case 0x06: // SRLV
			return srlv(&c.gpr, &instR)
		case 0x07: // SRAV
			return srav(&c.gpr, &instR)
		case 0x08: // JR
			return jr(&c.pc, &c.gpr, &instR)
		case 0x09: // JALR
			return jalr(&c.pc, &c.gpr, &instR)
		case 0x0D:
			util.TODO("BREAK")
		case 0x0F:
			util.TODO("SYNC")
		case 0x10: /// MFHI
			return mfhi(c.hi, &instR)
		case 0x11: // MTHI
			return mthi(&c.gpr, &c.hi, &instR)
		case 0x12: // MFLO
			return mflo(c.lo, &instR)
		case 0x13: // MTLO
			return mtlo(&c.gpr, &c.lo, &instR)
		case 0x14: // DSLLV
			return dsllv(&c.gpr, &instR)
		case 0x16: // DSRLV
			return dsrlv(&c.gpr, &instR)
		case 0x17: // DSRAV
			return dsrav(&c.gpr, &instR)
		case 0x18: // MULT
			return mult(&c.gpr, &c.hi, &c.lo, &instR)
		case 0x19: // MULTU
			return multu(&c.gpr, &c.hi, &c.lo, &instR)
		case 0x1A: // DIV
			return div(&c.gpr, &c.hi, &c.lo, &instR)
		case 0x1B: // DIVU
			return divu(&c.gpr, &c.hi, &c.lo, &instR)
		case 0x1C: // DMULT
			return dmult(&c.gpr, &c.hi, &c.lo, &instR)
		case 0x1D: // DMULTU
			return dmultu(&c.gpr, &c.hi, &c.lo, &instR)
		case 0x1E: // DDIV
			return ddiv(&c.gpr, &c.hi, &c.lo, &instR)
		case 0x1F: // DDIVU
			return ddivu(&c.gpr, &c.hi, &c.lo, &instR)
		case 0x20: // ADD
			output := add(&c.gpr, &instR)
			if output != nil {
				c.trapIntegerOverflow()
			}
			return output
		case 0x21: // ADDU
			return addu(&c.gpr, &instR)
		case 0x22:
			output := sub(&c.gpr, &instR)
			if output != nil {
				c.trapIntegerOverflow()
			}
			return output
		case 0x23: // SUBU
			return subu(&c.gpr, &instR)
		case 0x24: // AND
			return and(&c.gpr, &instR)
		case 0x25: // OR
			return or(&c.gpr, &instR)
		case 0x26: // XOR
			return xor(&c.gpr, &instR)
		case 0x27: // NOR
			return nor(&c.gpr, &instR)
		case 0x2A: // SLT
			return slt(&c.gpr, &instR)
		case 0x2B: // SLTU
			return sltu(&c.gpr, &instR)
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
		switch instI.Rt {
		case 0x00:
			util.TODO("BLTZ")
		case 0x01:
			util.TODO("BGEZ")
		case 0x02:
			util.TODO("BLTZL")
		case 0x03:
			util.TODO("BGEZL")
		case 0x08:
			util.TODO("TGEI")
		case 0x09:
			util.TODO("TGEIU")
		case 0x0A:
			util.TODO("TLTI")
		case 0x0B:
			util.TODO("TLTIU")
		case 0x0C:
			util.TODO("TEQI")
		case 0x0E:
			util.TODO("TNEI")
		case 0x10:
			util.TODO("BLTZAL")
		case 0x11:
			util.TODO("BGEZAL")
		case 0x12:
			util.TODO("BLTZALL")
		case 0x13:
			util.TODO("BGEZALL")
		}
	case 0x02:
		util.TODO("J")
	case 0x03:
		util.TODO("JAL")
	case 0x04:
		util.TODO("BEQ")
	case 0x05:
		util.TODO("BNE")
	case 0x06:
		util.TODO("BLEZ")
	case 0x07:
		util.TODO("BGTZ")
	case 0x08:
		util.TODO("ADDI")
	case 0x09:
		util.TODO("ADDIU")
	case 0x0A:
		util.TODO("SLTI")
	case 0x0B:
		util.TODO("SLTIU")
	case 0x0C:
		util.TODO("ANDI")
	case 0x0D:
		util.TODO("ORI")
	case 0x0E:
		util.TODO("XORI")
	case 0x0F:
		util.TODO("LUI")
	case 0x10:
		util.TODO("COP0")
	case 0x11:
		util.TODO("COP1")
	case 0x12:
		util.TODO("COP2")
	case 0x14:
		util.TODO("BEQL")
	case 0x15:
		util.TODO("BNEL")
	case 0x16:
		util.TODO("BLEZL")
	case 0x17:
		util.TODO("BGTZL")
	case 0x18:
		util.TODO("DADDI")
	case 0x19:
		util.TODO("DADDIU")
	case 0x1A:
		util.TODO("LDL")
	case 0x1B:
		util.TODO("LDR")
	case 0x20: // LB
		return lb(&c.gpr, &instI)
	case 0x21: // LH
		return lh(&c.gpr, &instI)
	case 0x22:
		util.TODO("LWL")
	case 0x23: // LW
		return lw(&c.gpr, &instI)
	case 0x24:
		util.TODO("LBU")
	case 0x25:
		util.TODO("LHU")
	case 0x26:
		util.TODO("LWR")
	case 0x27:
		util.TODO("LWU")
	case 0x28:
		util.TODO("SB")
	case 0x29:
		util.TODO("SH")
	case 0x2A:
		util.TODO("SWL")
	case 0x2B:
		util.TODO("SW")
	case 0x2C:
		util.TODO("SDL")
	case 0x2D:
		util.TODO("SDR")
	case 0x2E:
		util.TODO("SWR")
	case 0x2F:
		util.TODO("CACHE")
	case 0x30:
		util.TODO("LL")
	case 0x31:
		util.TODO("LWC1")
	case 0x32:
		util.TODO("LWC2")
	case 0x34:
		util.TODO("LLD")
	case 0x35:
		util.TODO("LDC1")
	case 0x36:
		util.TODO("LDC2")
	case 0x37:
		util.TODO("LD")
	case 0x38:
		util.TODO("SC")
	case 0x39:
		util.TODO("SWC1")
	case 0x3A:
		util.TODO("SWC2")
	case 0x3C:
		util.TODO("SCD")
	case 0x3D:
		util.TODO("SDC1")
	case 0x3E:
		util.TODO("SDC2")
	case 0x3F:
		util.TODO("SD")
	}
	return nil
}
