package model

// 文件扫描模型
type FileScan struct {
	// 文件名
	FileName string
	// 文件路径
	FilePath string
	// 文件大小
	FileSize int64
	// 文件Key
	FileKey string
}

func NewFileScan(fileName, filePath, fileKey string, fileSize int64) *FileScan {
	return &FileScan{
		FileName: fileName,
		FilePath: filePath,
		FileSize: fileSize,
		FileKey:  fileKey,
	}
}
