/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := TreeNode{}
	root.Val = preorder[0]

	var rootIndex int
	for i:=0; i<len(inorder); i++ {
		if inorder[i] == root.Val {
			rootIndex = i
		}
	}

	var leftInorder, leftPreorder, rightInorder, rightPreorder []int
	leftInorder  = inorder[:rootIndex]
	
	if rootIndex < len(inorder)-1  {
		rightInorder = inorder[rootIndex+1:]
	}
	
	if len(leftInorder) > 0 {
		leftPreorder = preorder[1:1+len(leftInorder)]
	}

	if len(rightInorder) > 0 {
		rightPreorder = preorder[1+len(leftInorder):]
	}

	root.Left  = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	
	return &root
}
