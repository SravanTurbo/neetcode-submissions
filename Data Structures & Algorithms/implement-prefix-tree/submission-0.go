type TrieNode struct {
	children map[rune]*TrieNode
	endOfWord bool
}

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
	root := &TrieNode{children: make(map[rune]*TrieNode)}
    return PrefixTree{root: root}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for i, b := range word {
		child, exists := cur.children[b]
		if !exists {
			child = &TrieNode{children: make(map[rune]*TrieNode)}
			cur.children[b] = child
		}

		cur = child

		if i == len(word)-1 {
			cur.endOfWord = true
		}
	}
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	for i, b := range word {
		child, exists := cur.children[b]
		if !exists{
			return false
		}

		cur = child

		if i == len(word)-1 && !cur.endOfWord {
			return false
		}
	}

	return true
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.root
	for _, b := range prefix {
		child, exists := cur.children[b]
		if !exists {
			return false
		}
		cur = child
	}

	return true
}
