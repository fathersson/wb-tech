package main

import (
	"strings"
)

// func main() {
// 	s := "ехал грека через реку, видит грека в реке рак"
// 	fmt.Println(reverseWords(s))
// }

func reverseWords(s string) string {
	// Builder позволяет использовать буфер и не создавать новые строки для каждого слова
	var res strings.Builder
	// Сразу задаем размер буфера, равный исходной строке,
	// Чтобы была только одна аллокация памяти
	res.Grow(len(s))

	// Создаем слайс слов. Сами слова при этом не копируются,
	// слайс просто хранит указатели на части исходной строки s
	words := strings.Split(s, " ")

	// Идем в обратном порядке по слайсу слов
	for i := len(words) - 1; i >= 0; i-- {
		res.WriteString(words[i])
		// Добавляем пробел между словами, но не после последнего
		if i > 0 {
			res.WriteByte(' ') // Быстрее, чем WriteString(" ")
		}
	}

	// Возращаем конечную строку из буфера с помощью String()
	return res.String()
}
