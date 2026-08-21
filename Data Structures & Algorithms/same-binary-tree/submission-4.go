/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
Idea: recursive DFS.

Trees are identical iff:
- p, q are both nil 
- or have the same non-nil value 
- and their left/right subtrees are identical.

Time: O(min(n, m)), where n and m are the number of nodes in the two trees.
Space: O(h), where h is the maximum recursion depth.

Trade-off: For a Skewed deep tree, recursion stack might overflow, might have to persist using iterative DFS/BFS

Idea2: Iterative BFS
Time: O(min(n,m))
Space: O(w), where w is the maximum width of the tree; in the worst case, w = O(n), so the worst-case space complexity is O(n).
*/

type Pair struct {
	P *TreeNode
	Q *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := []Pair{Pair{P:p, Q:q}}

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]

		if pair.P == nil && pair.Q == nil {
			continue
		}

		if pair.P == nil || pair.Q == nil || pair.P.Val != pair.Q.Val {
			return false
		}

		queue = append(queue, Pair{P:pair.P.Left, Q:pair.Q.Left})
		queue = append(queue, Pair{P:pair.P.Right, Q:pair.Q.Right})
	}

	return true
}
