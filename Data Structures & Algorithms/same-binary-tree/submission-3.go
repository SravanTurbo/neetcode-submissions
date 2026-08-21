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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pQueue := []*TreeNode{p}
	qQueue := []*TreeNode{q}

	for len(pQueue) > 0 {
		pItem := pQueue[0]
		pQueue = pQueue[1:]
		qItem := qQueue[0]
		qQueue = qQueue[1:]

		if pItem == nil && qItem == nil {
			continue
		}

		if pItem == nil || qItem == nil || pItem.Val != qItem.Val {
			return false
		}

		pQueue = append(pQueue, []*TreeNode{pItem.Left, pItem.Right}...)
		qQueue = append(qQueue, []*TreeNode{qItem.Left, qItem.Right}...)
	}

	return true
}
