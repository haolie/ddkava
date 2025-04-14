package sFile

import (
	"time"
)

type SFileModel struct {
	// 文件Id
	Id int64 `gorm:"column:Id;primary_key"`
	// 文件唯一标识
	FKey string `gorm:"column:FKey"`
	// 文件Id
	ParentID int64 `gorm:"column:ParentID"`
	// 文件Id
	FName string `gorm:"column:FName"`
	// downInfoId
	DownId int64 `gorm:"column:DownId"`
	// 文件Id
	DownUrl string `gorm:"column:DownUrl"`
	// VName
	VName string `gorm:"column:VName"`
	// 更新时间
	UpdateTime time.Time `gorm:"column:UpdateTime"`
	// 创建时间
	CrtTime time.Time `gorm:"column:CrtTime"`
}

func (df *SFileModel) GetTableName() string {
	return "save_info"
}

func NewSFileModel(id, parentId, downId int64, fkey, fname, downUrl, vname string, updateTime, CrtTime time.Time) *SFileModel {
	return &SFileModel{
		Id:         id,
		FKey:       fkey,
		ParentID:   parentId,
		FName:      fname,
		DownId:     downId,
		DownUrl:    downUrl,
		VName:      vname,
		UpdateTime: updateTime,
		CrtTime:    CrtTime,
	}
}
