func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i:=0; i<len(nums)-2; i++{
		if i>0 && nums[i] == nums[i-1] {
			continue //nums[i-1]'s triplets are superset of nums[i] when equal - deduplication
		}

		res := getTriplets(i, nums)
		result = append(result, res...)
	}
	
	return result
}

func getTriplets(i int, nums []int) [][]int{
	res := twoSum(nums[i], nums[i+1:])

	for j:=0; j<len(res); j++{
		res[j]=append(res[j], nums[i])
	}
	
	return res
}

func twoSum(k int, n []int) [][]int{
	var result [][]int
	i, j := 0, len(n)-1

	for i<j{
		sum := k+n[i]+n[j]
		if sum < 0{
			i++	
		}else if sum > 0{
			j--
		}else{
			result = append(result, []int{n[i], n[j]})
			i++
			j--

			for i<j && n[i]==n[i-1]{
				i++
			}

			for i<j && n[j]==n[j+1]{
				j--
			}
		}
	}
	return result
}
