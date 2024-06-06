package main

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Task структура задачи, содержит число и канал для результата.
type Task struct {
	Number int
	Result chan<- int
}

// worker представляет собой воркера, который обрабатывает задачи.
func worker(log *logrus.Logger, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		log.Infof("Worker started processing %d", task.Number)
		// Имитация задержки для демонстрации работы в конкурентной среде.
		time.Sleep(time.Second)

		square := task.Number * task.Number

		task.Result <- square
		log.Infof("Worker finished processing %d", task.Number)
	}
}

func main() {
	// Инициализация логгера.
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Количество воркеров.
	const numWorkers = 3

	// Канал для передачи задач воркерам.
	tasks := make(chan Task, numWorkers)

	// Канал для сбора результатов
	results := make(chan int, numWorkers)

	// WaitGroup для синхронизации завершения воркеров.
	var wg sync.WaitGroup

	// Инициализация воркеров.
	wg.Add(numWorkers)

	for range numWorkers {
		go worker(log, tasks, &wg)
	}

	// Массив чисел для вычисления квадратов.
	numbers := []int{2, 4, 6, 8, 10}

	// Отправка задач воркерам.
	for _, number := range numbers {
		tasks <- Task{Number: number, Result: results}
	}

	// Закрытие канала задач после отправки всех задач.
	close(tasks)

	// Горутина для закрытия канала результатов после завершения всех воркеров.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Чтение и вывод результатов.
	for result := range results {
		log.Infoln("Square:", result)
	}
}
