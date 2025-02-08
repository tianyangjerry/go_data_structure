package mergeSort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMergeSortMultiple(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		// 生成100个随机数
		arr := make([]int, 100)
		for j := range arr {
			arr[j] = rand.Intn(1000) // 生成0到999之间的随机整数
		}

		// 使用 Go 的 sort 包进行排序
		expected := make([]int, len(arr))
		copy(expected, arr)
		sort.Ints(expected)

		// 使用 mergeSort 进行排序
		d := Data{}
		data := d.NewData(arr)
		data.MergeSort(0, len(arr))

		// 比对结果
		for k, v := range arr {
			if v != expected[k] {
				t.Errorf("Test %d, Expected %d, got %d", i, expected[k], v)
			}
		}
	}
}
