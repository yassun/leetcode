import (
	"container/heap"
)

type KthLargest struct {
	k    int
	heap Heap
}

func (k *KthLargest) Add(val int) int {
	heap.Push(&k.heap, val)
	for k.k < len(k.heap) {
		heap.Pop(&k.heap)
	}
	return k.heap[0]
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	h := Heap(nums)
	heap.Init(&h)
	return KthLargest{
		k:    k,
		heap: h,
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
