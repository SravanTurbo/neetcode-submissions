/*
In:
- string
output:
- bool

q:
- ")(" valid? No

req:
- every open should be closed
- every close should have open before
- order to maintain

BF:

Idea: stack

Pseudo:
func isValid(s string) bool {
	bracePairs := map{"]":"[", "}":"{", ")":"("}
	openBraces := []byte{}

	pop := func(t []byte) (byte, bool) {
		if len(t)==0{
			return "",false
		}

		return t[len(t)]
	}

	for i in len(s){
		b := s[i]
		if b is not in bracePairs {
			push to openBraces
			continue
		}
		
		if len(openBraces) == 0 {
			return false
		}

		if bracePairs[b] != openBraces.pop() {
			return false
		}
	}

	return len(openBraces)>0 
}
*/

func isValid(s string) bool {
	bracePairs := map[string]string{"]":"[", "}":"{", ")":"("}
	openBraces := []string{}

	pop := func(t []string) ([]string, string) {
		return t[:len(t)-1], t[len(t)-1]
	}

	for i:=0; i<len(s); i++{
		v, ok := bracePairs[string(s[i])]
		if !ok {
			openBraces = append(openBraces, string(s[i]))
			continue
		}

		if len(openBraces) == 0 {
			return false
		}
		
		poppedItem := ""
		openBraces, poppedItem = pop(openBraces)
		if !ok{
			return false
		}

		if v != poppedItem {
			return false
		}
	}

	return len(openBraces)==0
}
