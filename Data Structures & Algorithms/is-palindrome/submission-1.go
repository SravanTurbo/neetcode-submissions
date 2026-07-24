func isPalindrome(s string) bool {
	alphaNumeric := "abcdefghijklmnopqrstuvwxyz0123456789"
	alphaNumMap := make(map[byte]bool)
	for i:=0; i<len(alphaNumeric); i++{
		alphaNumMap[alphaNumeric[i]] = true
	}

	isAlphaNumeric := func(c byte) bool {
		if _, ok := alphaNumMap[c]; ok {
			return true
		}
		return false
	}

	s = strings.ToLower(s)
	i:=0
	j:=len(s)-1
	for i<len(s) && j>=0 && i<j {
		for i<len(s)-1 && !isAlphaNumeric(s[i]){
			i++
		}

		for j>0 && !isAlphaNumeric(s[j]){
			j--
		}

		if i<j && s[i]!=s[j]{
			return false
		}

		i++
		j--
	}

	return true
}
