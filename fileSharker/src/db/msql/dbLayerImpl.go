package msql

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"fileSharker/src/db/dbManager"
	"fileSharker/src/db/dbModel"
	. "fileSharker/src/model"
)

const (
	con_db_table = "v_file_info"
)

func init() {
	dbManager.RegisterDBLayer(dbModel.DbTypeEnum_Mysql, new(mysqlDbLayer))
}

type mysqlDbLayer struct {
	conn *gorm.DB
}

func (mdb *mysqlDbLayer) GetDbType() dbModel.DbTypeEnum {
	return dbModel.DbTypeEnum_Mysql
}

func (mdb *mysqlDbLayer) AddFile(tbName string, fm interface{}) error {
	db := mdb.getConn().Table(tbName).Create(fm)
	return db.Error
}

func (mdb *mysqlDbLayer) UpdateFile(tbName string, fm interface{}) error {
	db := mdb.getConn().Table(tbName).Update(fm)
	return db.Error
}

func (mdb *mysqlDbLayer) Remove(tbName string, fm interface{}) error {
	db := mdb.getConn().Table(tbName).Delete(fm)
	return db.Error
}

func (mdb *mysqlDbLayer) GetDFModelWithKey(keyName, key string, isString bool) (dFModel *DFileModel, exists bool, err error) {
	var dm DFileModel

	var whereStr string
	if isString {
		whereStr = fmt.Sprintf("%s = \"%s\"", keyName, key)
	} else {
		whereStr = fmt.Sprintf("%s = %s", keyName, key)
	}

	db := mdb.getConn().Table(SysCon_Table_Name_Down).First(&dm, whereStr)
	if gorm.IsRecordNotFoundError(db.Error) {
		return
	}

	err = db.Error
	if err != nil {
		return
	}

	exists = true
	dFModel = &dm

	return
}

func (mdb *mysqlDbLayer) GetDFListWithKey(keyName, key string) (list []*DFileModel, exists bool, err error) {
	list = make([]*DFileModel, 0)
	db := mdb.getConn().Table(SysCon_Table_Name_Down).Find(&list, fmt.Sprintf("%s = %s", keyName, key))
	if gorm.IsRecordNotFoundError(db.Error) {
		return
	}

	err = db.Error
	if err != nil {
		return
	}

	exists = db.RowsAffected > 0

	return
}

func (mdb *mysqlDbLayer) GetDFListWithSameKey(keyName, key string) (list []*DFileModel, exists bool, err error) {
	list = make([]*DFileModel, 0)
	db := mdb.getConn().Table(SysCon_Table_Name_Down).Where(fmt.Sprintf("%s LIKE %s", keyName, "\"%"+key+"%\"")).Find(&list)
	if gorm.IsRecordNotFoundError(db.Error) {
		return
	}

	err = db.Error
	if err != nil {
		return
	}

	exists = db.RowsAffected > 0

	return
}

func (mdb *mysqlDbLayer) GetSFModelWithKey(keyName, key string, isStr bool) (sFModel *SFileModel, exists bool, err error) {
	var dm SFileModel

	var whereStr string
	if isStr {
		whereStr = fmt.Sprintf("%s = \"%s\"", keyName, key)
	} else {
		whereStr = fmt.Sprintf("%s = %s", keyName, key)
	}

	db := mdb.getConn().Table(SysCon_Table_Name_Save).First(&dm, whereStr)
	if gorm.IsRecordNotFoundError(db.Error) {
		return
	}

	err = db.Error
	if err != nil {
		return
	}

	exists = true
	sFModel = &dm

	return
}

func (mdb *mysqlDbLayer) GetSFListWithKey(keyName, key string) (list []*SFileModel, exists bool, err error) {
	list = make([]*SFileModel, 0)
	db := mdb.getConn().Table(SysCon_Table_Name_Save).Find(&list, fmt.Sprintf("%s = %s", keyName, key))
	if gorm.IsRecordNotFoundError(db.Error) {
		return
	}

	err = db.Error
	if err != nil {
		return
	}

	exists = db.RowsAffected > 0

	return
}

func (mdb *mysqlDbLayer) GetSFListWithSameKey(keyName, key string) (list []*SFileModel, exists bool, err error) {
	list = make([]*SFileModel, 0)
	db := mdb.getConn().Table(SysCon_Table_Name_Save).Find(&list, fmt.Sprintf("%s like %s", keyName, key))
	if gorm.IsRecordNotFoundError(db.Error) {
		return
	}

	err = db.Error
	if err != nil {
		return
	}

	exists = db.RowsAffected > 0

	return
}

func (mdb *mysqlDbLayer) getConn() *gorm.DB {
	if mdb.conn == nil {
		var err error
		mdb.conn, err = createCon()
		if err != nil {
			panic(err)
		}
	}

	return mdb.conn
}
