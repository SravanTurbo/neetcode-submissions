/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	resp := []int{}
	if root == nil {
		return resp
	}

	left  := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	if len(left) > 0 {
		resp = append(resp, left...)
	}

	resp = append(resp, root.Val)

	if len(right) > 0 {
		resp = append(resp, right...)
	}

	return resp
}
