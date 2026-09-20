/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

    left := height(root.Left)
	right := height(root.Right)
	if abs(left-right) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := height(node.Left)
	right := height(node.Right)

	return max(left, right) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
} 