package types

type Byte byte
type HalfWord uint16
type Word uint32
type DoubleWord uint64

type Endianness byte

const (
	Big Endianness = iota
	Little
)
