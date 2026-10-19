/*
start: 9:19
idea: dp
t:O(n*n) s:O(n)
*/
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	dp := make(map[int]int)
	gMax := 0

	timings := make([][]int, n)
	for i:=0; i<n; i++ {
		timings[i] = append(timings[i], []int{endTime[i], startTime[i], profit[i]}...)
	}

	sort.Slice(timings, func(i, j int) bool {
		return timings[i][0] < timings[j][0]
	})

	for i, timing := range timings {
		tEnd, tStart, tProfit := timing[0], timing[1], timing[2]
		maxi := 0
		for j:=0; j<=i; j++ {
			prft := tProfit
			jEnd := timings[j][0]
			if jEnd <= tStart {
				prft += dp[jEnd]
			}
			maxi = max(maxi, prft)
		}
		dp[tEnd] = max(dp[tEnd], maxi)
		gMax = max(gMax, dp[tEnd])
	}

	return gMax
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}