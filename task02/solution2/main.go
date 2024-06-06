package main

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// calculateSquare вычисляет квадрат числа и отправляет результат в канал.
func calculateSquare(num int, results chan<- int) {
	logrus.Infof("Calculating square for %d", num)
	time.Sleep(time.Second) // Имитация задержки для демонстрации конкурентности

	square := num * num
	results <- square
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Массив чисел для вычисления квадратов
	numbers := []int{2, 4, 6, 8, 10}

	// Создание канала для передачи результатов
	results := make(chan int, len(numbers))

	// WaitGroup для синхронизации завершения всех горутин
	var wg sync.WaitGroup

	// Запуск горутин для каждого числа в массиве
	for _, num := range numbers {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			calculateSquare(n, results)
		}(num)
	}

	// Горутина для закрытия канала по завершению всех вычислений
	go func() {
		wg.Wait()

		close(results)
	}()

	// Чтение из канала и вывод результатов
	for square := range results {
		log.Infof("Square: %d", square)
	}
}
