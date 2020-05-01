package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := &Trie{map[rune]*Trie{}, false}
	trie.Insert("hello")
	trie.Insert("hellos")
	assert.True(t, trie.Search("hello"))
	assert.False(t, trie.Search("hell"))
	assert.True(t, trie.Search("hellos"))
	assert.False(t, trie.Search("helloss"))
}
