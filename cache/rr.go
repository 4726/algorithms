package cache

//discards random item

type RRCache struct {
	limit  int
	values map[string]string
}

func NewRRCache(limit int) *RRCache {
	return &RRCache{
		limit:  limit,
		values: map[string]string{},
	}
}

func (c *RRCache) Get(k string) (string, bool) {
	v, ok := c.values[k]
	return v, ok
}

func (c *RRCache) Add(k, v string) {
	if len(c.values) >= c.limit {
		for k := range c.values {
			delete(c.values, k)
			break
		}
	}
	c.values[k] = v
}
