package cart

// Game Cartridge
// EEPROM or Battery-Backed RAM depending on whether the cassette is equipped with
type Cart struct {
	ROM    *ROM
	EEPROM *EEPROM
	NVSRAM *NVSRAM
}

// Read from file and initialize
// eepromPath and nvsramPath are optional.
func NewCart(romPath, eepromPath, nvsramPath string) (*Cart, error) {
	c := Cart{}
	// Read ROM
	rom, err := NewRom(romPath)
	if err != nil {
		return &c, err
	}
	c.ROM = rom

	// Read EEPROM
	eeprom, err := NewEEPROM(romPath)
	if err != nil {
		return &c, err
	}
	c.EEPROM = eeprom

	// Read Battery-backed RAM or FlashRAM
	nvsram, err := NewNVSRAM(nvsramPath)
	if err != nil {
		return &c, err
	}
	c.NVSRAM = nvsram

	return &c, nil
}

// TODO: read/write function
// TODO: import/export function
