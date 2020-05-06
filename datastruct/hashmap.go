package datastruct

type HashMap struct {
	data [][]HashMapEntry
}

type HashMapEntry struct {
	k, v string
}

func NewHashMap() *HashMap {
	return &HashMap{
		data: make([][]HashMapEntry, 100),
	}
}

func (hm *HashMap) Insert(k, v string) {
	if _, ok := hm.Get(k); ok {
		h := hm.hash(k)
		for i, entry := range hm.data[h] {
			if entry.k == k {
				hm.data[h][i] = HashMapEntry{k, v}
			}
		}
	}
	h := hm.hash(k)
	hm.data[h] = append(hm.data[h], HashMapEntry{k, v})
}

func (hm *HashMap) Get(k string) (string, bool) {
	h := hm.hash(k)
	var value string
	var found bool
	for _, entry := range hm.data[h] {
		if entry.k == k {
			value = entry.v
			found = true
		}
	}
	return value, found
}

func (hm *HashMap) hash(k string) int {
	return len(k) % len(hm.data)
}
