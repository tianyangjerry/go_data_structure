package array

func CompareArrIsSame[T comparable](arr1 []T, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
