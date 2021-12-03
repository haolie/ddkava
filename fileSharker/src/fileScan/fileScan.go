package fileScan

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"fileSharker/config"
	"fileSharker/src/model"
)

const (
	con_buff_size = 1024
	con_off_size  = 1024 * 1024
	con_max_num   = 64
)

// 扫描指定路径下的所有文件
func ScanDir(dirPath string, deep bool) (fileList []*model.FileScan, err error) {

	list, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, item := range list {

		tempName := path.Join(dirPath, item.Name())
		// 处理文件
		if !item.IsDir() {
			if !config.CheckFileType(path.Ext(item.Name())) {
				continue
			}

			scanObj, err := scan(tempName)
			if err != nil {
				return nil, err
			}

			fileList = append(fileList, scanObj)

			continue
		}

		if !deep {
			continue
		}

		tempList, err := ScanDir(tempName, deep)
		if err != nil {
			return nil, err
		}

		fileList = append(fileList, tempList...)
	}

	return
}

// 扫描指定文件路径
func ScanFile(filePath string) (fileObj *model.FileScan, err error) {
	return scan(filePath)
}

func scan(filePath string) (fileOjb *model.FileScan, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()

	fStat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	fSize := fStat.Size()
	num := fSize / con_off_size
	if fSize%con_off_size > 0 {
		num += 1
	}

	if num > con_max_num {
		num = con_max_num
	}

	totalBuff := make([]byte, 0, num*con_buff_size)
	buff := make([]byte, con_buff_size, con_buff_size)
	var i int64
	for ; i < num; i++ {
		_, err = f.ReadAt(buff, con_buff_size*i)
		if err != nil {
			return nil, err
		}

		totalBuff = append(totalBuff, buff...)
	}

	h := sha1.New()
	h.Write(totalBuff)

	strBuff := h.Sum(nil)
	key := fmt.Sprintf("%x\n", strBuff)

	return model.NewFileScan(fStat.Name(), filePath, key, fSize), nil
}
