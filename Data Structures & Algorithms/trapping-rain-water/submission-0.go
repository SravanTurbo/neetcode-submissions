/*
start: 1:57
idea: 2 pointer
t: s:
end:
*/
func trap(height []int) int {
	total := 0
	
	l, r := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for l < r {
		leftMax  = max(leftMax, height[l])
		rightMax = max(rightMax, height[r])
		
		if leftMax <= rightMax {
			total += leftMax - height[l]
			l++
		} else {
			total += rightMax - height[r]
			r--
		}
	}

	return total
}
