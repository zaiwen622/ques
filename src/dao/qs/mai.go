package qs

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"ques/grpc/service"
	baseDao "ques/src/dao"
	"ques/src/model"
)

type MainDao struct {
}

const (
	ModuletypeKeySingle =iota+1
	ModuletypeKeyMinitor
	ModuletypeKeyTrain

	ModuletypeValSingle ="单独"
	ModuletypeValMinitor ="监课"
	ModuletypeValTrain="训练"

)

const (
	StatusKeyDown= iota+1
	StatusKeyUp

	StatusValDown ="挺用"
	StatusValUp ="启用"
)


var (
	m    model.QsMain
	list []model.QsMain

	ModuleTypeMap map[int]string
	StatusMap map[int]string
)

func init()  {
	ModuleTypeMap = map[int]string{
		ModuletypeKeySingle:ModuletypeValSingle,
		ModuletypeKeyMinitor:ModuletypeValMinitor,
		ModuletypeKeyTrain:ModuletypeValTrain,
	}

	StatusMap = map[int]string{
		StatusKeyDown:StatusValDown,
		StatusKeyUp: StatusValUp,
	}
}
//Table 获取表数据库操作对象
func (dao MainDao) Table() *gorm.DB {
	return baseDao.GetDb().Table(m.TableName())
}

func (dao MainDao) Model() *gorm.DB {
	return baseDao.GetDb().Model(&m)
}


// GetList 获取列表
// params 请求参数
// fun 回调方法处理获取的列表数据
// count 返回获取的列表总数
func (dao MainDao) GetList(params service.QuesListParams,fun func(rows *sql.Rows)error) (count int,err error){
	query := dao.Table()
	if params.GetGroupId() !=0 {
		query = query.Where("template_group =?",params.GetGroupId())
	}
	if params.GetCreatorId() !=0 {
		query = query.Where("oper_id = ?",params.GetCreatorId())
	}
	if params.GetTitle() !="" {
		query = query.Where("template_title like ?","%"+params.GetTitle()+"%")
	}
	if params.GetTemplateType() !=0 {
		query = query.Where("template_type = ?",params.GetTemplateType())
	}
	if params.GetStatus() !=0 {
		query = query.Where("status = ?",params.GetStatus())
	}

	query = query.Select("id, template_title").Order("created_at desc")

	count,row,err :=baseDao.GetPage(query,params.GetPage(),params.GetLimit())

	if err !=nil {
		return 0,err
	}

	if count ==0 {
		return count,nil
	}

	err = fun(row)
	if err !=nil {
		return count ,err
	}
	return count,nil
}
