package binarySearch

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const testTime = 100

func TestBinarySearch(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < testTime; i++ {
		arr := make([]int, 100)
		for j := 0; j < 100; j++ {
			arr[j] = rand.Intn(1000)
		}
		sort.Ints(arr)
		target := rand.Intn(1000)
		got := BinarySearch(arr, target)
		if got != -1 {
			got = arr[got]
		}
		ans := -1
		for _, element := range arr {
			if element == target {
				ans = element
			}
		}
		if ans != got {
			t.Errorf("BinarySearch(%v, %v) = %v, want %v", arr, target, got, ans)
		}
	}
}
