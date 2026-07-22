type Solution struct{
	encodingDelimiter string "-"
	positionDelimiter string ":"
}

func (s *Solution) Encode(strs []string) string {
	var encodedStr, positionStr string
	for _, str := range strs {
		l := len(str)
		lStr := strconv.Itoa(l)
		positionStr += lStr + ":"
		encodedStr += str
	}
	encodedStr = positionStr + "-" + encodedStr
	return encodedStr
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{""}
	}

	var result []string
	var positionStr, pStr string

	positionStr, encoded = splitBySymbol(encoded, "-")
	for len(positionStr)>0 {
		res := ""
		pStr, positionStr = splitBySymbol(positionStr, ":")
		position, _ := strconv.Atoi(pStr)
		res, encoded = splitByPosition(encoded, position)
		result = append(result, res)
	}

	return result
}

func splitBySymbol(str, symbol string) (string, string) {
	i := 0
	for i<len(str) {
		if string(str[i]) == symbol {
			break
		}
		i++
	}
	return str[:i], str[i+1:]
}

func splitByPosition(str string, position int) (string, string) {
	i := 0
	for i<len(str) {
		if i == position {
			break
		}
		i++
	}
	return str[:i], str[i:]
}
