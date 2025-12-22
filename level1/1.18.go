package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// --- Решение с использованием ATOMIC ---
// Этот подход использует низкоуровневые инструкции процессора
// Он быстрее мьютексов, так как не блокирует выполнение горутин на уровне ОС

type CounterA struct {
	value int64 // Для atomic нужно использовать типы с фиксированным размером int32/int64
}

// Inc безопасно увеличивает значение на 1 на уровне CPU
func (c *CounterA) Inc() {
	atomic.AddInt64(&c.value, 1)
}

// Value безопасно читает значение, гарантируя актуальность данных из кэша процессора
func (c *CounterA) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func workAtomic(c *CounterA, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		c.Inc()
	}
}

// ------------------------------------------------------------------------------

// --- Решение с использованием MUTEX ---
// Универсаленьный подход, RWMutex позволяет разделять доступ
// на чтение (одновременно многим) и запись (только одному)

type CounterM struct {
	value int
	mu    sync.RWMutex // Используем RWMutex для оптимизации производительности при чтении
}

// Inc блокирует структуру для полной записи
func (c *CounterM) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Value позволяет множеству читателей обращаться к данным одновременно
func (c *CounterM) Value() int {
	c.mu.RLock() // Читающая блокировка (разрешает другие RLock, но запрещает Lock)
	defer c.mu.RUnlock()
	return c.value
}

func workMutex(c *CounterM, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		c.Inc()
	}
}

// ------------------------------------------------------------------------------

func main() {
	// 1. Демонстрация работы ATOMIC
	ca := CounterA{}
	var wgA sync.WaitGroup

	wgA.Add(3) // Инициализируем счетчик WaitGroup для ATOMIC
	go workAtomic(&ca, &wgA)
	go workAtomic(&ca, &wgA)
	go workAtomic(&ca, &wgA)

	wgA.Wait() // Главный поток спит, пока счетчик wgA не станет 0
	fmt.Printf("Atomic результат: %d\n", ca.Value())

	// 2. Демонстрация работы MUTEX
	cm := CounterM{}
	var wgM sync.WaitGroup

	wgM.Add(3) // Инициализируем WaitGroup для MUTEX
	go workMutex(&cm, &wgM)
	go workMutex(&cm, &wgM)
	go workMutex(&cm, &wgM)

	wgM.Wait() // Ожидаем завершения всех горутин
	fmt.Printf("Mutex результат:  %d\n", cm.Value())
}
