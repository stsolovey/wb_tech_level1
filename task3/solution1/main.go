package main

import (
	"fmt"
	"sync"
)

// calculateSquare вычисляет квадрат числа.
func calculateSquare(number int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := number * number
	results <- square
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(numbers))
	var wg sync.WaitGroup

	// Создание горутин для каждого числа в массиве
	for _, number := range numbers {
		wg.Add(1)
		go calculateSquare(number, results, &wg)
	}

	// Закрытие канала results после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Суммирование результатов
	sum := 0
	for result := range results {
		sum += result
	}

	fmt.Printf("Сумма квадратов: %d\n", sum)
}
