package data_structures

import "fmt"

type queue struct {
	front *node
	back  *node
}

type node struct {
	val  int
	next *node
	prev *node
}

func (q *queue) Enqueue(val int) {
	if q.front == nil {
		n := &node{val, nil, nil}
		q.front = n
		q.back = n
		return
	}

	n := &node{val, nil, q.back}
	q.back.next = n
	q.back = n

}

func (q *queue) Dequeue() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("Empty queue")
	}

	n := q.front
	q.front = n.next
	if q.front != nil {
		q.front.prev = nil
	}

	return n.val, nil

}
