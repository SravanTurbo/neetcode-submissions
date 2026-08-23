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
 Idea2: Iterative DFS
 */

func inorderTraversal(root *TreeNode) []int {
	res   := []int{}
	stack := []*TreeNode{}
	curr  := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr  = curr.Left
		}
		curr   = stack[len(stack)-1]
		stack  = stack[:len(stack)-1]
		res    = append(res, curr.Val)
		curr   = curr.Right
	}

	return res
}
