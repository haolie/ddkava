package model

import (
	"time"
)

type SFileModel struct {
	// 文件Id
	Id int64 `gorm:"column:Id;primary_key"`
	// 文件唯一标识
	FKey string `gorm:"column:FKey"`
	// 文件Id
	FName string `gorm:"column:FName"`
	// 文件Id
	FSize int64 `gorm:"column:FSize"`
	// 文件Id
	DownUrl string `gorm:"column:DownUrl"`
	// VName
	VName string `gorm:"column:VName"`
	// 更新时间
	UpdateTime time.Time `gorm:"column:UpdateTime"`
	// 创建时间
	CrtTime time.Time `gorm:"column:CrtTime"`
}
