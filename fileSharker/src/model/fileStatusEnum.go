package model

type FileStatusEnum int32

const (
	FileStatusEnum_NoRecord FileStatusEnum = iota
	FileStatusEnum_Record
	FileStatusEnum_Down
)
