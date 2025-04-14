package fileInfo

import (
	"fmt"

	"fileSharker/src/db/dbManager"
)

var (
	emptyObj = new(FInfo)
)

func GetItem(key string) (item *FInfo, exists bool, err error) {
	item = new(FInfo)
	exists, err = dbManager.GetDbLayer().GetOne(item, fmt.Sprintf("FKey = \"%s\"", key))

	return
}

func AddItem(item *FInfo) error {
	return dbManager.GetDbLayer().Add(item)
}
