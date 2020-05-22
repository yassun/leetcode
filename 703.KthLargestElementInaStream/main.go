type KthLargest struct {
	k    int
	heap Heap
}

func (k *KthLargest) Add(val int) int {
	k.heap.Push(val)

	t := []int{}
	for i := 0; i < k.k; i++ {
		t = append(t, k.heap.Pop())
	}

	for _, v := range t {
		k.heap.Push(v)
	}

	return t[k.k-1]
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h *Heap) Swap(i, j int) {
	temp := (*h)[i]
	(*h)[i] = (*h)[j]
	(*h)[j] = temp
}

func (h Heap) IsBigger(i, j int) bool {
	return h[i] > h[j]
}

func (h *Heap) Push(val int) {
	*h = append(*h, val)
	h.UpHeap()
}

func (h *Heap) Pop() int {
	val := (*h)[0]
	*h = (*h)[1:]
	h.UpHeap()
	return val
}

func (h *Heap) UpHeap() {
	for i := 0; i < h.Len(); i++ {
		for c := i; ; {
			if c == 0 {
				break
			}
			p := (c - 1) / 2
			if h.IsBigger(c, p) {
				h.Swap(c, p)
				c = p
			} else {
				break
			}
		}
	}
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		k:    k,
		heap: Heap(nums),
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

