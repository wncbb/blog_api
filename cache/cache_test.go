package cache

import (
	"testing"
	"time"

	"wncbb.cn/log"
)

func Test_Normal(t *testing.T) {
	cacheConfig := &CacheConfig{
		CacheType:   CacheTypeBigCache,
		MarshalType: MarshalTypeJSON,
	}
	cache := NewCache(cacheConfig)
	value, isExist, isExpired, err := cache.GetBytes("name")
	log.DefaultLog().Debugf(
		"LINE17: value:%v, isExist:%v, isExpired:%v, err:%v\n",
		value,
		isExist,
		isExpired,
		err,
	)
}

func Test_A1(t *testing.T) {
	cacheConfig := &CacheConfig{
		CacheType: CacheTypeBigCache,
	}
	cache := NewCache(cacheConfig)
	err := cache.Set("name", "todd", time.Duration(1)*time.Second)
	log.DefaultLog().Debugf("LINE25: err:%v\n", err)
	time.Sleep(time.Second * 2)
	var a string
	isExist, isExpired, err := cache.Get("name", &a)
	log.DefaultLog().Debugf(
		"LINE30: value:%v, isExist:%v, isExpired:%v, err:%v\n",
		a, isExist, isExpired, err,
	)
}
