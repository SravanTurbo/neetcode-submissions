/*
start: 4:00pm
idea: cycle path using DFS, return the connection with highest index in input
t: O(V+E) s: O(V+E) 
end: 4:50pm
*/

type connection struct {
	src int
	dest int
}

func findRedundantConnection(edges [][]int) []int {
    adj := make(map[int][]int)
	edgeIdx := make(map[connection]int)
	for idx, edge := range edges {
		src, dest := edge[0], edge[1]
		adj[src] = append(adj[src], dest)
		adj[dest] = append(adj[dest], src)

		edgeIdx[connection{src, dest}] = idx
	}

	visited := map[int]bool{}
	var dfs func(node, prevNode int, path []int) ([]int, bool)
	dfs = func(node, prevNode int, path []int) ([]int, bool) {
		path = append(path, node)
		if visited[node] {
			return path, true
		}

		visited[node] = true
		for _, neighbor := range adj[node] {
			if neighbor == prevNode {
				continue
			}

			path, cycle := dfs(neighbor, node, path) 
			if cycle {
				return path, true
			}
		}

		return nil, false
	}

	cyclePath, cycleExists := dfs(1, -1, []int{})
	if cycleExists {
		maxIdx := -1
		startPointFound := false
		for i:=0; i<len(cyclePath)-1; i++ {
			if startPointFound || cyclePath[i] == cyclePath[len(cyclePath)-1] {
				startPointFound = true
				maxIdx = max(maxIdx, edgeIdx[connection{cyclePath[i], cyclePath[i+1]}])
				maxIdx = max(maxIdx, edgeIdx[connection{cyclePath[i+1], cyclePath[i]}])
			}
		}

		return edges[maxIdx]
	}

	return []int{}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

