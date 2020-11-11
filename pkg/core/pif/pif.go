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
	"n64emu/pkg/types"
	"n64emu/pkg/util"
	"n64emu/pkg/util/assert"
)

const (
	PifRomSize = 0x7c0
	PifRamSize = 64
	// 0~3: joystick, 4: EEPROM, 5: EEPROM(Option)
	PifNumOfChannels = 6
)

// StatusByte: PIF RAM[0x3f]
type PifStatus types.Byte

const (
	Idle       = PifStatus(0x00)
	NewCommand = PifStatus(0x01)
	Busy       = PifStatus(0x80)
	// TODO: I'll need to find some more information.
	//  PifStatus(0x02)
	//  PifStatus(0x08)
	//  PifStatus(0x10)
	//  PifStatus(0x30)
	//  PifStatus(0xc0)
)

type ChannelType types.Byte

const (
	SkipChannel  = ChannelType(0x00)
	DummyData    = ChannelType(0xff)
	EndOfSetup   = ChannelType(0xfe)
	ChannelReset = ChannelType(0xfd)
)

// Command processing types
type CommandType types.Byte

const (
	RequestInfo         = CommandType(0x00)
	ReadButtonValues    = CommandType(0x01)
	ReadFromMempackSlot = CommandType(0x02)
	WriteToMempackSlot  = CommandType(0x03)
	ReadEeprom          = CommandType(0x04)
	WriteEeprom         = CommandType(0x05)
	RTCStatuQuery       = CommandType(0x06)
	ReadRTCBlock        = CommandType(0x07)
	WriteRTCBlock       = CommandType(0x08)
	Reset               = types.Byte(0xff)
)

// Command processing result
type CommandResult types.Byte

const (
	// no error, operation successful.
	Suscess = CommandResult(0x00)
	// error, device not present for specified command.
	DeviceNotPresent = CommandResult(0x80)
	// error, unable to send/recieve the number bytes for command type.
	UnableToTransferDatas = CommandResult(0x40)
)

type PIF struct {
	// PIF ROM
	rom [PifRomSize]types.Byte
	// PIF RAM
	ram [PifRamSize]types.Byte
}

// Emulate PIF Boot Rom
func (pif *PIF) EmulateBoot() {
	util.TODO("unimplemented.")
}

// Communicate to joystick
func (pif *PIF) communicateJoystick(channel types.Byte, txBuf, rxBuf []types.Byte) CommandResult {
	cmd := CommandType(txBuf[0])
	assert.Assert((cmd != ReadEeprom) && (cmd != WriteEeprom) && (cmd != RTCStatuQuery) && (cmd != ReadRTCBlock) && (cmd != WriteRTCBlock), "Do not send eeprom commands to joyStick")

	// TODO: Keep a Joystick instance and send commands
	return DeviceNotPresent
}

// Communicate to EEPROM
func (pif *PIF) communicateEEPROM(channel types.Byte, txBuf, rxBuf []types.Byte) CommandResult {
	cmd := CommandType(txBuf[0])
	assert.Assert((cmd != ReadButtonValues) && (cmd != ReadFromMempackSlot) && (cmd != WriteToMempackSlot), "Do not send joystick commands to eeprom")

	// TODO: Keep a EEPROM instance and send commands
	return DeviceNotPresent
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
	result := DeviceNotPresent
	switch channel {
	case 0: // joystick0
	case 1: // joystick1
	case 2: // joystick2
	case 3: // joystick3
		result = pif.communicateJoystick(channel, txBuf, rxBuf)
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

// Update PIF Status
func (pif *PIF) Update() {
	// check status bytes
	// statusByte := pif.ram[len(pif.ram)-1]
	// TODO: check statusByte

	// scan command in ram
	channel := types.Byte(0)
	offset := types.Byte(0)
	for (offset < PifRamSize) && (channel < PifNumOfChannels) {
		switch ChannelType(pif.ram[offset]) {
		case SkipChannel:
		case DummyData:
			offset++
			channel++
			break
		case EndOfSetup:
		case ChannelReset:
			// abort scan
			offset = PifRamSize
			break
		default:
			// do command
			arrangeBytes := pif.runCmd(channel, offset)
			offset += arrangeBytes
			break
		}
	}
}
