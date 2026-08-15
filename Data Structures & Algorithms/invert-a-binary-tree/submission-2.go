/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	right := root.Right
	root.Right = invertTree(root.Left)
	root.Left = invertTree(right)

	return root
}
