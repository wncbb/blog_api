package user

import (
	"net/http"

	"blog_api/api/define"
	sessionMw "blog_api/api/middleware/session"

	"github.com/gin-gonic/gin"
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
