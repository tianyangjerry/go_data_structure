package heap

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type BinaryHeap[T constraints.Ordered] []T

func NewBinaryHeap[T constraints.Ordered]() *BinaryHeap[T] {
	return &BinaryHeap[T]{}
}

func (h *BinaryHeap[T]) len() int {
	return len(*h)
}

func (h *BinaryHeap[T]) parent(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeap[T]) leftChild(i int) int {
	return i*2 + 1
}

func (h *BinaryHeap[T]) rightChild(i int) int {
	return i*2 + 2
}

func (h *BinaryHeap[T]) hasLeftChild(i int) bool {
	return h.leftChild(i) < h.len()
}

func (h *BinaryHeap[T]) hasRightChild(i int) bool {
	return h.rightChild(i) < h.len()
}

func (h *BinaryHeap[T]) up(i int) {
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

func (h *BinaryHeap[T]) down(i int) {
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

func (h *BinaryHeap[T]) Insert(value T) {
	*h = append(*h, value)
	h.up(h.len() - 1)
}

func (h *BinaryHeap[T]) GetMIN() T {
	return (*h)[0]
}

func (h *BinaryHeap[T]) DeleteMin() (T, error) {
	if h.len() == 0 {
		return T(0), errors.New("heap is empty")
	}
	m := (*h)[0]
	(*h)[0] = (*h)[h.len()-1]
	*h = (*h)[:h.len()-1]
	h.down(0)
	return m, nil
}
