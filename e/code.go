package e

const (
	SUCCESS           = 200
	ERROR             = 500
	InvalidParams     = 400
	FoundResultIsNull = 9000

	// 数据库错误
	SqlFindError   = 20001
	SqlUpdateError = 20002
	SqlCreateError = 20003
	SqlDeleteError = 20004

	// 缓存错误
	CacheGetError = 30001

	// Token错误
	TokenGenerateError         = 40001
	AuthCheckTokenFailError    = 40002
	AuthCheckTokenTimeoutError = 40003
	TokenNotNullError          = 40004

	// http请求错误
	HttpPostRequestError = 80001

	// 公共处理
	AesEncodeError = 90001
	AesDecodeError = 90002

	// AI绘图错误
	AICn2EnError         = 11001
	AIPromptMaxLength    = 11002
	AINotFoundError      = 11003
	AIImageNotNullError  = 11004
	AIGeneralDrawError   = 11005
	AIModelNotFoundError = 11006
	AIBase642ImgError    = 11008
)
