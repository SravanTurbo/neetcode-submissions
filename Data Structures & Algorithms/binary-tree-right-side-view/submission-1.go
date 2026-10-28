/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
start: 2:40pm
idea: return last element of every level where a traversal for level is from left to right
t: O(n) s: O(n)
end: 2:54pm
*/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
	res := []int{}
	cur := []*TreeNode{root} //current level
	for len(cur) > 0 {
		next := []*TreeNode{} //next level
		for i:=0; i<len(cur); i++ {
			node := cur[i]
			if node.Left != nil {
				next = append(next, node.Left)
			}

			if node.Right != nil {
				next = append(next, node.Right)
			}

			if i == len(cur)-1 {
				res = append(res, node.Val)
			}
		}
		cur = next
	}

	return res
}
