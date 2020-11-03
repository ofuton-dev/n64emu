/*
.N64 .Z64 ROM Reader
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
	"errors"
	"io/ioutil"
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
	Cart MediaFormat = 'N'
	// 64DD
	Dd MediaFormat = 'D'
	// cartridge part of expandable game
	CartEx MediaFormat = 'C'
	// 64DD expansion for cart
	DdEx MediaFormat = 'E'
	// Aleck64 Cartridge
	ACart MediaFormat = 'Z'
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
	European1    CountryCode = 0x58
	European2    CountryCode = 0x59
)

type Rom struct {
	// filepath
	RomPath string
	// 0x04, 4 bytes
	ClockRateOverride uint32
	// 0x08, 4 bytes
	ProgramCounter uint32
	// 0x0c, 4 bytes
	ReleaseAddress uint32
	// 0x10, 4 bytes
	Crc1 uint32
	// 0x14, 4 bytes
	Crc2 uint32
	// 0x20, 20 bytes
	ImageName string
	// 0x38, 4 bytes
	MediaFormat uint32
	// 0x3c, 2 bytes
	CartridgeId uint16
	// 0x3e, 1 byte
	CountryCode CountryCode
	// 0x3f, 1 byte
	Version byte
	// 0x40, 4032 bytes
	BootCode [BootCodeSize]byte
	// 0x1000 ~ File End
	Data []byte
}

// Swap the values of A and B.
func swap(a *byte, b *byte) {
	tmp := *a
	*a = *b
	*b = tmp
}

// Restore a byte-swapped array
func convertByteSwapped(src *[]byte) error {
	if len(*src)%2 != 0 {
		return errors.New("ROM size not divisible by two")
	}

	loopCount := len(*src) / 2
	for i := 0; i < loopCount; i++ {
		index := i * 2
		swap(*src[index], *src[index+1])
	}

	return nil
}

// Convert little-endian arrays to big-endian arrays
func convertLittle(src *[]byte) error {
	if len(*src)%4 != 0 {
		return errors.New("ROM size not divisible by four")
	}

	loopCount := len(*src) / 4
	for i := 0; i < loopCount; i++ {
		index := i * 4
		swap(*src[index], *src[index+3])
		swap(*src[index+1], *src[index+2])
	}

	return nil
}

// Repairing array order
func repairOrder(src *[]byte) error {
	if len(*src) < 4 {
		return errors.New("The size is less than 4 bytes")
	}

	header := (uint32(*src[0]) << 24) | (uint32(*src[1]) << 16) | (uint32(*src[2]) << 8) | (uint32(*src[3]) << 0)
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
func NewRom(romPath string) (Rom, error) {
	dst := Rom{
		RomPath: romPath,
	}
	// Check file
	info, err := os.Stat(romPath)
	// not found
	if os.IsNotExist(err) {
		return dst, err
	}
	// not file
	if info.IsDir() {
		return dst, errors.New("romPath is directory")
	}
	// No data for the ROM Header
	if info.Size() < RomHeaderSize {
		return dst, errors.New("The size is less than 4096 bytes")
	}

	// Read from file
	src, err := ioutil.ReadFile(romPath)
	if err != nil {
		return dst, err
	}

	// Detect identifier. repair rom endian and byte-swapped.
	if err := repairOrder(src); err != nil {
		return dst, err
	}

	// Parse cartridge rom header and data
	dst.ClockRateOverride = (uint32(src[0x04]) << 24) | (uint32(src[0x05]) << 16) | (uint32(src[0x06]) << 8) | (uint32(src[0x07]) << 0)
	dst.ProgramCounter = (uint32(src[0x08]) << 24) | (uint32(src[0x09]) << 16) | (uint32(src[0x0a]) << 8) | (uint32(src[0x0b]) << 0)
	dst.ReleaseAddress = (uint32(src[0x0c]) << 24) | (uint32(src[0x0d]) << 16) | (uint32(src[0x0e]) << 8) | (uint32(src[0x0f]) << 0)
	dst.Crc1 = (uint32(src[0x10]) << 24) | (uint32(src[0x11]) << 16) | (uint32(src[0x12]) << 8) | (uint32(src[0x13]) << 0)
	dst.Crc2 = (uint32(src[0x14]) << 24) | (uint32(src[0x15]) << 16) | (uint32(src[0x16]) << 8) | (uint32(src[0x17]) << 0)
	dst.ImageName = string(src[0x20 : 0x20+ImageNameSize]) // 0x20 ~ 0x34
	dst.MediaFormat = (uint32(src[0x38]) << 24) | (uint32(src[0x39]) << 16) | (uint32(src[0x3a]) << 8) | (uint32(src[0x3b]) << 0)
	dst.CartridgeId = (uint16(src[0x3c]) << 8) | (uint16(src[0x3d]) << 0)
	dst.CountryCode = uint8(src[0x3e])
	dst.Version = uint8(src[0x3f])
	dst.BootCode = src[0x40 : 0x40+BootCodeSize] // 0x40 ~ 0x1000
	dst.Data = src[RomHeaderSize:]               // 0x1000 ~ File End

	// CRC Check
	// If the check fails, the data may still be usable, so set up the data and then perform the CRC check.
	// TODO: impl here

	// done.
	return dst, nil
}
