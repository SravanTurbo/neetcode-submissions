/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var maxPath func(node *TreeNode) int
	maxPath = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		pathMax := node.Val
		left    := maxPath(node.Left)
		right   := maxPath(node.Right)

		pathMax = maximum([]int{pathMax, node.Val + left, node.Val + right})

		maxSum = maximum([]int{maxSum, pathMax, node.Val + left + right})

		return pathMax
	}

	maxPath(root)
	return maxSum
}

func maximum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxVal := math.MinInt32
	for i:=0; i<len(nums); i++ {
		maxVal = max(maxVal, nums[i])
	}

	return maxVal
}


