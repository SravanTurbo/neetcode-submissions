/*
start: 9:07
idea: topological trimming
t: O(V+E) s: O(V+E)
end: 9:28
*/

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)

	adj     := make(map[int][]int, n+1)
	degree := make(map[int]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		degree[u] += 1
		degree[v] += 1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	queue := make([]int, n+1)
	for i:=1; i<=n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	for i:=0; i<len(queue); i++ {
		u := queue[i]
		for _, v := range adj[u] {
			val, exists := degree[v]
			if !exists || val <= 1 {
				continue
			}

			degree[v] -= 1
			if degree[v] == 1 {
				queue = append(queue, v)
			}
		}

		delete(degree, u)
	}

	for i:=len(edges)-1; i>=0; i-- {
		u, v := edges[i][0], edges[i][1]
		_, uExists := degree[u]
		_, vExists := degree[v]
		if uExists && vExists {
			return []int{u, v}
		}
	}

	return nil
}
