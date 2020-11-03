package reg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCP0_WriteRead(t *testing.T) {
	testData := uint32(0xaa995566)

	for i := 0; i < NumOfRegsInCp0; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			cp0 := CP0{}

			// write testData and read verify
			cp0.Write(i, testData)
			got := cp0.Read(i)

			assert.Equal(t, testData, got)
		})
	}
}
