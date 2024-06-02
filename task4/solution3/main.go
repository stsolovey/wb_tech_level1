package main

import (
	"context"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	numJobs             = 100 // Количество задач
	delayJobsGen        = 500 * time.Millisecond
	delayJobsProcessing = 5 * time.Second
)

func worker(ctx context.Context, log *logrus.Logger, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				log.Infof("worker %d: channel closed\n", id)

				return
			}

			log.Infof("worker %d processing job %d\n", id, job)
			time.Sleep(time.Second) // Имитация работы

		case <-ctx.Done():
			log.Infof("worker %d: stopping as per context signal\n", id)

			return
		}
	}
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		log.Infoln("Please provide a valid number of workers.")

		return
	}

	jobs := make(chan int, numJobs) // Канал с буфером
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	// Создание пула воркеров
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)

		go worker(ctx, log, w, jobs, &wg)
	}

	// Постоянная отправка данных в канал
	go func() {
		for j := 1; j <= 100; j++ { // Ограничиваем количество задач для демонстрации
			jobs <- j

			time.Sleep(delayJobsGen) // Более быстрая генерация задач
		}

		close(jobs)
	}()

	// Даем время на обработку
	time.Sleep(delayJobsProcessing)

	cancel() // Отправляем сигнал остановки воркерам

	// Ожидание завершения всех воркеров
	wg.Wait()
	log.Infoln("All workers stopped gracefully.")
}
