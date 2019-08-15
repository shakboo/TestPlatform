package e

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR: "fail",
	INVALID_PARAMS: "请求参数错误",
	ERROR_AUTH_TOKEN: "账号密码错误",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}