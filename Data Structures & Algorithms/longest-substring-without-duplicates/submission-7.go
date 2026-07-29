/*
// requirements:
- s.length = 0 
- s.length = 1

//idea: Sliding Window

// pseudo-code:
maxi=0
M=map()
l=0
	for r in len(s):
		if s[r] exits in M:
			l=M[s[r]]+1
			
		M[s[r]] = r
		maxi = max(maxi, r-l+1)
return maxi
*/

func lengthOfLongestSubstring(s string) int {
	M := make(map[byte]int)
	maxi := 0
	l := 0
	for r:=0; r<len(s); r++{
		v, ok := M[s[r]]
		if ok && l<=v{
			l = v+1
		}
		
		M[s[r]] = r
		maxi = max(maxi, r-l+1)
	}

	return maxi
}
