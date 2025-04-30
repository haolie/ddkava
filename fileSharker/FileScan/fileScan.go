package FileScan

import (
	"context"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path"
	"sync/atomic"
	"time"

	"lyh/ddkava/fileSharker/Common/Def"
	"lyh/ddkava/fileSharker/Config"
	"lyh/ddkava/fileSharker/Model"
)

const (
	con_buff_size = 1024        // 采样大小
	con_off_size  = 1024 * 1024 // 采样区间（每个con_off_size采样con_buff_size）
	con_max_num   = 64          // 最大采样次数
)

var (
	scanInfo *Model.ScanInfo
	cancelFn context.CancelFunc
)

func init() {
	scanInfo = &Model.ScanInfo{}
}

// 重置扫描信息
func resetInfo() {
	scanInfo.Msg = ""
	scanInfo.FileList = make([]*Model.FileScanModel, 0, 8)
	scanInfo.FileCount = 0
	scanInfo.Root = ""
	scanInfo.Status = int32(Model.ScanStatusEnum_Free)
	scanInfo.StartTime = time.Now()
	scanInfo.CurDir = ""

	cancelFn = nil
}

type scanCb func(d string, isFile bool, m *Model.FileScanModel) bool

func StopScan() Model.ScanInfo {
	status := atomic.LoadInt32(&scanInfo.Status)
	if status == int32(Model.ScanStatusEnum_Running) {
		cancelFn()
	}

	return GetScanInfo()
}

func StartScanDir(ctx context.Context, dirPath string, deep bool) (err error) {
	status := atomic.LoadInt32(&scanInfo.Status)
	if status == int32(Model.ScanStatusEnum_Running) {
		err = fmt.Errorf("scan is running")
		return
	}

	f, err := os.Stat(dirPath)
	if err != nil {
		return
	}

	if !f.IsDir() {
		err = fmt.Errorf("%s is not a dir", dirPath)
		return
	}

	resetInfo()
	scanInfo.Status = int32(Model.ScanStatusEnum_Running)
	scanInfo.Root = dirPath

	var scanCtx context.Context
	scanCtx, cancelFn = context.WithCancel(ctx)
	fn := func(d string, isFile bool, m *Model.FileScanModel) bool {
		scanInfo.CurDir = d
		if m != nil {
			atomic.AddInt32(&scanInfo.FileCount, 1)
		}

		select {
		case <-scanCtx.Done():
			return false
		default:
			return true
		}
	}

	go func() {
		scanInfo.FileList, err = scanDir(dirPath, deep, fn)
		if err != nil {
			scanInfo.Status = int32(Model.ScanStatusEnum_Failed)
		} else {
			scanInfo.Status = int32(Model.ScanStatusEnum_Completed)
		}

	}()

	return nil
}

func scanDir(dirPath string, deep bool, cb scanCb) (fileList []*Model.FileScanModel, err error) {

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
		var scanObj *Model.FileScanModel
		// 处理文件
		if !item.IsDir() {
			exeName := path.Ext(item.Name())
			if _, exists := typeMap[exeName]; exists {
				scanObj, err = scan(tempName)
				if err != nil {
					return nil, err
				}

				fileList = append(fileList, scanObj)
			}

			if !cb(tempName, true, scanObj) {
				err = fmt.Errorf("scan stopped")
				return nil, err
			}

			continue
		}

		// 处理文件夹
		if !deep {
			continue
		}

		if !cb(tempName, false, nil) {
			err = fmt.Errorf("scan stopped")
			return nil, err
		}

		tempList, err := scanDir(tempName, deep, cb)
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

	var totalBuff []byte
	fSize := fStat.Size()
	if fSize > con_buff_size {
		// 确定取样数量 （每间隔 con_off_size 取样一次  有余值 +1）
		num := fSize / con_off_size
		if fSize%con_off_size > 0 {
			num += 1
		}

		// 限制取样最大数量
		if num > con_max_num {
			num = con_max_num
		}

		totalBuff = make([]byte, 0, con_buff_size*num)
		buff := make([]byte, con_buff_size, con_buff_size)
		var i int64
		for ; i < num; i++ {
			_, err := f.ReadAt(buff, i*con_off_size)
			if err == io.EOF {
				break
			}

			if err != nil {
				return nil, err
			}

			totalBuff = append(totalBuff, buff...)
		}
	} else {
		// todo 未验证
		totalBuff = make([]byte, 0, fSize)
		_, err := f.Read(totalBuff)
		if err != nil && err != io.EOF {
			return nil, err
		}
	}

	h := sha1.New()
	h.Write(totalBuff)
	strBuff := h.Sum(nil)
	key := fmt.Sprintf("%x", strBuff)

	fModel = Model.NewFileScan(fStat.Name(), fPath, key, fSize)
	return
}

/*
返回当前状态信息
*/
func GetScanInfo() Model.ScanInfo {
	temp := *scanInfo
	if temp.Status == int32(Model.ScanStatusEnum_Running) {
		temp.FileList = make([]*Model.FileScanModel, 0, 8)
	}

	return temp
}
