type MinStack struct {
	gmin int
	storage []int
	min []int
}

func Constructor() MinStack {
	return MinStack{
		gmin: math.MaxInt32,
		storage: make([]int, 0),
		min: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.storage = append(this.storage, val)

	this.gmin = min(this.gmin, val)

	this.min = append(this.min, this.gmin)
}

func (this *MinStack) Pop() {
	if len(this.storage) == 0 {
		panic("Pop should be called on non-empty stack")
	}

	this.storage = this.storage[:len(this.storage)-1]

	this.min = this.min[:len(this.min)-1]

	if len(this.min) == 0 {
		this.gmin = math.MaxInt32
	} else {
		this.gmin = this.min[len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.storage) == 0 {
		panic("Top should be called on non-empty stack")
	}

	return this.storage[len(this.storage)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		panic("GetMin should be called on non-empty stack")
	}

	return this.min[len(this.min)-1]
}
