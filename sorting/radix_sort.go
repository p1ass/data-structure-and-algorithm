package sorting

import (
	"strconv"
)

func radixSort(ary []int) []int {
	if len(ary) <= 1 {
		return ary
	}

	buckets := make([][]int, 10)

	n := len(ary)

	maxVal := -105000000
	for i := 0; i < n; i++ {
		if maxVal < ary[i] {
			maxVal = ary[i]
		}
	}
	maxDigit := len([]rune(strconv.Itoa(maxVal)))

	for digit := 0; digit <= maxDigit; digit++ {
		for i := 0; i < n; i++ {
			runeVal := []rune(strconv.Itoa(ary[i]))
			var digitVal int
			if digit < len(runeVal) {
				digitVal, _ = strconv.Atoi(string(runeVal[len(runeVal)-digit-1]))
			} else {
				digitVal = 0
			}
			buckets[digitVal] = append(buckets[digitVal], ary[i])
		}

		ary = []int{}
		for i, bucket := range buckets {
			ary = append(ary, bucket...)
			buckets[i] = nil
		}
	}
	return ary
}
