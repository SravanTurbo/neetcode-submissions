/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
start: 11:34
idea: dfs
t: O(n) s: O(log n) (stack length)
end: 11:46
*/

func goodNodes(root *TreeNode) int {
	count := 0

	var dfs func(node *TreeNode, maxi int)
	dfs = func(node *TreeNode, maxi int) {
		if node == nil {
			return
		}

		if node.Val >= maxi {
			count++
			maxi = node.Val
		}

		dfs(node.Left, maxi)
		dfs(node.Right, maxi)
	}


	dfs(root, -101)

	return count
}
