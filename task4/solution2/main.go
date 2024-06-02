package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	numWorkers       = 5               // Количество воркеров
	dataSendInterval = 1 * time.Second // Интервал отправки данных
)

// worker - функция горутины, обрабатывающая входящие данные из канала.
func worker(log *logrus.Logger, done <-chan bool, id int, data <-chan int) {
	for {
		select {
		case <-done:
			// Если получен сигнал завершения, выходим из горутины
			log.Infof("Worker %d is stopping\n", id)

			return
		case val := <-data:
			// Обработка полученных данных
			log.Infof("Worker %d received data: %d\n", id, val)
		}
	}
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Канал done используется для сигнализации о необходимости завершения работы воркеров
	done := make(chan bool)
	// Канал data используется для передачи данных воркерам
	data := make(chan int)

	// Запуск заданного количества воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(log, done, i, data)
	}

	// Горутина для генерации данных
	go func() {
		for i := 0; ; i++ {
			data <- i
			// Пауза перед следующей отправкой данных
			time.Sleep(dataSendInterval)
		}
	}()

	// Настройка для перехвата сигналов прерывания (Ctrl+C) и завершения программы
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // Ожидание сигнала

	// Закрытие канала done для уведомления всех воркеров о необходимости завершения работы
	close(done)
}
