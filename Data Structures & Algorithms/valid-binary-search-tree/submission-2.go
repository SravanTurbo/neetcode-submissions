/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
Idea: 
- split into sub problem at each node and verify bottom up - Recursive DFS
- find (min, max) pair for every node where min is min of tree starting at that node and max is max
- BST becomes:
	- root > min of left child
	- root < max of right child
	- (min, max) at root = (min of left, max of right)

EC:
- nil nodes
- equal val nodes as failure
- all nodes on subtree should satify the condition, not just at immediate descendants


Time: O(n)
Space: O(h) - stack
*/

func isValidBST(root *TreeNode) bool {
	result := validate(root)
	return result.isValid
}

type bstResponse struct {
	isValid bool
	hasValue bool
	min int
	max int
}

func validate(node *TreeNode) (resp bstResponse) {	
	if node == nil {
		resp.isValid = true
		resp.hasValue = false
		return resp
	}

	left := validate(node.Left)
	right := validate(node.Right)

	if !left.isValid || !right.isValid {
		resp.isValid = false
		return resp
	}

	if left.hasValue && left.max >= node.Val {
		resp.isValid = false
		return resp
	}

	if right.hasValue && right.min <= node.Val {
		resp.isValid = false
		return resp
	}

	if !left.hasValue {
		resp.min = node.Val
	} else {
		resp.min = left.min
	}

	if !right.hasValue {
		resp.max = node.Val
	} else {
		resp.max = right.max
	}

	resp.isValid = true
	resp.hasValue = true
	
	return resp
}

