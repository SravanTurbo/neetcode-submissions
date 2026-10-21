/*
start: 12:33
idea: 
t: O() s: O()
end:
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}
	merged  = append(merged, intervals[0])

	for i:=1; i<len(intervals); i++ {
		if intervals[i][0] <= merged[len(merged)-1][1] {
			lastIdx := len(merged)-1
			merged[lastIdx][0] = min(merged[lastIdx][0], intervals[i][0])
			merged[lastIdx][1] = max(merged[lastIdx][1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}
