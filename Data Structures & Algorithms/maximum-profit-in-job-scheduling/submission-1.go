
/*
idea: dp
t: O(nlogn) s:O(n)
*/
type job struct {
	start   int
	end     int
	profit int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n    := len(startTime)
	jobs := make([]job, n)

	for i := 0; i < n; i++ {
		jobs[i] = job{
			start:  startTime[i],
			end:    endTime[i],
			profit: profit[i],
		}
	}

	sort.Slice(jobs, func(i,j int) bool {
		return jobs[i].end < jobs[j].end
	})

	//max profit using first n jobs
	dp   := make([]int, n+1)
	dp[0] = 0

	for i := 0; i < n; i++ {

		k := sort.Search(i, func(j int) bool {
			return jobs[j].end > jobs[i].start
		})

		take := jobs[i].profit + dp[k]
		skip := dp[i]

		dp[i+1] = max(take, skip)
	}

	return dp[n]
}
