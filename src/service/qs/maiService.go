package qs
// 问卷主表服务层
import (
	"database/sql"
	"errors"
	"ques/grpc/service"
	"ques/src/dao"
	"ques/src/dao/qs"
	baseService "ques/src/service"
)
// GetList 获取问卷列表
func GetList(params service.QuesListParams)(result service.QuesList,err error)  {
	err=baseService.CheckPageInput(&params.Page, &params.Limit)

	if err !=nil {
		return result,err
	}
	var mainDao qs.MainDao
	var data []*service.QuesListData
	count ,err:=mainDao.GetList(params, func(rows *sql.Rows) error{
		if rows.Next() {
			var row service.QuesListData
			err = rows.Scan(&row.Id,&row.Title)

			if err !=nil {
				return err
			}

			//row.ModuleTypeName = qs.ModuleTypeMap[int(row.GetModuleType())]

			data = append(data,&row)

		}
		return nil
	})
	if err !=nil {
		return result,err
	}

	if count !=0 {
		result.Page = baseService.MakePager(int32(count),params.GetPage(),params.GetLimit())
	}
	if data !=nil {
		result.Data = data
	}


	return result,nil
}

// ChangeStatus 更新状态
func ChangeStatus(params service.ChangeStatusParams)(err error)  {
	if params.GetId()<=0 {
		return errors.New("id is error")
	}

	if _,ok :=qs.StatusMap[int(params.GetStatus())];!ok{
		return errors.New("status is error")
	}

	var mainDao qs.MainDao
	if dao.UpdateById(mainDao,params.GetId(), map[string]any{
		"template_status":params.GetStatus(),
	}) == false {
		return errors.New("sql execute is error")
	}
	return nil
}