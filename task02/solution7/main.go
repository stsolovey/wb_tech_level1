package main

import (
	"reflect"
	"sync"

	"github.com/sirupsen/logrus"
)

// calculateSquare вычисляет квадрат числа и отправляет результат в один из каналов.
func calculateSquare(log *logrus.Logger, num int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	square := num * num

	results <- square

	log.Infof("Finished processing %d", num)
}

func main() {
	// Инициализация логгера.
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Массив чисел для вычисления квадратов.
	numbers := []int{2, 4, 6, 8, 10}

	// Создание каналов для каждого числа
	channels := make([]chan int, len(numbers))

	var wg sync.WaitGroup

	// Инициализация каналов и запуск горутин.
	for i, num := range numbers {
		channels[i] = make(chan int, 1)

		wg.Add(1)

		go calculateSquare(log, num, channels[i], &wg)
	}

	// Создание slice из reflect.SelectCase.
	cases := make([]reflect.SelectCase, len(channels))
	for i, ch := range channels {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	// Ожидание результатов из всех каналов.
	for len(cases) > 0 {
		chosen, value, ok := reflect.Select(cases)
		if !ok {
			// Канал закрыт, убираем его из списка.
			cases = append(cases[:chosen], cases[chosen+1:]...)

			continue
		}

		log.Infof("Received %d from channel %d\n", value.Int(), chosen)
	}

	// Закрытие всех каналов и ожидание завершения всех горутин.
	wg.Wait()

	for _, ch := range channels {
		close(ch)
	}
}
