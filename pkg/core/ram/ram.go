/*
	N64 memory map
	00000000-03EFFFFF   RDRAM Memory (RDRAM = Rambus DRAM)
	03F00000-03FFFFFF   RDRAM Registers
	04000000-040FFFFF   SP Registers
	04100000-041FFFFF   DP Command Registers
	04200000-042FFFFF   DP Span Registers
	04300000-043FFFFF   MIPS Interface (MI) Registers
	04400000-044FFFFF   Video Interface (VI) Registers
	04500000-045FFFFF   Audio Interface (AI) Registers
	04600000-046FFFFF   Peripheral Interface (PI) Registers
	04700000-047FFFFF   RDRAM Interface (RI) Registers
	04800000-048FFFFF   Serial Interface (SI) Registers
	04900000-04FFFFFF   Unused
	05000000-05FFFFFF   Cartridge Domain 2 Address 1
	06000000-07FFFFFF   Cartridge Domain 1 Address 1
	08000000-0FFFFFFF   Cartridge Domain 2 Address 2
	10000000-1FBFFFFF   Cartridge Domain 1 Address 2
	1FC00000-1FC007BF   PIF Boot ROM
	1FC007C0-1FC007FF   PIF RAM
	1FC00800-1FCFFFFF   Reserved
	1FD00000-7FFFFFFF   Cartridge Domain 1 Address 3
	80000000-FFFFFFFF   External SysAD Device

	Reference:
	- http://n64.icequake.net/mirror/www.jimb.de/Projects/N64TEK.htm#memorymapoverview
	- https://web.archive.org/web/20200429103221/http://en64.shoutwiki.com/wiki/Memory_map_detailed
	- https://ultra64.ca/files/tools/DETAILED_N64_MEMORY_MAP.txt
*/

package ram

import "n64emu/pkg/types"

// RAM N64 memory map
type RAM struct {
	// 0x00000000 to 0x003FFFFF R/W RDRAM range 0 (4MB)
	RDRAM0 [0x400000]types.Byte

	// 0x00400000 to 0x007FFFFF R/W RDRAM range 1 (4MB) Red Expansion PAK
	RDRAM1 [0x400000]types.Byte

	// 0x00800000 to 0x03EFFFFF  * Unused

	RDRAMReg     RDRAMReg
	SPReg        SPReg
	DPCommandReg DPCommandReg
	DPSpanReg    DPSpanReg
	MIReg        MIReg
	VIReg        VIReg
	AI           AIReg
	PIReg        PIReg
	RIReg        RIReg
	SIReg        SIReg
	CartDomain1  CartDomain1
	CartDomain2  CartDomain2

	// external SysAD device(0x80000000 to 0xFFFFFFFF) is effectively mirror of lower addresses
}

// RDRAMReg RDRAM Registers 0x03F00000 to 0x03FFFFFF
type RDRAMReg struct {
}

// SPReg Signal Processor Registers 0x04000000 to 0x0400FFFF
type SPReg struct {
}

// DPCommandReg Display Processor Command Registers 0x04100000 to 0x041FFFF
type DPCommandReg struct{}

// DPSpanReg Display Processor Span Registers 0x04200000 to 0x042FFFFF
type DPSpanReg struct{}

// MIReg MIPS Interface (MI) Registers 0x04300000 to 0x043FFFFF
type MIReg struct{}

// VIReg Video Interface (VI) Registers 0x04400000 to 0x044FFFFF
type VIReg struct{}

// AIReg Audio Interface (AI) Registers 0x04500000 to 0x045FFFFF
type AIReg struct{}

// PIReg Peripheral/Parallel Interface (PI) Registers 0x04600000 to 0x046FFFFF
type PIReg struct{}

// RIReg RDRAM Interface (RI) Registers 0x04700000 to 0x047FFFFF
type RIReg struct{}

// SIReg Serial Interface (SI) Registers 0x04800000 to 0x048FFFFF
type SIReg struct{}

// CartDomain2 Cartridge Domain 2
type CartDomain2 struct {
	// 0x05000000 to 0x05FFFFFF
	Address1 struct{}

	// 0x08000000 to 0x0FFFFFFF
	Address2 struct{}
}

// CartDomain1 Cartridge Domain 1
type CartDomain1 struct {
	// 0x06000000 to 0x07FFFFFF
	Address1 struct{}

	// 0x10000000 to 0x1FBFFFFF
	Address2 struct{}

	// 0x1FD00000 to 0x7FFFFFFF Unknown
	Address3 struct{}
}
