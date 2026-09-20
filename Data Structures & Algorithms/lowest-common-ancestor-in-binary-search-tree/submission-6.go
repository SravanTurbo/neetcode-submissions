/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
idea: bst is sorted in inorder, compare with root and if p and q lie on either side of any node - that is their common ancestor when checking from top down
t: O(h)
s: O(1)
*/

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	curr := root

	for curr != nil {
		if p.Val < curr.Val && q.Val < curr.Val {
			curr = curr.Left
			continue
		}

		if p.Val > curr.Val && q.Val > curr.Val {
			curr = curr.Right
			continue
		}

		return curr
	}

	return nil
}
