package tool

import (
	"fmt"
	"testing"
)

func TestCreateId(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(CreateId())
	}
}
