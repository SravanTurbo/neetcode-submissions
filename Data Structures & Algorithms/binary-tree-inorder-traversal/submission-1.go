/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 Idea: Recursive DFS using helper function
 */

func inorderTraversal(root *TreeNode) []int {
	resp := []int{}
	
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		resp = append(resp, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	return resp
}
