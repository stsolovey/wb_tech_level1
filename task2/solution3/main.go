package main

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Job структура, содержащая число для возведения в квадрат и канал для результата.
type Job struct {
	Number int

	Result chan<- int
}

// worker функция, выполняющая вычисления квадратов чисел из канала jobs.
func worker(jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		logrus.Infof("Calculating square for %d", job.Number)
		time.Sleep(time.Second) // Имитация задержки для проверки конкурентности

		square := job.Number * job.Number

		job.Result <- square
	}
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Массив чисел для вычисления квадратов
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(numbers))
	jobs := make(chan Job, len(numbers))

	// WaitGroup для синхронизации завершения всех воркеров
	var wg sync.WaitGroup

	// Запуск воркеров
	numWorkers := 3
	for range numWorkers {
		wg.Add(1)

		go worker(jobs, &wg)
	}

	// Отправка заданий в канал jobs
	for _, num := range numbers {
		jobs <- Job{Number: num, Result: results}
	}

	close(jobs)

	// Закрытие канала results после завершения всех воркеров
	go func() {
		wg.Wait()

		close(results)
	}()

	// Чтение и вывод результатов
	for square := range results {
		log.Infoln(square)
	}
}
