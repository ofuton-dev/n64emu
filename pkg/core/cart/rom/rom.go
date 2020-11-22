/*
ROM Reader
http://en64.shoutwiki.com/wiki/ROM#Cartridge_ROM_Header

Cartridge ROM Header format:
	Offset Bytes  Explanation
	0x00   x	  indicator for endianness (nybble)*
	0x01   x	  initial PI_BSB_DOM1_LAT_REG (nybble)*
	0x01   x	  initial PI_BSD_DOM1_PGS_REG (nybble)*
	0x02   1	  initial PI_BSD_DOM1_PWD_REG*
	0x03   1	  initial PI_BSB_DOM1_PGS_REG*
	0x04   4	  ClockRate Override(0 uses default)*
	0x08   4	  Program Counter*
	0x0C   4	  Release Address
	0x10   4	  CRC1 (checksum)
	0x14   4	  CRC2
	0x18   8	  Unknown/Not used
	0x20   20	  Image Name/Internal name*
	0x34   4	  Unknown/Not used
	0x38   4	  Media format
	0x3C   2	  Cartridge ID (alphanumeric)
	0x3E   1	  Country Code
	0x3F   1	  Version
	0x40   4032	  Boot code/strap

*/

package rom

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
	"n64emu/pkg/types"
	"os"
)

const (
	// (Native) .N64, z64
	RomHeaderBigEndian = 0x80371240
	// .rom, .u64, .v64
	RomHeaderBigEndianByteSwapped = 0x37804012
	// .n64
	RomHeaderLittleEndian            = 0x40123780
	RomHeaderLittleEndianByteSwapped = 0x12408037
)

const (
	ImageNameSize = 20
	BootCodeSize  = 4032
	RomHeaderSize = 0x1000
)

type MediaFormatFirstByte byte

const (
	// Cartridge
	Cart MediaFormatFirstByte = 'N'
	// 64DD
	Dd MediaFormatFirstByte = 'D'
	// cartridge part of expandable game
	CartEx MediaFormatFirstByte = 'C'
	// 64DD expansion for cart
	DdEx MediaFormatFirstByte = 'E'
	// Aleck64 Cartridge
	ACart MediaFormatFirstByte = 'Z'
)

type CountryCode byte

const (
	Beta         CountryCode = 0x37
	Asian        CountryCode = 0x41
	Brazillian   CountryCode = 0x42
	Chinese      CountryCode = 0x43
	German       CountryCode = 0x44
	NorthAmerica CountryCode = 0x45
	French       CountryCode = 0x46
	GatewayNtsc  CountryCode = 0x47
	Dutch        CountryCode = 0x48
	Italian      CountryCode = 0x49
	Japanese     CountryCode = 0x4a
	Korean       CountryCode = 0x4b
	GatewayPal   CountryCode = 0x4c
	Canadian     CountryCode = 0x4e
	European     CountryCode = 0x50
	Spanish      CountryCode = 0x53
	Australian   CountryCode = 0x55
	Scandinavian CountryCode = 0x57
	EuropeanX    CountryCode = 0x58
	EuropeanY    CountryCode = 0x59
)

type ROM struct {
	// filepath
	RomPath string
	// 0x04, 4 bytes
	ClockRateOverride types.Word
	// 0x08, 4 bytes
	ProgramCounter types.Word
	// 0x0c, 4 bytes
	ReleaseAddress types.Word
	// 0x10, 4 bytes
	Crc1 types.Word
	// 0x14, 4 bytes
	Crc2 types.Word
	// 0x20, 20 bytes
	ImageName string
	// 0x38, 4 bytes
	MediaFormat types.Word
	// 0x3c, 2 bytes
	CartridgeID types.HalfWord
	// 0x3e, 1 byte
	CountryCode CountryCode
	// 0x3f, 1 byte
	Version types.Byte
	// 0x40, 4032 bytes(BootCodeSize)
	BootCode []types.Byte
	// 0x1000 ~ File End
	Data []types.Byte
}

