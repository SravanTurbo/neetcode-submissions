/*
start: 8:37am
idea: linkedlist + floyd's algo
t: O(n) s: O(1)
end: 
*/
func findDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	slow, fast := 0, 0 //index
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0 //index
	for {
		slow  = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			return slow
		}
	}

	return -1
} 
