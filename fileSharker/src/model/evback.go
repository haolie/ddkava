package model

import (
	"fmt"
)

type EVBack struct {
	callbackMap map[string]func(interface{}) []error
}

func NewEVBack() *EVBack {
	return &EVBack{callbackMap: make(map[string]func(interface{}) []error, 4)}
}

func (ev *EVBack) AddCallBack(key string, callback func(interface{}) []error) error {

	if _, exists := ev.callbackMap[key]; exists {
		return fmt.Errorf("%s has register", key)
	}

	ev.callbackMap[key] = callback

	return nil
}

func (ev *EVBack) CallAll(param interface{}) (errList []error) {
	for _, cb := range ev.callbackMap {
		errList = append(errList, cb(param)...)
	}

	return
}
