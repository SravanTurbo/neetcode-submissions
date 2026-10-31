/*
start: 10:00pm
idea: iteration, copy for replaying
t: O(n ) s: O()
end: 11:02 pm
*/

func subsets(nums []int) [][]int {
	res := [][]int{{}}

	for _, num := range nums {
		n := len(res)
		for i:=0; i<n; i++ {
			newSubset := make([]int, len(res[i]))
			copy(newSubset, res[i])
			newSubset = append(newSubset, num)
			res = append(res, newSubset)
		}
	}

	return res
}
