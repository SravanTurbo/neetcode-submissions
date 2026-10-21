/*
start: 9:15
idea: 
t: s:
end:
*/

func characterReplacement(s string, k int) int {
	state := make(map[byte]int) //active-window
	maxf  := 0
	start := 0 
	maxLen := 0
	for end := 0; end < len(s); end++ {
		state[s[end]] += 1
		maxf = max(maxf, state[s[end]]) //in any window

		if end-start+1-maxf > k {
			state[s[start]] -= 1
			start += 1
		}

		if end-start+1 > maxLen {
			maxLen = end-start+1
		}
	}

	return maxLen
}
