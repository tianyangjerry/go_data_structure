package binaryTree

import (
	"math/rand"
	"testing"
	"time"
)

func TestBST(t *testing.T) {
	testNumbers := make([]int, 100)
	cnt := make(map[int]int)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	minExpVal := int(1e10)
	maxExpVal := int(-1e10)
	for i := 0; i < 100; i++ {
		testNumbers[i] = rand.Intn(1000)
		cnt[testNumbers[i]] += 1
	}

	bst := NewBST()
	for _, num := range testNumbers {
		bst.Insert(num)
		minExpVal = min(minExpVal, num)
		maxExpVal = max(maxExpVal, num)
	}

	// 验证所有数是否存在
	for _, num := range testNumbers {
		numCnt := bst.Search(num)
		if numCnt == 0 {
			t.Errorf("Search error: Expected value to be true, but got false")
		}
	}

	// 验证存在次数是否正确
	for _, num := range testNumbers {
		numCnt := bst.Search(num)
		if numCnt != cnt[num] {
			t.Errorf("Search error: Expected value to be %d, but got %d", cnt[num], numCnt)
		}
	}

	// 验证最大值和最小值是否正确
	minVal, _ := bst.FindMin()
	maxVal, _ := bst.FindMax()
	if minVal != minExpVal || maxVal != maxExpVal {
		t.Errorf("FindMin error: Expected min value to be %d, but got %d", minExpVal, minVal)
		t.Errorf("FindMax error: Expected max value to be %d, but got %d", maxExpVal, maxVal)
	}
}
