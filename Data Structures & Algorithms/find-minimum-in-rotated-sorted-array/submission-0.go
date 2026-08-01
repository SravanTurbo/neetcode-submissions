/*
req: 
- unique
- sorted in ascending

Idea: 
- binary search
- move towards low
*/
func findMin(nums []int) int {
	l:=0
	r:=len(nums)-1
	for nums[l]>nums[r]{
		mid:=(l+r)/2
		if nums[mid]>nums[l]{
			l=mid
		}else if nums[mid]<nums[r]{
			r=mid
		}else{
			break
		}
	}

	return min(nums[l], nums[r])
}
