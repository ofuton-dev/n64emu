package reg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFPR_WriteRead(t *testing.T) {
	testData := -0.1234567890

	for i := 0; i < NumOfRegsInFpr; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			fpr := FPR{}

			// write testData and read verify
			fpr.Write(i, testData)
			got := fpr.Read(i)

			assert.Equal(t, testData, got)
		})
	}
}
