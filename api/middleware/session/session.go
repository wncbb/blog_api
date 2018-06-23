package session

import (
	"blog_api/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSessionRedis(sessionRedisConfig *config.SessionRedisConfig) redis.Store {
	store, err := redis.NewStoreWithDB(sessionRedisConfig.MaxIdlNum, "tcp", sessionRedisConfig.GetConnString(), sessionRedisConfig.Password, sessionRedisConfig.Db, []byte(sessionRedisConfig.Secret))
	if err != nil {
		panic(err)
	}
	return store
}

func Session(name string, sessionRedisConfig *config.SessionRedisConfig) gin.HandlerFunc {
	store := InitSessionRedis(sessionRedisConfig)
	return sessions.Sessions(name, store)
}

func Default(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}
