/*
start: 11:49pm
idea: maxheap
t: O(n log n) s: O(n)
end: 12:07
*/

type maxHeap []int

func(h maxHeap) Len() int { return len(h) }
func(h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func(h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func(h *maxHeap) Push(x any) { (*h) = append((*h), x.(int)) }
func(h *maxHeap) Pop() any { 
	lastItem := (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	return lastItem
}

func lastStoneWeight(stones []int) int {
	mh := &maxHeap{}
	heap.Init(mh)
	for i:=0; i<len(stones); i++ {
		heap.Push(mh, stones[i])
	}

	for (*mh).Len() > 1 {
		max1 := heap.Pop(mh)
		max2 := heap.Pop(mh)

		newStone := max1.(int) - max2.(int)
		if newStone != 0 {
			heap.Push(mh, newStone)
		}
	}

	if (*mh).Len() > 0 {
		return (*mh)[0]
	}

	return 0
}