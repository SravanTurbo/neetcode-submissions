/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 Time: O(n)
 Space: O(n)
 */

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
    currLevel := []*TreeNode{root}
	nextLevel := []*TreeNode{}

	for i:=0; i<len(currLevel); i++ {
		node := currLevel[i]
		if node == nil {
			continue
		}

		if node.Left != nil {
			nextLevel = append(nextLevel, node.Left)
		}

		if node.Right != nil {
			nextLevel = append(nextLevel, node.Right)
		}

		if i == len(currLevel)-1 {
			result = append(result, getVals(currLevel))
			currLevel = nextLevel
			nextLevel = []*TreeNode{}
			i = -1
		}
	}

	return result
}

func getVals(nodes []*TreeNode) []int {
	result := []int{}

	for i:=0; i<len(nodes); i++ {
		node := nodes[i]
		result = append(result, node.Val)
	}

	return result
}
