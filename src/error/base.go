package error
// 业务错误代码配置
const (
	SucceedCode = 20000+iota
	ErrorCode = 20001
)

const (
	succeedCodeMsg = "成功"
	errorCodeMsg = "异常错误"
)

var errorMap = map[int]string{
	SucceedCode:succeedCodeMsg,
	ErrorCode :errorCodeMsg,
}

func GetErrorMsg(code int) string {
	if str,ok :=errorMap[code];ok{
		return str
	}
	return "未知错误"
}