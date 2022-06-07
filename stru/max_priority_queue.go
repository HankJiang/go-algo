package stru

import "fmt"

// @see https://labuladong.github.io/algo/2/21/62/

// MaxPriorityQueue 优先级队列（最大堆）
type MaxPriorityQueue struct {
	// 存储元素的数组
	Values []int
	// 当前队列中元素个数
	Size int
}

// Parent 父节点
func (a *MaxPriorityQueue) Parent(x int) int {
	return x / 2
}

// Left 左子节点
func (a *MaxPriorityQueue) Left(x int) int {
	return x * 2
}

// Right 右子节点
func (a *MaxPriorityQueue) Right(x int) int {
	return x*2 + 1
}

// Max 当前队列中的最大元素
func (a *MaxPriorityQueue) Max() int {
	return a.Values[1]
}

// Insert 插入元素
func (a *MaxPriorityQueue) Insert(value int) {
	// 插入到堆底，然后上浮
	a.Size += 1
	a.Values[a.Size] = value
	a.Swim(a.Size)
}

// Delete 删除并返回当前队列中的最大元素
func (a *MaxPriorityQueue) Delete() int {
	value := a.Max()
	a.Swap(1, a.Size)
	// 和堆底交换，删除堆底，堆顶下沉
	a.Values[a.Size] = 0
	a.Size -= 1
	a.Sink(1)
	return value
}

// Swim 上浮第 x 个元素，以维护最大堆性质
func (a *MaxPriorityQueue) Swim(x int) {
	for {
		if x > 1 && a.Less(a.Parent(x), x) {
			a.Swap(a.Parent(x), x)
			a.Swim(a.Parent(x))
		} else {
			break
		}
	}
}

// Sink 下沉第 x 个元素，以维护最大堆性质
func (a *MaxPriorityQueue) Sink(x int) {
	for {
		if x <= a.Size {
			// 找到左右最大值
			max := a.Left(x)
			if a.Less(max, a.Right(x)) {
				max = a.Right(x)
			}
			// 和左右最大值交换
			if a.Less(x, max) {
				a.Swap(x, max)
			}
			x = max
		} else {
			break
		}
	}
}

// Swap 交换数组的两个元素
func (a *MaxPriorityQueue) Swap(i, j int) {
	tmp := a.Values[i]
	a.Values[i] = a.Values[j]
	a.Values[j] = tmp
}

// Less pq[i] 是否比 pq[j] 小？
func (a *MaxPriorityQueue) Less(i, j int) bool {
	return a.Values[i] < a.Values[j]
}

func BuildPQ(arr []int) *MaxPriorityQueue {
	pq := &MaxPriorityQueue{
		Values: make([]int, len(arr)*2+1),
		Size:   0,
	}
	for _, value := range arr {
		pq.Insert(value)
	}
	return pq
}

/***************** demo ******************/

func TestPQ() {
	var arr = []int{1, 3, 8, 6, 4, 2, 7}
	pq := BuildPQ(arr)
	fmt.Println(pq)
	for {
		if pq.Size > 0 {
			fmt.Println(pq.Delete())
		} else {
			break
		}
	}
}
