package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	errorCode "ques/src/error"
	"reflect"
)
// ErrorOutPut 业务异常输出
func ErrorOutPut(c *gin.Context,code int,msg string)  {
	if code ==0 {
		code = errorCode.ErrorCode
	}
	if msg =="" {
		msg = errorCode.GetErrorMsg(code)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": msg,
	})
}
// SucceedOutPut 成功输出
// c gin框架的上下文
// objPtr protobuf的返回结构体 指针对象
func SucceedOutPut(c *gin.Context, objPtr any) {
	//获得值
	objRef := reflect.ValueOf(objPtr)
	if objRef.Kind() != reflect.Ptr && objRef.Elem().Kind() != reflect.Struct {
		panic(any("objPtr error"))
	}
	codeRef :=objRef.Elem().FieldByName("Code")
	MsgRef :=objRef.Elem().FieldByName("Msg")
	if codeRef.CanSet() {
		codeRef.SetInt(errorCode.SucceedCode)
	}
	if MsgRef.CanSet() {
		MsgRef.SetString(errorCode.GetErrorMsg(errorCode.SucceedCode))
	}

	c.JSON(http.StatusOK,objPtr)
}