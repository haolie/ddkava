package dbManager

import (
	"sync"
	"time"

	"fileSharker/src/model"
)

const (
	con_buff_size = 8
	con_timeout   = time.Minute * 5
)

var (
	cashMap = make(map[string]*downCash, con_buff_size)
	locker  = sync.Mutex{}
)

type downCash struct {
	downD    *model.DFileModel
	cashTime time.Time
}

func CashDownInfo(downUrl, vname string, size int64) {
	locker.Lock()
	defer locker.Unlock()

	_, exists := cashMap[downUrl]
	if exists {
		cashMap[downUrl].cashTime = time.Now()
	} else {
		d := model.NewDFileModel(0, "", "", size, downUrl, vname, 0, 0, time.Now(), time.Now())
		cashMap[downUrl] = &downCash{d, time.Now()}
	}

	for k, d := range cashMap {
		if time.Now().Sub(d.cashTime) > con_timeout {
			delete(cashMap, k)
		}
	}
}

func GetCash(downUrl string) (df *model.DFileModel, exists bool) {
	locker.Lock()
	defer locker.Unlock()

	d, exists := cashMap[downUrl]
	if exists {
		df = d.downD
	}

	return
}

func RemoveChase(downUrl string) {
	locker.Lock()
	defer locker.Unlock()

	delete(cashMap, downUrl)
}
