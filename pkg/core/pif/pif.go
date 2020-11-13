/*

MemoryMap:
	0x1FC00000-0x1FC007BF	PIF_ROM_START
	0x1FC007C0-0x1FC007FF	PIF_RAM

Reference:
	- https://web.archive.org/web/20200429103221/http://en64.shoutwiki.com/wiki/Memory_map_detailed#PIF
	- https://web.archive.org/web/20180809011637/http://en64.shoutwiki.com/wiki/SI_Registers_Detailed#PIF_Usage
*/

package pif

import (
	"n64emu/pkg/core/joybus"
	"n64emu/pkg/types"
	"n64emu/pkg/util"
	"n64emu/pkg/util/assert"
)

const (
	PIFROMSize         = 0x7c0
	PIFRAMSize         = 64
	PIFNumOfController = 4
	PIFNumOfEEPROM     = 2
	PIFNumOfChannels   = PIFNumOfController + PIFNumOfEEPROM
)

// StatusByte: PIF RAM[0x3f]
type PIFStatus types.Byte

const (
	Idle       = PIFStatus(0x00)
	NewCommand = PIFStatus(0x01)
	Busy       = PIFStatus(0x80)
	// TODO: I'll need to find some more information.
	//  PifStatus(0x02)
	//  PifStatus(0x08)
	//  PifStatus(0x10)
	//  PifStatus(0x30)
	//  PifStatus(0xc0)
)

// Definition of the case where TX byte has a special meaning
type ChannelType types.Byte

const (
	SkipChannel  = ChannelType(0x00)
	DummyData    = ChannelType(0xff)
	EndOfSetup   = ChannelType(0xfe)
	ChannelReset = ChannelType(0xfd)
)

// Peripheral Interface
type PIF struct {
	// PIF ROM
	rom [PIFROMSize]types.Byte
	// PIF RAM
	ram [PIFRAMSize]types.Byte
	// Controllers
	controllers [PIFNumOfController]*joybus.JoyBus
	// EEPROM
	eeproms [PIFNumOfEEPROM]*joybus.JoyBus
}

// Emulate PIF Boot ROM
func (pif *PIF) EmulateBoot() {
	util.TODO("unimplemented")
}

// Communicate to controller
func (pif *PIF) communicateController(channel types.Byte, txBuf, rxBuf []types.Byte) joybus.CommandResult {
	cmd := joybus.CommandType(txBuf[0])
	assert.Assert((cmd != joybus.ReadEeprom) && (cmd != joybus.WriteEeprom) && (cmd != joybus.RTCStatuQuery) && (cmd != joybus.ReadRTCBlock) && (cmd != joybus.WriteRTCBlock), "Do not send eeprom commands to joyStick")

	// channel=0 ,1 ,2 ,3
	assert.Assert(channel < PIFNumOfController, "channel is out of bounds")

	// No connection
	if pif.controllers[channel] == nil {
		return joybus.DeviceNotPresent
	}
	// Do command
	return (*pif.controllers[channel]).Run(cmd, txBuf, rxBuf)
}

// Communicate to EEPROM
func (pif *PIF) communicateEEPROM(channel types.Byte, txBuf, rxBuf []types.Byte) joybus.CommandResult {
	cmd := joybus.CommandType(txBuf[0])
	assert.Assert((cmd != joybus.ReadButtonValues) && (cmd != joybus.ReadFromMempackSlot) && (cmd != joybus.WriteToMempackSlot), "Do not send controller commands to eeprom")

	// channel=4 or 5
	assert.Assert((PIFNumOfController < channel) && (channel < PIFNumOfController+PIFNumOfEEPROM), "channel is out of bounds.")
	eepromIndex := channel - PIFNumOfController

	// No connection
	if pif.eeproms[eepromIndex] == nil {
		return joybus.DeviceNotPresent
	}

	// Do command
	return (*pif.eeproms[eepromIndex]).Run(cmd, txBuf, rxBuf)
}

// Run Command on PIF RAM
// Command Block Format:
//   tx, rx, [cmd, ...], [a, b, ...]
//           ^^^^^^^^^^  ^^^^^^^^^^^
//           |tx buf     |rx buf
//
// `rx` byte Format:
//   { result[7:6], num_of_read_bytes[5:0] }
func (pif *PIF) runCmd(channel, offset types.Byte) types.Byte {
	// extract buf
	data := pif.ram[offset:]
	assert.Assert(len(data) > 2, "format error")
	txLen, rxLen := data[0], (data[1] & 0x3f)
	assert.AssertNe(txLen, 0, "send tx num of bytes is zero")

	txBuf := data[2 : 2+(txLen-1)]         // [cmd, ...]
	rxBuf := data[2+txLen : 2+txLen+rxLen] // [a, b, ...]

	// communicate
	result := joybus.DeviceNotPresent
	switch channel {
	case 0: // controller0
	case 1: // controller1
	case 2: // controller2
	case 3: // controller3
		result = pif.communicateController(channel, txBuf, rxBuf)
		break
	case 4: // eeprom
	case 5: // eeproom(option)
		result = pif.communicateEEPROM(channel, txBuf, rxBuf)
		break
	default: // NC
		assert.Assert(false, "unexpected channel")
		break
	}

	// Write the result to rx byte
	data[1] |= types.Byte(result)

	// number of bytes read and write from RAM
	return 2 + txLen + rxLen
}

// Reset PIF HW
func (pif *PIF) Reset() {
	// clear RAM
	pif.ram = [PIFRAMSize]types.Byte{}

	// this status byte will be set to 0x80 after holding down the reset button.
	// set to 0x00 after you let go of the reset button or .5 seconds passes(whichever comes first)
	pif.ram[PIFRAMSize-1] = types.Byte(Busy)
	pif.ram[PIFRAMSize-1] = types.Byte(Idle) // TODO: It might be better to transfer the call to EmulateBoot or NMI interrupts
}

// Register JoyBus device
func (pif *PIF) Register(channel types.Byte, device *joybus.JoyBus) {
	if channel < PIFNumOfController {
		pif.controllers[channel] = device
	} else if channel < PIFNumOfController+PIFNumOfEEPROM {
		pif.eeproms[channel-PIFNumOfController] = device
	} else {
		assert.Assert(false, "channel is out of bounds")
	}
}

// Unregister JoyBus device
func (pif *PIF) Unregister(channel types.Byte) {
	pif.Register(channel, nil)
}

// Update PIF Status
func (pif *PIF) Update() {
	// check status byte
	switch PIFStatus(pif.ram[PIFRAMSize-1]) {
	case NewCommand:
		break
	case Busy:
	case Idle:
	default:
		// TODO: Investigate what else to do with the status
		return
	}

	// scan command in ram
	channel := types.Byte(0)
	offset := types.Byte(0)
	for (offset < PIFRAMSize) && (channel < PIFNumOfChannels) {
		switch ChannelType(pif.ram[offset]) {
		case SkipChannel:
		case DummyData:
			offset++
			channel++
			break
		case EndOfSetup:
		case ChannelReset:
			// abort scan
			offset = PIFRAMSize
			break
		default:
			// do command
			arrangeBytes := pif.runCmd(channel, offset)
			offset += arrangeBytes
			break
		}
	}

	// update status byte
	pif.ram[PIFRAMSize-1] = types.Byte(Idle)
}
