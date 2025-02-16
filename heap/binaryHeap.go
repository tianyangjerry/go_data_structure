package heap

import "errors"

type BinaryHeap []int

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{}
}

func (h *BinaryHeap) len() int {
	return len(*h)
}

func (h *BinaryHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeap) leftChild(i int) int {
	return i*2 + 1
}

func (h *BinaryHeap) rightChild(i int) int {
	return i*2 + 2
}

func (h *BinaryHeap) hasLeftChild(i int) bool {
	return h.leftChild(i) < h.len()
}

func (h *BinaryHeap) hasRightChild(i int) bool {
	return h.rightChild(i) < h.len()
}

func (h *BinaryHeap) up(i int) {
	for {
		if h.parent(i) < 0 {
			break
		}
		if (*h)[i] < (*h)[h.parent(i)] {
			(*h)[i], (*h)[h.parent(i)] = (*h)[h.parent(i)], (*h)[i]
			i = h.parent(i)
		} else {
			break
		}
	}
}

func (h *BinaryHeap) down(i int) {
	for {
		if !h.hasLeftChild(i) {
			break
		}
		minChild := h.leftChild(i)
		if h.hasRightChild(i) && (*h)[h.rightChild(i)] < (*h)[h.leftChild(i)] {
			minChild = h.rightChild(i)
		}
		if (*h)[i] > (*h)[minChild] {
			(*h)[i], (*h)[minChild] = (*h)[minChild], (*h)[i]
			i = minChild
		} else {
			break
		}
	}
}

func (h *BinaryHeap) Insert(value int) {
	*h = append(*h, value)
	h.up(h.len() - 1)
}

func (h *BinaryHeap) GetMIN() int {
	return (*h)[0]
}

func (h *BinaryHeap) DeleteMin() (int, error) {
	if h.len() == 0 {
		return 0, errors.New("heap is empty")
	}
	m := (*h)[0]
	(*h)[0] = (*h)[h.len()-1]
	*h = (*h)[:h.len()-1]
	h.down(0)
	return m, nil
}
