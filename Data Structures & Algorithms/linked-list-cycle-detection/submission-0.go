/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
BF: 
- traverse the list and make the visited node val as -10001,
- if we visit a node whose value is -10001, return true
- else false

*/

func hasCycle(head *ListNode) bool {
	for head != nil {
		if head.Val == -1001 {
			return true
		}
		head.Val = -1001
		head = head.Next
	}
	
    return false
}
