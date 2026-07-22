import "slices"

func groupAnagrams(strs []string) [][]string {
	var output [][]string

	dict := make(map[string][]string)
	for k := range(strs){
		runes := []rune(strs[k])
		slices.Sort(runes)
		sortedString := string(runes)
		dict[sortedString] = append(dict[sortedString], strs[k])
	}

	for _, v := range(dict){
		output = append(output, v)
	}

	return output
}
