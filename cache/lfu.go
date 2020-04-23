package cache

//discards least frequently used item

type LFUCache struct {
	limit  int
	values map[string]LFUValue
}

type LFUValue struct {
	Value string
	Uses  int
}

func NewLFUCache(limit int) *LFUCache {
	return &LFUCache{
		limit:  limit,
		values: map[string]LFUValue{},
	}
}

func (c *LFUCache) Get(k string) (string, bool) {
	v, ok := c.values[k]
	if ok {
		v.Uses++
		c.values[k] = v
	}
	return v.Value, ok
}

func (c *LFUCache) Add(k, v string) {
	if len(c.values) >= c.limit {
		var lfu int
		var lfuKey string
		for k, v := range c.values {
			lfu = v.Uses
			lfuKey = k
			break
		}

		for k, v := range c.values {
			if v.Uses < lfu {
				lfu = v.Uses
				lfuKey = k
			}
		}
		delete(c.values, lfuKey)
	}
	c.values[k] = LFUValue{v, 0}
}
