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
