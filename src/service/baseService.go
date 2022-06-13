package service

import (
	"errors"
	"ques/grpc/common"
	"ques/src/tool"
)

const (
	PageNum10 = 10
	PageNum20 = 20
	PageNum50 = 50
	PageNum100 =100
)

var PageNumArr []int32
//初始化
func init()  {
	PageNumArr =[]int32{
		PageNum10,PageNum20,PageNum50,PageNum100,
	}
}
// CheckPageInput 检查过滤分页的基本参数
func CheckPageInput(page ,num *int32)error  {
	if *page<0 {
		return errors.New("page is error")
	}

	if *page ==0 {
		*page = 1
	}

	if *num<0 {
		return errors.New("num is error")
	}

	if *num ==0 {
		*num = PageNum10
	}

	if *num>0 && !tool.IsInArrayInt(PageNumArr,*num){
		return errors.New("limit is error")
	}

	return nil
}

func MakePager(num,pager,limit int32) *common.Pager {
	var page common.Pager
	page.Total = num
	page.Page = pager
	page.Limit = limit
	return &page
}
