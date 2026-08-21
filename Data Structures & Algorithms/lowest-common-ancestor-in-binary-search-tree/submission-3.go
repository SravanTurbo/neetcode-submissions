/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
idea: recursion

Time: O(h) where h is the height of the tree
Space: stack - O(h)

idea2: iterative traversal

Time: O(h)
Space: O(1)
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    curr := root

    for curr != nil {
        if p.Val < curr.Val && q.Val < curr.Val {
            curr = curr.Left
            continue
        } else if p.Val > curr.Val && q.Val > curr.Val {
            curr = curr.Right
            continue
        }

        return curr
    }

    return nil
}
