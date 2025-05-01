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

	groupList := make([]*FileGroup, 0, 8)
	tempMap := make(map[string]*FileGroup, 8)
	if info.Status == int32(Model.ScanStatusEnum_Completed) {
		for _, item := range info.FileList {
			group, exists := tempMap[item.FileKey]
			if !exists {
				group = &FileGroup{
					Key:  item.FileKey,
					List: make([]*Model.FileScanModel, 0, 4),
				}

				tempMap[item.FileKey] = group
				groupList = append(groupList, group)
			}

			group.List = append(group.List, item)
		}
	}

	m["Files"] = groupList

	ctx.JSON(200, HttpTools.CreateSuccessHSResponse(m))
}

type FileGroup struct {
	Key  string                 `json:"Key"`
	List []*Model.FileScanModel `json:"FileList"`
}
