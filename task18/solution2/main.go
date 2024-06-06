package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

// MutexCounter представляет счетчик, который может быть
// безопасно инкрементирован в конкурентной среде с использованием мьютекса.
type MutexCounter struct {
	// mu - мьютекс для защиты доступа к счетчику.
	mu sync.Mutex
	// value - внутреннее поле, хранящее значение счетчика.
	value int
}

// Increment увеличивает значение счетчика на единицу.
// Использует мьютекс для обеспечения взаимного исключения.
func (c *MutexCounter) Increment() {
	// Блокируем мьютекс перед изменением значения счетчика.
	c.mu.Lock()
	// Увеличиваем значение счетчика на 1.
	c.value++
	// Разблокируем мьютекс после изменения значения счетчика.
	c.mu.Unlock()
}

// Value возвращает текущее значение счетчика.
// Использует мьютекс для безопасного чтения значения.
func (c *MutexCounter) Value() int {
	// Блокируем мьютекс перед чтением значения счетчика.
	c.mu.Lock()
	// Используем defer для автоматической разблокировки мьютекса
	// после завершения функции.
	defer c.mu.Unlock()
	// Возвращаем текущее значение счетчика.
	return c.value
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Создаем новый экземпляр MutexCounter.
	counter := &MutexCounter{}

	// Используем WaitGroup для ожидания завершения всех горутин.
	var wg sync.WaitGroup

	// Количество горутин, которые будут инкрементировать счетчик.
	numGoroutines := 100
	// Количество инкрементов в каждой горутине.
	numIncrements := 1000

	// Добавляем numGoroutines в WaitGroup.
	wg.Add(numGoroutines)

	// Запускаем numGoroutines горутин, каждая из которых инкрементирует счетчик numIncrements раз.
	for range numGoroutines {
		go func() {
			defer wg.Done()

			for range numIncrements {
				counter.Increment()
			}
		}()
	}

	// Ожидаем завершения всех горутин.
	wg.Wait()

	// Выводим итоговое значение счетчика.
	// Ожидаемое значение: numGoroutines * numIncrements.
	log.Infof("Final Counter Value: %d\n", counter.Value())
}
