package cache

//discards most recently used item

type MRUCache struct {
	limit         int
	values        map[string]MRUValue
	lastUsedCount int
}

type MRUValue struct {
	Value    string
	LastUsed int
}

func NewMRUCache(limit int) *MRUCache {
	return &MRUCache{
		limit:  limit,
		values: map[string]MRUValue{},
	}
}

func (c *MRUCache) Get(k string) (string, bool) {
	v, ok := c.values[k]
	if ok {
		v.LastUsed = c.lastUsedCount
		c.lastUsedCount++
		c.values[k] = v
	}
	return v.Value, ok
}

func (c *MRUCache) Add(k, v string) {
	if len(c.values) >= c.limit {
		var mru int
		var mruKey string

		for k, v := range c.values {
			if v.LastUsed > mru {
				mru = v.LastUsed
				mruKey = k
			}
		}
		delete(c.values, mruKey)
	}
	c.values[k] = MRUValue{v, c.lastUsedCount}
	c.lastUsedCount++
}
