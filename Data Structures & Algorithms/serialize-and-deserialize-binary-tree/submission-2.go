/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    delimiter  string
	nilNodeVal string
}

func Constructor() Codec {
    return Codec {
		delimiter: ",",
		nilNodeVal: "N",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return this.nilNodeVal
	}

	queue := []*TreeNode{root}

	vals := []string{}
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		if node == nil {
			vals = append(vals, this.nilNodeVal)
			continue
		}

		vals = append(vals, strconv.Itoa(node.Val))
		queue = append(queue, []*TreeNode{node.Left, node.Right}...)
	}

	return strings.Join(vals, this.delimiter)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, this.delimiter)

	if vals[0] == this.nilNodeVal {
		return nil
	}

	rootVal, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}

	for i:=1; i<len(vals); i++ {
		node := queue[0]
		queue = queue[1:]

		if vals[i] != this.nilNodeVal {
			leftVal, _ := strconv.Atoi(vals[i])
			node.Left = &TreeNode{Val: leftVal}
			queue = append(queue, node.Left)
		}
		
		i++

		if vals[i] != this.nilNodeVal {
			rightVal, _ := strconv.Atoi(vals[i])
			node.Right = &TreeNode{Val: rightVal}
			queue = append(queue, node.Right)
		}
	}

	return root
}
