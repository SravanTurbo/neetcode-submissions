/*
pseudo-code:

maxLen=0
currLen=0
drop=-1
M=map[string]int [char:index]
s.lower()
for i:=0; i<len(s); i++:
	if s[i] not exists in M:
		M[s[i]] = i
		currLen=i-drop
		maxLen=max(maxLen, currLen)
	elif M[s[i]] < drop:
		M[s[i]] = i
		currLen=i-drop
		maxLen=max(maxLen, currLen)
	elsif M[s[i]] > drop:
		drop=M[s[i]]
		currLen=i-drop
		maxLen=max(maxLen, currLen)
		M[s[i]]=i
return maxLen
*/
func lengthOfLongestSubstring(s string) int {
	maxLen, currLen, drop := 0, 0, -1
	M := make(map[byte]int)

	for i:=0; i<len(s); i++{
		v, ok := M[s[i]]
		if !ok || v<drop{
			M[s[i]] = i
			currLen = i-drop
			maxLen = max(maxLen, currLen)
		}else{
			drop=M[s[i]]
			currLen=i-drop
			maxLen=max(maxLen, currLen)
			M[s[i]]=i
		}
	}

	return maxLen
}
