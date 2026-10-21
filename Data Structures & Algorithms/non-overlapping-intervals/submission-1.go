/*
start: 12:00
idea: sort by end time & greedy for max non overlapping, decrement from total
t: O(nlogn) s: O(1)
end: 12:22
*/

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	n := len(intervals)
	prevEnd := intervals[0][1]
	maxNonOverlap := 1

	for i:=1; i<n; i++ {
		if intervals[i][0] >= prevEnd {
			maxNonOverlap +=1
			prevEnd = intervals[i][1]
		}
	}

	return n-maxNonOverlap
}
