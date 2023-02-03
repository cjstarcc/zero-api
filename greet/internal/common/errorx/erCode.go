package errorx

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

/**全局错误码*/
//服务器开小差
const ServerCommonError uint32 = 100001

// 请求参数错误
const ReuqestParamError uint32 = 100002

// token过期
const TokenExpireError uint32 = 100003

// 生成token失败
const TokenGenerateError uint32 = 100004

// 数据库繁忙,请稍后再试
const DbError uint32 = 100005

// 更新数据影响行数为0
const DbUpdateAffectedZeroError uint32 = 100006

// 数据不存在
const DataNoExistError uint32 = 100007

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[ReuqestParamError] = "参数错误"
	message[TokenExpireError] = "token失效，请重新登陆"
	message[TokenGenerateError] = "生成token失败"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[DbUpdateAffectedZeroError] = "更新数据影响行数为0"
	message[DataNoExistError] = "数据不存在"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
