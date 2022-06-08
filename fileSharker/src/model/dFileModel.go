package model

import (
	"time"
)

type DFileModel struct {
	// 文件Id
	Id int64 `gorm:"column:Id;primary_key"`
	// 文件唯一标识
	FKey string `gorm:"column:FKey"`
	// 文件Id
	CompressName string `gorm:"column:CompressName"`
	// 文件Id
	CompressSize int64 `gorm:"column:CompressSize"`
	// 文件Id
	DownUrl string `gorm:"column:DownUrl"`
	// VName
	VName string `gorm:"column:VName"`
	// 文件Id
	SameLink int64 `gorm:"column:SameLink"`
	// 狀態
	Status int32 `gorm:"column:Status"`
	// 更新时间
	UpdateTime time.Time `gorm:"column:UpdateTime"`
	// 创建时间
	CrtTime time.Time `gorm:"column:CrtTime"`
}

func NewDFileModel(id int64, fkey, compressName string, compressSize int64, downUrl, vname string, link int64, status int32, updateTime, crtTime time.Time) *DFileModel {
	return &DFileModel{
		Id:           id,
		FKey:         fkey,
		CompressName: compressName,
		CompressSize: compressSize,
		DownUrl:      downUrl,
		VName:        vname,
		SameLink:     link,
		Status:       status,
		UpdateTime:   updateTime,
		CrtTime:      crtTime,
	}
}
