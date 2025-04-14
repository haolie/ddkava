package sFile

import (
	"fmt"

	"fileSharker/src/db/dbManager"
)

func AddItem(item *SFileModel) error {
	return dbManager.GetDbLayer().Add(item)
}

// 更新记录
func UpdateItem(item *SFileModel) error {

	return dbManager.GetDbLayer().Update(item)
}

// 删除记录
func DeleteItem(item *SFileModel) error {
	return dbManager.GetDbLayer().DeleteOne(item)
}

// 返回指定id
func GetItem(id int64) (item *SFileModel, exists bool, err error) {
	item = new(SFileModel)
	filter := fmt.Sprintf("id = %d", id)
	exists, err = dbManager.GetDbLayer().GetOne(item, filter)

	return
}

func GetListWithFKey(key string) (list []*SFileModel, exists bool, err error) {
	return getListByOneKey("FKey", key, true)
}

func GetListWithDUrl(url string) (list []*SFileModel, exists bool, err error) {
	return getListByOneKey("DownUrl", url, true)
}

func GetListWithFName(name string) (list []*SFileModel, exists bool, err error) {
	return getListByOneKey("FName", name, true)
}

func GetListWithVName(name string) (list []*SFileModel, exists bool, err error) {
	return getListByOneKey("VName", name, true)
}

func GetListWithParentId(parentId int64) (list []*SFileModel, exists bool, err error) {
	return getListByOneKey("ParentID", parentId, false)
}

func GetListWithDownId(downId int64) (list []*SFileModel, exists bool, err error) {
	return getListByOneKey("DownId", downId, false)
}

// 列表查询
func getListByOneKey(keyName string, key interface{}, isString bool) (list []*SFileModel, exists bool, err error) {
	filter := fmt.Sprintf("%s = %v", keyName, key)
	if isString {
		filter = fmt.Sprintf("%s = \"%v\"", keyName, key)
	}

	list = make([]*SFileModel, 0, 2)

	exists, err = dbManager.GetDbLayer().GetList(new(SFileModel), &list, filter)
	return
}
