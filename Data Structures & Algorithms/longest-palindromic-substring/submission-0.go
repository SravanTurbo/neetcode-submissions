/*
start: 9:50
t: O(n*n) s:O(1)
*/
func longestPalindrome(s string) string {
	maxLen := 0
	res := ""
	for i:=0; i<len(s); i++ {
		//odd-length-palindrome
		l, r := i, i
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}

			if r-l+1 > maxLen {
				maxLen = r-l+1
				res    = s[l:r+1]
			}

			l--
			r++
		}

		//even-length-palindrome
		l, r = i, i+1
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			
			if r-l+1 > maxLen {
				maxLen = r-l+1
				res    = s[l:r+1]
			}

			l--
			r++
		}
	}

	return res
}