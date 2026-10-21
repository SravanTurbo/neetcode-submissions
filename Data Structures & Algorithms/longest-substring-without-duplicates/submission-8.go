/*
start: 7:51
idea: sliding window
t: O(n) s:O(n)
end: 8:01
*/

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	start  := 0
	state  := make(map[byte]int)

	length := 0
	for end := 0; end < len(s); end++ {
		length += 1
		state[s[end]] += 1

		for len(state) < length {
			state[s[start]] -= 1
			if state[s[start]] == 0 {
				delete(state, s[start])
			}
			start  += 1
			length -= 1
		}

		maxLen = max(maxLen, length)
	}

	return maxLen
}
