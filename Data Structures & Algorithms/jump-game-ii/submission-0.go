/*
start:12:50
*/
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	jumps := 0
	maxReach := 0
	i:=0
	for i < len(nums) {
		if len(nums)-1 <= maxReach {
			break
		}

		if i > maxReach {
			break
		}

		newMaxReach := 0
		for i <= maxReach {
			newMaxReach = max(newMaxReach, i + nums[i])
			i++
		}
		jumps++
		maxReach = newMaxReach
	}

	return jumps
}
