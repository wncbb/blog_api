package cache

import (
	"encoding/binary"
	"time"
)

type Entry []byte

const (
	entryExpireTimeStart int = 0
	entryExpireTimeStop      = 8
	entryDataStart       int = 8
)

func NewEntry(value []byte, expireAt time.Time) Entry {
	entryLength := len(value) + entryExpireTimeStop - entryExpireTimeStart
	entry := make([]byte, entryLength)
	binary.LittleEndian.PutUint64(entry[entryExpireTimeStart:entryExpireTimeStop], uint64(expireAt.Unix()))
	copy(entry[entryDataStart:], value)
	return Entry(entry)
}

func (p Entry) Value() []byte {
	return p[entryDataStart:]
}

func (p Entry) ExpireAt() time.Time {
	timeInSec := int64(binary.LittleEndian.Uint64(p[entryExpireTimeStart:entryExpireTimeStop]))
	return time.Unix(timeInSec, 0)
}

func (p Entry) IsExpired() bool {
	return time.Now().After(p.ExpireAt())
}

func (p Entry) SetExpireAt(time time.Time) {
	timeInSec := time.Unix()
	binary.LittleEndian.PutUint64(p[entryExpireTimeStart:entryExpireTimeStart], uint64(timeInSec))
}
