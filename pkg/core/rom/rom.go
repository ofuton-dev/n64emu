/*
.N64 .Z64 ROM Reader
*/

package rom

import (
	"errors"
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

type Rom struct {
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
}
