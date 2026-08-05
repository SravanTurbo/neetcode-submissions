/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{}
	dummy.Next = head

	prev := dummy
	curr := head
	count := 0

	for head.Next != nil {
		if count >= n-1 {
			prev = curr
			curr = curr.Next
		}
		count++
		head = head.Next
	}

	prev.Next = curr.Next

	return dummy.Next
}
