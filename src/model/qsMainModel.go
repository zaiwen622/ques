package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"ques/conf"
	"time"
)

// QsMain 问卷列表主表
type QsMain struct {
	gorm.Model
	TemplateType   sql.NullInt16 //模板类型，1考核问卷 2调研问卷
	TemplateTitle  string        `gorm:"size:255"` //模板名
	TemplateCode   string        `gorm:"size:255"` //模板编号hash
	TemplateDesc   string        `gorm:"size:500"` //问卷说明
	TemplateStatus sql.NullInt16 //状态 1为不启用，2为启用
	TemplateGroup  sql.NullInt16 //归属分组
	ScoreType      sql.NullInt16 //0 未设置 1是自动计分 2是手动填写 3不计分
	ModuleType     sql.NullInt16 //投放渠道：单独使用=1;监课=2;教研师训=3
	ModuleRuleId   sql.NullInt64 //投放渠道的拓展规则id
	ReportTotal    sql.NullInt64 //已经收集报告数
	ReadTotal      sql.NullInt64 //浏览量
	Score          string        //卷面总分
	OperId         sql.NullInt64 //操作人
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TableName 表名
func (m QsMain) TableName() string {
	return conf.DbPrefix+"qs_"+"main"
}