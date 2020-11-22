package joybus

import "n64emu/pkg/types"

// Command processing types
type CommandType types.Byte

const (
	RequestInfo         = CommandType(0x00)
	ReadButtonValues    = CommandType(0x01)
	ReadFromMempackSlot = CommandType(0x02)
	WriteToMempackSlot  = CommandType(0x03)
	ReadEEPROM          = CommandType(0x04)
	WriteEEPROM         = CommandType(0x05)
	RTCStatuQuery       = CommandType(0x06)
	ReadRTCBlock        = CommandType(0x07)
	WriteRTCBlock       = CommandType(0x08)
	Reset               = CommandType(0xff)
)

// Device Identifier
type Identifier types.HalfWord

const (
	Controller           = Identifier(0x0500)
	VoiceRecognitionUnit = Identifier(0x0001)
	Mouse                = Identifier(0x0200)
	Keyboard             = Identifier(0x0002)
	TrainController      = Identifier(0x2004)
)

// Command processing result
type CommandResult types.Byte

const (
	// no error, operation successful.
	Success = CommandResult(0x00)
	// error, device not present for specified command.
	DeviceNotPresent = CommandResult(0x80)
	// error, unable to send/recieve the number bytes for command type.
	UnableToTransferDatas = CommandResult(0x40)
)

// JoyBus is interface of joybus accessor
// - PIF to EEPROM(in GamePak)
// - PIF to Controller
// - PIF to ControllerPak via Controller
type JoyBus interface {
	// Reset Device
	Reset()
	// Do Command
	Run(cmd CommandType, txBuf, rxBuf []types.Byte) CommandResult
}
