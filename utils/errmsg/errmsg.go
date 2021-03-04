package errmsg

const (
	SUCCSE = 200
	ERROR  = 400

	STUDENT_EXIST = 1001
	STUDENT_NOT_EXIST = 1002

	DATA_VALIDATE_ERROR = 2001



)

var codeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	STUDENT_EXIST:			"学生已存在",
	STUDENT_NOT_EXIST:		"学生不存在",
	DATA_VALIDATE_ERROR:	"数据校验错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
