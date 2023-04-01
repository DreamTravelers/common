package err_code

//所有错误码都在这里枚举

var (
	Success = buildError(0, "success")

	// 1xxxx 公共类错误
	InternalError  = buildError(1000, "系统异常，请稍后重试")
	UnknownError   = buildError(1001, "未知异常，请稍后重试")
	ParaError      = buildError(1002, "参数检查异常")
	UnDefinedError = buildError(1003, "未定义异常")

	// 2xxx 伍佰冻品商行错误
	GetUserInfoFromWXError = buildError(2000, "获取用户信息失败")
	UserStatusException    = buildError(2001, "用户状态异常")
	NotAllowCategoryDelete = buildError(2002, "当前类别下存在商品，不允许删除")

	// 3xxx LazyAI 错误码
	ConnectionTimedOut = buildError(3000, "连接超时")
)
