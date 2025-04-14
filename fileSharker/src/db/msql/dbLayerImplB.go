package msql

import (
	"github.com/jinzhu/gorm"

	"fileSharker/src/db/dbManager"
	. "fileSharker/src/db/dbModel"
)

const (
	con_db_table = "v_file_info"
)

func init() {
	dbManager.RegisterDBLayer(DbTypeEnum_Mysql, new(mysqlDbLayer))
}

type mysqlDbLayer struct {
	conn *gorm.DB
}

func (mdb *mysqlDbLayer) GetDbType() DbTypeEnum {
	return DbTypeEnum_Mysql
}

func (mdb *mysqlDbLayer) Add(obj ITable) error {
	db := mdb.getConn().Table(obj.GetTableName()).Create(obj)
	return db.Error
}

func (mdb *mysqlDbLayer) Update(obj ITable) error {
	db := mdb.getConn().Table(obj.GetTableName()).Model(obj).Update(obj)
	//db := mdb.getConn().Table(obj.GetTableName()).Update(obj)
	return db.Error
}

func (mdb *mysqlDbLayer) DeleteOne(obj ITable) error {
	db := mdb.getConn().Table(obj.GetTableName()).Delete(obj)
	return db.Error
}

func (mdb *mysqlDbLayer) GetOne(t ITable, filter string) (exists bool, err error) {
	db := mdb.getConn().Table(t.GetTableName()).First(t, filter)
	if gorm.IsRecordNotFoundError(db.Error) {
		return
	}

	err = db.Error
	if err != nil {
		return
	}
	exists = true

	return
}

func (mdb *mysqlDbLayer) GetList(t ITable, out interface{}, filter string) (exists bool, err error) {
	db := mdb.getConn().Table(t.GetTableName()).Find(out, filter)
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

func (mdb *mysqlDbLayer) GetConn() *gorm.DB {
	return mdb.getConn()
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
