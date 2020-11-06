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
	// 0x03F00000 to 0x03F00003 R/W RDRAM_CONFIG_REG or RDRAM_DEVICE_TYPE_REG
	Config types.Word

	// 0x03F00004 to 0x03F00007 R/W RDRAM_DEVICE_ID_REG
	DeviceID types.Word

	// 0x03F00008 to 0x03F0000B R/W RDRAM_DELAY_REG
	Delay types.Word

	// 0x03F0000C to 0x03F0000F R/W RDRAM_MODE_REG
	Mode types.Word

	// 0x03F00010 to 0x03F00013 R/W RDRAM_REF_INTERVAL_REG
	RefInterval types.Word

	// 0x03F00014 to 0x03F00017 R/W RDRAM_REF_ROW_REG
	RefRow types.Word

	// 0x03F00018 to 0x03F0001B R/W RDRAM_RAS_INTERVAL_REG
	RasInterval types.Word

	// 0x03F0001C to 0x03F0001F R/W RDRAM_MIN_INTERVAL_REG
	MinInterval types.Word

	// 0x03F00020 to 0x03F00023 R/W RDRAM_ADDR_SELECT_REG
	AddrSelect types.Word

	// 0x03F00024 to 0x03F00027 R/W RDRAM_DEVICE_MANUF_REG
	DeviceManuf types.Word

	// 0x03F00028 to 0x03FFFFFF *   Unknown
}

// SPReg Signal Processor Registers 0x04000000 to 0x0400FFFF
type SPReg struct {
	DMem [0x1000]types.Byte
	IMem [0x1000]types.Byte

	// Master, SP memory address 0x04040000 to 0x04040003
	MemAddr types.Word

	// Slave, SP DRAM DMA address 0x04040004 to 0x04040007
	DramAddr types.Word

	// SP read DMA length 0x04040008 to 0x0404000B
	RdLen types.Word

	// SP write DMA length 0x0404000C to 0x0404000F
	WrLen types.Word

	// SP status 0x04040010 to 0x04040013
	Status types.Word

	// SP DMA full 0x04040014 to 0x04040017
	DMAFull types.Word

	// SP DMA busy 0x04040018 to 0x0404001B
	DMABusy types.Word

	// SP semaphore 0x0404001C to 0x0404001F
	Semaphore types.Word

	// SP PC 0x04080000 to 0x04080003
	PC types.Word

	// SP IMEM BIST  0x04080004 to 0x04080007
	IBist types.Word
}

// DPCommandReg Display Processor Command Registers 0x04100000 to 0x041FFFF
type DPCommandReg struct {
	Start    types.Word
	End      types.Word
	Current  types.Word
	Status   types.Word
	Clock    types.Word
	BufBusy  types.Word
	PipeBusy types.Word
	TMem     types.Word
}

// DPSpanReg Display Processor Span Registers 0x04200000 to 0x042FFFFF
type DPSpanReg struct {
	TBist       types.Word
	TestMode    types.Word
	BufTestAddr types.Word
	BufTestData types.Word
}

// MIReg MIPS Interface (MI) Registers 0x04300000 to 0x043FFFFF
type MIReg struct {
	InitMode types.Word
	Version  types.Word
	Intr     types.Word
	IntrMask types.Word
}

// VIReg Video Interface (VI) Registers 0x04400000 to 0x044FFFFF
type VIReg struct {
	Status  types.Word
	Origin  types.Word
	Width   types.Word
	Intr    types.Word
	Current types.Word
	Burst   types.Word
	VSync   types.Word
	HSync   types.Word
	Leap    types.Word
	HStart  types.Word
	VStart  types.Word
	VBurst  types.Word
	XScale  types.Word
	YSCale  types.Word
}

// AIReg Audio Interface (AI) Registers 0x04500000 to 0x045FFFFF
type AIReg struct {
	DramAddr types.Word
	Len      types.Word
	Status   types.Word
	DACRate  types.Word
	BitRate  types.Word
}

// PIReg Peripheral/Parallel Interface (PI) Registers 0x04600000 to 0x046FFFFF
type PIReg struct {
	DramAddr types.Word
	CartAddr types.Word
	RdLen    types.Word
	WrLen    types.Word
	Status   types.Word
	Dom1Lat  types.Word
	Dom1Pwd  types.Word
	Dom1Pgs  types.Word
	Dom1Rls  types.Word
	Dom2Lat  types.Word
	Dom2Pwd  types.Word
	Dom2Pgs  types.Word
	Dom2Rls  types.Word
}

// RIReg RDRAM Interface (RI) Registers 0x04700000 to 0x047FFFFF
type RIReg struct {
	Mode    types.Word
	Config  types.Word
	Current types.Word
	Refresh types.Word
	Latency types.Word
	RdError types.Word
	WrError types.Word
}

// SIReg Serial Interface (SI) Registers 0x04800000 to 0x048FFFFF
type SIReg struct {
	DramAddr     types.Word
	PIFAddrRd64B types.Word
	PIFAddrWr64B types.Word
	Status       types.Word
}

// CartDomain1 Cartridge Domain 1
type CartDomain1 struct {
	// 0x06000000 to 0x07FFFFFF This address seems to be where the n64ddrive would be addressed
	Address1 [0x2000000]types.Byte

	// 0x10000000 to 0x1FBFFFFF
	Address2 struct {
		// 0x10000000 to 0x1000003F ROM header
		// NOTE: not define ROM header because ROM header info is in rom.ROM struct

		// 0x10000040 to 0x10000B6F
		RAMROMBootstrapOffset [2864]types.Byte

		// 0x10000B70 to 0x10000FEF
		RAMROMFontDataOffset [1152]types.Byte

		// 0x10001000 to 0x10FF9FFF
		RAMROMGameOffset [16748544]types.Byte

		// 0x10FFA000 to 0x10FFAFFF
		RAMROMAppReadAddr [4096]types.Byte

		// 0x10FFB000 to 0x10FFBFFF
		RAMROMAppWriteAddr [4096]types.Byte

		// 0x10FFC000 to 0x10FFCFFF
		RAMROMRmonReadAddr [4096]types.Byte

		// 0x10FFD000 to 0x10FFDFFF
		RAMROMRmonWriteAddr [4096]types.Byte

		// 0x10FFE000 to 0x10FFEFFF
		RAMROMPrintfAddr [4096]types.Byte

		// 0x10FFF000 to 0x10FFFFFF
		RAMROMLogAddr [4096]types.Byte
	}

	// Address3 0x1FD00000 to 0x7FFFFFFF Unknown
}

// CartDomain2 Cartridge Domain 2
type CartDomain2 struct {
	// 0x05000000 to 0x05FFFFFF
	Address1 [0x1000000]types.Byte

	// 0x08000000 to 0x0FFFFFFF SRAM could be here
	Address2 [0x8000000]types.Byte
}
