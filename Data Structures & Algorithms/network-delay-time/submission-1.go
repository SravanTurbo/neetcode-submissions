/*
start: 12:56pm
idea: bfs 
t: O(V+E) s: O(V)
end: 2:31pm
*/

func networkDelayTime(times [][]int, n int, k int) int {
	travelTime := map[int]int{}
	adj := map[int][][]int{}
	for i:=0; i<len(times); i++ {
		src := times[i][0]
		adj[src] = append(adj[src], times[i])
	}

	queue := []int{k}
	travelTime[k] = 0
	for i := 0; i<len(queue); i++ {
		node := queue[i]

		edges, exists := adj[node]
		if !exists {
			continue
		}

		for _, edge := range edges {
			src, dest, time := edge[0], edge[1], edge[2]
			
			newTime := travelTime[src] + time
			oldTime, exists := travelTime[dest]
			if !exists {
				travelTime[dest] = newTime
				queue = append(queue, dest)
			} else if newTime < oldTime {
				travelTime[dest] = newTime
				queue = append(queue, dest)
			}
		}
	}

	if len(travelTime) != n {
		return -1
	}

	maxi := -1
	for _, v := range travelTime {
		maxi = max(maxi, v)
	}

	fmt.Println(adj)
	fmt.Println(travelTime)

	return maxi
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
