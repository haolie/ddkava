package FileScan

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path"

	"lyh/ddkava/fileSharker/Common/Def"
	"lyh/ddkava/fileSharker/Config"
	"lyh/ddkava/fileSharker/Model"
)

const (
	con_buff_size = 1024        // 采样大小
	con_off_size  = 1024 * 1024 // 采样区间（每个con_off_size采样con_buff_size）
	con_max_num   = 64          // 最大采样次数
)

func ScanDir(dirPath string, deep bool) (fileList []*Model.FileScanModel, err error) {

	typeMapValue, exists := Config.GetValue(Def.Config_File_TypeMap)
	if !exists {
		panic("typeMap not exists")
	}

	typeMap := typeMapValue.(map[string]struct{})
	fileList = make([]*Model.FileScanModel, 0, 8)
	list, err := os.ReadDir(dirPath)
	if err != nil {
		return
	}

	for _, item := range list {
		tempName := path.Join(dirPath, item.Name())
		// 处理文件
		if !item.IsDir() {
			exeName := path.Ext(item.Name())
			if _, exists := typeMap[exeName]; !exists {
				continue
			}

			scanObj, err := scan(tempName)
			if err != nil {
				return nil, err
			}

			fileList = append(fileList, scanObj)

			continue
		}

		// 处理文件夹
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

/*
扫描生成文件模型
*/
func scan(fPath string) (fModel *Model.FileScanModel, err error) {
	f, err := os.Open(fPath)
	if err != nil {
		return
	}

	defer f.Close()

	fStat, err := f.Stat()
	if err != nil {
		return
	}

	fSize := fStat.Size()
	// 确定取样数量 （每间隔 con_off_size 取样一次  有余值 +1）
	num := fSize / con_off_size
	if fSize%con_off_size > 0 {
		num += 1
	}

	// 限制取样最大数量
	if num > con_max_num {
		num = con_max_num
	}

	totalBuff := make([]byte, 0, con_buff_size)
	buff := make([]byte, con_buff_size, con_buff_size)
	var i int64
	for ; i < num; i++ {
		_, err := f.ReadAt(buff, i*con_buff_size)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		totalBuff = append(totalBuff, buff...)
	}

	h := sha1.New()
	h.Write(totalBuff)
	strBuff := h.Sum(nil)
	key := fmt.Sprintf("%x", strBuff)

	fModel = Model.NewFileScan(fStat.Name(), fPath, key, fSize)
	return
}
