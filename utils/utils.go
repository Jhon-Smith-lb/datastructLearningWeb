package utils

import (
	"dataStructLearningWeb/dm"
)

func SetResp(code int64, data map[string]interface{}, errMsg string) *dm.Resp {
	resp := &dm.Resp{}
	resp.Code = code
	resp.Data = data
	resp.Msg = errMsg
	return resp
}