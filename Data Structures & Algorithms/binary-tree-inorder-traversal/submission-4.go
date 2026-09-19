/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 Idea1: Normal DFS 
		t: O(n)
		s: O(n) : recursion stack + variables (left, right, result)
 */

func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	left := inorderTraversal(root.Left)
	if len(left) > 0 {
		result = append(result, left...)
	}

	result = append(result, root.Val)

	right := inorderTraversal(root.Right)
	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
