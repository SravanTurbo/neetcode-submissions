/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 /*
 start: 12pm, 08-15-26

 req:
 
 edge: 
 - ending soon
 - EC: take care if a list ends

 BF: create a map to remember heads and tails of list of same value 
 - map: head : node
 - map: tail : node
 - t: O(n*m) + O(log n*m)[sorting]
 - s: O(1)[map] + O(n*m)[keys]

 BF-implement: 12:15pm - 1:15pm
 */

func mergeKLists(lists []*ListNode) *ListNode {
	head := make(map[int]*ListNode)
	tail := make(map[int]*ListNode)
	
	for i:=0; i<len(lists); i++ {
		node := lists[i]
		for node != nil {
			next := node.Next
			node.Next = nil

			t, ok := tail[node.Val]
			if !ok {
				tail[node.Val] = node
				head[node.Val] = node
			} else {
				t.Next = node
				tail[node.Val] = node
			}

			node = next
		}
	}

	var keys []int
	for k, _ := range head {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	dummyHead := ListNode{}
	prev := &dummyHead
	for i:=0; i<len(keys); i++ {
		key := keys[i]

		prev.Next = head[key]
		prev = tail[key]
	}
	
	return dummyHead.Next
}
