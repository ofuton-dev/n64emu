package reg

import (
	"fmt"
	"n64emu/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGPR_WriteRead(t *testing.T) {
	testData := types.DoubleWord(0xaa995566_33ccbbdd)

	for i := 0; i < NumOfRegsInGpr; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			gpr := GPR{}

			// write testData and read verify
			gpr.Write(types.Byte(i), testData)
			got := gpr.Read(types.Byte(i))

			if i == 0 {
				// always zero
				assert.Equal(t, types.DoubleWord(0x0), got)
			} else {
				assert.Equal(t, testData, got)
			}
		})
	}
}
