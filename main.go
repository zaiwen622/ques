package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ques/src/common"
	"ques/src/router"
)

func main() {
	//关闭数据库连接
	defer common.Db.Close()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// 添加路由
	router.AddApiRouter(r)


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
