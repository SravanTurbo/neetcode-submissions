func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix, suffix, output := make([]int, n), make([]int, n), make([]int, n)
	p, s := 1, 1
	for i := range(nums){
		p *= nums[i]
		prefix[i] = p
		s *= nums[n-1-i]
		suffix[n-1-i]=s
	}

	for i := range(nums){
		out := 1
		if i>0{
			out *= prefix[i-1]
		}
		if i<n-1{
			out *= suffix[i+1]
		}
		output[i] = out
	}

	return output
}
