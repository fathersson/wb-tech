package main

// Создаем множество строк с помощью ключей map[string]struct{}
// После заполнения мапы, формируем слайс уникальных слов
// func main() {
// 	words := []string{"cat", "cat", "dog", "cat", "tree"}
// 	res := make([]string, 0, len(words))

// 	set := make(map[string]struct{})

// 	for _, value := range words {
// 		set[value] = struct{}{}
// 	}

// 	for value := range set {
// 		res = append(res, value)
// 	}

// 	fmt.Println(res)

// }
