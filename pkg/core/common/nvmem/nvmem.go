package nvmem

import (
	"errors"
	"io/ioutil"
	"n64emu/pkg/core/joybus"
	"n64emu/pkg/types"
	"n64emu/pkg/util/assert"
	"os"
)

const (
	// 1 Block = 8 bytes
	BlockSize = 8
	// CmdId + Block Offset
	TxHeaderOffset = 2
)

// External non-volatile memory connected with joybus.
type NVMem struct {
	// Battery-backed RAM
	Data []types.Byte
}

// Allocate Data Buffer
func (m *NVMem) Init(size uint) {
	m.Data = make([]types.Byte, size)
}

// Read from binary file
func (m *NVMem) FromFile(binPath string) error {
	assert.AssertNe(0, len(m.Data), "Data is not initialized.")

	// Check file
	info, err := os.Stat(binPath)
	// not found
	if os.IsNotExist(err) {
		return err
	}
	// check file size
	if info.Size() > int64(len(m.Data)) {
		return errors.New("The file size is gather than allocated size")
	}
	// Read from file
	src, err := ioutil.ReadFile(binPath)
	if err != nil {
		return err
	}
	// copy from file to ram
	for i := 0; i < len(src); i++ {
		m.Data[i] = src[i]
	}

	// done.
	return nil
}

func (m *NVMem) Read(blockOffset types.Byte, rxBuf []types.Byte) joybus.CommandResult {
	assert.AssertNe(0, len(m.Data), "Data is not initialized.")

	byteOffset := blockOffset / BlockSize
	for rxIndex := 0; rxIndex < len(rxBuf); rxIndex++ {
		address := int(byteOffset) + rxIndex

		// boundary check
		if address >= len(m.Data) {
			assert.Assert(false, "read address is out of bounds")
			return joybus.UnableToTransferDatas
		}

		// data copy
		rxBuf[rxIndex] = m.Data[address]
	}
	return joybus.Success
}

func (m *NVMem) Write(blockOffset types.Byte, txBuf []types.Byte) joybus.CommandResult {
	assert.AssertNe(0, len(m.Data), "Data is not initialized.")

	byteOffset := blockOffset / BlockSize
	for txIndex := 0; txIndex+TxHeaderOffset < len(txBuf); txIndex++ {
		address := int(byteOffset) + txIndex

		// boundary check
		if address >= len(m.Data) {
			assert.Assert(false, "write address is out of bounds")
			return joybus.UnableToTransferDatas
		}

		// data copy
		m.Data[address] = txBuf[TxHeaderOffset+txIndex] // 2 byte = [cmd + block offset]
	}
	return joybus.Success
}
