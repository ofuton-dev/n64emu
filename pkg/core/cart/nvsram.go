package cart

import (
	"errors"
	"io/ioutil"
	"n64emu/pkg/types"
	"os"
)

const (
	// SRAM or FlashRAM: [0x0800_0000 - 0x0800_ffff]
	RAMSize = 0x1000
)

// Battery-Backed RAM or FlashRAM
type NVSRAM struct {
	Data [RAMSize]types.Byte
}

func NewNVSRAM(binPath string) (*NVSRAM, error) {
	n := NVSRAM{}

	// Allocate only
	info, err := os.Stat(binPath)
	if os.IsNotExist(err) {
		return &n, nil
	}

	// check file size
	if info.Size() > int64(len(n.Data)) {
		return &n, errors.New("The file size is gather than 4096 bytes")
	}
	// Read from file
	src, err := ioutil.ReadFile(binPath)
	if err != nil {
		return &n, err
	}
	// copy from file to ram
	for i := 0; i < len(src); i++ {
		n.Data[i] = src[i]
	}

	// done.
	return &n, nil
}

// TODO: import/export function
