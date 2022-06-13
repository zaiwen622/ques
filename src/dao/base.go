package dao

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"ques/src/common"
)
// GetDb 获取Db对象
func GetDb() *gorm.DB {
	return common.Db
}

type DaoBase interface {
	Table() *gorm.DB
	Model() *gorm.DB
}

// UpdateById 根据主键更新字段
func UpdateById(d DaoBase, id int32 ,data map[string]any) bool {
	resultDb:=d.Model().Where("id= ?",id).Updates(data)
	rowsAff:=resultDb.RowsAffected
	if rowsAff !=0 {
		return true
	}
	return false
}
// GetPage 获取分页
func GetPage(query *gorm.DB,page int32,limit int32)(count int,row *sql.Rows,err error)  {
	query = query.Where("deleted_at is null")
	query.Count(&count)
	if count ==0 {
		return 0,nil,nil
	}

	row,err =query.Offset((page-1)*limit).Limit(limit).
		Rows()
	return
}