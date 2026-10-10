/*
start: 10:23
idea: DP : dp[i][j] -> is s[i:j+1] palindrome, i is start, j is end of string 
dp[i][j] = true if dp[i+1][j-1] && s[i] == s[j]
t:O(n*n) s:O(n*n)
*/
func longestPalindrome(s string) string {
	maxLen := 0
	res := ""

	rows, cols := len(s), len(s)
    dp := make([][]bool, rows)
	for r := range rows {
		dp[r] = make([]bool, cols)
	}

	for i := rows-1; i >= 0; i-- {
		for j := i; j < cols; j++ {
			if s[i] != s[j] {
				continue
			}

			if j-i <= 2 || (i+1 < rows && j-1 >= 0 && dp[i+1][j-1]) {
				dp[i][j] = true
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j-i+1
				res = s[i:j+1]
			}
		}
	}

	return res
}
