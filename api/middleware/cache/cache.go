package cache

import (
	"time"

	"blog_api/api/define"
	"blog_api/log"

	utilsCache "blog_api/cache.v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/protobuf/proto"
)

func CacheMw(ttl time.Duration, cacheKey string, respPb proto.Message) func(*gin.Context) {
	inMemoryCache := utilsCache.NewInMemoryCache()
	return func(c *gin.Context) {
		var entry utilsCache.Entry
		var err error
		entry, err = inMemoryCache.Get(cacheKey)
		log.DefaultLog().Debugf("LINE20 entry:%v err:%v", entry, err)
		var cacheStatus utilsCache.CacheStatus

		if err != nil {
			cacheStatus = utilsCache.CacheStatusMiss
		} else {
			if entry.IsExpired() {
				cacheStatus = utilsCache.CacheStatusExpired
			} else {
				cacheStatus = utilsCache.CacheStatusHit
			}

			err = proto.Unmarshal(entry.GetData(), respPb)
			if err != nil {
				cacheStatus = utilsCache.CacheStatusMiss
			}
		}

		log.DefaultLog().Debugf("status: %v", cacheStatus)

		switch cacheStatus {
		case utilsCache.CacheStatusMiss:
			log.DefaultLog().Debugf("cache miss")
			fallthrough
		case utilsCache.CacheStatusExpired:
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

			data, err := proto.Marshal(resp)
			if err != nil {
				return
			}
			entry := utilsCache.NewEntry(binding.MIMEJSON, time.Now().Add(ttl), data)
			err = inMemoryCache.Set(cacheKey, entry)
			if err != nil {
				// TODO: log
				//log.DefaultLog().Errorf()
			}

		case utilsCache.CacheStatusHit:
			log.DefaultLog().Debugf("cache hit")
			log.DefaultLog().Debugf("LINE69: %v", respPb)
			c.Set(define.CtxRespKey, respPb)
			c.Abort()
		}

	}
}
