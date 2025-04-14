package filein

import (
	"testing"
)

func TestFileIn(t *testing.T) {
	err := FileIn("E:/Temp/FSTest")
	if err != nil {
		t.Error(err)
	}
}
