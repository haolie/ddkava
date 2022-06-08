package dbModel

import (
	"fileSharker/src/model"
)

type IDbLayer interface {
	// GetDbType
	// @description: 返回数据存储类型
	// parameter:
	// return:
	//		@DbTypeEnum: 存储类型
	GetDbType() DbTypeEnum

	// AddFile
	// @description: 添加文件信息
	// parameter:
	//		@fm: 文件数据模型
	// return:
	//		@error: 错误
	AddFile(tbName string, fm interface{}) error

	// UpdateFile
	// @description: 更新文件信息
	// parameter:
	//		@fm: 文件信息
	// return:
	//		@error: 错误
	UpdateFile(tbName string, fm interface{}) error

	// Remove
	// @description: 删除文件信息
	// parameter:
	//		@fm: 文件对象
	// return:
	//		@error: 错误
	Remove(tbName string, fm interface{}) error

	// GetDFModelWithKey
	// @description: 获取下载文件模型
	// parameter:
	//		@keyName: 字段名称
	//		@key: 字段值
	// return:
	//		@dFModel: 下载文件对象
	//		@exists: 是否存在
	//		@err: 错误
	GetDFModelWithKey(keyName, key string, isString bool) (dFModel *model.DFileModel, exists bool, err error)

	// GetDFListWithKey
	// @description: 返回列表
	// parameter:
	//		@keyName: 字段名称
	//		@key: 字段值
	// return:
	//		@list: 列表
	//		@exists: 是否存在
	//		@err: 错误
	GetDFListWithKey(keyName, key string) (list []*model.DFileModel, exists bool, err error)

	// GetDFListWithSameKey
	// @description: 返回列表
	// parameter:
	//		@keyName: 字段名称
	//		@key: 字段值
	// return:
	//		@list: 列表
	//		@exists: 是否存在
	//		@err: 错误
	GetDFListWithSameKey(keyName, key string) (list []*model.DFileModel, exists bool, err error)

	// GetSFModelWithKey
	// @description: 获取保存文件数据
	// parameter:
	//		@keyName: 字段名称
	//		@key: 字段值
	// return:
	//		@sFModel: 保存文件信息
	//		@exists: 是否存在
	//		@err: 错误
	GetSFModelWithKey(keyName, key string, isString bool) (sFModel *model.SFileModel, exists bool, err error)

	// GetSFListWithKey
	// @description: 获取保存文件信息列表
	// parameter:
	//		@keyName: 字段名称
	//		@key: 字段值
	// return:
	//		@list: 列表
	//		@exists: 是否存在
	//		@err: 错误
	GetSFListWithKey(keyName, key string) (list []*model.SFileModel, exists bool, err error)

	// GetSFListWithSameKey
	// @description: 获取保存文件信息列表
	// parameter:
	//		@keyName: 字段名称
	//		@key: 字段值
	// return:
	//		@list: 列表
	//		@exists: 是否存在
	//		@err: 错误
	GetSFListWithSameKey(keyName, key string) (list []*model.SFileModel, exists bool, err error)
}
