package main

// 9. Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.`

import (
	"sync"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	numbers := []int{1, 2, 3, 4, 5} // Массив чисел для обработки
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	var wg sync.WaitGroup

	// Горутина для записи чисел в первый канал
	wg.Add(1)

	go func() {
		defer wg.Done()

		for _, num := range numbers {
			inputChannel <- num
		}

		close(inputChannel)
	}()

	// Горутина для обработки чисел из первого канала и записи во второй канал
	wg.Add(1)

	go func() {
		defer wg.Done()

		for num := range inputChannel {
			outputChannel <- num * 2
		}

		close(outputChannel)
	}()

	// Горутина для чтения результатов из второго канала и вывода в stdout
	wg.Add(1)

	go func() {
		defer wg.Done()

		for result := range outputChannel {
			log.Infoln(result)
		}
	}()

	// Ожидание завершения всех горутин
	wg.Wait()
}
