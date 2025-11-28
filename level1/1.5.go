// Таймаут на канал

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения.
// По истечении N секунд программа должна завершаться.

// Подсказка: используйте time.After или таймер для ограничения времени работы.

package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// Создаем контекст WithTimeout, чтобы через n секунд завершить программу
// Создаем WaitGroup, чтобы перед выходом дождаться пока горутины завершат работу
// в отдельных горутинах запускаем писателя и читателя и ожидаем завершения горутин
// func main() {
// 	var n time.Duration
// 	n = 3

// 	ctx, cancel := context.WithTimeout(context.Background(), n*time.Second)
// 	defer cancel()

// 	ch := make(chan int, 1)

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go writer(&wg, ctx, ch)
// 	go reader(&wg, ctx, ch)

// 	wg.Wait()
// 	fmt.Println("Время вышло")
// }

// // writer отправляет данные в канал, пока контекст не будет отменён.
// // После таймаута writer завершится через <-ctx.Done().
// func writer(wg *sync.WaitGroup, ctx context.Context, ch chan int) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Время вышло, завершаем запись")
// 			return
// 		case ch <- 1:
// 		}
// 	}
// }

// // reader получает данные из канала, пока контекст не будет отменён,
// // после чего завершается через <-ctx.Done().
// func reader(wg *sync.WaitGroup, ctx context.Context, ch chan int) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Время вышло, завершаем чтение")
// 			return
// 		case v := <-ch:
// 			fmt.Println(v)
// 		}
// 	}
// }
