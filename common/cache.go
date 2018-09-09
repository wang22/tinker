package common

import (
	"github.com/coocood/freecache"
)

var Cache Cacher

type Cacher interface {
	Set(key string, val interface{})
	SetTTL(key string, val interface{}, expire int)
	Get(key string)
	Del(key string)
}

type RedisCacher struct{}

func (self *RedisCacher) Set(key string, val interface{}) {
}

func (self *RedisCacher) SetTTL(key string, val interface{}, expire int) {
}

func (self *RedisCacher) Get(key string) {
}

func (self *RedisCacher) Del(key string) {
}

type GOCacher struct {
	Cache *freecache.Cache
}

func (self *GOCacher) Set(key string, val interface{}) {
	// self.SetTTL(key,val, math.MaxInt32)
}

func (self *GOCacher) SetTTL(key string, val interface{}, expire int) {
	self.Cache.Set([]byte(key),[]byte(key), expire)
}

func (self *GOCacher) Get(key string) {
}

func (self *GOCacher) Del(key string) {
}
