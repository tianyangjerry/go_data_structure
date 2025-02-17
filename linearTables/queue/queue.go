package queue

type Queue struct {
	data []int
	size int
	head int
	tail int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0, 10),
		size: 0,
		head: 0,
		tail: 0,
	}
}

func (q *Queue) Enqueue(v int) {
	q.data = append(q.data, v)
	q.size++
}

func (q *Queue) Dequeue() int {
	if q.size == 0 {
		return -1
	}
	res := q.data[q.head]
	q.head++
	q.size--
	return res
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) toString() string {
	QS := "["
	for i := 0; i < q.size; i++ {
		QS += string(rune(q.data[i]))
		if i != q.size-1 {
			QS += ", "
		}
	}
	QS += "]"
	return QS
}
