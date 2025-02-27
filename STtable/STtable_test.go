package STtable

import (
	"math/rand"
	"testing"
	"time"
)

func TestSTtable(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, 100)

	mx := int(-1e10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
		mx = max(mx, arr[i])
	}

}
