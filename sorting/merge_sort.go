package sorting

func mergeSort(ary []int) []int {
	if len(ary) <= 1 {
		return ary
	}

	if len(ary) == 2 {
		if ary[0] < ary[1] {
			return ary
		} else {
			return []int{ary[1], ary[0]}
		}
	}

	n := len(ary)

	sub1 := mergeSort(ary[:n/2])
	sub2 := mergeSort(ary[n/2:])

	i := 0
	j := 0
	sorted := []int{}
	for k := 0; k < n; k++ {
		if i < len(sub1) && j < len(sub2) {
			if sub1[i] < sub2[j] {
				sorted = append(sorted, sub1[i])
				i++
			} else {
				sorted = append(sorted, sub2[j])
				j++
			}
			continue
		}

		if i >= len(sub1) {
			sorted = append(sorted, sub2[j])
			j++
		} else {
			sorted = append(sorted, sub1[i])
			i++
		}
	}
	return sorted
}
