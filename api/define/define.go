package define

import (
	pb "blog_api/api/protobuf"
)

const CtxRespKey = "resp"
const CtxNeedCacheKey = "needCache"
const CtxUserIdKey = "userId"

const SessionUserIdKey = "uid"

const ShowMsgSuccess = "success"

type RespFormatType int

const (
	RespFormatTypeUnset RespFormatType = 0
	RespFormatTypeJSON  RespFormatType = 1
	RespFormatTypePb    RespFormatType = 2
)

const QueryFormat = "format"

type GeneralQuery struct {
	// 必须在tag里加form
	Format RespFormatType `json:"format" form:"format"`
}

func NewDefaultGeneralQuery() *GeneralQuery {
	return &GeneralQuery{
		Format: RespFormatTypeJSON,
	}
}

var statusCodeMsgMap = map[pb.ResponseCode]string{
	pb.ResponseCode_Success:             "success",
	pb.ResponseCode_Empty:               "empty",
	pb.ResponseCode_QueryArgumentsError: "query arguments error",
	pb.ResponseCode_ShouldLogoutFirst:   "should logout first",
	pb.ResponseCode_ShouldLoginFirst:    "should login first",
	pb.ResponseCode_InternalError:       "internal error",
	pb.ResponseCode_GenRandomError:      "get random error",
}

func GetCodeMsg(code pb.ResponseCode) (msg string) {
	var ok bool
	if msg, ok = statusCodeMsgMap[code]; ok {
		return
	}
	msg = "unknown error msg"
	return
}
