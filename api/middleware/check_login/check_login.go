package check_login

import (
	api_define "blog_api/api/define"
	middleware_session "blog_api/api/middleware/session"
	pb "blog_api/api/protobuf"

	"github.com/gin-gonic/gin"
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
