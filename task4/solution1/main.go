package main

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	minArgs = 2
)

func produce(data chan<- int) {
	// Генерация данных в бесконечном цикле
	count := 0

	for {
		data <- count

		count++

		time.Sleep(1 * time.Second) // Имитация задержки для читабельности вывода
	}
}

func consume(log *logrus.Logger, id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range data {
		log.Infof("Worker %d received data: %d\n", id, num)
	}
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	if len(os.Args) < minArgs {
		log.Infoln("Usage: go run main.go <number of workers>")

		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Infof("Invalid number of workers: %s\n", err)

		return
	}

	// Канал для передачи данных от продюсера к воркерам
	data := make(chan int)

	var wg sync.WaitGroup

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)

		go consume(log, i, data, &wg)
	}

	// Запуск производителя данных
	go produce(data)

	wg.Wait() // Ожидание завершения работы всех воркеров (теоретически не наступит)
}
