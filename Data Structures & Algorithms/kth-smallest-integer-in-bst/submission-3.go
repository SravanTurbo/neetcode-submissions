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
- find position of node using recursive DFS
- traverse again with position data for kth smallest

Time: O(n) - traversing all nodes twice
Space: O(n) - position map for all nodes
*/


type meta struct {
	count int
	position  int
}

func kthSmallest(root *TreeNode, k int) int {
    metaData := make(map[*TreeNode]meta)

	compute(root, metaData)
	
	return getKthSmallest(root, k, metaData)
}

func compute(node *TreeNode, metaData map[*TreeNode]meta) (resp meta) {
	if node == nil {
		return resp
	}

	left  := compute(node.Left, metaData)
	right := compute(node.Right, metaData)

	resp.position  = left.count + 1
	resp.count     = left.count + right.count + 1
	metaData[node] = resp

	return resp
}

func getKthSmallest(node *TreeNode, k int, metaData map[*TreeNode]meta) int {
	nodeMeta := metaData[node]
	
	if k > nodeMeta.position {
		return getKthSmallest(node.Right, k-nodeMeta.position, metaData)
	} else if k < nodeMeta.position {
		return getKthSmallest(node.Left, k, metaData)
	}
	
	return node.Val
}