package filein

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"fileSharker/src/db/fileInfo"
	"fileSharker/src/db/sFile"
	"fileSharker/src/fileScan"
	"fileSharker/src/tool"
)

type fileInfoEx struct {
	fileInfo fs.FileInfo
	dir      string
	key      string
}

func init() {
	//Server.RegisterStartFun("filein", func(ctx context.Context) []error {
	//	list:=make([]error,2)
	//	err:= FileIn("E:/Temp/FSTest")
	//	if err!=nil{
	//		list=append(list,err)
	//	}
	//
	//	return list
	//})
}

func FileIn(dir string) error {
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

	for _, item := range list {
		oneIn(filepath.Join(item.dir, item.fileInfo.Name()), "", "", 0)
	}

	return nil
}

func oneIn(oldFilePath, downUrl, vName string, downId int64) (modelObj *sFile.SFileModel, err error) {
	// 获取文件扫描几个
	scanObj, err := fileScan.ScanFile(oldFilePath)
	if err != nil {
		return
	}

	// 判断文件是否已经保存
	_, existsFile, err := fileInfo.GetItem(scanObj.FileKey)
	if err != nil {
		return
	}

	// 文件不存在 先入库文件
	if !existsFile {
		savePath := CreateSavePath(scanObj.FileKey)
		// 文件已拷贝但未入库处理
		_, fileErr := os.Stat(savePath)
		if fileErr != nil {
			// 拷贝文件
			_, err = copy(oldFilePath, savePath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		err = fileInfo.AddItem(fileInfo.NewFInfo(scanObj.FileKey, savePath, scanObj.FileSize, time.Now()))
		if err != nil {
			return
		}
	}

	// 删除原文件
	os.Remove(oldFilePath)

	// 查找相同key的文件
	fList, exists, err := sFile.GetListWithFKey(scanObj.FileKey)
	if err != nil {
		return
	}

	// 如果key和vname都相同  就无需再入库
	if exists {
		for _, f := range fList {
			if f.VName == vName {
				return f, nil
			}
		}
	}

	modelObj = sFile.NewSFileModel(tool.CreateId(), 0, downId, scanObj.FileKey, scanObj.FileName, downUrl, vName, time.Now(), time.Now())
	err = sFile.AddItem(modelObj)

	return
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)

	return nBytes, err
}

func CreateSavePath(fKey string) string {
	return fmt.Sprintf("E:/Temp/DistTest/%s.db", fKey)
}
