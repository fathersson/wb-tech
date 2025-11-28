package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// 1) завершение горутины по условию
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
func chanSignal() {
	quit := make(chan struct{})

	go func() {
		select {
		case <-quit:
			fmt.Println("Горутина завершается")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Работаю...")
		}
	}()

	time.Sleep(5 * time.Second)
	quit <- struct{}{}
}

// 3) Горутина завершается, когда канал закрыт и чтение возвращает ok = false
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

// Остановка горутины через context.WithTimeout
func contTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
}

// завершаем работу горутины по истечению времени time.After
func after() {
	go func() {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("Завершаемся после time.After")
			return
		default:
			fmt.Println("Работа...")
		}
	}()
}
