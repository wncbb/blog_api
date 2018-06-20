package cache

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"wncbb.cn/api/define"
	"wncbb.cn/cache"
	"wncbb.cn/log"
)

type CacheStatus int32

const (
	cacheStatusMiss    CacheStatus = 1
	cacheStatusExpired             = 2
	cacheStatusHit                 = 3
)

func CacheMw(ttl time.Duration, cacheKey string, respPb proto.Message) func(*gin.Context) {
	workCache := cache.NewCache(&cache.CacheConfig{
		CacheType:   cache.CacheTypeBigCache,
		MarshalType: cache.MarshalTypeJSON,
	})
	return func(c *gin.Context) {
		var err error
		isExist, isExpired, err := workCache.Get(cacheKey, respPb)
		var cacheStatus CacheStatus

		if err != nil || !isExist {
			cacheStatus = cacheStatusMiss
		} else if isExpired {
			cacheStatus = cacheStatusExpired
		} else {
			cacheStatus = cacheStatusHit
		}

		log.DefaultLog().Debugf("status: %v", cacheStatus)

		switch cacheStatus {
		case cacheStatusMiss:
			log.DefaultLog().Debugf("cache miss")
			fallthrough
		case cacheStatusExpired:
			log.DefaultLog().Debugf("cache expired")
			c.Next()
			needCache := c.GetBool(define.CtxNeedCacheKey)
			if !needCache {
				return
			}

			ctxResp, exists := c.Get(define.CtxRespKey)
			if !exists {
				return
			}
			resp, ok := ctxResp.(proto.Message)
			if !ok {
				return
			}

			workCache.Set(cacheKey, resp, ttl)

			if err != nil {
				// TODO: log
				//log.DefaultLog().Errorf()
			}

		case cacheStatusHit:
			log.DefaultLog().Debugf("cache hit")
			log.DefaultLog().Debugf("LINE69: %v", respPb)
			c.Set(define.CtxRespKey, respPb)
			c.Abort()
		}

	}
}
