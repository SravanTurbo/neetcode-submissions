/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left  := dfs(root.Left)
		right := dfs(root.Right)
		res    = max(res, left + right)

		return max(left, right) + 1
	}

	dfs(root)
	return res
}
