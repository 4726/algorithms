package cache

//discards least recently added item

type FIFOCache struct {
	limit  int
	values []FIFOValue
}

type FIFOValue struct {
	Key, Value string
}

func NewFIFOCache(limit int) *FIFOCache {
	return &FIFOCache{
		limit: limit,
	}
}

func (c *FIFOCache) Get(k string) (string, bool) {
	for _, v := range c.values {
		if v.Key == k {
			return v.Value, true
		}
	}
	return "", false
}

func (c *FIFOCache) Add(k, v string) {
	if len(c.values) >= c.limit {
		c.values = c.values[1:]
	}
	c.values = append(c.values, FIFOValue{k, v})
}
