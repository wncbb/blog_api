package cache_bigcache

import (
	"time"

	"blog_api/cache/define"

	"github.com/allegro/bigcache"
	"github.com/pkg/errors"
)

type InnerBigCache struct {
	bigCache *bigcache.BigCache
}

func NewInnerBigCache() *InnerBigCache {
	var cache, _ = bigcache.NewBigCache(bigcache.Config{
		Shards: 256,
		// time after which entry can be evicted
		LifeWindow: 30 * time.Minute,
		// rps * lifeWIndow, used only initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,
		// prints information about additional memory allocation
		Verbose: true,
	})
	return &InnerBigCache{
		bigCache: cache,
	}
}

func (p *InnerBigCache) Get(key string) ([]byte, error) {
	val, err := p.bigCache.Get(key)
	if err != nil {
		if _, ok := err.(*bigcache.EntryNotFoundError); ok {
			return nil, define.NotFound(key)
		}
		return nil, errors.WithStack(err)
	}
	return val, nil
}

func (p *InnerBigCache) Set(key string, val []byte) error {
	return p.bigCache.Set(key, val)
}

func (p *InnerBigCache) Delete(key string) error {
	return p.bigCache.Delete(key)
}
