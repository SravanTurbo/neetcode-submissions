/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		item.Right, item.Left = item.Left, item.Right

		if item.Left != nil {
			queue = append(queue, item.Left)
		}

		if item.Right != nil {		
			queue = append(queue, item.Right)
		}
	}

	return root
}
