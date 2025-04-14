package dbModel

import (
	"github.com/jinzhu/gorm"
)

type IDbLayer interface {
	Add(obj ITable) error
	Update(obj ITable) error
	DeleteOne(obj ITable) error
	GetOne(obj ITable, filter string) (bool, error)
	GetList(t ITable, out interface{}, filter string) (exists bool, err error)
	GetConn() *gorm.DB
}
