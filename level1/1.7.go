package main

import (
	"fmt"
	"sync"
)

// Mutex защищает map от одновременной записи из нескольких горутин.
// WaitGroup используется, чтобы дождаться завершения всех горутин.
// и чтобы программа не завершилась, пока они не закончат свою работу
// суть работы программу - конкурентно с помощью двух горутин
// записать в цикле данные в мапу и завершить программу
func main() {
	m := make(map[int]int, 100)

	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			m[i] = 1
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			m[i] = 2
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println(m)
}
