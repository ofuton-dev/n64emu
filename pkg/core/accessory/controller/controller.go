package controller

import (
	"n64emu/pkg/core/joybus"
	"n64emu/pkg/types"
	"n64emu/pkg/util/assert"
)

type ButtonType types.HalfWord

const (
	CameraRight ButtonType = 1 << iota
	CameraLeft
	CameraDown
	CameraUp
	RightTrigger
	LeftTrigger
	Reserved // do not use
	Rst
	DirectionalRight
	DirectionalLeft
	DirectionalDown
	DirectionalUp
	Start
	Z
	B
	A
)

// N64 Standard Controller
type Controller struct {
	// byte0:[A,B,Z,S,dU,dD,dL,DR], byte1:[]
	buttonStatus types.HalfWord
	// byte2:[X-Axis]
	XAxis types.Byte
	// byte3:[Y-Axis]
	yAxis types.Byte
	// accessory port in the controller (e.g. RumblePak, ControllerPak,...)
	accessory *joybus.JoyBus
}

// Controller constructor
func NewController() *Controller {
	c := &Controller{
		buttonStatus: 0,
		XAxis:        0,
		yAxis:        0,
		accessory:    nil,
	}
	return c
}

// Attach accessory
func (c *Controller) AttachAccessory(accessory *joybus.JoyBus) {
	c.accessory = accessory
}

// Remove accessory
func (c *Controller) RemoveAccessory(accessory *joybus.JoyBus) {
	c.accessory = accessory
}

// Initialize the variables
func (c *Controller) Reset() {
	c.buttonStatus = 0x0
	c.XAxis = 0
	c.yAxis = 0

	if c.accessory != nil {
		(*c.accessory).Reset()
	}
}

// Update the user input
func (c *Controller) Input(buttonStatus types.HalfWord, xAxis, yAxis types.Byte) {
	c.buttonStatus = buttonStatus
	c.XAxis = xAxis
	c.yAxis = yAxis
}

// Responding Device Identifier
func (c *Controller) readInfo(rxBuf []types.Byte) joybus.CommandResult {
	rxLen := len(rxBuf)

	// byte0
	if rxLen < 2 {
		return joybus.Success
	}
	rxBuf[0] = types.Byte((joybus.Controller >> 16) & 0xff)

	// byte1
	if rxLen < 3 {
		return joybus.Success
	}
	rxBuf[1] = types.Byte((joybus.Controller >> 8) & 0xff)

	// byte2
	if rxLen < 4 {
		return joybus.Success
	}
	rxBuf[2] = types.Byte((joybus.Controller >> 0) & 0xff)

	// No more data can respond
	if rxLen >= 4 {
		return joybus.UnableToTransferDatas
	}

	return joybus.Success
}

// Responding Input Status
func (c *Controller) readInputStatus(rxBuf []types.Byte) joybus.CommandResult {
	rxLen := len(rxBuf)

	// byte0
	if rxLen < 2 {
		return joybus.Success
	}
	rxBuf[0] = types.Byte((c.buttonStatus >> 8) & 0xff)

	// byte1
	if rxLen < 3 {
		return joybus.Success
	}
	rxBuf[0] = types.Byte((c.buttonStatus >> 0) & 0xff)

	// byte2
	if rxLen < 4 {
		return joybus.Success
	}
	rxBuf[2] = c.XAxis

	// byte3
	if rxLen < 5 {
		return joybus.Success
	}
	rxBuf[3] = c.XAxis

	// No more data can respond
	if rxLen >= 5 {
		return joybus.UnableToTransferDatas
	}

	return joybus.Success
}

// Do Command
func (c *Controller) Run(cmd joybus.CommandType, txBuf, rxBuf []types.Byte) joybus.CommandResult {
	// Check tx len
	if len(txBuf) != 1 {
		assert.Assert(false, "only cmd should be sent.")
		return joybus.UnableToTransferDatas
	}

	switch cmd {
	case joybus.Reset: // Reset and Info
		assert.AssertEq(len(rxBuf), 3, "rx length is specified 3.")
		c.Reset()
		return c.readInfo(rxBuf)
	case joybus.RequestInfo: // Info
		assert.AssertEq(len(rxBuf), 3, "rx length is specified 3.")
		return c.readInfo(rxBuf)
	case joybus.ReadButtonValues:
		assert.AssertEq(len(rxBuf), 4, "rx length is specified 4.")
		return c.readInputStatus(rxBuf)

	case joybus.ReadFromMempackSlot: // Connection to the accessory port
		fallthrough
	case joybus.WriteToMempackSlot: // Connection to the accessory port
		if c.accessory == nil {
			return joybus.DeviceNotPresent
		}
		return (*c.accessory).Run(cmd, txBuf, rxBuf)

	default:
		assert.Assert(false, "Unsupported command")
		return joybus.UnableToTransferDatas
	}
}
