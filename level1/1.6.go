package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// 1) завершение горутины по условию
// горутина проверяет условие и делает return
func condition() {
	go func() {
		for i := 0; i < 5; i++ {
			if i == 2 {
				fmt.Println("Горутина завершается")
				return
			}
		}
	}()
}

// 2) завершение горутины по сигналу из канала
// выполняем полезную работу и ждем, пока в канал поступит сигнал, после чего горутина закрывается
func chanSignal() {
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Горутина завершается")
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("Работаю...")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	quit <- struct{}{}
}

// 3) Горутина завершается, когда канал закрыт и чтение возвращает ok = false
// Горутина завершится автоматически, когда range обнаружит закрытие канала, закроется когда запись прекратится
func chanClose() {
	ch := make(chan int)

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	ch <- 1
	ch <- 2
	close(ch)
}

// 4) мгновенно завершает только текущую горутину, игнорируя return
func goexit() {
	go func() {
		fmt.Println("Работа")
		runtime.Goexit()
	}()
}

// 5) Остановка горутины через context.WithCancel
// Горутина завершится, когда cancel() вызовет завершение контекста и ctx.Done() станет читаемым
func contCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена context.Cancel()")
				return
			default:
				fmt.Println("Работа...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()

}

// 6) Остановка горутины через context.WithTimeout
// Горутина завершится, когда истечёт таймаут и ctx.Done() подаст сигнал о закрытии контекста
func contTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Таймаут, завершаем работу")
				return
			default:
				fmt.Println("Работаем...")
			}
		}
	}()
	time.Sleep(5 * time.Second)
}

// 7) завершаем работу горутины по истечению времени time.After
// Горутина завершится, когда сработает таймер из time.After
func after() {
	timer := time.After(5 * time.Second)
	go func() {
		for {
			select {
			case <-timer:
				fmt.Println("Завершаемся после time.After")
				return
			default:
				fmt.Println("Работа...")
			}
		}
	}()
}
