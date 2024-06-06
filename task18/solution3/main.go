package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

// ChannelCounter представляет счетчик, который может быть
// безопасно инкрементирован в конкурентной среде с использованием каналов.
type ChannelCounter struct {
	value   int            // Внутреннее поле, хранящее значение счетчика.
	channel chan int       // Канал для передачи инкрементных запросов.
	done    chan struct{}  // Канал для завершения работы горутины.
	mu      sync.Mutex     // Мьютекс для защиты доступа к значению счетчика.
	wg      sync.WaitGroup // WaitGroup для отслеживания активных инкрементов.
}

// NewChannelCounter создает новый экземплят ChannelCounter и запускает горутину для обработки запросов.
func NewChannelCounter() *ChannelCounter {
	cc := &ChannelCounter{
		channel: make(chan int),
		done:    make(chan struct{}),
	}
	go cc.run()

	return cc
}

// run обрабатывает инкрементные запросы, поступающие в канал.
func (c *ChannelCounter) run() {
	for {
		select {
		case delta := <-c.channel:
			c.mu.Lock()
			c.value += delta
			c.mu.Unlock()
		case <-c.done:
			return
		}
	}
}

// Increment увеличивает значение счетчика на 1, отправляя запрос в канал и учитывая WaitGroup.
func (c *ChannelCounter) Increment() {
	c.wg.Add(1) // Увеличиваем счетчик активных инкрементов.

	go func() {
		c.channel <- 1
		c.wg.Done() // Уменьшаем счетчик после отправки инкремента.
	}()
}

// Close завершает работу горутины.
func (c *ChannelCounter) Close() {
	c.wg.Wait()   // Дожидаемся завершения всех инкрементов.
	close(c.done) // Закрываем канал done.
}

// Value возвращает текущее значение счетчика.
func (c *ChannelCounter) Value() int {
	c.mu.Lock() // Блокируем мьютекс для безопасного доступа к value.
	defer c.mu.Unlock()

	return c.value
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Создаем новый экземпляр ChannelCounter.
	counter := NewChannelCounter()

	// Используем WaitGroup для ожидания завершения всех горутин.
	var wg sync.WaitGroup

	// Количество горутин, которые будут инкрементировать счетчик.
	numGoroutines := 100
	// Количество инкрементов в каждой горутине.
	numIncrements := 1000

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

	// Завершаем работу горутины.
	counter.Close()

	// Выводим итоговое значение счетчика.
	// Ожидаемое значение: numGoroutines * numIncrements.
	log.Infof("Final Counter Value: %d\n", counter.Value())
}
