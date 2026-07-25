func maxArea(heights []int) int {
	maxi := 0
	l, r := 0, len(heights)-1 //indexes
	for l<r{
		if heights[l]<heights[r]{
			maxi = max(maxi, (r-l)*heights[l])
			l++
		}else{
			maxi = max(maxi, (r-l)*heights[r])
			r--
		}
	}
	return maxi
}
