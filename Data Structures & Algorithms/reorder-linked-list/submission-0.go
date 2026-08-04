/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
/*
BF: O(n*n); O(1)
- for every node i :
  - j = i.Next
  - tranverse nodes from j - find the last and insert after i
  - i = j.Next

Idea:
- reverse the second and use 2 pointer, one at head and other at tail

*/

func reorderList(head *ListNode) {
	slowP := head
	fastP := head

	for fastP != nil && fastP.Next != nil && fastP.Next.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}

	if fastP.Next != nil {
		fastP = fastP.Next
	}

	var prev *ListNode //this makes the mid node.Next to nil
	for slowP != nil {
		next := slowP.Next
		slowP.Next = prev
		prev = slowP //first node in reverse is saved in prev
		slowP = next
	}


	for head != nil {
		next := head.Next
		rnext := prev.Next

		head.Next = prev
		prev.Next = next

		head = next
		prev = rnext
	}
}
