package httpSv

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"fileSharker/log"
	"fileSharker/src/db/dFile"
	"fileSharker/src/db/dbManager"
	"fileSharker/src/model"
	"fileSharker/src/tool"
)

func init() {
	RegisterRoot("sisPageFile", sisPageFile)
}

// s界面接口  捕获下载地址、名称  如果已记录返回状态，未记录则添加记录
func sisPageFile(c *gin.Context) {
	downUrl, existsDownUrl := c.GetQuery(con_param_downUrl)
	vname, existsName := c.GetQuery(con_param_vname)

	// 没有downUrl 和 名称  返回错误
	if !existsDownUrl && !existsName {
		c.JSON(http.StatusOK, CreateErrHSResponse("param err"))
		return
	}

	// 查询是否有记录

	var dm *dFile.DFileModel
	var exists bool
	var err error
	var recordType int // 0 無效  1：下載鏈接  2： vname

	if existsDownUrl {
		_, exists, err = dFile.GetListWithUrl(downUrl)
		if exists {
			recordType = 1
		}
		if err != nil {
			log.LogError(err)
			c.JSON(http.StatusOK, CreateErrHSResponse("db err"))
			return
		}
	}

	if !exists {
		_, exists, err = dFile.GetListWithVName(vname)
		if err != nil {
			log.LogError(err)
			c.JSON(http.StatusOK, CreateErrHSResponse("db err"))
			return
		}
	}

	result := make(map[string]interface{}, 2)
	result[con_param_status] = model.FileStatusEnum_NoRecord
	result[con_param_record] = recordType

	// 未记录  添加记录返回未下载
	if !exists {
		// 已记录  返回下载状态
		dm = model.NewDFileModel(tool.CreateId(), "", "", 0, 0, downUrl, vname, 0, int32(model.FileStatusEnum_Record), time.Now(), time.Now())
		dbManager.GetDbLayer().AddFile(dm.GetTableName(), dm)
	} else {
		result[con_param_status] = dm.Status
	}

	c.JSON(http.StatusOK, CreateSuccessHSResponse(result))
	return
}
