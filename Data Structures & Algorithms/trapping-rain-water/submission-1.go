/*
start: 1:57
idea: 2 pointer
t:O(n) s:O(1)
end:
*/
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	total := 0

	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]

	for left < right {
		if leftMax < rightMax {
			left += 1
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
		} else {
			right -=1 
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
		}
	}

	return total
}
