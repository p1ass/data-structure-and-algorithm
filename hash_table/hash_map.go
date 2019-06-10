package hash_table

import "fmt"

type node struct {
	key, val int
	next     *node
}

type hashMap struct {
	N     int
	table []*node
}

func NewHashTable(N int) *hashMap {
	data := make([]*node, N)
	return &hashMap{N, data}
}

func (h *hashMap) Insert(key, val int) {
	hash := h.hash(key)

	n := h.table[hash]
	for n != nil {
		if n.key == key {
			n.val = val
			fmt.Println(key, val)
			return
		}
		n = n.next
	}

	h.table[hash] = &node{key, val, h.table[hash]}
}

func (h *hashMap) Erase(key int) {
	hash := h.hash(key)

	n := h.table[hash]
	if n == nil {
		return
	}
	if n.key == key {
		h.table[hash] = n.next
		return
	}

	for n.next != nil {
		if n.next.key == key {
			n.next = n.next.next
			h.table[hash] = n
			break
		}

		n = n.next
	}
}

func (h *hashMap) Get(key int) (int, bool) {
	hash := h.hash(key)

	n := h.table[hash]
	for n != nil {
		if n.key == key {
			return n.val, true
		}
		n = n.next
	}
	return 0, false
}

func (h *hashMap) hash(val int) int {
	return val % h.N
}
