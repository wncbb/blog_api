package cors

import (
	"blog_api/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	AllowOrigins []string
}

func Cors(config *Config) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.AllowOrigins
	log.DefaultLog().Debugf("Cors Config: %v", corsConfig)
	return cors.New(corsConfig)
}
