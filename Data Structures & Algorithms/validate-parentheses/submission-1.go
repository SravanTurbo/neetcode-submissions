func isValid(s string) bool {
	for strings.Contains(s, "()") || 
	strings.Contains(s, "[]") || 
	strings.Contains(s, "{}"){
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)
	}

	return s==""
}
