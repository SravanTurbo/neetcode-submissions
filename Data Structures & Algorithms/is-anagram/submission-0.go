func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphaMap := make(map[string]int)
	for i := range(s) {
		alphaMap[string(s[i])]+=1
	}

	for j := range(t) {
		alphaMap[string(t[j])]-=1
		if alphaMap[string(t[j])] == -1 {
			return false
		}
	}

	return true
}
