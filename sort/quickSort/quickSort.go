package quickSort

type Data struct {
	value []int
}

func (d *Data) NewData(arr []int) *Data {
	d.value = arr
	return d
}

func (d *Data) swap(i, j int) {
	d.value[i], d.value[j] = d.value[j], d.value[i]
}

func (d *Data) partition(l, r int) int {
	pivot := d.value[r]
	for l < r {
		for l < r && d.value[l] <= pivot {
			l++
		}
		d.swap(l, r)
		for l < r && d.value[r] >= pivot {
			r--
		}
		d.swap(l, r)
	}
	return l
}

// QuickSort l: 左边界 r: 右边界
// 注意：左右边界都是闭区间 eg: [0, len(arr)-1]
func (d *Data) QuickSort(l, r int) {
	if l < r {
		p := d.partition(l, r)
		d.QuickSort(l, p-1)
		d.QuickSort(p+1, r)
	}
}
