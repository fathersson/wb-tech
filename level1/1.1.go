// Встраивание структур

// Дана структура Human (с произвольным набором полей и методов).

// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

// Подсказка: используйте композицию (embedded struct), чтобы Action
// имел все методы Human.

package main

import "fmt"

// Дочерняя структура
type Human struct {
	name string
	age  int
}

// 1 Метод работающий с обоими структурами
func (h Human) SayMyName() {
	fmt.Println("Привет, меня зовут:", h.name)
}

// 2 Метод работающий с обоими структурами
func (h Human) SayMyAge() {
	fmt.Println("Мне", h.age, "лет")
}

// Родительская структура
type Action struct {
	Human
}

// func main() {
// 	a := Action{Human{name: "Роман", age: 20}}
// 	a.SayMyName()
// 	a.SayMyAge()
// }
