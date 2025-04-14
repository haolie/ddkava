package fileInfo

import (
	"time"
)

const tName = "file_info"

type FInfo struct {
	FKey    string    `gorm:"column:FKey;primary_key"`
	Path    string    `gorm:"column:Path"`
	CrtTime time.Time `gorm:"column:CrtTime"`
	FSize   int64     `gorm:"column:FSize"`
}

func NewFInfo(key, path string, size int64, t time.Time) *FInfo {
	return &FInfo{
		FKey:    key,
		Path:    path,
		CrtTime: t,
		FSize:   size,
	}
}

func (f *FInfo) GetTableName() string {
	return tName
}
