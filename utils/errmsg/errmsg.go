package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code = 1000 用户模块错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKNE_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007


	ERROR_ART_NOT_EXIST = 2001

	// 文章
	ERROR_CATENAME_USED = 3001
	ERROR_CATE_NOT_EXIST = 3002

)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "token不存在",
	ERROR_TOKNE_RUNTIME:    "token已过期",
	ERROR_TOKEN_WRONG:      "token不正确",
	ERROR_TOKEN_TYPE_WRONG: "token格式错误",

	ERROR_ART_NOT_EXIST:	"文章不存在",

	ERROR_CATENAME_USED:	"分类已经存在",
	ERROR_CATE_NOT_EXIST:	"分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
