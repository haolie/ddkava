package impl

import (
	"github.com/gin-gonic/gin"
	"lyh/ddkava/fileSharker/Common/Def"
	"lyh/ddkava/fileSharker/FileScan"
	"lyh/ddkava/fileSharker/Model"
	"lyh/ddkava/fileSharker/httpSv/ginSev"
	HttpTools "lyh/ddkava/fileSharker/httpSv/httpTools"
)

func init() {
	ginSev.RegisterRoot(Def.Http_GetInfo, false, getInfo)
}

func getInfo(ctx *gin.Context) {

	var info = FileScan.GetScanInfo()
	m := make(map[string]interface{}, 8)
	m["Root"] = info.Root
	m["Status"] = info.Status
	m["Msg"] = info.Msg
	m["FileCount"] = info.FileCount
	m["StartTime"] = info.StartTime.Unix()
	m["CurDir"] = info.CurDir

	fileMap := make(map[string][]*Model.FileScanModel, 8)
	if info.Status == int32(Model.ScanStatusEnum_Completed) {
		for _, item := range info.FileList {
			_, exists := fileMap[item.FileKey]
			if !exists {
				fileMap[item.FileKey] = make([]*Model.FileScanModel, 0)
			}

			fileMap[item.FileKey] = append(fileMap[item.FileKey], item)
		}
	}

	m["Files"] = fileMap

	ctx.JSON(200, HttpTools.CreateSuccessHSResponse(m))
}
