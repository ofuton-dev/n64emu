package types

type Byte = byte
type SByte = int8
type HalfWord = uint16
type SHalfWord = int16
type Word = uint32
type SWord = int32
type DoubleWord = uint64
type SDoubleWord = int64

type Endianness byte

const (
	Big Endianness = iota
	Little
)
