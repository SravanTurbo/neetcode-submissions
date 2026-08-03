/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
BF: use a map to save nodes and check 


Idea-1
- traverse the list and make the visited node val as -1001,
- if we visit a node whose value is -10001, return true
- else false
- but we are changing values of node, and one trick to be able to revert back is instead of -1001, do (node.Val)-10000, values will be <-9000 and can be reverted by adding 10000

Idea-2:slow and fast pointer

*/

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	
    return false
}
