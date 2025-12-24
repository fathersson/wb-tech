package main

import (
	"strings"
)

// func main() {
// 	fmt.Println(check("abcd"))      // true
// 	fmt.Println(check("abCdefAaf")) // false
// 	fmt.Println(check("aabcd"))     // false
// }

func check(s string) bool {
	// Приводим к нижнему регистру
	lowerStr := strings.ToLower(s)

	// Используем map для отслеживания встреченных символов
	// struct{} не занимает места в памяти, в отличие от int или bool
	m := make(map[rune]struct{})

	for _, char := range lowerStr {
		// Проверяем, есть ли уже такой символ в карте
		// if значение, существует := карта[ключ]; существует { ... }
		if _, ok := m[char]; ok {
			return false // Нашли повтор - сразу выходим
		}
		// Если нет - добавляем в карту
		m[char] = struct{}{}
	}

	return true // Если дошли до конца и не нашли повторов
}
