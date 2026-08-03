/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 /*
 idea: pointers at each of the list, compare and move the lower one to next
 t: O(n+m)
 s: O(1)
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}

	curr := dummy
	for list1 != nil && list2 != nil{
		if list1.Val <= list2.Val{
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 == nil{
		curr.Next = list2
	}

	if list2 == nil{
		curr.Next = list1
	}

	return dummy.Next
}
