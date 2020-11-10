/*

MemoryMap:
	0x1FC00000-0x1FC007BF	PIF_ROM_START
	0x1FC007C0-0x1FC007FF	PIF_RAM


Command structure:
	[64byte block] at 0xbfc007c0 (1fc007c0)
	{
	00 00 00 00 : 00 00 00 00 - 8 bytes
	00 00 00 00 : 00 00 00 00 - 8 bytes
	00 00 00 00 : 00 00 00 00 - 8 bytes
	00 00 00 00 : 00 00 00 00 - 8 bytes
	00 00 00 00 : 00 00 00 00 - 8 bytes
	00 00 00 00 : 00 00 00 00 - 8 bytes
	00 00 00 00 : 00 00 00 00 - 8 bytes
	00 00 00 00 : 00 00 00 00 - 8 bytes
	}                       ^^pif status/control byte

	Commands are processed from any byte in pifram. ie: The pif chip steps thru each byte to load commands.

	Each command has a structure like so:

	byte 1 (t) = x number of bytes to send to pif chip
	byte 2 (r) = x number of bytes to recieve from pif chip
	byte 3 (c) = Command type

Command Types:
	| Command |       Description        |t |r |
	+---------+--------------------------+-----+
	|   00    |   request info (status)  |01|03|
	|   01    |   read button values     |01|04|
	|   02    |   read from mempack slot |03|21|
	|   03    |   write to mempack slot  |23|01|
	|   04    |   read eeprom            |02|08|
	|   05    |   write eeprom           |10|01|
	|   ff    |   reset                  |01|03|

Status of controller:
	%X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X
		A  B  Z  ST U  D  L  R  ?  ?  PL PR CU CD CL CR
	A,B,Z,ST    = A,B,Z, Start buttons
	U,D,L,R     = Joypad directions
	?,?         = Unknown
	PL,PR       = Pan left, Pan right buttons
	CU,CD,CL,CR = C buttons (up,down,left,right)

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
	PifRomSize       = 0x7c0
	PifRamSize       = 64
	PifNumOfChannels = 8
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
	RequestInfo         = CommandType(0x00) // send:  1 byte,  recv:  3 bytes
	ReadButtonValues    = CommandType(0x01) // send:  1 byte,  recv:  4 bytes
	ReadFromMempackSlot = CommandType(0x02) // send:  3 bytes, recv: 33 bytes
	WriteToMempackSlot  = CommandType(0x03) // send: 35 bytes, recv:  1 byte
	ReadEeprom          = CommandType(0x04) // send:  2 bytes, recv:  8 bytes
	WriteEeprom         = CommandType(0x05) // send: 16 bytes, recv:  1 byte
	Reset               = types.Byte(0xff)  // send:  1 byte,  recv:  3 bytes
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

// Parse the contents of the command
func parseCmd(data []types.Byte) (cmd CommandType, tx, rx types.Byte) {
	assert.Assert(len(data) > 4, "At least four bytes")
	return CommandType(data[2]), data[0], data[1]
}

func (pif *PIF) runCmd(offset types.Byte) types.Byte {
	// parse command
	data := pif.ram[offset:]
	// cmd, tx, rx := parseCmd(data)

	return 8 // TODO:
}

func (pif *PIF) update() {
	offset := types.Byte(0)

	for offset < PifRamSize; {
		switch ChannelType(pif.ram[i]) {
		case SkipChannel:
		case DummyData:
			offset++
			break
		case EndOfSetup:
			// abort scan
			offset = PifRamSize
			break
		case ChannelReset:
			util.TODO("unimplemented")
			i++
			break
		default:
			// do command
			arrangeBytes := pif.runCmd(i)
			i += arrangeBytes
			break
		}

	}
}
