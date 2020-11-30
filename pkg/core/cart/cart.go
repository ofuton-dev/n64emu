package cart

import (
	"n64emu/pkg/core/cart/eeprom"
	"n64emu/pkg/core/cart/nvsram"
	"n64emu/pkg/core/cart/rom"
)

// Game Cartridge
type Cart struct {
	EEPROM eeprom.EEPROM
	NVSRAM nvsram.NVSRAM
	ROM    rom.ROM
}
