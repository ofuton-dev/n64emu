package cpu

import (
	"n64emu/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockBus struct {
	MockMemory [0x10000]types.Byte
}

func (b *MockBus) WriteByte(e types.Endianness, addr types.Word, data types.Byte) {
}

func (b *MockBus) WriteHalfWord(e types.Endianness, addr types.Word, data types.HalfWord) {
}

func (b *MockBus) WriteWord(e types.Endianness, addr types.Word, data types.Word) {
}

func (b *MockBus) WriteDoubleWord(e types.Endianness, addr types.Word, data types.DoubleWord) {
}

func (b *MockBus) ReadByte(e types.Endianness, addr types.Word) types.Byte {
	return 0
}

func (b *MockBus) ReadHalfWord(e types.Endianness, addr types.Word) types.HalfWord {
	return 0
}

func (b *MockBus) ReadWord(e types.Endianness, addr types.Word) types.Word {
	// TODO:  For now, fixed by BIG endian
	return types.Word(b.MockMemory[addr]) | types.Word(b.MockMemory[addr+1])<<8 | types.Word(b.MockMemory[addr+2])<<16 | types.Word(b.MockMemory[addr+3])<<24
}

func (b *MockBus) ReadDoubleWord(e types.Endianness, addr types.Word) types.DoubleWord {
	return 0
}

func (b *MockBus) SetMemory(offset types.Word, data []types.Byte) {
	for i, d := range data {
		b.MockMemory[offset+types.Word(i)] = d
	}
}

func setupCPU(offset types.Word, data []types.Byte) (*CPU, *MockBus) {
	b := MockBus{}
	b.SetMemory(offset, data)
	return NewCPU(&b), &b
}

func TestSLL(t *testing.T) {
	assert := assert.New(t)
	// SLL rd=3, rt=2, sa=3
	cpu, _ := setupCPU(0, []types.Byte{0xC0, 0x18, 0x02, 0x00})
	cpu.gpr.Write(2, 0x2)
	cpu.Step()
	assert.Equal(types.DoubleWord(0x10), cpu.gpr.Read(3), "should shifted value stored")
	// SLL rd=3, rt=2, sa=0 with sign extended
	cpu, _ = setupCPU(0, []types.Byte{0x00, 0x18, 0x02, 0x00})
	cpu.gpr.Write(2, 0x00000000FFFFFFFF)
	cpu.Step()
	assert.Equal(types.DoubleWord(0xFFFFFFFFFFFFFFFF), cpu.gpr.Read(3), "should shifted value stored")
}

func TestOR(t *testing.T) {
	assert := assert.New(t)
	// OR rd=3, rs=1, rt=2
	cpu, _ := setupCPU(0, []types.Byte{0x25, 0x18, 0x22, 0x00})
	cpu.gpr.Write(1, 0x00000000AAAAAAAA)
	cpu.gpr.Write(2, 0x5555555500000000)
	cpu.Step()
	assert.Equal(types.DoubleWord(0x55555555AAAAAAAA), cpu.gpr.Read(3), "should ORed value stored")
}
