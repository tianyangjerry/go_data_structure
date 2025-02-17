package selectionSort

import (
	"golang.org/x/exp/constraints"
)

func SelectionSort[T constraints.Ordered](arr []T) {
	lenArr := len(arr)
	for i := 0; i < lenArr-1; i++ {
		for j := 0; j < lenArr-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

func swap[T constraints.Ordered](arr []T, index1, index2 int) {
	arr[index1], arr[index2] = arr[index2], arr[index1]
}
