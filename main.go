package main

import "go_data_structure/sort/quickSort"

func main() {
	arr := [3]int{2, 1, 4}

	quickSort.QuickSort(arr[:])

	for _, v := range arr {
		println(v)
	}
}
