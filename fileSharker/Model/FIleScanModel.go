package Model

/*
文件扫描模型
*/
type FileScanModel struct {
	// 文件名
	FileName string
	// 文件路径
	FilePath string
	// 文件大小
	FileSize int64
	// 文件Key
	FileKey string
}

/*
初始文件扫描模型
*/
func NewFileScan(fileName, filePath, fileKey string, fileSize int64) *FileScanModel {
	return &FileScanModel{
		FileName: fileName,
		FilePath: filePath,
		FileSize: fileSize,
		FileKey:  fileKey,
	}
}
