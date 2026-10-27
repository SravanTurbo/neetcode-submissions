/*
start: 9:55
idea: minheap
t: O(nlogk) s: O(k) where n is len(stream)
end: 10:22
mistake: iterating nums using range
*/

type minHeap []int

func (mh minHeap) Len() int { return len(mh) }
func (mh minHeap) Less(i, j int) bool { return mh[i] < mh[j] }
func (mh minHeap) Swap(i, j int) { mh[i], mh[j] = mh[j], mh[i] }
func (mh *minHeap) Push(x any) { (*mh) = append((*mh), x.(int)) }
func (mh *minHeap) Pop() any {
	lastItem := (*mh)[len(*mh)-1]
	(*mh) = (*mh)[:len(*mh)-1]
	return lastItem
}

type KthLargest struct {
	capacity int
	storage *minHeap
}

func Constructor(k int, nums []int) KthLargest {
	mh := &minHeap{}
	heap.Init(mh)

	this := KthLargest{
		capacity: k,
		storage: mh,
	}

	for i:=0; i<len(nums); i++ {
		if len(*this.storage) == this.capacity && nums[i] <= (*this.storage)[0] {
			continue
		} else if len(*this.storage) == this.capacity {
			heap.Pop(this.storage)
		}

		heap.Push(this.storage, nums[i])
	}

	return this
}


func (this *KthLargest) Add(val int) int {
	if len(*this.storage) == this.capacity && val <= (*this.storage)[0] {
		return (*this.storage)[0]
	} else if len(*this.storage) == this.capacity {
		heap.Pop(this.storage)
	}

	heap.Push(this.storage, val)

	return (*this.storage)[0]
}
