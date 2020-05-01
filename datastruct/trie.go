package datastruct

type Trie struct {
	children map[rune]*Trie
	end      bool
}

func Constructor() *Trie {
	return &Trie{
		children: map[rune]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	curNode := t
	for i, r := range word {
		last := i+1 == len(word)
		next, ok := curNode.children[r]
		if !ok {
			newNode := &Trie{
				children: map[rune]*Trie{},
				end:      last,
			}
			curNode.children[r] = newNode
			next = newNode
		} else {
			if last {
				next.end = true
			}
		}
		curNode = next
	}
}

func (t *Trie) Search(word string) bool {
	curNode := t
	for i, r := range word {
		last := i+1 == len(word)
		next, ok := curNode.children[r]
		if !ok {
			return false
		} else {
			if last {
				return next.end
			}
		}
		curNode = next
	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	curNode := t
	for _, r := range prefix {
		next, ok := curNode.children[r]
		if !ok {
			return false
		}
		curNode = next
	}
	return true
}
