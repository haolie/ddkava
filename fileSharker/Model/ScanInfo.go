package Model

import (
	"time"
)

/*
扫描信息
*/
type ScanInfo struct {
	// 扫描根路径
	Root string
	// 当前状态
	Status int32
	// 信息
	Msg string
	// 已扫描文件数量
	FileCount int32
	// 扫描结果
	FileList []*FileScanModel
	// 开始时间
	StartTime time.Time
	// 当前路径
	CurDir string
}

type ScanStatusEnum int32

const (
	// 空闲中
	ScanStatusEnum_Free ScanStatusEnum = iota

	/*
		正在扫描
	*/
	ScanStatusEnum_Running

	/*
		已完成
	*/
	ScanStatusEnum_Completed

	/*
		扫描失败
	*/
	ScanStatusEnum_Failed
)
