/*
start: 8:51am
idea: dfs for cycle detection + path constructed in reverse recursion
t: O(V+E) s: O(V+E)
end: 9:01
*/

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)

	adj := make(map[int][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make(map[int]bool, n+1)
	cycle   := make(map[int]bool, n+1)
	cycleStart := -1

	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {
		if visited[node] {
			cycleStart = node
			return true
		}

		visited[node] = true
		for _, neighbor := range adj[node] {
			if neighbor == parent {
				continue
			}

			if dfs(neighbor, node) {
				if cycleStart != -1 {
					cycle[node] = true
				}
				if node == cycleStart {
					cycleStart = -1
				}
				return true
			}
		}

		return false
	}

	dfs(1, -1)

	for i:=len(edges)-1; i>=0; i-- {
		u, v := edges[i][0], edges[i][1]
		if cycle[u] && cycle[v] {
			return []int{u, v}
		}
	}

	return nil
}
