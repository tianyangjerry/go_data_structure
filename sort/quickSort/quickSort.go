package quickSort

import (
	"golang.org/x/exp/constraints"
)

type quickArr[T constraints.Ordered] []T

func QuickSort[T constraints.Ordered](arr []T) {
	a := quickArr[T](arr)
	a.fastSort(0, len(arr)-1)
}

func (arr quickArr[T]) fastSort(l, r int) {
	if l < r {
		p := arr.partition(l, r)
		arr.fastSort(l, p-1)
		arr.fastSort(p+1, r)
	}
}

func (arr quickArr[T]) partition(l, r int) int {
	pivot := arr[r]
	i := l - 1
	for j := l; j < r; j++ {
		if arr[j] < pivot {
			i++
			arr.swap(i, j)
		}
	}

	i++
	arr.swap(i, r)
	return i
}

func (arr quickArr[T]) swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
