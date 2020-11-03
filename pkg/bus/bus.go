package bus

import "n64emu/pkg/types"

// Bus is interface of bus accessor
type Bus interface {
	WriteByte(e types.Endianness, addr types.Word, data types.Byte)
	WriteHalfWord(e types.Endianness, addr types.Word, data types.HalfWord)
	WriteWord(e types.Endianness, addr types.Word, data types.Word)
	WriteDoubleWord(e types.Endianness, addr types.Word, data types.DoubleWord)
	ReadByte(e types.Endianness, addr types.Word) types.Byte
	ReadHalfWord(e types.Endianness, addr types.Word) types.HalfWord
	ReadWord(e types.Endianness, addr types.Word) types.Word
	ReadDoubleWord(e types.Endianness, addr types.Word) types.DoubleWord
}
