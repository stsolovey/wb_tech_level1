package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

func calculateSquare(num int, wg *sync.WaitGroup, results *sync.Map) {
	defer wg.Done()

	logrus.Infof("Calculating square for %d", num)
	square := num * num

	results.Store(num, square)
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Массив чисел для вычисления квадратов
	numbers := []int{2, 4, 6, 8, 10}

	// `sync.Map` для хранения результатов
	var results sync.Map

	// WaitGroup для синхронизации завершения всех горутин
	var wg sync.WaitGroup

	// Запуск горутин для вычисления квадратов чисел
	for _, num := range numbers {
		wg.Add(1)

		go calculateSquare(num, &wg, &results)
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод результатов
	results.Range(func(key, value any) bool {
		log.Infof("Square of %d is %d", key, value)

		return true
	})
}
