/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 idea:
 t: O(n)
 s: O(n)
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    pStack := []*TreeNode{p}
	qStack := []*TreeNode{q}

	for i:=0; i<len(pStack); i++ {
		pNode := pStack[i]
		qNode := qStack[i]

		if pNode == nil && qNode == nil {
			continue
		}

		if pNode == nil || qNode == nil {
			return false
		}

		if pNode.Val != qNode.Val {
			return false
		}

		pStack = append(pStack, pNode.Left, pNode.Right)
		qStack = append(qStack, qNode.Left, qNode.Right)
	}

	return true
}
