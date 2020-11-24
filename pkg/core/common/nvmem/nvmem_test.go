package nvmem

import (
	"n64emu/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllRangeWR(t *testing.T) {
	size := 256
	m := NVMem{}
	m.Init(size)

	// All range seq write
	initData := make([]types.Byte, size)
	for i := 0; i < len(initData); i++ {
		initData[i] = types.Byte(i)
	}
	m.Write(0, initData)

	// All range seq read
	readData := make([]types.Byte, size)
	m.Read(0, readData)
	assert.Equal(t, initData, readData)
}
