package quickSort

import (
	"go_data_structure/tools/array"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const testTime = 100

func TestQuickSortMultiple(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < testTime; i++ {
		// 生成100个随机数
		targetArr := make([]int, 100)
		arr := make([]int, 100)
		for j := 0; j < len(targetArr); j++ {
			targetArr[j] = rand.Intn(1000) // 生成0到999之间的随机整数
			arr[j] = targetArr[j]
		}

		// 使用 Go 的 sort 包进行排序
		sort.Ints(targetArr)

		// 使用 quickSort 进行排序
		QuickSort(arr)

		// 比对结果
		if array.CompareArrIsSame(arr, targetArr) {
			t.Log("排序成功")
		} else {
			t.Error("排序失败")
			println(arr)
		}
	}
}
