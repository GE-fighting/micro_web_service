package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Init(userName string, password string, ipaddress string, port int, dbName string) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, ipaddress, port, dbName)
	GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
