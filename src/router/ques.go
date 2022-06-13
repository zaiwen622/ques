package router

import (
	"github.com/gin-gonic/gin"
	"ques/src/controller/api/ques"
)
// AddApiRouter 添加路由
func AddApiRouter(r *gin.Engine)  {
	api := r.Group("/api")
	{
		//问卷
		quesRouter(api)
	}
}
// quesRouter 问卷路由
func quesRouter(r *gin.RouterGroup)  {
	//问卷
	quesGroup := r.Group("/ques")
	{
		//列表
		quesGroup.POST("/list", ques.List)
		quesGroup.POST("/changeStatus",ques.ChangeStatus)
	}
}