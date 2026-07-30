/*
Requirements:
- replace with only one char atmost k times

edge-cases:
- k=0 -> longest 
- k=s.length -> s.length
- s.length = 1 -> 1

idea: 
- sliding window: window with most repeting character is the answer.
- if M is most repeating char with count as c in window W at k replacements, 
this window will only increase if count c increases even for another char N.
- moving left pointer means we should not could that moved out char in count.


pseudo:
counter map[string]int
mostRepeated=0
l=0
res=0
for r:=0; <len(s); r++:
	counter[s[r]]++
	mostRepeated = max(mostRepeated, counter[s[r]])
	if (r-l+1)-mostRepeated > k:
		counter[s[l]]--
		l++
	res = max(res, r-l+1)

return res
*/


func characterReplacement(s string, k int) int {
	maxWindow := 0
	counter := make(map[byte]int)
	mostRepeated := 0 //in any window
	
	l := 0
	for r:=0; r<len(s); r++{
		counter[s[r]]++
		mostRepeated = max(mostRepeated, counter[s[r]])

		if (r-l+1)-mostRepeated > k{
			counter[s[l]]--
			l++
		}

		maxWindow = max(maxWindow, r-l+1)
	}

	return maxWindow
}
