package common

import (
	api_define "blog_api/api/define"
	pb "blog_api/api/protobuf"
	"blog_api/log"

	"github.com/gin-gonic/gin"
)

func SetErrorResponse(c *gin.Context, code pb.ResponseCode, msg string, err error, logMsg string) {
	pbResp := &pb.BaseResponse{}
	log.DefaultLogError(c.GetString(api_define.CtxLogIdKey), logMsg, err)
	pbResp.Code = pb.ResponseCode_InternalError
	pbResp.Msg = msg
	c.Set(api_define.CtxRespKey, pbResp)
}
