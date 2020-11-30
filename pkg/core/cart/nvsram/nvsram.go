package nvsram

import (
	"n64emu/pkg/types"
)

const (
	// SRAM or FlashRAM: [0x0800_0000 - 0x0800_ffff]
	RamSize = 0x1000
)

// Battery-Backed RAM or FlashRAM
type NVSRAM struct {
	Data [RamSize]types.Byte
}
