package main

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	jobQueueBuffer = 100
)

func worker(log *logrus.Logger, id int, jobs <-chan int, wg *sync.WaitGroup) {
	for j := range jobs {
		log.Infof("worker %d processing job %d\n", id, j)
		time.Sleep(time.Second) // Имитация работы
	}

	wg.Done()
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		log.Infoln("Please provide a valid number of workers.")

		return
	}

	jobs := make(chan int, jobQueueBuffer) // Канал с буфером
	var wg sync.WaitGroup

	// Создание пула воркеров
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)

		go worker(log, w, jobs, &wg)
	}

	// Постоянная отправка данных в канал
	go func() {
		for j := 1; ; j++ {
			jobs <- j

			time.Sleep(time.Second) // Имитация новых данных
		}
	}()

	// Ожидание завершения воркеров (теоретически не наступит)
	wg.Wait()
}
