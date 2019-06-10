package data_structures

import "fmt"

type stack struct {
	data []int
}

func (s *stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("empty stack")
	}

	n := len(s.data)
	val := s.data[n-1]
	s.data = s.data[:n-1]
	return val, nil
}
