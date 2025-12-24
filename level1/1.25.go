package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Запускаем таймер и блокируем горутину!")
	Sleep(3 * time.Second)
	fmt.Println("Конец программы!")

}

func Sleep(d time.Duration) {
	timer := time.NewTimer(d)
	ticker := time.NewTicker(1 * time.Second) // Один тикер на весь цикл
	defer ticker.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("Горутина разблокирована!")
			return
		case <-ticker.C:
			fmt.Println("Таймер еще тикает...")
		}
	}
}
