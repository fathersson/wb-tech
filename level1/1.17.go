package main

import "fmt"

func main() {
	s := []int{1, 12, 23, 34, 56, 67, 78, 90, 120, 340, 450, 567}
	x := 567
	fmt.Println(binSearch(s, x))
}

func binSearch(s []int, x int) (index int) {
	low := 0           // начальный индекс
	high := len(s) - 1 // конечный индекс

	for low <= high {
		mediumInd := (low + high) / 2 // находим средний индекс
		mediumVal := s[mediumInd]     // находим значение среднего индекса

		if mediumVal == x {
			return mediumInd // нашли
		}

		if mediumVal < x {
			low = mediumInd + 1 // Искомое в правой половине
		} else {
			high = mediumInd - 1 // Искомое в левой половине
		}

	}

	return -1
}
