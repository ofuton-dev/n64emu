package nvmem

import (
	"io/ioutil"
	"n64emu/pkg/types"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllRangeWR(t *testing.T) {
	size := 256
	m := NVMem{}
	m.Init(size)

	// Write incremental pattern
	initData := make([]types.Byte, size)
	for i := 0; i < len(initData); i++ {
		initData[i] = types.Byte(i)
	}
	m.Write(0, initData)

	// Read all range
	readData := make([]types.Byte, size)
	m.Read(0, readData)
	assert.Equal(t, initData, readData)
}

func TestScatterRead(t *testing.T) {
	size := 256
	readBytes := 8
	m := NVMem{}
	m.Init(size)

	// Write incremental pattern
	initData := make([]types.Byte, size)
	for i := 0; i < len(initData); i++ {
		initData[i] = types.Byte(i)
	}
	m.Write(0, initData)

	// Split and read all range
	for i := 0; i < len(initData)/BlockSize; i++ {
		blockOffset := types.HalfWord(i)
		readData := make([]types.Byte, readBytes)
		m.Read(blockOffset, readData)
		assert.Equal(t, initData[i*BlockSize:i*BlockSize+readBytes], readData)
	}
}

func TestScatterWrite(t *testing.T) {
	size := 256
	writeBytes := 8
	m := NVMem{}
	m.Init(size)

	// Split and write incremental pattern
	initData := make([]types.Byte, size)
	for i := 0; i < len(initData)/BlockSize; i++ {
		blockOffset := types.HalfWord(i)
		writeData := make([]types.Byte, writeBytes)
		for j := 0; j < writeBytes; j++ {
			writeData[j] = types.Byte(i*BlockSize + j)
			initData[i*BlockSize+j] = writeData[j]
		}
		m.Write(blockOffset, writeData)
	}

	// Read all range
	readData := make([]types.Byte, size)
	m.Read(0, readData)
	assert.Equal(t, initData, readData)
}

func TestFromFile(t *testing.T) {
	dummyFile := "testfromfile.tmp"
	size := 256

	// Create test file
	initData := make([]types.Byte, size)
	for i := 0; i < len(initData); i++ {
		initData[i] = types.Byte(i)
	}
	ioutil.WriteFile(dummyFile, initData, 0644)
	defer os.Remove(dummyFile)

	// Init and read from file
	m := NVMem{}
	m.Init(size)
	m.FromFile(dummyFile)

	// Read all range
	readData := make([]types.Byte, size)
	m.Read(0, readData)
	assert.Equal(t, initData, readData)
}
