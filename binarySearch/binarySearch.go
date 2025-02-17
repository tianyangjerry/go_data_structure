package binarySearch

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](array []T, target T) int {
	begin := 0
	end := len(array) - 1

	for begin <= end {
		mid := (begin + end) >> 1
		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}

	return -1
}
