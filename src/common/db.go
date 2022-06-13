package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"ques/conf"
)

var Db *gorm.DB

// init 初始化连接db
func init() {
	dbHost := conf.DbHost
	dbPort := conf.DbPort
	dbName := conf.DbName
	dbUser := conf.DbUser
	dbPassword := conf.DbPassword

	dbBase, err := gorm.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		log.Fatal(err)
	}
	Db = dbBase
	Db.LogMode(true)
	Db.SingularTable(true)
}
