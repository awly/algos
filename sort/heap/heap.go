// Heapsort
// measures time to sort 1Mil random elements via max-heap
package heap

func Sort(data []int) {
	h := heap(data)
	h.buildMaxHeap()
	for i := 0; i < len(data); i++ {
		(&h).extractMax()
	}
}

type heap []int

func (h heap) buildMaxHeap() {
	for i := len(h) / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
}

func (h heap) maxHeapify(i int) {
	lefti, righti := i*2+1, i*2+2
	if lefti < len(h) && righti < len(h) {
		if h[lefti] > h[i] && h[lefti] > h[righti] {
			h[i], h[lefti] = h[lefti], h[i]
			h.maxHeapify(lefti)
		} else if h[righti] > h[i] {
			h[i], h[righti] = h[righti], h[i]
			h.maxHeapify(righti)
		}
	} else if lefti < len(h) && h[lefti] > h[i] {
		h[i], h[lefti] = h[lefti], h[i]
		h.maxHeapify(lefti)
	} else if righti < len(h) && h[righti] > h[i] {
		h[i], h[righti] = h[righti], h[i]
		h.maxHeapify(righti)
	}
}

func (h *heap) extractMax() int {
	res := (*h)[0]
	(*h)[0], (*h)[len((*h))-1] = (*h)[len((*h))-1], (*h)[0]
	(*h) = (*h)[:len((*h))-1]
	(*h).maxHeapify(0)
	return res
}
