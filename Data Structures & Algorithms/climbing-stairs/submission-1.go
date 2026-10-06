/*
start: 10 am 
*/
func climbStairs(n int) int {
    res := make([]int, n)
	res[0] = 1
	if n == 1 {
		return res[0]
	}

	res[1] = 2
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res[n-1]
}
