/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type nodeMeta struct {
	count int
	position  int
}

func kthSmallest(root *TreeNode, k int) int {
    meta := make(map[*TreeNode]nodeMeta)
	compute(root, meta)
	return getKthSmallest(root, k, meta)
}

func compute(node *TreeNode, meta map[*TreeNode]nodeMeta) (resp nodeMeta) {
	if node == nil {
		return resp
	}

	left := compute(node.Left, meta)
	right := compute(node.Right, meta)

	resp.position = left.count + 1
	resp.count = left.count + right.count + 1
	meta[node] = resp

	return resp
}

func getKthSmallest(node *TreeNode, k int, meta map[*TreeNode]nodeMeta) int {
	nodenodeMeta := meta[node]
	
	if k > nodenodeMeta.position {
		return getKthSmallest(node.Right, k-nodenodeMeta.position, meta)
	} else if k < nodenodeMeta.position {
		return getKthSmallest(node.Left, k, meta)
	}
	
	return node.Val
}