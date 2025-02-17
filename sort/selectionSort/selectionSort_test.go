package selectionSort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const testTime = 100

func TestSelectionSort(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < testTime; i++ {
		arr := make([]int, 100)
		targetArr := make([]int, 100)
		for j := 0; j < 100; j++ {
			arr[j] = rand.Intn(100)
			targetArr[j] = arr[j]
		}
		sort.Ints(targetArr)

		SelectionSort(arr)

		for index, element := range arr {
			if element != targetArr[index] {
				t.Errorf("expected %d, got %d", targetArr[index], element)
			}
		}
	}
}
