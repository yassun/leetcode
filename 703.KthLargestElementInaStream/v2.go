type KthLargest struct {
	k    int
	heap Heap
}

func (k *KthLargest) Add(val int) int {
	k.heap.Push(val)
	for k.k < len(k.heap) {
		k.heap.Pop()
	}
	return k.heap[0]
}

type Heap []int

func (h *Heap) UpHeap() {
	for i := 0; i < len(*h); i++ {
		for child := i; ; {
			if child == 0 {
				break
			}
			parent := (child - 1) / 2
			if (*h)[parent] > (*h)[child] {
				(*h)[parent], (*h)[child] = (*h)[child], (*h)[parent]
				child = parent
			} else {
				break
			}
		}
	}
}

func (h Heap) DownHeap(startIndex, maxIndex int) {
	parent := startIndex
	for {
		leftChild := 2*parent + 1
		if maxIndex < leftChild {
			break
		}

		child := leftChild
		if rightChild := leftChild + 1; rightChild <= maxIndex && h[rightChild] < h[leftChild] {
			child = rightChild
		}

		if h[parent] < h[child] {
			break
		}

		h[parent], h[child] = h[child], h[parent]
		parent = child
	}
}

func (h *Heap) Push(val int) {
	*h = append(*h, val)
	h.UpHeap()
}

func (h *Heap) Pop() int {
	val := (*h)[0]
	*h = (*h)[1:len(*h)]
	h.DownHeap(0, len(*h)-1)
	return val
}

func Constructor(k int, nums []int) KthLargest {
	h := Heap(nums)
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		h.DownHeap(i, n-1)
	}

	return KthLargest{
		k:    k,
		heap: h,
	}
}
