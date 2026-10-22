func largestRectangleArea(heights []int) int {
	maxarea := 0
	stack := []int{}
	i := 0
	for i<len(heights) {
		if len(stack) == 0 || heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			right := i
			left  := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			area := heights[top] * ((right-1)-(left+1)+1)
			maxarea = max(maxarea, area)
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		right := i
		left := -1
		if len(stack) > 0 {
			left = stack[len(stack)-1]
		}
		area := heights[top] * ((right-1)-(left+1)+1)
		maxarea = max(maxarea, area)
	}

	return maxarea
}
