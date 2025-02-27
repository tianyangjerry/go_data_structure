package mergeSort

import "golang.org/x/exp/constraints"

type mergeSortArr[T constraints.Ordered] []T

func MergeSort[T constraints.Ordered](arr []T) {
	a := mergeSortArr[T](arr)
	a.fastSort(a, 0, len(arr))
}

func (a mergeSortArr[T]) fastSort(arr []T, l, r int) {
	if r-l <= 1 {
		return
	}

	mid := (l + r) >> 1
	a.fastSort(arr, l, mid)
	a.fastSort(arr, mid, r)

	merged := a.merge(arr[l:mid], arr[mid:r])
	copy(arr[l:r], merged)
}

func (a mergeSortArr[T]) merge(leftArr, rightArr []T) []T {
	result := make([]T, 0, len(leftArr)+len(rightArr))

	i, j := 0, 0

	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			result = append(result, leftArr[i])
			i += 1
		} else {
			result = append(result, rightArr[j])
			j += 1
		}
	}

	result = append(result, leftArr[i:]...)
	result = append(result, rightArr[j:]...)
	return result
}
