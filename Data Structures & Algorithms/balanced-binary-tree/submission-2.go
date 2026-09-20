/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 idea: recursive dfs
 t: O(n) 
 s: O(n) : recursive stack
 */

type result struct {
	height int
	balanced bool
}

func isBalanced(root *TreeNode) bool {
    var dfs func(node *TreeNode) result
	dfs = func(node *TreeNode) result {
		if node == nil {
			return result{0, true}
		}

		left := dfs(node.Left)
		if !left.balanced {
			return result{0, false}
		}

		right := dfs(node.Right)
		if !right.balanced {
			return result{0, false}
		}

		if abs(left.height-right.height) > 1 {
			return result{0, false}
		}

		return result{max(left.height, right.height) + 1, true}
	}

	output := dfs(root)
	return output.balanced
}



func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
