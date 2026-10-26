/*
start: 11:34pm
idea: stack
t: O(n) s: O(n)
end: 11:51pm
*/

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)

	monoStack := []int{}

	for i:=0; i<n; i++ {
		temp := temperatures[i]
		for len(monoStack) > 0 {
			idx := monoStack[len(monoStack)-1] //top-index
			if temperatures[idx] >= temp {
				break
			}
			res[idx] = i - idx
			monoStack = monoStack[:len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}


	return res
}