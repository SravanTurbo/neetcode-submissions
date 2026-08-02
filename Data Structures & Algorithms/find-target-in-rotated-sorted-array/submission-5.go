/*
req:
- find target index else -1

constraints:
- unique
- sorted in asceding order & rotated
- O(log n)

idea: 
- find the sorted half in which answer lies
- find the index
*/
func search(nums []int, target int) int {
	l := 0
	r := len(nums)-1
	
	for nums[l]>nums[r]{
		mid := (l+r)/2
		if nums[mid]>nums[l]{
			if target>=nums[l] && target<=nums[mid]{
				r=mid
				break
			}else{
				l=mid+1
			}
		}else if nums[mid]<nums[r]{
			if target>=nums[mid] && target<=nums[r]{
				l=mid
				break
			}else{
				r=mid-1
			}
		}else{
			//only 2 elements are left & nums[l]<nums[r]
			if nums[l]==target{
				r=mid
			}else{
				l=mid+1
			}
		}
	}

	for l<=r{
		mid:=(l+r)/2
		if target==nums[mid]{
			return mid
		}else if target<nums[mid]{
			r=mid-1
		}else if target>nums[mid]{
			l=mid+1
		}
	}

	return -1
}
