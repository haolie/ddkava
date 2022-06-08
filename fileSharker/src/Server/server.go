package Server

import (
	"fmt"
	"sync"

	"fileSharker/config"
	"fileSharker/src/model"
)

var (
	loadEv      = model.NewEVBack()
	completedEv = model.NewEVBack()
)

func loadConfig() error {
	return config.Load("")
}

func Start(wg *sync.WaitGroup) bool {
	err := loadConfig()
	if err != nil {
		fmt.Println(err)

		return false
	}

	list := loadEv.CallAll(nil)
	if len(list) > 0 {
		fmt.Println(list)

		return false
	}

	list = completedEv.CallAll(nil)
	if len(list) > 0 {
		fmt.Println(list)

		return false
	}

	return true
}

func RegisterLoad(key string, cb func(interface{}) []error) {
	err := loadEv.AddCallBack(key, cb)
	if err != nil {
		panic(err)
	}
}

func RegisterCompleted(key string, cb func(interface{}) []error) {
	err := completedEv.AddCallBack(key, cb)
	if err != nil {
		panic(err)
	}
}
