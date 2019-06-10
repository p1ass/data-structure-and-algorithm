package data_structures

import "fmt"

type priorityQueue struct {
	data []int
}

func (pq *priorityQueue) Less(i, j int) bool {
	//ここに優先順位を書く
	// iの方が先に並んでほしかったらtrueを返す

	if pq.data[i] < pq.data[j] {
		return true
	}
	return false
}

func (pq *priorityQueue) Push(val int) {
	pq.data = append(pq.data, val)
	pq.up(len(pq.data) - 1)
}

func (pq *priorityQueue) Pop() (int, error) {
	if len(pq.data) == 0 {
		return 0, fmt.Errorf("emptry priority queue")
	}

	val := pq.data[0]

	pq.data[0], pq.data[len(pq.data)-1] = pq.data[len(pq.data)-1], pq.data[0]
	pq.data = pq.data[:len(pq.data)-1]
	pq.down(0)

	return val, nil
}

func (pq *priorityQueue) up(idx int) {
	if idx == 0 {
		return
	}

	parent := (idx - (2 - idx%2)) / 2

	if pq.Less(idx, parent) {
		pq.data[idx], pq.data[parent] = pq.data[parent], pq.data[idx]
		pq.up(parent)
	}
}

func (pq *priorityQueue) down(idx int) {
	left := 2*idx + 1
	right := 2*idx + 2

	if left >= len(pq.data) {
		return
	}

	if right >= len(pq.data) {
		if pq.data[idx] < pq.data[left] {
			pq.data[idx], pq.data[left] = pq.data[left], pq.data[idx]
		}
		return
	}

	biggerChild := left
	if pq.data[left] < pq.data[right] {
		biggerChild = right
	}

	if !pq.Less(idx, biggerChild) {
		pq.data[idx], pq.data[biggerChild] = pq.data[biggerChild], pq.data[idx]
		pq.down(biggerChild)
	}
}
