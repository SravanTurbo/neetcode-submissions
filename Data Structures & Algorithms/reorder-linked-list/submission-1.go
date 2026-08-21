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
- find the mid point in list
- reverse the second half in direction 
- use 2 pointer, one at head and other at tail
- traverse and keep appending
- take care of edge case at mid point to point to nil
*/

func reorderList(head *ListNode) {
	slowP := head
	fastP := head

	for fastP != nil && fastP.Next != nil && fastP.Next.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}

	//slowP is at mid

	var prev *ListNode
	for slowP != nil {
		next := slowP.Next
		slowP.Next = prev
		prev = slowP
		slowP = next
	}

	//prev is at tail

	for head != nil {
		next := head.Next
		tnext := prev.Next

		head.Next = prev
		prev.Next = next

		head = next
		prev = tnext
	}

}
