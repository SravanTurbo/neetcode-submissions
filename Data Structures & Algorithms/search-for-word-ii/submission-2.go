/*
start: 10:42pm
idea: trie + backtracking 
t: O(n * m * ) s: O(s) where s = sum(len(words))
end: 11:27pm
*/

type TrieNode struct {
	children map[byte]*TrieNode
	endOfWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: map[byte]*TrieNode{},
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) addWord(word string) {
	cur := t.root
	for i:=0; i<len(word); i++ {
		char := word[i]
		_, ok := cur.children[char]
		if !ok {
			cur.children[char] = NewTrieNode()
		}

		cur = cur.children[char]
	}
	cur.endOfWord = true
}


func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return []string{}
	}

	trie := NewTrie()
	for _, word := range words {
		trie.addWord(word)
	}

	res := []string{}
	wordSet := make(map[string]bool)
	rows, cols := len(board), len(board[0])
	var dfs func(r, c int, node *TrieNode, word []byte) 

	dfs = func(r, c int, node *TrieNode, word []byte) {
		if r<0 || c<0 || r>=rows || c>=cols {
			return
		}

		char := board[r][c]
		if char == '#' { //check-visited
			return
		}

		nextNode, ok := node.children[char]
		if !ok {
			return
		}

		word = append(word, char)
		if nextNode.endOfWord {
			wordSet[string(word)] = true
		}

		temp := char
		board[r][c] = '#'

		dfs(r+1, c, nextNode, word)
		dfs(r-1, c, nextNode, word)
		dfs(r, c+1, nextNode, word)
		dfs(r, c-1, nextNode, word)

		board[r][c] = temp
	}

	for r := range rows {
		for c := range cols {
			dfs(r, c, trie.root, []byte{})
		}
	}

	for word := range wordSet {
		res = append(res, word)
	}

	return res
}
