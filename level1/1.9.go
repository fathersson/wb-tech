package main

import (
	"fmt"
	"sync"
)

// Создаем два канала, слайс исходных чисел и WaitGroup,
// чтобы дождаться завершения всех горутин конвейера
//
// 1-я горутина пишет числа из слайса в ch1 и закрывает его,
// т.к новых значений не будет
//
// 2-я горутина читает числа из ch1, умножает на 2,
// отправляет в ch2 и закрывает канал, когда числа заканчиваются
//
// 3-я горутина читает готовые значения из ch2 и выводит в stdout
//
// Цикл range по каналу будет продолжать блокироваться и читать значения
// из канала до тех пор, пока канал не будет закрыт
func main() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	var wg sync.WaitGroup

	// x := []int{1,2,3,4,5}
	x := make([]int, 0, 5)
	for i := 0; i < cap(x); i++ {
		x = append(x, i+1)
	}

	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < len(x); i++ {
			ch1 <- x[i]
		}
		close(ch1)
	}()

	go func() {
		defer wg.Done()
		for v := range ch1 {
			ch2 <- v * 2
		}
		close(ch2)
	}()

	go func() {
		defer wg.Done()
		for v := range ch2 {
			fmt.Printf("%d ", v)
		}
	}()

	wg.Wait()
}
