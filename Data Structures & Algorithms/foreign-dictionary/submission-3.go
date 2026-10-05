/*
start: 8:55am
*/
func foreignDictionary(words []string) string {
	graph := make(map[byte][]byte)	//directed in lex order
	inDegrees := make(map[byte]int)

	for _, word := range words {
		for i:=0; i<len(word); i++ {
			char := word[i]
			inDegrees[char] = 0
		}
	}

	var add func(src, dest byte)
	add = func(src, dest byte) {
		for _, nei := range graph[src] {
			if nei == dest {
				return
			}
		}
		graph[src] = append(graph[src], dest)
		inDegrees[dest] += 1
	}

	for i:=0; i<len(words)-1; i++{
		cur  := words[i]
		next := words[i+1]

		for j:=0; j<len(cur); j++ {
			if j == len(next) {
				return ""
			}

			if cur[j] == next[j] {
				continue
			}
			add(cur[j], next[j])
			break
		}
		//j<len(next) case - needed for indegrees - handled above
	}

	queue := []byte{}
	for k, v := range inDegrees {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	order := []byte{}
	for i:=0; i<len(queue); i++ {
		cur := queue[i]

		order = append(order, cur)
		for _, nei := range graph[cur] {
			inDegrees[nei]--
			if inDegrees[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	for _, indegree := range inDegrees {
		if indegree != 0 {
			return ""
		}
	}

	return string(order)
}
