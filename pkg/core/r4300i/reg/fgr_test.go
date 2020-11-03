package reg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFGR_WriteRead(t *testing.T) {
	testData := -0.1234567890

	for i := 0; i < NumOfRegsInFgr; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			fgr := FGR{}

			// write testData and read verify
			fgr.Write(i, testData)
			got := fgr.Read(i)

			assert.Equal(t, testData, got)
		})
	}
}
