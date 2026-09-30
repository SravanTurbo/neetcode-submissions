type Maxheap []int

func (h Maxheap) Len() int { return len(h) }
func (h Maxheap) Less(i, j int) bool { return h[i] > h[j] }
func (h Maxheap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Maxheap) Push(i interface{}) { *h = append(*h, i.(int)) }
func (h *Maxheap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[:n-1]
    return x
}

type Minheap []int

func (h Minheap) Len() int { return len(h) }
func (h Minheap) Less(i, j int) bool { return h[i] < h[j] }
func (h Minheap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Minheap) Push(i interface{}) { *h = append(*h, i.(int)) }
func (h *Minheap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[:n-1]
    return x
}

type MedianFinder struct {
    small *Maxheap
    large *Minheap
}

func Constructor() MedianFinder {
    small := &Maxheap{}
    large := &Minheap{}
    heap.Init(small)
    heap.Init(large)
    return MedianFinder{small: small, large: large}
}

func (this *MedianFinder) AddNum(num int)  {
    if this.small.Len() == 0 {
        heap.Push(this.small, num)
        return
    }

    if num < (*this.small)[0] {
        heap.Push(this.small, num)
    } else {
        heap.Push(this.large, num)
    }

    if this.small.Len() > this.large.Len() + 1 {
        smax := heap.Pop(this.small)
        heap.Push(this.large, smax)
    }else if this.large.Len() > this.small.Len() + 1 {
        lmin := heap.Pop(this.large)
        heap.Push(this.small, lmin)
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if (this.small.Len() + this.large.Len()) %2 == 0 {
        return float64((*this.small)[0] + (*this.large)[0])/float64(2)
    }

    if this.small.Len() > this.large.Len() {
        return float64((*this.small)[0])
    } else {
        return float64((*this.large)[0])
    }
}
