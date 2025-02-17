package queue

import "testing"

func TestQueue(t *testing.T) {
	Q := NewQueue()

	for i := 0; i < 5; i++ {
		Q.Enqueue(i)
	}

	println(Q.toString())

	for i := 0; i < 5; i++ {
		print(Q.Dequeue(), " ")
	}
}
