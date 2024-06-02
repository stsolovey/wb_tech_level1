package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	numJobs = 100
)

func worker(log *logrus.Logger, id int, jobs <-chan int, done chan<- bool) {
	for job := range jobs {
		log.Infof("Worker %d started job %d\n", id, job)
		log.Infof("Worker %d finished job %d\n", id, job)
	}

	done <- true
}

func fanOut(jobs chan<- int, numJobs int) {
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)
}

func fanIn(done <-chan bool, numWorkers int) {
	for range numWorkers {
		<-done
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

	jobs := make(chan int, numJobs)
	done := make(chan bool, numWorkers)

	// Fan-out: запуск воркеров
	for w := 1; w <= numWorkers; w++ {
		go worker(log, w, jobs, done)
	}

	// Отправка работы воркерам
	go fanOut(jobs, numJobs)

	// Fan-in: ожидание завершения всех воркеров
	fanIn(done, numWorkers)

	log.Infoln("All jobs processed.")
}
