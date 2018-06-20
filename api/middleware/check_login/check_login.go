package check_login

import (
	"github.com/gin-gonic/gin"
	api_define "wncbb.cn/api/define"
	middleware_session "wncbb.cn/api/middleware/session"
	pb "wncbb.cn/api/protobuf"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := middleware_session.Default(c)
		userId := s.Get(api_define.SessionUserIdKey)
		if userId == nil {
			resp := &pb.BaseResponse{
				Code: pb.ResponseCode_ShouldLoginFirst,
			}
			c.Set(api_define.CtxRespKey, resp)
			c.Abort()
			return
		}
		c.Set(api_define.CtxUserIdKey, userId)
		c.Next()
	}
}
