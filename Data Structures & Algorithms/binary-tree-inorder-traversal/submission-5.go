/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 Idea1: Recursive DFS 
	t: O(n)
    s: O(n) : recursion stack + variables (left, right, result)

 Idea2: Recursive DFS with a helper function
	t: O(n)
	s: O(n) : recursion stack, optimising variable space at every iteration

 */

func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return 
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	return result
}
