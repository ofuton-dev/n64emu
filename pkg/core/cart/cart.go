package cart

import (
	"n64emu/pkg/core/cart/eeprom"
	"n64emu/pkg/core/cart/nvsram"
	"n64emu/pkg/core/cart/rom"
	"n64emu/pkg/util"
)

// Game Cartridge
// EEPROM or Battery-Backed RAM depending on whether the cassette is equipped with
type Cart struct {
	EEPROM eeprom.EEPROM
	NVSRAM nvsram.NVSRAM
	ROM    rom.ROM
}

// Read from file and initialize
func NewCart() {
	util.TODO("Read from file and initialize")
}

// TODO: import/export function
