package main

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// calculateSquare вычисляет квадрат числа и отправляет результат в канал.
func calculateSquare(num int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()

	logrus.Infof("Calculating square for %d", num)
	time.Sleep(time.Second) // Имитация задержки для проверки конкурентности

	square := num * num
	results <- square
}

func main() {
	// Инициализация логгера
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Массив чисел для вычисления квадратов
	numbers := []int{2, 4, 6, 8, 10}

	// Канал для передачи результатов
	results := make(chan int, len(numbers))

	// WaitGroup для синхронизации завершения всех горутин
	var wg sync.WaitGroup

	// Запуск горутин для вычисления квадратов чисел
	for _, num := range numbers {
		wg.Add(1)

		go calculateSquare(num, &wg, results)
	}

	// Закрытие канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Чтение и вывод результатов
	for square := range results {
		log.Infoln(square)
	}
}
