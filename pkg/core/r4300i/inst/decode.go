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

// I-Type Instruction format
type InstI struct {
	// src[31:26] 6 bits
	Opcode byte
	// src[25:21] 5 bits
	Rs byte
	// src[20:16] 5 bits
	Rt byte
	// src[15:0] 16 bits
	Immediate uint16
}

// J-Type Instruction format
type InstJ struct {
	// src[31:26] 6 bits
	Opcode byte
	// src[25:0] 26 bits
	Address uint32
}

// R-Type Instruction format
type InstR struct {
	// src[31:26] 6 bits
	Opcode byte
	// src[25:21] 5 bits
	Rs byte
	// src[20:16] 5 bits
	Rt byte
	// src[15:11] 5 bits
	Rd byte
	// src[10:6] 5 bits
	Sa byte
	// src[5:0] 6 bits
	Funct byte
}

// Decode I-Type Instruction
func DecodeI(src uint32) InstI {
	return InstI{
		Opcode:    byte((src >> 26) & 0x3f),
		Rs:        byte((src >> 21) & 0x1f),
		Rt:        byte((src >> 16) & 0x1f),
		Immediate: uint16((src >> 0) & 0xffff),
	}
}

// Decode J-Type Instruction
func DecodeJ(src uint32) InstJ {
	return InstJ{
		Opcode:  byte((src >> 26) & 0x3f),
		Address: (src >> 0) & 0x03ff_ffff,
	}
}

// Decode R-Type Instruction
func DecodeR(src uint32) InstR {
	return InstR{
		Opcode: byte((src >> 26) & 0x3f),
		Rs:     byte((src >> 21) & 0x1f),
		Rt:     byte((src >> 16) & 0x1f),
		Sa:     byte((src >> 6) & 0x1f),
		Funct:  byte((src >> 0) & 0x3f),
	}
}
