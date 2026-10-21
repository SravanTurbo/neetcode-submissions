/*
start: 11:46am
idea: 
t: s:
*/
func partitionLabels(s string) []int {
	lastIdx := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		lastIdx[s[i]] = i
	}

	l, r     := 0, 0
	segments := []int{}
	farthest := -1
	for r < len(s) {
		farthest = max(farthest, lastIdx[s[r]])
		if farthest == r {
			segments = append(segments, r-l+1)
			farthest = -1
			r++
			l=r
		} else {
			r++
		}
	}

	return segments
}
