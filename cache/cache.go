package cache

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"wncbb.cn/cache/cache_bigcache"
	"wncbb.cn/cache/marshal_json"

	"wncbb.cn/cache/define"
	"wncbb.cn/log"
)

// type Locker interface {
// 	TryLock() bool
// 	Unlock()
// }

type CacheType int32

const (
	CacheTypeNoop     CacheType = 1
	CacheTypeBigCache           = 2
)

type MarshalType int32

const (
	MarshalTypeJSON = 1
)

type InnerCache interface {
	Get(key string) ([]byte, error)
	//MGet(keys []string) (map[string][]byte, error)
	Set(key string, val []byte) error
	Delete(key string) error
	//IsExist(key string) bool
	//ClearAll() error
	//Locker(key string) Locker
}

type Cache struct {
	innerCache   InnerCache
	innerMarshal InnerMarshal
	cacheConfig  *CacheConfig
}

type CacheConfig struct {
	CacheType   CacheType
	MarshalType MarshalType
}

func NewCache(cfg *CacheConfig) *Cache {
	var innerCache InnerCache
	switch cfg.CacheType {
	case CacheTypeBigCache:
		fallthrough
	default:
		innerCache = cache_bigcache.NewInnerBigCache()
	}

	var innerMarshal InnerMarshal
	switch cfg.MarshalType {
	case MarshalTypeJSON:
		fallthrough
	default:
		innerMarshal = marshal_json.NewInnerJSON()
	}

	return &Cache{
		innerCache:   innerCache,
		innerMarshal: innerMarshal,
		cacheConfig:  cfg,
	}
}

func (p *Cache) Marshal(v interface{}) ([]byte, error) {
	var retBytes []byte
	var err error
	switch p.cacheConfig.MarshalType {
	case MarshalTypeJSON:
		fallthrough
	default:
		retBytes, err = json.Marshal(v)
	}
	return retBytes, err
}

func (p *Cache) Unmarshal(data []byte, v interface{}) error {
	var err error
	switch p.cacheConfig.MarshalType {
	case MarshalTypeJSON:
		fallthrough
	default:
		err = json.Unmarshal(data, v)
	}
	return err
}

// value, isExist, isExpired, error
func (p *Cache) GetBytes(key string) ([]byte, bool, bool, error) {
	var entry Entry
	var err error
	entry, err = p.innerCache.Get(key)
	if err == nil {
		log.DefaultLog().Debugf("Cache: [%s] cache hit", key)
		if entry.IsExpired() {
			log.DefaultLog().Debugf("Cache: [%s] cache expired", key)
			// TODO
			return entry.Value(), true, true, nil
		}
		return entry.Value(), true, false, nil
	}
	_, ok := errors.Cause(err).(define.CacheNotFoundError)
	if ok {
		return nil, false, false, nil
	}
	return nil, false, false, errors.WithStack(err)
}

func (p Cache) SetBytes(key string, value []byte, ttl time.Duration) error {
	entry := NewEntry(value, time.Now().Add(ttl))
	return p.innerCache.Set(key, []byte(entry))
}

func (p Cache) Set(key string, value interface{}, ttl time.Duration) error {
	valueBytes, err := p.Marshal(value)
	if err != nil {
		return errors.WithStack(err)
	}
	entry := NewEntry(valueBytes, time.Now().Add(ttl))
	return p.innerCache.Set(key, []byte(entry))
}

func (p Cache) Get(key string, value interface{}) (bool, bool, error) {
	valueBytes, isExist, isExpired, err := p.GetBytes(key)
	if isExist {
		err := p.Unmarshal(valueBytes, value)
		if err != nil {
			return false, false, errors.WithStack(err)
		}
		return isExist, isExpired, nil
	}
	return isExist, isExpired, err
}
