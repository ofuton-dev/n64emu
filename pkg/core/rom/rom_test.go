package rom

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadRom(t *testing.T) {
	// big-endian, no byteswap
	header := []byte{
		0x80, 0x37, 0x12, 0x40, 0x00, 0x00, 0x00, 0x0F, 0x80, 0x00, 0x10, 0x00, 0x00, 0x00, 0x14, 0x44,
		0xB5, 0xAF, 0x61, 0xE1, 0xB7, 0xEF, 0x19, 0xE2, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x6f, 0x66, 0x75, 0x74, 0x6f, 0x6e, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x6e, 0x36, 0x34, 0x64, 0x65, // "ofuton-dev/n64dev   "
		0x76, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	// dummy data
	bootCode := bytes.Repeat([]byte{0x0}, BootCodeSize)
	// 0, 1, 2, 3, 4, 5, 6, 7, ...
	bigEndianSrc := append(header, bootCode...)

	// create other order array
	bigEndianByteSwapSrc := make([]byte, RomHeaderSize)
	littleEndianSrc := make([]byte, RomHeaderSize)
	littleEndianByteSwapSrc := make([]byte, RomHeaderSize)
	for i := 0; i < len(bigEndianSrc); i++ {
		// 1, 0, 3, 2, 5, 4, 7, 6, ...
		bigEndianByteSwapSrc[i] = bigEndianSrc[i^0x1]
		// 3, 2, 1, 0, 7, 6, 5, 4, ...
		littleEndianSrc[i] = bigEndianSrc[(i & ^0x3)|(3-(i%4))]
		// 2, 3, 0, 1, 6, 7, 4, 5, ...
		littleEndianByteSwapSrc[i] = bigEndianSrc[(i & ^0x3)|(3-(i%4))^0x1]
	}

	// check conversion process
	assert.Equal(t, uint32(RomHeaderBigEndian), (uint32(bigEndianSrc[0])<<24)|(uint32(bigEndianSrc[1])<<16)|(uint32(bigEndianSrc[2])<<8)|(uint32(bigEndianSrc[3])<<0))
	assert.Equal(t, uint32(RomHeaderBigEndianByteSwapped), (uint32(bigEndianByteSwapSrc[0])<<24)|(uint32(bigEndianByteSwapSrc[1])<<16)|(uint32(bigEndianByteSwapSrc[2])<<8)|(uint32(bigEndianByteSwapSrc[3])<<0))
	assert.Equal(t, uint32(RomHeaderLittleEndian), (uint32(littleEndianSrc[0])<<24)|(uint32(littleEndianSrc[1])<<16)|(uint32(littleEndianSrc[2])<<8)|(uint32(littleEndianSrc[3])<<0))
	assert.Equal(t, uint32(RomHeaderLittleEndianByteSwapped), (uint32(littleEndianByteSwapSrc[0])<<24)|(uint32(littleEndianByteSwapSrc[1])<<16)|(uint32(littleEndianByteSwapSrc[2])<<8)|(uint32(littleEndianByteSwapSrc[3])<<0))

	// setup tests
	want := "ofuton-dev/n64dev   "
	tests := []struct {
		name string
		src  []byte
	}{
		{name: "BigEndian", src: bigEndianSrc},
		{name: "BigEndianByteSwap", src: bigEndianByteSwapSrc},
		{name: "LittleEndian", src: littleEndianSrc},
		{name: "LittleEndianByteSwap", src: littleEndianByteSwapSrc},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dummyFile := fmt.Sprintf("%s.tmp", tt.name)

			// write dummy file
			ioutil.WriteFile(dummyFile, tt.src, 0644)
			defer os.Remove(dummyFile)

			// read rom
			rom, err := NewRom(dummyFile)
			if err != nil {
				t.Fatal(err)
			}
			// check read data
			assert.Equal(t, want, rom.ImageName)
		})
	}
}
