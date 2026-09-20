/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    var diameter int 

	var depthOfBinaryTree func(node *TreeNode) int 
	depthOfBinaryTree = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := depthOfBinaryTree(node.Left)
		right := depthOfBinaryTree(node.Right)

		diameter = max(diameter, right + left)

		return max(right, left) + 1
	}

	depthOfBinaryTree(root)
	return diameter
}
