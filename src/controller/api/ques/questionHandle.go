package ques

import (
	"github.com/gin-gonic/gin"
	"io"
	"ques/grpc/common"
	"ques/grpc/service"
	"ques/src/controller/api"
	"ques/src/service/qs"
)

// List 问卷列表
func List(c *gin.Context)  {
	// 获取参数对象
	var params service.QuesListParams
	err :=c.ShouldBindJSON(&params)
	if err!=nil && err != io.EOF{
		api.ErrorOutPut(c,0,err.Error())
		return
	}
	// 获取列表内容
	var result service.QuesList
	result,err =qs.GetList(params)
	if err!=nil {
		api.ErrorOutPut(c,0,err.Error())
		return
	}
	//返回结果
	api.SucceedOutPut(c,&result)

	return
}

// ChangeStatus 开启关闭问卷
func ChangeStatus(c *gin.Context)  {
	// 获取参数对象
	var params service.ChangeStatusParams
	err :=c.ShouldBindJSON(&params)
	if err!=nil && err != io.EOF{
		api.ErrorOutPut(c,0,err.Error())
		return
	}

	err =qs.ChangeStatus(params)
	if err!=nil {
		api.ErrorOutPut(c,0,err.Error())
		return
	}

	var result common.Reply
	//返回结果
	api.SucceedOutPut(c,&result)

	return
}
