package Stack

import "errors"

type Stack struct {
	val     []int
	len     int
	maxSize int
}

func NewStack(n int) *Stack {
	s := Stack{
		val:     make([]int, 0, n),
		len:     0,
		maxSize: n,
	}

	return &s
}

func (s *Stack) Push(n int) error {
	if s.len == s.maxSize {
		return errors.New("stack is full")
	}

	s.val = append(s.val, n)
	s.len++

	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.len == 0 {
		return -1, errors.New("stack is empty")
	}

	temp := s.val[s.len-1]
	s.val = s.val[:s.len-1]
	s.len--

	return temp, nil
}
