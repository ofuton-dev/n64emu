package mempak

import (
	"n64emu/pkg/core/common/nvmem"
	"n64emu/pkg/core/joybus"
	"n64emu/pkg/types"
	"n64emu/pkg/util"
	"n64emu/pkg/util/assert"
	"os"
)

const (
	// 32KByte Battery-backed RAM
	RAMSize = 0x8000
	// CmdId:1byte + { block_offset[15:5] ,crc[4:0] }:2byte
	TxHeaderOffset = 3
)

// External storage memory, available for connection to the controller's expansion port.
type MEMPAK struct {
	// Battery-backed RAM
	RAM nvmem.NVMem
}

// Read from binary file
// If you specify nil for binPath, it will be initialized with 0
func NewMEMPak(binPath string) (*MEMPAK, error) {
	dst := MEMPAK{}
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
// Data are not cleared because are backed up by batteries.
func (pak *MEMPAK) Reset() {
	// do nothing
}

// Do Command
func (pak *MEMPAK) Run(cmd joybus.CommandType, txBuf, rxBuf []types.Byte) joybus.CommandResult {
	// block offset and crc not found
	if len(txBuf) < TxHeaderOffset {
		assert.Assert(false, "block offset and crc are not included in the sent data.")
		return joybus.UnableToTransferDatas
	}
	// parse args: { block_offset[15:5] ,crc[4:0] }
	blockOffset := types.HalfWord((txBuf[1] << 3) | (txBuf[2] >> 5))
	// crc := (txBuf[1] & 0x1f)

	// Examine the CRC included in the command
	util.TODO("examine the crc")

	switch cmd {
	case joybus.ReadFromMempackSlot:
		// At the end of the readout data, the CRC is set to 1 byte
		ret := pak.RAM.Read(blockOffset, rxBuf[:len(rxBuf)-1]) // without crc byte
		util.TODO("check crc")
		return ret
	case joybus.WriteToMempackSlot:
		ret := pak.RAM.Write(blockOffset, txBuf[TxHeaderOffset:]) // without header
		// If 1 is specified in rx, calculate and write the CRC to the Read result
		util.TODO("check crc")
		return ret
	default:
		// via the controller's accesory port, so Reset and Info are not accepted
		assert.Assert(false, "Unsupported command")
		return joybus.UnableToTransferDatas
	}
}
