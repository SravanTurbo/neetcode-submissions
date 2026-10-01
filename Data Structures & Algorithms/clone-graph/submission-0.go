/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

 /*
 t: O(n*m) s:O(V)
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	queue   := []*Node{node}
	copyMap  := make(map[int]*Node)
	visited := make(map[int]bool)

	for i := 0; i < len(queue); i++ {
		cur := queue[i]

		if visited[cur.Val] {
			continue
		}

		copi, exists := copyMap[cur.Val]
		if !exists {
			copi = &Node{Val: cur.Val}
			copyMap[cur.Val] = copi
		}

		copiNeighbors := []*Node{}
		for _, neighbor := range cur.Neighbors {
			copiNeighbor, exists := copyMap[neighbor.Val]
			if !exists {
				copiNeighbor = &Node{Val: neighbor.Val}
				copyMap[neighbor.Val] = copiNeighbor
			}

			copiNeighbors = append(copiNeighbors, copiNeighbor)
		}

		queue = append(queue, cur.Neighbors...)

		copi.Neighbors = copiNeighbors

		visited[copi.Val] = true
	}

	return copyMap[node.Val]
}
