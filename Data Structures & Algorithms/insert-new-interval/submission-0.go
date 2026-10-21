/*
start: 10:30
idea: intervals
t: O(n) s:O(n)
end: 11:11
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	merged := [][]int{}
	n := len(intervals)
	i := 0
	for i<n && intervals[i][1] < newInterval[0] {
		merged = append(merged, intervals[i])
		i++
	}

	for i<n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}

	merged = append(merged, newInterval)

	for j:=i; j<n; j++ {
		merged = append(merged, intervals[j])
	}

	return merged
}
