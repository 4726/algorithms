package cache

//discards most recently added item

type LIFOCache struct {
	limit  int
	values []LIFOValue
}

type LIFOValue struct {
	Key, Value string
}

func NewLIFOCache(limit int) *LIFOCache {
	return &LIFOCache{
		limit: limit,
	}
}

func (c *LIFOCache) Get(k string) (string, bool) {
	for _, v := range c.values {
		if v.Key == k {
			return v.Value, true
		}
	}
	return "", false
}

func (c *LIFOCache) Add(k, v string) {
	if len(c.values) >= c.limit {
		c.values = c.values[:len(c.values)-1]
	}
	c.values = append(c.values, LIFOValue{k, v})
}
