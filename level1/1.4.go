package main

import (
	"context"
	"fmt"
	"sync"
)

// // Реализовано два способа передачи сигнала отмены, но основным выбран Context, т.к является
// // стандартом для языка и graceful shutdown, контекст также дает больше гибкости в будущем,
// // например таймауты, вложенные отмены, также с контекстом кода меньше и выглядит он лаконичнее и проще
// func main() {
// 	// Создаем Context, который ждет сигнал завершения  и принудительно завершаем в конце программы
// 	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
// 	defer stop()

// 	ch := make(chan int, 10)
// 	n := 5

// 	// WaitGroup, чтобы программа не завершлиась раньше того, как завершатся все горутины
// 	var wg sync.WaitGroup
// 	wg.Add(n)

// 	Workers(ch, n, ctx, &wg)

// 	// Бесконечно пишем в канал, пока не будет сигнала отмены (SIGINT),
// 	// после которого канал контекста Done() закрывается автоматически и запускается case <-ctx.Done()
// 	// далее мы закрываем канал в который писали, для сигнализации горутинам о завершении работы
// 	// и для предотвращения потенциальных проблем
// 	// wg.Wait() ожидает завершения горутин, мы выходим из цикла и программа завершается
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Завершаем запись в канал")
// 			close(ch)
// 			wg.Wait()
// 			return
// 		case ch <- 1:
// 		}
// 	}

// }

func Workers(ch chan int, count int, ctx context.Context, wg *sync.WaitGroup) {

	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()

			// бесконечно читаем данные из канал и ожидаем завершения контекста
			// после чего завершаем горутину
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Горутина завершается %v \n", i)
					return
				case v := <-ch:
					fmt.Println(v)
				}
			}
		}(i)
	}
}

// func main() {
// 	signalChan := make(chan os.Signal, 1)
// 	signal.Notify(signalChan, syscall.SIGINT)
// 	defer signal.Stop(signalChan)

// 	ch := make(chan int, 10)
// 	n := 5

// 	var wg sync.WaitGroup
// 	wg.Add(n)

// 	Workers(ch, n, signalChan, &wg)

// 	for {
// 		select {
// 		case <-signalChan:
// 			fmt.Println("Завершаем запись в канал")
// 			close(ch)
// 			wg.Wait()
// 			return
// 		case ch <- 1:
// 		}
// 	}

// }

// func Workers(ch chan int, count int, signalChan chan os.Signal, wg *sync.WaitGroup) {

// 	for i := 0; i < count; i++ {
// 		go func(i int) {
// 			defer wg.Done()

// 			for {
// 				v, ok := <-ch
// 				if !ok {
// 					fmt.Printf("Канал закрыт, воркер уходит %v \n", i)
// 					return
// 				}
// 				fmt.Println(v)
// 			}
// 		}(i)
// 	}
// }
