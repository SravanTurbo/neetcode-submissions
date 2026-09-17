/*
t: O(n) 
s: O(1)
*/

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l<r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			break
		} else if sum > target {
			r--
		} else if sum < target {
			l++
		}
	}

	return []int{l+1,r+1}

}
