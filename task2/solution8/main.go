package main

import (
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

// SquareResult структура для хранения результатов.
type SquareResult struct {
	Number int
	Square int64
}

func main() {
	// Инициализация логгера.
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Массив чисел для вычисления квадратов.
	numbers := []int{2, 4, 6, 8, 10}

	// Слайс для хранения результатов.
	results := make([]SquareResult, len(numbers))

	// WaitGroup для ожидания завершения всех горутин.
	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for i, num := range numbers {
		go func(index, number int) {
			defer wg.Done()
			// Вычисляем квадрат числа.
			square := int64(number * number)

			// Записываем результат атомарно
			atomic.StoreInt64(&results[index].Square, square)

			results[index].Number = number
		}(i, num)
	}

	// Ожидаем завершения всех горутин.
	wg.Wait()

	// Выводим результаты.
	for _, result := range results {
		log.Infof("The square of %d is %d\n", result.Number, result.Square)
	}
}
