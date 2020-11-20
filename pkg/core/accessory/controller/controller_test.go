package controller

import (
	"n64emu/pkg/core/joybus"
	"n64emu/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReset(t *testing.T) {
	c := NewController()

	// initial state
	c.Input(types.HalfWord(A|B|Z|Start), 12, -34)
	assert.Equal(t, c.buttonStatus, types.HalfWord(0xf000))
	assert.Equal(t, c.xAxis, types.Byte(12))
	assert.Equal(t, c.yAxis, types.Byte(0xde)) // -34

	// reset
	txBuf := []types.Byte{types.Byte(joybus.Reset)}
	rxBuf := make([]types.Byte, 3)
	c.Run(joybus.Reset, txBuf, rxBuf)

	// verify read result
	assert.Equal(t, []types.Byte{0x05, 0x00, 0x00}, rxBuf)

	// verify reset state
	assert.Equal(t, c.buttonStatus, types.HalfWord(0x0))
	assert.Equal(t, c.xAxis, types.Byte(0))
	assert.Equal(t, c.yAxis, types.Byte(0))
}

func TestRequestInfo(t *testing.T) {
	c := NewController()

	// initial state
	var pak joybus.JoyBus = NewController() // dummy device
	c.AttachPak(&pak)

	// read info
	txBuf := []types.Byte{types.Byte(joybus.RequestInfo)}
	rxBuf := make([]types.Byte, 3)
	c.Run(joybus.RequestInfo, txBuf, rxBuf)

	// verify read result
	assert.Equal(t, []types.Byte{0x05, 0x00, 0x01}, rxBuf)
}

func TestReadButtonValues(t *testing.T) {
	c := NewController()
	txBuf := []types.Byte{types.Byte(joybus.ReadButtonValues)}
	rxBuf := make([]types.Byte, 4)

	// push all
	c.Input(types.HalfWord(A|B|Z|Start|DirectionalUp|DirectionalDown|DirectionalLeft|DirectionalRight|Rst|LeftTrigger|RightTrigger|CameraUp|CameraDown|CameraLeft|CameraRight), 12, -34)
	c.Run(joybus.ReadButtonValues, txBuf, rxBuf)
	assert.Equal(t, []types.Byte{0xff, 0xbf, 0x0c, 0xde}, rxBuf)

	// release all
	c.Input(types.HalfWord(0), 0, 0)
	c.Run(joybus.ReadButtonValues, txBuf, rxBuf)
	assert.Equal(t, []types.Byte{0x00, 0x00, 0x00, 0x00}, rxBuf)
}
