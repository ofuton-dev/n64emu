package cpu

import (
	"encoding/binary"
	"n64emu/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockBus struct {
	MockMemory [0x10000]types.Byte
}

func (b *MockBus) WriteByte(e types.Endianness, addr types.Word, data types.Byte) {
	b.SetMemory(addr, []byte{data})
}

func (b *MockBus) WriteHalfWord(e types.Endianness, addr types.Word, data types.HalfWord) {
}

func (b *MockBus) WriteWord(e types.Endianness, addr types.Word, data types.Word) {
	// TODO:  For now, fixed by BIG endian
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, data)
	b.SetMemory(addr, bytes)
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
	return binary.BigEndian.Uint32(b.MockMemory[addr : addr+4])
}

func (b *MockBus) ReadDoubleWord(e types.Endianness, addr types.Word) types.DoubleWord {
	return 0
}

func (b *MockBus) SetMemory(offset types.Word, data []types.Byte) {
	for i, d := range data {
		b.MockMemory[offset+types.Word(i)] = d
	}
}

func beOpcodes2bytes(opecodes ...types.Word) []types.Byte {
	res := []types.Byte{}
	for _, o := range opecodes {
		bytes := make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, o)
		res = append(res, bytes...)
	}
	return res
}

func setupCPU(offset types.Word, data []types.Byte) (*CPU, *MockBus) {
	b := MockBus{}
	b.SetMemory(offset, data)
	return NewCPU(&b), &b
}

func TestSLL(t *testing.T) {
	assert := assert.New(t)
	// SLL rd=3, rt=2, sa=3
	cpu, _ := setupCPU(0, beOpcodes2bytes(0x000218C0))
	cpu.gpr.Write(2, 0x2)
	cpu.RunUntil(5)
	assert.Equal(types.DoubleWord(0x10), cpu.gpr.Read(3), "should shifted value stored")
	// SLL rd=3, rt=2, sa=0 with sign extended
	cpu, _ = setupCPU(0, beOpcodes2bytes(0x00021800))
	cpu.gpr.Write(2, 0x00000000FFFFFFFF)
	cpu.RunUntil(5)
	assert.Equal(types.DoubleWord(0xFFFFFFFFFFFFFFFF), cpu.gpr.Read(3), "should shifted value stored")
}

func TestSRL(t *testing.T) {
	assert := assert.New(t)
	// SLL rd=3, rt=2, sa=3
	cpu, _ := setupCPU(0, beOpcodes2bytes(0x000218C2))
	cpu.gpr.Write(2, 0x10)
	cpu.RunUntil(5)
	assert.Equal(types.DoubleWord(0x2), cpu.gpr.Read(3), "should shifted value stored")
}

func TestSRA(t *testing.T) {
	assert := assert.New(t)
	// SRA rd=3, rt=2, sa=3
	cpu, _ := setupCPU(0, beOpcodes2bytes(0x000218C3))
	cpu.gpr.Write(2, 0x00000000FFFFFFFF)
	cpu.RunUntil(5)
	assert.Equal(types.DoubleWord(0xFFFFFFFFFFFFFFFF), cpu.gpr.Read(3), "should shifted value stored")
}

func TestJR(t *testing.T) {
	assert := assert.New(t)
	// JR rs=4
	// SRA rd=3, rt=2, sa=3
	cpu, _ := setupCPU(0, beOpcodes2bytes(0x00800008, 0x000218C3))
	cpu.gpr.Write(2, 0x00000000FFFFFFFF)
	cpu.gpr.Write(4, 0x0000000000000100)
	cpu.RunUntil(6)
	// By delay slot SRA instruction should be executed.
	assert.Equal(types.DoubleWord(0xFFFFFFFFFFFFFFFF), cpu.gpr.Read(3), "should shifted value stored")
	assert.Equal(types.DoubleWord(0x110), cpu.pc, "should jumped value stored")
}

func TestMTHI(t *testing.T) {
	assert := assert.New(t)
	// MTHI rs=1
	cpu, _ := setupCPU(0, beOpcodes2bytes(0x00200011))
	cpu.gpr.Write(1, 0x5555AAAA5555AAAA)
	cpu.RunUntil(5)
	assert.Equal(types.DoubleWord(0x5555AAAA5555AAAA), cpu.hi, "should 0x5555AAAA5555AAAA in hi register")
}

func TestDSLLV(t *testing.T) {
	assert := assert.New(t)
	// DSLLV rd=3, rt=2, rs=1
	cpu, _ := setupCPU(0, beOpcodes2bytes(0x00221814))
	cpu.gpr.Write(2, 0x5555AAAA5555AAAA)
	cpu.gpr.Write(1, 0x1)
	cpu.RunUntil(5)
	assert.Equal(types.DoubleWord(0xAAAB5554AAAB5554), cpu.gpr.Read(3), "should shifted value stored")
}

func TestOR(t *testing.T) {
	assert := assert.New(t)
	// OR rd=3, rs=1, rt=2
	cpu, _ := setupCPU(0, beOpcodes2bytes(0x00221825))
	cpu.gpr.Write(1, 0x00000000AAAAAAAA)
	cpu.gpr.Write(2, 0x5555555500000000)
	cpu.RunUntil(5)
	assert.Equal(types.DoubleWord(0x55555555AAAAAAAA), cpu.gpr.Read(3), "should ORed value stored")
}

func TestLW(t *testing.T) {
	assert := assert.New(t)
	// LW base=1, rt=3, offset=0x0100
	cpu, bus := setupCPU(0, beOpcodes2bytes(0x8C230100))
	cpu.gpr.Write(1, 0x0000000000000004)
	bus.WriteWord(types.Big, 0x00000104, 0x5555AAAA)
	cpu.RunUntil(5)
	assert.Equal(types.DoubleWord(0x000000005555AAAA), cpu.gpr.Read(3), "should specified word data loaded")
}
