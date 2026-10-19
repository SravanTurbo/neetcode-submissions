/*
start: 5:36
idea: dp - dp[i] = dp[j] + wordDict(s[j:i])
t: O(n*3), s: n+m

*/
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)

	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	dp[0] = true
	for i:=1; i<=n; i++ {
		for j:=0; j<i; j++ {
			word := s[j:i]
			if _, ok := wordMap[word]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
