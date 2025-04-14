package httpSv

//
//func init() {
//	RegisterRoot("xnPageFile", xnPageFile)
//}
//
//// xunniu 页面请求  根据地址、压缩文件名查询  返回下载状态、或补充信息
//func xnPageFile(c *gin.Context) {
//	downUrl, existsDownUrl := c.GetQuery(con_param_downUrl)
//	cname, existsName := c.GetQuery(con_param_cname)
//
//	// 没有downUrl 和 名称  返回错误
//	if !existsDownUrl && !existsName {
//		c.JSON(http.StatusOK, CreateErrHSResponse("param err"))
//		return
//	}
//
//	var dm *model.DFileModel
//	var exists bool
//	var err error
//	var cnameFind bool
//
//	if existsDownUrl {
//		dm, exists, err = dbManager.GetDbLayer().GetDFModelWithKey("DownUrl", downUrl, true)
//		if err != nil {
//			log.LogError(err)
//			c.JSON(http.StatusOK, CreateErrHSResponse("db err"))
//			return
//		}
//	}
//
//	// 补充已查询到的信息
//	if exists {
//		if dm.CompressName == "" && len(cname) > 0 {
//			dm.CompressName = cname
//			err = dbManager.GetDbLayer().UpdateFile(dm.GetTableName(), dm)
//
//			if err != nil {
//				log.LogError(err)
//				c.JSON(http.StatusOK, CreateErrHSResponse("db err"))
//				return
//			}
//		}
//	}
//
//	if !exists {
//		dm, exists, err = dbManager.GetDbLayer().GetDFModelWithKey("CName", cname, true)
//		if err != nil {
//			log.LogError(err)
//			c.JSON(http.StatusOK, CreateErrHSResponse("db err"))
//			return
//		}
//
//		cnameFind = exists
//	}
//
//	result := make(map[string]interface{}, 2)
//	result[con_param_status] = model.FileStatusEnum_NoRecord
//	result["onlyName"] = cnameFind
//	result[con_param_compressSize] = 0
//
//	if exists {
//		result[con_param_compressSize] = dm.CompressSize
//		result[con_param_status] = dm.Status
//	}
//
//	c.JSON(http.StatusOK, CreateSuccessHSResponse(result))
//	return
//}
