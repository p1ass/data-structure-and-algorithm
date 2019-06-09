package sorting

import "fmt"

// あるpivotを選び、左右から挟むように走査し、値をswapしてソートする。
// 平均 O(n log n) 最悪 O(n^2)
func quickSort(ary []int) []int {
	fmt.Println(ary)
	n := len(ary)

	if n <= 1 {
		return ary
	}

	p := ary[n/2]
	for {
		var i int
		var j int

		for i = 0; ary[i] < p; i++ {
		}

		for j = n - 1; ary[j] > p; j-- {
		}

		if i < j {
			ary[i], ary[j] = ary[j], ary[i]
			i++
			j--
		} else {
			if i >= 1 {
				quickSort(ary[:i-1])
			}
			if i <= n-1 {
				quickSort(ary[i:])
				break
			}
		}
	}
	return ary
}
