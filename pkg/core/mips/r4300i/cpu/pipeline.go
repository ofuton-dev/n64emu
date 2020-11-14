package cpu

import (
	"n64emu/pkg/core/bus"
	"n64emu/pkg/types"
)

// Pipeline is vr4300 pipeline module
type Pipeline struct {
	bus                bus.Bus // Bus accessor
	registerFetchReady bool
	registerFetchLatch *types.Word
	executionLatch     *aluOutput
	dataCacheLatch     *aluOutput
}

// NewPipeline is Pipeline constructor
func NewPipeline(bus bus.Bus) *Pipeline {
	return &Pipeline{
		bus: bus,
	}
}

func (p *Pipeline) step(execute func(types.Word) *aluOutput, fetch func() types.Word, writeBack func(*aluOutput)) {
	// TODO: We need to consider about pipeline exception, branch delay, load delay and etc...
	p.writeBackStage(writeBack)

	p.dataCacheStage()

	p.executionStage(execute)

	p.registerFetchStage(fetch)

	p.instructionCacheFetchStage()
}

// WB - Write Back
func (p *Pipeline) writeBackStage(writeBack func(*aluOutput)) {
	if p.dataCacheLatch != nil {
		writeBack(p.dataCacheLatch)
	}
}

// DC - Data Cache Fetch
func (p *Pipeline) dataCacheStage() {
	p.dataCacheLatch = p.executionLatch
}

// EX - Execution
func (p *Pipeline) executionStage(execute func(types.Word) *aluOutput) {
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