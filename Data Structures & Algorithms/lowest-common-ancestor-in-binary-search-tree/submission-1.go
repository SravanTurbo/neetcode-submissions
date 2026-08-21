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

Time: O(h) where h is the height of the tree
Space: stack - O(h)
 */


func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    if root == nil || p == nil || q == nil {
        return nil
    }

    if (root.Val - p.Val) * (root.Val - q.Val) <= 0 {
        return root
    }
    
    if (p.Val < root.Val) {
        return lowestCommonAncestor(root.Left, p, q)
    } 
    
    return lowestCommonAncestor(root.Right, p, q)
}
