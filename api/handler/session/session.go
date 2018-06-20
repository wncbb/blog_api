package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wncbb.cn/api/middleware/session"
)

func TestSetName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.DefaultQuery("name", "default")
		s := session.Default(c)
		s.Set("name", name)
		err := s.Save()
		c.JSON(http.StatusOK, gin.H{
			"err":  err,
			"name": name,
		})
	}
}

func TestGetName() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := session.Default(c)
		name := s.Get("name")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	}
}
