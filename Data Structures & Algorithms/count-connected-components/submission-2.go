/*
t: O(V+E) s: O(V+E)
*/
func countComponents(n int, edges [][]int) int {
    adjacency := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjacency[u] = append(adjacency[u], v)
		adjacency[v] = append(adjacency[v], u)
	}

	visited := make(map[int]bool)
	var dfs func(n int)
	dfs = func(n int) {
		visited[n] = true
		for _, adj := range adjacency[n] {
			if !visited[adj] {
				dfs(adj)
			}
		}
	}

	count := 0
	for i := range n {
		if !visited[i] {
			dfs(i)
			count++
		}
	}

	return count
}
