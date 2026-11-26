/*
start: 10:00pm
idea: graph of words, dfs
t:O(n^2*m^2) s: O(n), n = len(wordList), m=len(word)
end: 10:40pm
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	endWordExists := false
	for _, word := range wordList {
		if word == endWord {
			endWordExists = true
			break
		}
	}
	if !endWordExists {
		return 0
	}


	n := len(wordList)
	adj := make(map[string][]string, n+1)
	for i:=-1; i<len(wordList); i++ { //O(n^2*m^2)
		var w1 string
		if i == -1 {
			w1 = beginWord
		} else {
			w1 = wordList[i]
		}

		for j:=0; j<len(wordList); j++ {
			w2 := wordList[j]
			if isOneCharDiff(w1, w2) {
				adj[w1] = append(adj[w1], w2)
				adj[w2] = append(adj[w2], w1)
			}
		}
	}

	lengths := make(map[string]int, n+1)
	lengths[beginWord] = 1 

	queue   := []string{beginWord}
	
	visited := make(map[string]bool, n+1)
	visited[beginWord] = true

	for i:=0; i<len(queue); i++ { //O(n*len(adj[i]))
		current := queue[i]
		currLen := lengths[current]
		for _, neighbor := range adj[current] {
			if visited[neighbor] {
				continue
			}

			lengths[neighbor] = currLen + 1
			visited[neighbor] = true
			if neighbor == endWord {
				return lengths[neighbor]
			}

			queue = append(queue, neighbor)
		}
	}

	return 0
}


func isOneCharDiff(w1, w2 string) bool { //O(m^2)
	if len(w1) != len(w2) {
		return false
	}

	count := 0
	for i:=0; i<len(w1); i++ {
		if w1[i] != w2[i] {
			count++
		}
	}

	if count == 1 {
		return true
	}

	return false
}
