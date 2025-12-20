package main

import (
	"fmt"
	"reflect"
)

func main() {
	// m := make(chan (string))
	test(nil)
}

func test(v any) {
	switch x := v.(type) {
	case int:
		fmt.Printf("Это целое число %d\n", x)
	case string:
		fmt.Printf("Это строка %s\n", x)
	case bool:
		fmt.Printf("Это булево значение %t\n", x)
	default:
		if x == nil {
			fmt.Println("Значение равно nil")
			return
		}
		if reflect.TypeOf(x).Kind() == reflect.Chan {
			fmt.Printf("Это канал %T\n", x)
		} else {
			fmt.Printf("Неизвестный тип %T\n", x)
		}
	}
}
