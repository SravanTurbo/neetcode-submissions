/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 /*
start: 1:46pm
idea:
t: s:
end: 
*/


func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{} //dummy
	dummy.Next = head

	prev := dummy
	first := head
    for first != nil {
		if first.Next == nil {
			break
		}

		second := first.Next
		third  := second.Next

		prev.Next   = second
		second.Next = first
		first.Next  = third

		prev  = first 
		first = first.Next
	}

	return dummy.Next
}
