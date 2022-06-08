package filein

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"fileSharker/config"
	"fileSharker/src/db/dbManager"
	"fileSharker/src/fileScan"
)

type fileInfoEx struct {
	fileInfo fs.FileInfo
	dir      string
	key      string
}

func FileIn(dir string) error {
	parentPath := path.Dir(dir)

	list := make([]*fileInfoEx, 0, 16)
	err := fileScan.EachDir(dir, true, func(dir string, info fs.FileInfo) error {
		if !info.IsDir() {
			list = append(list, &fileInfoEx{info, dir, ""})
		}

		return nil
	})

	if err != nil {
		return err
	}

	fileMap := make(map[string]*fileInfoEx, 16)
	for _, item := range list {
		// 获取文件扫描几个
		scanObj, err := fileScan.ScanFile(path.Join(item.dir, item.fileInfo.Name()))
		if err != nil {
			return err
		}

		// 忽略重复文件
		if _, exists := fileMap[scanObj.FileKey]; exists {
			continue
		}

		// 已入库文件
		_, exists, err := dbManager.GetDbLayer().GetFileWithKey(scanObj.FileKey)
		if exists {
			continue
		}

		fileMap[scanObj.FileKey] = item

	}

	// 文件已拷贝但未入库处理

	// 拷贝文件

	// 录入信息

	// 删除原文件

	return nil
}

func copyFile(file fileInfoEx, parentPath string) error {

	relPath, err := filepath.Rel(parentPath, path.Join(file.dir, file.fileInfo.Name()))
	if err != nil {
		return err
	}

	fileSavePath := path.Join(config.GetConfig().SaveRoot, relPath)

	// 如已拷贝不在重复
	_, fileErr := os.Stat(fileSavePath)
	if fileErr != nil {
		return nil
	}

}
