package heap

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestBinaryHeap(t *testing.T) {
	rand.New(rand.NewSource(time.Now().Unix()))
	m := int(1e10)

	bh := NewBinaryHeap()
	targetedArr := make([]int, 10)
	for i := 0; i < len(targetedArr); i++ {
		randNum := rand.Intn(100)
		targetedArr[i] = randNum
		bh.Insert(randNum)
		m = min(m, randNum)
	}

	sort.Ints(targetedArr)

	// 验证最小值
	if bh.GetMIN() != m {
		t.Error("GetMIN() error")
	}

	for i := 0; i < len(targetedArr); i++ {
		if bh.GetMIN() != targetedArr[i] {
			t.Error("GetMIN() error")
		}
		_, err := bh.DeleteMin()
		if err != nil {
			t.Error("Heap is empty!!!")
			return
		}
	}

	// 验证是否为空
}
