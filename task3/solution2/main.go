package main

import (
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

func calculateSquare(number int, sum *int64, wg *sync.WaitGroup) {
	defer wg.Done()

	square := int64(number * number)

	atomic.AddInt64(sum, square)
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	numbers := []int{2, 4, 6, 8, 10}

	var sum int64

	var wg sync.WaitGroup

	for _, number := range numbers {
		wg.Add(1)

		go calculateSquare(number, &sum, &wg)
	}

	wg.Wait()
	log.Infof("Сумма квадратов: %d\n", sum)
}
