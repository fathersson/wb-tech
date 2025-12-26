package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

// правильный вариант:
// func test() error { // Возвращаем интерфейс, а не указатель на структуру
// 	return nil // Возвращает "чистый" nil (nil тип, nil значение)
// }

// func main() {
// 	var err error
// 	err = test()
// 	if err != nil {
// 		println("error")
// 		return
// 	}
// 	println("ok")
// }

// Программа выведет: error

// Объяснение: Интерфейс (в данном случае error) - это структура, состоящая из двух
// скрытых полей: Type и Value, интерфейс считается равным nil только тогда,
// когда оба этих поля равны nil

// В функции main переменная err имеет тип интерфейса error, когда выполняется присваивание
// err = test(), происходит следующее: Функция test возвращает конкретный тип *customError
// со значением nil, интерфейс err упаковывает этот результат:
// Поле Type заполняется типом *customError
// Поле Value заполняется значением nil
// Поскольку поле Type теперь не пустое, проверка if err != nil возвращает true

// Как исправить код:
// функции всегда должны возвращать интерфейс error напрямую
// В этом случае, при возврате nil, оба поля интерфейса (тип и значение) будут пустыми
