package ctrlpak

import (
	"n64emu/pkg/core/common/nvmem"
	"n64emu/pkg/core/joybus"
	"n64emu/pkg/types"
	"n64emu/pkg/util/assert"
	"os"
)

const (
	// 32KByte Battery-backed RAM
	RAMSize = 0x8000
)

// External storage memory, available for connection to the controller's expansion port.
type CtrlPAK struct {
	// Battery-backed RAM
	RAM nvmem.NVMem
}

// Read from binary file
// If you specify nil for binPath, it will be initialized with 0
func NewCtrlPak(binPath string) (*CtrlPAK, error) {
	dst := CtrlPAK{}
	dst.RAM.Init(RAMSize)

	// Allocate only
	_, err := os.Stat(binPath)
	if os.IsNotExist(err) {
		return &dst, nil
	}

	// Read from file
	err = dst.RAM.FromFile(binPath)
	return &dst, err
}

// Do nothing
// Data are not cleared because they are written on non-volatile memory or are backed up by batteries.
func (pak *CtrlPAK) Reset() {
	// do nothing
}

// Do Command
func (pak *CtrlPAK) Run(cmd joybus.CommandType, txBuf, rxBuf []types.Byte) joybus.CommandResult {
	// block offset not found
	if len(txBuf) < nvmem.TxHeaderOffset {
		assert.Assert(false, "block offset is not included in the sent data.")
		return joybus.UnableToTransferDatas
	}
	blockOffset := txBuf[1]

	switch cmd {
	case joybus.ReadFromMempackSlot:
		return pak.RAM.Read(blockOffset, rxBuf)
	case joybus.WriteToMempackSlot:
		return pak.RAM.Write(blockOffset, txBuf)

	default:
		assert.Assert(false, "Unsupported command")
		return joybus.UnableToTransferDatas
	}
}
