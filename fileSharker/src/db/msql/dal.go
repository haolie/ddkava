package msql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fileSharker/config"
)

func createConStr() string {
	configObj := config.GetConfig()
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", configObj.DbUser, configObj.DbPw, configObj.DbIp, configObj.DbPort, configObj.DbName)
}

func createCon() (*gorm.DB, error) {
	return gorm.Open("mysql", createConStr())
}
