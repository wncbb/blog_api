package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wncbb.cn/api/define"
	sessionMw "wncbb.cn/api/middleware/session"
)

func ShowMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessionMw.Default(c)
		userId := s.Get(define.SessionUserIdKey)
		c.JSON(http.StatusOK, gin.H{
			"userId": userId,
		})
	}
}
