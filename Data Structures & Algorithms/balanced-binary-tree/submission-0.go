/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    var dfs func(node *TreeNode) (int, bool) 
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}

		left, balancedLeft  := dfs(node.Left)
		if !balancedLeft {
			return 0, false
		}

		right, balancedRight := dfs(node.Right)
		if !balancedRight {
			return 0, false
		}

		if abs(right-left) > 1 {
			return 0, false
		}

		return max(left, right) + 1, true
	}

	_, balanced := dfs(root)

	return balanced
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
} 