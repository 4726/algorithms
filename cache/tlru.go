package cache

import "time"

//lazy deletes ttl expired items

type TLRUCache struct {
	values map[string]TLRUValue
}

type TLRUValue struct {
	Value      string
	ExpireTime time.Time
}

func NewTLRUCache() *TLRUCache {
	return &TLRUCache{
		values: map[string]TLRUValue{},
	}
}

func (c *TLRUCache) Get(k string) (string, bool) {
	v, ok := c.values[k]
	if ok {
		if time.Now().After(v.ExpireTime) {
			delete(c.values, k)
			return "", false
		}
	}
	return v.Value, true
}

func (c *TLRUCache) Add(k, v string, expireTime time.Time) {
	c.values[k] = TLRUValue{
		Value:      v,
		ExpireTime: expireTime,
	}
}
