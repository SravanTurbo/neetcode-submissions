/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 Idea: check is subtree is sametree with root starting at all nodes.

 Time: O(m*n), where m is substree nodes, n tree nodes
 Space: O(m+n), queue of len n for tree nodes and stack of len m for subtree
 */

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    queue := []*TreeNode{root}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if isSameTree(item, subRoot) {
			return true
		}

		if item.Left != nil {
			queue = append(queue, item.Left)
		}

		if item.Right != nil {
			queue = append(queue, item.Right)
		}
	}

	return false
}
