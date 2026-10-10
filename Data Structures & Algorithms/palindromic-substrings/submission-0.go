/*
start: 11:18
idea: evaluating if i can be midpoint of palindrome using 2-pointer
t: O(n*n) s:O(1)
*/
func countSubstrings(s string) int {
    count := 0
	
	for i := 0; i < len(s); i++ {
		//odd-length palindrome
		l, r := i, i
		for l >=0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			count++
			l--
			r++
		}

		//even-length palindrome
		l, r = i, i+1
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			count++
			l--
			r++
		}
	}

	return count
}
