package until

const (
	Success = 0
	Error   = iota
	ParamsError
	NoKnown
	CaptchaError
	LoginError
	NoRole
	TimeParse
)

var Message = map[int]string{
	Success:      "成功",
	Error:        "返回失败",
	ParamsError:  "参数错误",
	NoKnown:      "未知错误",
	CaptchaError: "验证码错误",
	LoginError:   "登录失败",
	NoRole:       "没有权限",
	TimeParse:    "时间表达式错误",
}

func Return(code int, data interface{}, msg string) map[string]interface{} {
	mp := make(map[string]interface{})
	mp["code"] = code
	if msg == "" {
		mp["msg"] = Message[code]
	} else {
		mp["msg"] = msg
	}

	mp["data"] = data
	return mp

}

func ReturnCount(code int, count int64, data interface{}, msg string) map[string]interface{} {
	mp := make(map[string]interface{})
	mp["code"] = code
	if msg == "" {
		mp["msg"] = Message[code]
	} else {
		mp["msg"] = msg
	}
	mp["count"] = count
	mp["data"] = data
	return mp

}
