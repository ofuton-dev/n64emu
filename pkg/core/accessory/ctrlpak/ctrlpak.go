package ctrlpak

import (
	"errors"
	"io/ioutil"
	"n64emu/pkg/core/joybus"
	"n64emu/pkg/types"
	"n64emu/pkg/util/assert"
	"os"
)

const (
	// 32KByte Battery-backed RAM
	RAMSize = 0x8000
	// 1 Block = 8 bytes
	BlockSize = 8
	// CmdId + Block Offset
	TxHeaderOffset = 2
)

// External storage memory, available for connection to the controller's expansion port.
type CtrlPAK struct {
	// Battery-backed RAM
	Data [RAMSize]types.Byte
}

// Read from binary file
// If you specify nil for binPath, it will be initialized with 0
func NewCtrlPak(binPath string) (*CtrlPAK, error) {
	dst := CtrlPAK{}

	// Check file
	info, err := os.Stat(binPath)
	// not found
	if os.IsNotExist(err) {
		return &dst, err
	}
	// No data for the ROM Header
	if info.Size() > RAMSize {
		return &dst, errors.New("The size is gather than 32K bytes")
	}
	// Read from file
	src, err := ioutil.ReadFile(binPath)
	if err != nil {
		return &dst, err
	}
	// copy from file to ram
	for i := 0; i < len(src); i++ {
		dst.Data[i] = src[i]
	}

	// done.
	return &dst, nil
}

// Do nothing
// Data are not cleared because they are written on non-volatile memory or are backed up by batteries.
func (pak *CtrlPAK) Reset() {
	// do nothing
}

func (pak *CtrlPAK) readData(blockOffset types.Byte, rxBuf []types.Byte) joybus.CommandResult {
	byteOffset := blockOffset / BlockSize
	for rxIndex := 0; rxIndex < len(rxBuf); rxIndex++ {
		address := int(byteOffset) + rxIndex

		// boundary check
		if address >= len(pak.Data) {
			assert.Assert(false, "read address is out of bounds")
			return joybus.UnableToTransferDatas
		}

		// data copy
		rxBuf[rxIndex] = pak.Data[address]
	}
	return joybus.Success
}

func (pak *CtrlPAK) writeData(blockOffset types.Byte, txBuf []types.Byte) joybus.CommandResult {
	byteOffset := blockOffset / BlockSize
	for txIndex := 0; txIndex+TxHeaderOffset < len(txBuf); txIndex++ {
		address := int(byteOffset) + txIndex

		// boundary check
		if address >= len(pak.Data) {
			assert.Assert(false, "write address is out of bounds")
			return joybus.UnableToTransferDatas
		}

		// data copy
		pak.Data[address] = txBuf[TxHeaderOffset+txIndex] // 2 byte = [cmd + block offset]
	}
	return joybus.Success
}

// Do Command
func (pak *CtrlPAK) Run(cmd joybus.CommandType, txBuf, rxBuf []types.Byte) joybus.CommandResult {
	// block offset not found
	if len(txBuf) < 2 {
		assert.Assert(false, "block offset is not included in the sent data.")
		return joybus.UnableToTransferDatas
	}
	blockOffset := txBuf[1]

	switch cmd {
	case joybus.ReadFromMempackSlot:
		return pak.readData(blockOffset, rxBuf)
	case joybus.WriteToMempackSlot:
		return pak.writeData(blockOffset, txBuf)

	default:
		assert.Assert(false, "Unsupported command")
		return joybus.UnableToTransferDatas
	}
}
