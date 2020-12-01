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
		0xff,
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
		0x0,
	}
	rxBuf := make([]types.Byte, 3)
	rom.Run(joybus.RequestInfo, txBuf, rxBuf)

	assert.Equal(t, []types.Byte{0x00, 0x80, 0x00}, rxBuf)
}
