package mergeSort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const testTime = 100

func TestMergeSortMultiple(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < testTime; i++ {
		// 生成100个随机数
		targetArr := make([]int, 100)
		arr := make([]int, 100)
		for j := range targetArr {
			targetArr[j] = rand.Intn(1000) // 生成0到999之间的随机整数
		}
		copy(arr, targetArr)

		// 使用 Go 的 sort 包进行排序
		sort.Ints(targetArr)

		// 使用 mergeSort 进行排序
		MergeSort(arr)

		// 比对结果
		for k, v := range targetArr {
			if v != arr[k] {
				t.Errorf("Test %d, Expected %d, got %d", i, arr[k], v)
			}
		}
	}
}
