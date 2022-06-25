package Queue

import "errors"

type Queue struct {
	val     []int
	front   int
	rear    int
	maxSize int
}

func NewQueue(n int) *Queue {
	q := Queue{
		val:     make([]int, 0, n),
		front:   -1,
		rear:    -1,
		maxSize: n,
	}

	return &q
}

func (q *Queue) Push(n int) error {
	if q.rear == q.maxSize-1 {
		return errors.New("queue is full")
	}

	if q.front == -1 && q.rear == -1 {
		q.val[0] = n
		q.front = 0
		q.rear = 0

		return nil
	}

	q.rear++
	q.val[q.rear] = n

	return nil

}

func (q *Queue) DeleteFront() (int, error) {
	if q.front == -1 && q.rear == -1 {
		return 0, errors.New("queue is empty")
	}

	if q.front == q.rear {
		temp := q.val[q.front]
		q.val = q.val[:0]

		return temp, nil
	}

	temp := q.val[q.front]
	q.front++
	q.val = q.val[q.front:]

	return temp, nil
}
