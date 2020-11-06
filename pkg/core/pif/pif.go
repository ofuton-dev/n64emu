/*

MemoryMap:
	0x1FC00000-0x1FC007BF	PIF_ROM_START
	0x1FC007C0-0x1FC007FF	PIF_RAM


0x1FC007C4 (16 low bits)   Status of controller:
	%X  X  X  X  X  X  X  X  X  X  X  X  X  X  X  X
		A  B  Z  ST U  D  L  R  ?  ?  PL PR CU CD CL CR
	A,B,Z,ST    = A,B,Z, Start buttons
	U,D,L,R     = Joypad directions
	?,?         = Unknown
	PL,PR       = Pan left, Pan right buttons
	CU,CD,CL,CR = C buttons (up,down,left,right)

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

Reference:
	- https://web.archive.org/web/20200429103221/http://en64.shoutwiki.com/wiki/Memory_map_detailed#PIF
	- https://web.archive.org/web/20180809011637/http://en64.shoutwiki.com/wiki/SI_Registers_Detailed#PIF_Usage
*/

package pif

import "n64emu/pkg/types"

const (
	PifRomSize = 0x7c0
	PifRamSize = 0x800
)

// Command processing types
type CommandTypes types.Byte

const (
	RequestInfo         = CommandTypes(0x00) // send:  1 byte,  recv:  3 bytes
	ReadButtonValues    = CommandTypes(0x01) // send:  1 byte,  recv:  4 bytes
	ReadFromMempackSlot = CommandTypes(0x02) // send:  3 bytes, recv: 33 bytes
	WriteToMempackSlot  = CommandTypes(0x03) // send: 35 bytes, recv:  1 byte
	ReadEeprom          = CommandTypes(0x04) // send:  2 bytes, recv:  8 bytes
	WriteEeprom         = CommandTypes(0x05) // send: 16 bytes, recv:  1 byte
	Reset               = types.Byte(0xff)   // send:  1 byte,  recv:  3 bytes
)

type Pif struct {
	// PIF ROM
	// No definition because of software emulation.

	// PIF RAM
	ram [PifRamSize]types.Byte
}
