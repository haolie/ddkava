package Park

import (
	"context"
)

type PackCB func(ctx context.Context) []error

var (
	loadMap  = make(map[string]PackCB, 8)
	startMap = make(map[string]PackCB, 8)
)

func RegisterLoad(k string, cb PackCB) {
	_, exists := loadMap[k]
	if exists {
		panic("Load Already Exists")
	}

	loadMap[k] = cb
}

func RegisterStart(k string, cb PackCB) {
	_, exists := startMap[k]
	if exists {
		panic("Load Already Exists")
	}

	startMap[k] = cb
}

func Load(ctx context.Context) []error {
	errList := make([]error, 0, 4)

	for k, v := range loadMap {
		tempList := v(ctx)
		if len(tempList) > 0 {
			outputErrList("load", k, tempList)
			errList = append(errList, tempList...)
		}
	}

	return errList
}

func Start(ctx context.Context) []error {
	errList := make([]error, 0, 4)
	for k, v := range startMap {
		tempList := v(ctx)
		if len(tempList) > 0 {
			outputErrList("Start", k, tempList)
			errList = append(errList, tempList...)
		}
	}

	return errList
}

func outputErrList(name string, k string, errList []error) {

}
