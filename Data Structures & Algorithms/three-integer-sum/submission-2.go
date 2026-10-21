/*
start: 12:32
idea: 2 pointer
t: O(n*n) s:O(n)
end: 12:48
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	i:=0
	for i<len(nums) {
		element := nums[i]

		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l] + nums[r] < -element {
				l++
			} else if nums[l] + nums[r] > -element {
				r--
			} else {
				result = append(result, []int{element, nums[l], nums[r]})

				prevLVal := nums[l]
				prevRVal := nums[r]
				for l < len(nums) && nums[l] == prevLVal {
					l++
				}
				for r >= 0 && nums[r] == prevRVal {
					r--
				}
			}
		}

		for i < len(nums) && nums[i] == element {
			i++
		}
	}

	return result
}
