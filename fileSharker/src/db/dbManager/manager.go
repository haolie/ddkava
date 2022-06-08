package dbManager

import (
	"fileSharker/src/db/dbModel"
	"fileSharker/src/model"
	"fileSharker/src/tool"
)

var (
	dataLayer = make(map[dbModel.DbTypeEnum]dbModel.IDbLayer, 2)
)

func RegisterDBLayer(typeEnum dbModel.DbTypeEnum, layer dbModel.IDbLayer) {
	dataLayer[typeEnum] = layer
}

func GetDbLayer() dbModel.IDbLayer {
	return dataLayer[dbModel.DbTypeEnum_Mysql]
}

func SIPageView(downUrl, vName string, size int64) (hasDown bool, err error) {
	_, cashed := GetCash(downUrl)
	if cashed {
		return false, nil
	}

	dm, exists, err := GetDbLayer().GetDFModelWithKey("DownUrl", downUrl, true)
	if err != nil {
		return false, err
	}

	if exists && dm.Status == int32(model.FileStatusEnum_Record) {
		return true, nil
	}

	if !exists {
		CashDownInfo(downUrl, vName, size)
	}

	return false, nil

}

func XYPageView(downUrl, compressName string, size int64) (hasDown bool, err error) {

	dm, exists, err := GetDbLayer().GetDFModelWithKey("DownUrl", downUrl, true)
	if err != nil {
		return false, err
	}
	if exists {
		return dm.Status == int32(model.FileStatusEnum_Record), nil
	}

	cashDM, cashed := GetCash(downUrl)
	if !cashed {
		return false, nil
	}

	cashDM.Id = tool.CreateId()
	cashDM.CompressSize = size
	cashDM.CompressName = compressName

	RemoveChase(downUrl)

	err = GetDbLayer().AddFile(model.SysCon_Table_Name_Down, cashDM)

	return
}
