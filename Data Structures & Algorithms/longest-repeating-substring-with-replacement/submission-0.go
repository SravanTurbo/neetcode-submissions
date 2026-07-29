/*
Requirements:
- replace with only one char atmost k times

edge-cases:
- k=0 -> longest 
- k=s.length -> s.length
- s.length = 1 -> 1

idea: start a window from the left and keep increasing it on the right until k is exhausted, if exhausted - slide window from the left to bring back k and go on

pseudo:
counter map[string]int
mostRepeated=0
l=0
res=0
for r:=0; <len(s); r++:
	counter[s[r]]++
	mostRepeated = max(mostRepeated, counter[s[r]])

	if (r-l+1)-mostRepeated > k:
		if s[l]==s[r]:
			mostRepeated--
		counter[s[l]]--
		l++

	
	res = max(res, r-l+1)

return res
*/


func characterReplacement(s string, k int) int {
	maxWindow := 0
	counter := make(map[byte]int)
	mostRepeated := 0
	
	l := 0
	for r:=0; r<len(s); r++{
		counter[s[r]]++
		mostRepeated = max(mostRepeated, counter[s[r]])

		if (r-l+1)-mostRepeated > k{
			counter[s[l]]--
			if s[l]==s[r]{
				mostRepeated--
			}
			l++
		}

		maxWindow = max(maxWindow, r-l+1)
	}

	return maxWindow
}
