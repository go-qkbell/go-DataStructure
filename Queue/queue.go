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
		q.val = append(q.val, n)
		q.front = 0
		q.rear = 0

		return nil
	}

	q.val = append(q.val, n)
	q.rear++

	return nil

}

func (q *Queue) DeleteFront() (int, error) {
	if q.front == -1 && q.rear == -1 {
		return 0, errors.New("queue is empty")
	}

	if q.front == q.rear {
		temp := q.val[0]
		q.val = q.val[:0]
		q.front, q.rear = -1, -1

		return temp, nil
	}

	temp := q.val[0]
	q.val = q.val[1:]
	q.front++

	return temp, nil
}
