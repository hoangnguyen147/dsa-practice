package queue

type Queue struct {
	items []int
	head int
	tail int
	size int
	capacity int
}

func NewQueue(k int) Queue {
	return Queue{
		items: make([]int, k),
		head: -1,
		tail: -1,
		size: 0,
		capacity: k,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}

func (q *Queue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}

	if q.IsEmpty() {
		q.head = 0
		q.tail = 0
	} else {
		q.tail = (q.tail + 1) % q.capacity
	}

	q.items[q.tail] = value
	q.size++
	return true
}

func (q *Queue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	if q.head == q.tail {
		q.head = -1
		q.tail = -1
	}

	q.head = (q.head + 1) % q.capacity
	q.size--
	return true	
}

func (q *Queue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.items[q.head]
}

func (q *Queue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.items[q.tail]
}
