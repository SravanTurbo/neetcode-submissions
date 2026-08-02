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
		//find sorted array in which target exists
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
			if nums[l]==target{
				return l
			}else if nums[r]==target{
				return r
			}else{
				return -1
			}
		}
	}

	fmt.Println(l,r)

	for l<r{
		mid:=(l+r)/2
		if target<nums[mid]{
			r=mid-1
		}else if target>nums[mid]{
			l=mid+1
		}else{
			return mid
			break
		}
	}

	fmt.Println(l,r)

	if nums[l]==target{
		return l
	}

	return -1
}
