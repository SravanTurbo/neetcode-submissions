/*
In: s, t
Req: 
- min substring of s
- order it not important - only the count
- else return ""

EC:
- case-sensitive
- t.len > s.len

BF:
- create map with counts from t : m
- from every i as starting of the substring, iterate j to find the substring 
that matches or more than m's count - (i-j+1) is len and stop once found
- return the min of all those valid lens
func minWindow(s string, t string) string {
	res := ""
	mini := len(s)+1
	for every i in len(s):
		cm = copyMap()
		for every j from i to len(s):
			if s[j] exists in cm:
				decrement cm[s[j]]
				if cm[s[j]] is empty:
					delete(cm, s[j])
			if len(cm)==0:
				if j-i+1 < mini{
					res = s[i:j+1]
					mini = j-i+1
				}
				
	return res
}
- 

- time-complexity: O(n*n); space-complexity: O(m)

Idea: sliding window

Pseudo:
count := map[byte]int
for i in len(t):
	count[t[i]]++

res := ""
mini := len(s)+1

matchedChars = 0
l:=0
for r:=0; r<len(s); r++:
	_, ok := count[s[r]]
	if !ok{
		continue
	}

	count[s[r]]--
	if count[s[r]]==0{
		matchedChars++
	}
	
	if matchedChars != len(count):
		continue
	
	for l<r:
		v, ok := count[s[l]]
		if !ok{
			l++
		}else if ok && v<0{
			l++
			count[s[l]]++
		}else{
			break
		}

	if r-l+1 < mini{
		res = s[i:j+1]
		mini = j-i+1
	}

return res

*/
func minWindow(s string, t string) string {
	count := make(map[byte]int)
	for i:=0; i<len(t); i++{
		count[t[i]]++
	}
		
	res:=""
	mini:=len(s)+1

	matchedChars:=0
	l:=0
	for r:=0; r<len(s); r++{
		_, ok := count[s[r]]
		if !ok{
			continue
		}

		count[s[r]]--
		if count[s[r]]==0{
			matchedChars++
		}
		
		if matchedChars != len(count){
			continue
		}
		
		for l<r{
			v, ok := count[s[l]]
			if !ok{
				l++
			}else if ok && v<0{
				count[s[l]]++
				l++
			}else{
				break
			}
		}
			

		if r-l+1 < mini{
			mini = r-l+1
			res = s[l:r+1]
		}
	}
	
	return res
}
