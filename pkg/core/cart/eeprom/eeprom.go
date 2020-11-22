package ctrlpak

import (
	"n64emu/pkg/core/common/nvmem"
	"n64emu/pkg/core/joybus"
	"n64emu/pkg/types"
	"n64emu/pkg/util/assert"
	"os"
)

const (
	// 16K bytes
	EEPROMSize = 0x4000
	// CmdId + Block Offset
	TxHeaderOffset = 2
)

// External storage memory, available for connection to the controller's expansion port.
type EEPROM struct {
	// Battery-backed RAM
	ROM nvmem.NVMem
}

// Read from binary file
// If you specify nil for binPath, it will be initialized with 0
func NewEEPROM(binPath string) (*EEPROM, error) {
	dst := EEPROM{}
	dst.ROM.Init(EEPROMSize)

	// Allocate only
	_, err := os.Stat(binPath)
	if os.IsNotExist(err) {
		return &dst, nil
	}

	// Read from file
	err = dst.ROM.FromFile(binPath)
	return &dst, err
}

// Do nothing
// Data are not cleared because they are written on non-volatile memory.
func (e *EEPROM) Reset() {
	// do nothing
}

// Responding Device Identifier
// rxBuf = { CartEEPROM ID High(0x00), CartEEPROM ID Low(0x80), 0x00 }
func (e *EEPROM) readInfo(rxBuf []types.Byte) joybus.CommandResult {
	rxLen := len(rxBuf)

	// byte0
	if rxLen < 1 {
		return joybus.Success
	}
	rxBuf[0] = types.Byte((joybus.CartEEPROM >> 8) & 0xff)

	// byte1
	if rxLen < 2 {
		return joybus.Success
	}
	rxBuf[1] = types.Byte((joybus.CartEEPROM >> 0) & 0xff)

	// byte2
	if rxLen < 3 {
		return joybus.Success
	}
	rxBuf[2] = 0x0

	// No more data can respond
	if rxLen >= 4 {
		return joybus.UnableToTransferDatas
	}

	return joybus.Success
}

// Do Command
func (e *EEPROM) Run(cmd joybus.CommandType, txBuf, rxBuf []types.Byte) joybus.CommandResult {
	// block offset not found
	if len(txBuf) < TxHeaderOffset {
		assert.Assert(false, "block offset is not included in the sent data.")
		return joybus.UnableToTransferDatas
	}
	blockOffset := types.HalfWord(txBuf[1])

	switch cmd {
	case joybus.ReadEEPROM:
		return e.ROM.Read(blockOffset, rxBuf)
	case joybus.WriteEEPROM:
		return e.ROM.Write(blockOffset, txBuf[TxHeaderOffset:]) // skip header offset

	default:
		assert.Assert(false, "Unsupported command")
		return joybus.UnableToTransferDatas
	}
}
