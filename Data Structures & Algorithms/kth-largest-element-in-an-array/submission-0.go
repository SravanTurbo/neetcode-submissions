/*
start: 5:30pm
idea: minheap of len k
t: O(k log k) s: O(k)
end: 5:45
*/

type minHeap []int

func(mh minHeap) Len() int { return len(mh) }
func(mh minHeap) Less(i, j int) bool { return mh[i] < mh[j] }
func(mh minHeap) Swap(i, j int) { mh[i], mh[j] = mh[j], mh[i] }
func(mh *minHeap) Push(x any) { (*mh) = append((*mh), x.(int)) }
func(mh *minHeap) Pop() any { 
	last := (*mh)[len(*mh)-1]
	(*mh) = (*mh)[:len(*mh)-1]
	return last 
}

func findKthLargest(nums []int, k int) int {
	mh := &minHeap{}
	heap.Init(mh)

	for i:=0; i<len(nums); i++ {
		if mh.Len() < k {
			heap.Push(mh, nums[i])
		} else if nums[i] > (*mh)[0] {
			heap.Pop(mh)
			heap.Push(mh, nums[i])
		}
	}

	if mh.Len() > 0 {
		return (*mh)[0]
	}

	return -1
}
