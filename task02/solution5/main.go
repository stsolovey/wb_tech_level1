package main

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// calculateSquare вычисляет квадрат числа и сохраняет результат в срез.
func calculateSquare(num int, wg *sync.WaitGroup, results []int, index int) {
	defer wg.Done()

	start := time.Now()

	logrus.Infof("Started calculating square for %d", num)
	time.Sleep(time.Second) // Фиксированная задержка для проверки конкурентности

	results[index] = num * num
	logrus.Infof("Finished calculating square for %d, took %v", num, time.Since(start))
}

func main() {
	// Инициализация логгера
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Массив чисел для вычисления квадратов
	numbers := []int{2, 4, 6, 8, 10}
	results := make([]int, len(numbers))

	// WaitGroup для синхронизации завершения всех горутин
	var wg sync.WaitGroup

	// Запуск горутин для вычисления квадратов чисел
	for i, num := range numbers {
		wg.Add(1)

		go calculateSquare(num, &wg, results, i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод результатов
	for _, square := range results {
		log.Infoln(square)
	}
}
