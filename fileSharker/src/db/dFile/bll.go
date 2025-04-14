package dFile

import (
	"fmt"

	"fileSharker/src/db/dbManager"
)

// 添加记录
func AddItem(item *DFileModel) error {
	return dbManager.GetDbLayer().Add(item)
}

// 更新记录
func UpdateItem(item *DFileModel) error {

	return dbManager.GetDbLayer().Update(item)
}

// 删除记录
func DeleteItem(item *DFileModel) error {
	return dbManager.GetDbLayer().DeleteOne(item)
}

// 返回指定id
func GetItem(id int64) (item *DFileModel, exists bool, err error) {
	item = new(DFileModel)
	filter := fmt.Sprintf("id = %d", id)
	exists, err = dbManager.GetDbLayer().GetOne(item, filter)

	return
}

// 根据FKey返回列表
func GetListWithKey(key string) (list []*DFileModel, exists bool, err error) {
	return getListByOneKey("FKey", key, true)
}

// 根据Url返回列表
func GetListWithUrl(url string) (list []*DFileModel, exists bool, err error) {
	return getListByOneKey("DownUrl", url, true)
}

// 根据CompressName返回列表
func GetListWithCompressName(CompressName string) (list []*DFileModel, exists bool, err error) {
	return getListByOneKey("CompressName", CompressName, true)
}

// 根据VName返回列表
func GetListWithVName(VName string) (list []*DFileModel, exists bool, err error) {
	return getListByOneKey("VName", VName, true)
}

func getListByOneKey(keyName string, key interface{}, isString bool) (list []*DFileModel, exists bool, err error) {
	filter := fmt.Sprintf("%s = %v", keyName, key)
	if isString {
		filter = fmt.Sprintf("%s = \"%v\"", keyName, key)
	}

	list = make([]*DFileModel, 0, 2)
	exists, err = dbManager.GetDbLayer().GetList(new(DFileModel), &list, filter)
	return
}
