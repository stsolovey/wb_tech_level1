package main

import (
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

// AtomicCounter - структура для счетчика, который поддерживает конкурентный доступ.
type AtomicCounter struct {
	value int64 // Используем int64, чтобы позволить использование атомарных операций
}

// Increment - атомарно увеличивает значение счетчика на 1.
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1) // Атомарно добавляем 1 к value
}

// Value - возвращает текущее значение счетчика.
func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value) // Атомарно загружаем текущее значение value
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	counter := AtomicCounter{} // Создаем экземпляр счетчика

	// Количество горутин и инкрементов
	numGoroutines := 100
	numIncrements := 1000

	// Используем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	wg.Add(numGoroutines)

	// Запускаем 100 горутин, каждая из которых инкрементирует счетчик 1000 раз
	for range numGoroutines {
		go func() {
			for range numIncrements {
				counter.Increment()
			}

			wg.Done()
		}()
	}

	wg.Wait() // Ожидаем завершения всех горутин
	log.Infof("Final Counter Value: %d\n", counter.Value())
}
