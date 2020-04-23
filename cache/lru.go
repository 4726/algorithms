package cache

//discards least recently used item

type LRUCache struct {
	limit         int
	values        map[string]LRUValue
	lastUsedCount int
}

type LRUValue struct {
	Value    string
	LastUsed int
}

func NewLRUCache(limit int) *LRUCache {
	return &LRUCache{
		limit:  limit,
		values: map[string]LRUValue{},
	}
}

func (c *LRUCache) Get(k string) (string, bool) {
	v, ok := c.values[k]
	if ok {
		v.LastUsed = c.lastUsedCount
		c.lastUsedCount++
		c.values[k] = v
	}
	return v.Value, ok
}

func (c *LRUCache) Add(k, v string) {
	if len(c.values) >= c.limit {
		var lru int
		var lruKey string
		for k, v := range c.values {
			lru = v.LastUsed
			lruKey = k
			break
		}

		for k, v := range c.values {
			if v.LastUsed < lru {
				lru = v.LastUsed
				lruKey = k
			}
		}
		delete(c.values, lruKey)
	}
	c.values[k] = LRUValue{v, c.lastUsedCount}
	c.lastUsedCount++
}
