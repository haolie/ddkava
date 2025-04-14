package httpSv

import (
	"time"

	"fileSharker/src/tool"
)

const (
	Con_HS_Status_Fail    = 0
	Con_HS_Status_Success = 200
)

func baseDataMap() map[string]interface{} {
	dataMap := make(map[string]interface{}, 5)
	dataMap["status"] = Con_HS_Status_Success
	dataMap["content"] = ""
	dataMap["time"] = time.Now().Unix()
	dataMap["timeStr"] = tool.ToDateTimeStr(time.Now())

	return dataMap
}

func CreateErrHSResponse(errInfo string) map[string]interface{} {
	dataMap := baseDataMap()
	dataMap["status"] = Con_HS_Status_Fail
	dataMap["content"] = errInfo

	return dataMap
}

func CreateSuccessHSResponse(info interface{}) map[string]interface{} {

	dataMap := baseDataMap()
	dataMap["status"] = Con_HS_Status_Success
	dataMap["content"] = info

	return dataMap
}
