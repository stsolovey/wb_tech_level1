package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	numWorkers       = 5               // количество воркеров
	dataSendInterval = 1 * time.Second // интервал отправки данных
)

// worker функция горутины, которая обрабатывает входящие данные.
func worker(log *logrus.Logger, wg *sync.WaitGroup, id int, data <-chan int) {
	defer wg.Done() // Пометить завершение горутины в WaitGroup по завершению функции

	for val := range data { // Бесконечный цикл, читающий данные из канала
		log.Infof("Worker %d received data: %d\n", id, val) // Логирование полученных данных
	}
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{}) // Настройка форматирования логов

	var wg sync.WaitGroup // Инициализация группы ожидания для синхронизации завершения горутин

	data := make(chan int) // Создание канала для передачи данных воркерам

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Уведомление WaitGroup о новой задаче

		go worker(log, &wg, i, data) // Запуск горутины воркера
	}

	// Горутина для генерации данных
	go func() {
		for i := 0; ; i++ { // Бесконечный цикл для генерации данных
			data <- i // Отправка данных в канал

			time.Sleep(dataSendInterval) // Пауза перед следующей отправкой
		}
	}()

	// Настройка перехвата системных сигналов для корректного завершения программы
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // Перехват сигналов прерывания и завершения
	<-c                                             // Ожидание поступления сигнала

	close(data) // Закрытие канала, что приведет к завершению чтения воркерами
	wg.Wait()   // Ожидание завершения всех воркеров
}
