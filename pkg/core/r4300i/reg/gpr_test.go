package reg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGPR_WriteRead(t *testing.T) {
	testData := uint64(0xaa995566_33ccbbdd)

	for i := 0; i < NumOfRegsInGpr; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			gpr := GPR{}

			// write testData and read verify
			gpr.Write(i, testData)
			got := gpr.Read(i)

			if i == 0 {
				// always zero
				assert.Equal(t, uint64(0x0), got)
			} else {
				assert.Equal(t, testData, got)
			}
		})
	}
}
