type TrieNode struct {
	children  map[int]*TrieNode
	endOfWord bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
    root := &TrieNode{
		children: make(map[int]*TrieNode),
	}

	return WordDictionary{root: root}
}

//t:O(n) s:O(n)
func (this *WordDictionary) AddWord(word string)  {
	cur := this.root
    for i, b := range word {
		index := int(b-'a')
		child, ok := cur.children[index]
		if !ok {
			child = &TrieNode{children: make(map[int]*TrieNode)}
			cur.children[index] = child
		}

		cur = child
		
		if i == len(word)-1 {
			cur.endOfWord = true
		}
	}
}

//t:O() s:O()
func (this *WordDictionary) Search(word string) bool {
	
	var subSearch func(cur *TrieNode, word string) bool
	subSearch = func(cur *TrieNode, word string) bool {
		for i := 0; i < len(word); i++ {
			b := word[i]
			index := int(b-'a')
			child, ok := cur.children[index]
			if !ok && b != '.' {
				return false
			}

			if b == '.' {
				for _, child = range cur.children {
					if i == len(word)-1 {
						if child.endOfWord {
							return true
						} else {
							continue
						}
					}

					if subSearch(child, word[i+1:]) {
						return true
					}
				}
				return false
			}

			cur = child

			if i == len(word)-1 && !cur.endOfWord {
				return false
			}		
		}

		return true
	}

	res := subSearch(this.root, word)

	return res
}
