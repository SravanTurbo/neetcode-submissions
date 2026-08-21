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
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

    if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
