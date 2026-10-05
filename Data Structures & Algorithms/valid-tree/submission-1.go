/*
t: O(V), s: O(V)   
*/
func validTree(n int, edges [][]int) bool {
   if len(edges) != n-1 {
      return false
   }

   adjacency := make(map[int][]int)
   for _, edge := range edges {
      n1, n2 := edge[0], edge[1]
      adjacency[n1] = append(adjacency[n1], n2)
      adjacency[n2] = append(adjacency[n2], n1)
   }

   visited := map[int]bool{}
   var dfs func(n, parent int) bool
   dfs = func(n, parent int) bool {
      if visited[n] {
         return false
      }

      visited[n] = true
      for _, adj := range adjacency[n] {
         if adj == parent {
            continue
         }
         if !dfs(adj, n) {
            return false
         }
      }

      return true
   }

   return dfs(0, -1) && len(visited) == n
}
