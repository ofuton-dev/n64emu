/*
Decoding MIPS Instructions

I-Instruction format:
	| opcode | rs     | rt     | immediate                |
	| ------ | ------ | ------ | ------------------------ |
	| 6 bit  | 5 bits | 5 bits | 16 bits                  |

J-Instruction format:
	| opcode | address                                    |
	| ------ | ------------------------------------------ |
	| 6 bit  | 26 bits                                    |

R-Instruction format:
	| opcode | rs     | rt     | rd     | sa     | funct  |
	| ------ | ------ | ------ | ------ | ------ | ------ |
	| 6 bit  | 5 bits | 5 bits | 5 bits | 5 bits | 6 bits |

*/

package inst

import "n64emu/pkg/types"

// I-Type Instruction format
type InstI struct {
	// src[31:26] 6 bits
	Opcode types.Byte
	// src[25:21] 5 bits
	Rs types.Byte
	// src[20:16] 5 bits
	Rt types.Byte
	// src[15:0] 16 bits
	Immediate types.HalfWord
}

// J-Type Instruction format
type InstJ struct {
	// src[31:26] 6 bits
	Opcode types.Byte
	// src[25:0] 26 bits
	Address types.Word
}

// R-Type Instruction format
type InstR struct {
	// src[31:26] 6 bits
	Opcode types.Byte
	// src[25:21] 5 bits
	Rs types.Byte
	// src[20:16] 5 bits
	Rt types.Byte
	// src[15:11] 5 bits
	Rd types.Byte
	// src[10:6] 5 bits
	Sa types.Byte
	// src[5:0] 6 bits
	Funct types.Byte
}

// Decode I-Type Instruction
func DecodeI(src types.Word) InstI {
	return InstI{
		Opcode:    GetOp(src),
		Rs:        types.Byte((src >> 21) & 0x1f),
		Rt:        types.Byte((src >> 16) & 0x1f),
		Immediate: types.HalfWord((src >> 0) & 0xffff),
	}
}

// Decode J-Type Instruction
func DecodeJ(src types.Word) InstJ {
	return InstJ{
		Opcode:  GetOp(src),
		Address: (src >> 0) & 0x03ff_ffff,
	}
}

// Decode R-Type Instruction
func DecodeR(src types.Word) InstR {
	return InstR{
		Opcode: GetOp(src),
		Rs:     types.Byte((src >> 21) & 0x1f),
		Rt:     types.Byte((src >> 16) & 0x1f),
		Rd:     types.Byte((src >> 11) & 0x1f),
		Sa:     types.Byte((src >> 6) & 0x1f),
		Funct:  types.Byte((src >> 0) & 0x3f),
	}
}

func GetOp(src types.Word) types.Byte {
	return types.Byte((src >> 26) & 0x3f)
}
