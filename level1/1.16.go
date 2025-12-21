package main

import "fmt"

func main() {
	s := []int{9, 4, 6, 3, 1, 8, 2, 5, 7, 0}
	fmt.Println(s)
	quickSort(s)
	fmt.Println(s)
}

func quickSort(s []int) {
	if len(s) < 2 {
		return
	}

	pivot := s[len(s)-1] // последний элемент как опорный
	i := 0               // индекс для разделения элементов

	// все что меньше pivot кидаем в левую часть
	for j := 0; j < len(s)-1; j++ {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}

	// Ставим опорный элемент сразу после меньших, а большие после опорного
	s[i], s[len(s)-1] = s[len(s)-1], s[i]

	// Рекурсивно сортируем левую и правые части относительно i
	quickSort(s[:i])
	quickSort(s[i+1:])
}
