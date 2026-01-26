package errors

import "libong/common/server/http/code"

var (
	ParamErrorOrEmpty = code.Error(10011, "参数为空或错误")
	NetworkError      = code.Error(10012, "网络错误，请稍后再试")
	AccountNotExist   = code.Error(90002, "用户不存在")
	ShoeNotExist      = code.Error(90003, "鞋子不存在")
)
