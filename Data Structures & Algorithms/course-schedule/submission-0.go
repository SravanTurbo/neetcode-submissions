/*
idea: to take all courses, a cycle anywhere will impact atleast 2 courses - cannot finish!
using DFS to detect cycles, 
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	//build adjacency list to represent directed graph for course prereqs
	preReq := make(map[int][]int)
	for c := range numCourses {
		preReq[c] = []int{}
	}

	for _, p := range prerequisites {
		crs, pre := p[0], p[1]
		preReq[crs] = append(preReq[crs], pre)
	}

	//dfs for cycle detection
	visited := map[int]bool{}
	var dfs func(crs int) bool
	dfs = func(crs int) bool {
		if visited[crs] {
			return false
		}

		if len(preReq[crs]) == 0 { //no prereqs
			return true
		}

		visited[crs] = true
		for _, pre := range preReq[crs] {
			if !dfs(pre) {
				return false
			}
		}
		visited[crs] = false //backtracking since its directed graph
		preReq[crs] = []int{} //already taken - crs & prereqs - should lead to preReq dfs

		return true
	}

	//apply dfs for all courses
	for c := range numCourses {
		if !dfs(c) {
			return false
		}
	}

	return true
}
