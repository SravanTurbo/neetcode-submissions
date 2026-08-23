/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 idea: recursive DFS
 */

func postorderTraversal(root *TreeNode) []int {
    res := []int{}

	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		postOrder(node.Left)
		postOrder(node.Right)
		res = append(res, node.Val)
	}

	postOrder(root)
	return res
}
