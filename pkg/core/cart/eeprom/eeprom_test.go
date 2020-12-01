package eeprom

import (
	"n64emu/pkg/core/joybus"
	"n64emu/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	// test target
	rom, err := NewEEPROM("")
	if err != nil {
		t.Fatal(err)
	}

	// send command
	txBuf := []types.Byte{
		0xff, // Reset
	}
	rxBuf := make([]types.Byte, 3)
	rom.Run(joybus.Reset, txBuf, rxBuf)

	assert.Equal(t, []types.Byte{0x00, 0x80, 0x00}, rxBuf)
}

func TestReadInfo(t *testing.T) {
	// test target
	rom, err := NewEEPROM("")
	if err != nil {
		t.Fatal(err)
	}

	// send command
	txBuf := []types.Byte{
		0x0, // RequestInfo
	}
	rxBuf := make([]types.Byte, 3)
	rom.Run(joybus.RequestInfo, txBuf, rxBuf)

	assert.Equal(t, []types.Byte{0x00, 0x80, 0x00}, rxBuf)
}

func TestReadEEPROM(t *testing.T) {
	// test target
	rom, err := NewEEPROM("")
	if err != nil {
		t.Fatal(err)
	}
	// want data
	rom.ROM.Data[0x08] = 0xaa
	rom.ROM.Data[0x09] = 0x99
	rom.ROM.Data[0x0a] = 0x55
	rom.ROM.Data[0x0b] = 0x66
	rom.ROM.Data[0x0c] = 0x33
	rom.ROM.Data[0x0d] = 0xcc
	rom.ROM.Data[0x0e] = 0x77
	rom.ROM.Data[0x0f] = 0xee

	// send command
	txBuf := []types.Byte{
		0x4, 0x1, // ReadEEPROM, block=0x1
	}
	rxBuf := make([]types.Byte, 8)
	rom.Run(joybus.ReadEEPROM, txBuf, rxBuf)

	assert.Equal(t, []types.Byte{0xaa, 0x99, 0x55, 0x66, 0x33, 0xcc, 0x77, 0xee}, rxBuf)
}

func TestWriteEEPROM(t *testing.T) {
	// test target
	rom, err := NewEEPROM("")
	if err != nil {
		t.Fatal(err)
	}

	// send command
	txBuf := []types.Byte{
		0x5, 0x1, // WriteEEPROM, block=0x1
		0xaa, 0x99, 0x55, 0x66, 0x33, 0xcc, 0x77, 0xee, // WriteData
	}
	rxBuf := make([]types.Byte, 1)
	rom.Run(joybus.WriteEEPROM, txBuf, rxBuf)

	assert.Equal(t, []types.Byte{0xaa, 0x99, 0x55, 0x66, 0x33, 0xcc, 0x77, 0xee}, rom.ROM.Data[0x08:0x10])
	assert.Equal(t, []types.Byte{0x00}, rxBuf)
}
