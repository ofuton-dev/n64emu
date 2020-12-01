package cpu

import (
	"n64emu/pkg/core/bus"
	"n64emu/pkg/core/mips/r4300i/reg"
	"n64emu/pkg/types"
)

// Pipeline is vr4300 pipeline module
type Pipeline struct {
	bus                        bus.Bus // Bus accessor
	instructionCacheFetchLatch types.DoubleWord
	registerFetchReady         bool
	registerFetchLatch         *types.Word
	executionLatch             *aluOutput
	dataCacheLatch             *dataCacheOutput
}

type dataCacheOutput struct {
	op     Op
	dest   types.Byte
	result types.DoubleWord
}

func newDataChacheOutput(op Op,
	dest types.Byte,
	result types.DoubleWord) *dataCacheOutput {
	return &dataCacheOutput{
		op:     op,
		dest:   dest,
		result: result,
	}
}

// NewPipeline is Pipeline constructor
func NewPipeline(bus bus.Bus) *Pipeline {
	return &Pipeline{
		bus: bus,
	}
}

// TODO: Refactor later.
func (p *Pipeline) step(endian types.Endianness, pc *types.DoubleWord, gpr *reg.GPR, execute func(types.Word) *aluOutput, fetch func(addr types.DoubleWord) types.Word) {
	// TODO: We need to consider about pipeline exception, branch delay, load delay and etc...
	p.writeBackStage(gpr)

	p.dataCacheStage(endian)

	p.executionStage(execute)

	p.registerFetchStage(fetch)

	p.instructionCacheFetchStage(pc)
}

// WB - Write Back
func (p *Pipeline) writeBackStage(gpr *reg.GPR) {
	if p.dataCacheLatch != nil {
		gpr.Write(p.dataCacheLatch.dest, p.dataCacheLatch.result)
	}
}

// DC - Data Cache Fetch
func (p *Pipeline) dataCacheStage(endian types.Endianness) {
	if p.executionLatch != nil {
		switch p.executionLatch.op {
		case LB:
			data := p.bus.ReadWord(endian, types.Word(p.executionLatch.result))
			result := types.DoubleWord(types.SByte(data))
			p.dataCacheLatch = newDataChacheOutput(p.executionLatch.op, p.executionLatch.dest, result)
		case LH:
			data := p.bus.ReadWord(endian, types.Word(p.executionLatch.result))
			result := types.DoubleWord(types.SHalfWord(data))
			p.dataCacheLatch = newDataChacheOutput(p.executionLatch.op, p.executionLatch.dest, result)
		case LW:
			data := p.bus.ReadWord(endian, types.Word(p.executionLatch.result))
			// In 64-bit mode, the loaded word is sign-extended to 64 bits.
			result := types.DoubleWord(types.SWord(data))
			p.dataCacheLatch = newDataChacheOutput(p.executionLatch.op, p.executionLatch.dest, result)
		default:
			p.dataCacheLatch = p.executionLatch.toDataChacheOutput()
		}
	}
}

// EX - Execution
func (p *Pipeline) executionStage(execute func(types.Word) *aluOutput) {
	if p.registerFetchLatch != nil {
		p.executionLatch = execute(*p.registerFetchLatch)
	}
}

// RF - Register Fetch
func (p *Pipeline) registerFetchStage(fetch func(addr types.DoubleWord) types.Word) {
	if p.registerFetchReady {
		opcode := fetch(p.instructionCacheFetchLatch)
		p.registerFetchLatch = &opcode
	}

}

// IC - Instruction Cache Fetch
func (p *Pipeline) instructionCacheFetchStage(pc *types.DoubleWord) {
	p.instructionCacheFetchLatch = *pc
	p.registerFetchReady = true
	*pc += 4
}
