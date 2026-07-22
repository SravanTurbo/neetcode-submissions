func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	for i := range(nums){
		frequency[nums[i]]+=1
	}

	n := len(nums)
	buckets := make([][]int, n)
	for key, value := range(frequency){
		buckets[value-1] = append(buckets[value-1], key)
	}

	var output []int
	to_fill := k
	for i := range(n) {
		bucket_size := len(buckets[n-1-i])
		if bucket_size == 0 {
			continue
		}

		fill := min(to_fill, bucket_size)
		output = append(output, buckets[n-1-i][:fill]...)
		to_fill -= fill
	}

	return output
} 
