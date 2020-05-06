package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap(t *testing.T) {
	hm := NewHashMap()
	hm.Insert("hello", "world")
	v, ok := hm.Get("hello")
	assert.Equal(t, "world", v)
	assert.True(t, ok)

	hm.Insert("hello2", "world2")
	v, ok = hm.Get("hello2")
	assert.Equal(t, "world2", v)
	assert.True(t, ok)

	v, ok = hm.Get("hello3")
	assert.Equal(t, "", v)
	assert.False(t, ok)

	hm.Insert("hello", "world2")
	v, ok = hm.Get("hello")
	assert.Equal(t, "world2", v)
	assert.True(t, ok)
}
