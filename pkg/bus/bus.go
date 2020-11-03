package bus

import "n64emu/pkg/types"

type Endianness byte

const (
	Big Endianness = iota
	Little
)

// Bus is interface of bus accessor
type Bus interface {
	WriteByte(e Endianness, addr types.Word, data types.Byte)
	WriteHalfWord(e Endianness, addr, data types.HalfWord)
	WriteWord(e Endianness, addr types.Word, data types.Word)
	WriteDoubleWord(e Endianness, addr types.Word, data types.DoubleWord)
	ReadByte(e Endianness, addr types.Word) types.Byte
	ReadHalfWord(e Endianness, addr types.Word) types.HalfWord
	ReadWord(e Endianness, addr types.Word) types.Word
	ReadDoubleWord(e Endianness, addr types.Word) types.DoubleWord
}
