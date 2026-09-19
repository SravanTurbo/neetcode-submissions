/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
Idea: Recursive DFS
t: O(n)
s: O(n) : output + recursive stack
*/

func preorderTraversal(root *TreeNode) []int {
    result := []int{}

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)
	return result
}
