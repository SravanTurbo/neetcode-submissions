/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
idea: BST val check

Time: O(log n)
Space: stack - O(log n)
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    if (root.Val - p.Val) * (root.Val - q.Val) <= 0 {
        return root
    }
    
    if (p.Val < root.Val) {
        return lowestCommonAncestor(root.Left, p, q)
    } 
    
    return lowestCommonAncestor(root.Right, p, q)
}