// Swap the values of A and B.
func swap(a *types.Byte, b *types.Byte) {
	tmp := *a
	*a = *b
	*b = tmp
}

// Restore a byte-swapped array
func convertByteSwapped(src []types.Byte) error {
	if len(src)%2 != 0 {
		return errors.New("ROM size not divisible by two")
	}

	loopCount := len(src) / 2
	for i := 0; i < loopCount; i++ {
		index := i * 2
		swap(&src[index], &src[index+1])
	}

	return nil
}

// Convert little-endian arrays to big-endian arrays
func convertLittle(src []types.Byte) error {
	if len(src)%4 != 0 {
		return errors.New("ROM size not divisible by four")
	}

	loopCount := len(src) / 4
	for i := 0; i < loopCount; i++ {
		index := i * 4
		swap(&src[index], &src[index+3])
		swap(&src[index+1], &src[index+2])
	}

	return nil
}

// Repairing array order
func repairOrder(src []types.Byte) error {
	if len(src) < 4 {
		return errors.New("The size is less than 4 bytes")
	}

	header := (types.Word(src[0]) << 24) | (types.Word(src[1]) << 16) | (types.Word(src[2]) << 8) | (types.Word(src[3]) << 0)
	switch header {
	case RomHeaderBigEndian:
		break
	case RomHeaderBigEndianByteSwapped:
		if err := convertByteSwapped(src); err != nil {
			return err
		}
		break
	case RomHeaderLittleEndian:
		if err := convertLittle(src); err != nil {
			return err
		}
		break
	case RomHeaderLittleEndianByteSwapped:
		if err := convertLittle(src); err != nil {
			return err
		}
		if err := convertByteSwapped(src); err != nil {
			return err
		}
		break
	default:
		return errors.New("Invalid Header")
	}

	return nil
}

// Read from ROM file
func NewRom(romPath string) (*ROM, error) {
	dst := ROM{
		RomPath: romPath,
	}
	// Check file
	info, err := os.Stat(romPath)
	// not found
	if os.IsNotExist(err) {
		return &dst, err
	}
	// not file
	if info.IsDir() {
		return &dst, errors.New("romPath is directory")
	}
	// No data for the ROM Header
	if info.Size() < RomHeaderSize {
		return &dst, errors.New("The size is less than 4096 bytes")
	}

	// Read from file
	src, err := ioutil.ReadFile(romPath)
	if err != nil {
		return &dst, err
	}

	// Detect identifier. repair rom endian and byte-swapped.
	if err := repairOrder(src); err != nil {
		return &dst, err
	}

	// Parse cartridge rom header and data
	dst.ClockRateOverride = types.Word(binary.BigEndian.Uint32(src[0x4:0x8]))
	dst.ProgramCounter = types.Word(binary.BigEndian.Uint32(src[0x8:0xc]))
	dst.ReleaseAddress = types.Word(binary.BigEndian.Uint32(src[0xc:0x10]))
	dst.Crc1 = types.Word(binary.BigEndian.Uint32(src[0x10:0x14]))
	dst.Crc2 = types.Word(binary.BigEndian.Uint32(src[0x14:0x18]))
	dst.ImageName = string(src[0x20 : 0x20+ImageNameSize]) // 0x20 ~ 0x34
	dst.MediaFormat = types.Word(binary.BigEndian.Uint32(src[0x38:0x3c]))
	dst.CartridgeID = types.HalfWord(binary.BigEndian.Uint16(src[0x3c:0x3e]))
	dst.CountryCode = CountryCode(src[0x3e])
	dst.Version = types.Byte(src[0x3f])
	dst.BootCode = src[0x40 : 0x40+BootCodeSize] // 0x40 ~ 0x1000
	dst.Data = src[RomHeaderSize:]               // 0x1000 ~ File End

	// done.
	return &dst, nil
}
