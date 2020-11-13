package cpu

import (
	"n64emu/pkg/core/mips/r4300i/reg"
	"n64emu/pkg/types"
	"n64emu/pkg/util"
)

// Pipeline is vr4300 pipeline module
type Pipeline struct {
	registerFetchReady bool
	registerFetchLatch *types.Word
	executionLatch     *aluOutput
	dataCacheLatch     *aluOutput
}

// NewPipeline is Pipeline constructor
func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) step(gpr *reg.GPR, execute func(opcode types.Word) *aluOutput, fetch func() types.Word) {
	// TODO: We need to consider about pipeline exception, branch delay, load delay and etc...
	p.writeBackStage(gpr)

	p.dataCacheStage()

	p.executionStage(execute)

	p.registerFetchStage(fetch)

	p.instructionCacheFetchStage()
}

// WB - Write Back
func (p *Pipeline) writeBackStage(gpr *reg.GPR) {
	if p.dataCacheLatch != nil {
		gpr.Write(p.dataCacheLatch.dest, p.dataCacheLatch.result)
	}
}

// DC - Data Cache Fetch
func (p *Pipeline) dataCacheStage() {
	p.dataCacheLatch = p.executionLatch
}

// EX - Execution
func (p *Pipeline) executionStage(execute func(opcode types.Word) *aluOutput) {
	if p.registerFetchLatch != nil {
		p.executionLatch = execute(*p.registerFetchLatch)
	}
}

// RF - Register Fetch
func (p *Pipeline) registerFetchStage(fetch func() types.Word) {
	if p.registerFetchReady {
		opcode := fetch()
		p.registerFetchLatch = &opcode
	}

}

// IC - Instruction Cache Fetch
func (p *Pipeline) instructionCacheFetchStage() {
	p.registerFetchReady = true
	// TODO: NOP for now
}

func (p *Pipeline) execute(gpr *reg.GPR, opcode types.Word) {
	op := GetOp(opcode)
	switch op {
	// R type instructions
	// SPECIAL
	case 0x00:
		instR := DecodeR(opcode)
		switch instR.Funct {
		case 0x00: // SLL
			p.executionLatch = sll(gpr, &instR)
		case 0x02: // SRL
			p.executionLatch = srl(gpr, &instR)
		case 0x03: // SRA
			p.executionLatch = sra(gpr, &instR)
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
		case 0x25: // OR
			p.executionLatch = or(gpr, &instR)
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
		instI := DecodeI(opcode)
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
	case 0x20:
		util.TODO("LB")
	case 0x21:
		util.TODO("LH")
	case 0x22:
		util.TODO("LWL")
	case 0x23:
		util.TODO("LW")
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
}
