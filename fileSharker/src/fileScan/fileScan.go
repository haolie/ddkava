package fileScan

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"fileSharker/config"
	"fileSharker/src/model"
)

const (
	con_buff_size = 1024        // 采样大小
	con_off_size  = 1024 * 1024 // 采样区间（每个con_off_size采样con_buff_size）
	con_max_num   = 64          // 最大采样次数
)

// 扫描指定路径下的所有文件
func ScanDir(dirPath string, deep bool) (fileList []*model.FileScan, err error) {

	fileList = make([]*model.FileScan, 0, 8)
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

			for i := range fileList {
				if fileList[i] == nil {
					fmt.Println(22323)
				}
			}

			if scanObj == nil {
				fmt.Println(22323)
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

		for i := range tempList {
			if tempList[i] == nil {
				fmt.Println(22323)
			}
		}

		for i := range fileList {
			if fileList[i] == nil {
				fmt.Println(22323)
			}
		}

		fileList = append(fileList, tempList...)

		for i := range fileList {
			if fileList[i] == nil {
				fmt.Println(22323)
			}
		}
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
		errc := f.Close()
		if err == nil {
			err = errc
		}
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

	// 采样最大次数限制
	if num > con_max_num {
		num = con_max_num
	}

	totalBuff := make([]byte, 0, num*con_buff_size)
	buff := make([]byte, con_buff_size, con_buff_size)
	var i int64
	for ; i < num; i++ {
		n, err := f.ReadAt(buff, con_buff_size*i)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(n)
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

// EachDir 遍历目录
// @description:
// parameter:
//		@dirPath: 目录
//		@deep: 是否深层递归
//		@callBack: 回调
// return:
//		@error: 错误
func EachDir(dirPath string, deep bool, callBack func(dir string, info fs.FileInfo) error) error {
	if callBack == nil {
		return nil
	}

	dirStar, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil || !dirStar.IsDir() {
		return err
	}

	infoList, err := ioutil.ReadDir(dirPath)
	if err != nil || len(infoList) == 0 {
		return err
	}

	for _, item := range infoList {
		err = callBack(dirPath, item)
		if err != nil {
			return err
		}

		if !deep {
			continue
		}

		err = EachDir(path.Join(dirPath, item.Name()), deep, callBack)
		if err != nil {
			return err
		}
	}

	return nil
}
