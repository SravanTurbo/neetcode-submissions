/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    res   := []int{}
	stack := []*TreeNode{}
	cur   := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res   = append(res, cur.Val)
			stack = append(stack, cur)
			cur   = cur.Left
		}

		cur   = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur   = cur.Right
	}

	return res
}
