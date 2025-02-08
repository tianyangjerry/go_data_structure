package mergeSort

type Data struct {
	Value []int
}

func (d *Data) NewData(arr []int) *Data {
	d.Value = arr
	return d
}

func (d *Data) Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// MergeSort l: 左边界 r: 右边界
// 注意：左边界是闭区间，右边界是开区间 Eg: [0, len(arr))
func (d *Data) MergeSort(l, r int) {
	if r-l <= 1 {
		return
	}

	mid := (l + r) >> 1
	d.MergeSort(l, mid)
	d.MergeSort(mid, r)

	merged := d.Merge(d.Value[l:mid], d.Value[mid:r])
	copy(d.Value[l:r], merged)
}
