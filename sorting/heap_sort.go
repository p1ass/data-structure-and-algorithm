package sorting

func heapSort(ary []int) []int {
	n := len(ary)

	//まずヒープを作る
	for i := 1; i < n; i++ {
		up(ary, i)
	}

	//ソート
	for i := 0; i < n; i++ {
		ary[0], ary[n-i-1] = ary[n-i-1], ary[0]
		down(ary[:n-i-1], 0)
	}

	return ary
}

func up(ary []int, idx int) {
	if idx == 0 {
		return
	}

	parent := (idx - (2 - idx%2)) / 2

	if ary[idx] > ary[parent] {
		ary[idx], ary[parent] = ary[parent], ary[idx]
		up(ary, parent)
	}
}

func down(ary []int, idx int) {
	left := 2*idx + 1
	right := 2*idx + 2

	if left >= len(ary) {
		return
	}

	if right >= len(ary) {
		if ary[idx] < ary[left] {
			ary[idx], ary[left] = ary[left], ary[idx]
		}
		return
	}

	biggerChild := left
	if ary[left] < ary[right] {
		biggerChild = right
	}

	if ary[idx] < ary[biggerChild] {
		ary[idx], ary[biggerChild] = ary[biggerChild], ary[idx]
		down(ary, biggerChild)
	}
}
