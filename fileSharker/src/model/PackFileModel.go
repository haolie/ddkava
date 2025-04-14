package model

type PackFileModel struct {
	// 文件Id
	Id int64 `gorm:"column:Id;primary_key"`
	// 下载压缩包Id
	ParkId int64 `gorm:"column:ParkId;primary_key"`
	// 文件Id
	FileId int64 `gorm:"column:FileId;primary_key"`
	// 相对压缩包路径
	ParkPath string `gorm:"column:ParkPath"`
	// 文件Id
	FName string `gorm:"column:FName"`
}
