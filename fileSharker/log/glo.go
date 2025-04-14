package log

import (
	"fmt"
)

func LogError(err error) {
	fmt.Println(err)
}

func LogInfo(info string) {
	fmt.Println(info)
}

func DebugInfo(info string) {

}

func LogErrs(errList []error) {
	for _, err := range errList {
		LogError(err)
	}
}
