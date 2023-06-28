package e

var MsgFlags = map[int]string{
	SUCCESS:           "ok",
	InvalidParams:     "invalid params",
	ERROR:             "fail",
	FoundResultIsNull: "查询结果为空",

	// 数据库错误
	SqlFindError:   "数据库操作错误",
	SqlUpdateError: "数据库操作错误",
	SqlCreateError: "数据库操作错误",
	SqlDeleteError: "数据库操作错误",

	// 缓存错误
	CacheGetError: "缓存操作错误",

	// Token错误
	TokenGenerateError:         "Token生成失败",
	TokenNotNullError:          "Token不能为空",
	AuthCheckTokenFailError:    "Token鉴权失败",
	AuthCheckTokenTimeoutError: "Token已超时",

	// 公共处理
	AesEncodeError: "Aes加密失败",
	AesDecodeError: "Aes解密失败",

	// http请求错误
	HttpPostRequestError: "post请求错误",

	// AI绘图错误
	AICn2EnError:         "文本翻译错误",
	AIPromptMaxLength:    "字数超限制",
	AINotFoundError:      "数据不存在",
	AIImageNotNullError:  "图片不能为空",
	AIGeneralDrawError:   "生成框架图失败",
	AIModelNotFoundError: "风格未找到",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
