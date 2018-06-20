package define

import (
	"fmt"
	"time"
)

const (
	NoExpiration      time.Duration = -1
	DefaultExpiration               = 24 * time.Hour
)

type CacheNotFoundError struct {
	message string
}

func (c CacheNotFoundError) Error() string {
	return c.message
}

func NotFound(key string) CacheNotFoundError {
	return CacheNotFoundError{
		message: fmt.Sprintf("cache [%s] not found", key),
	}
}
