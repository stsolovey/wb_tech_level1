package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	numWorkers   = 5               // Количество воркеров
	dataInterval = 1 * time.Second // Интервал отправки данных
)

// worker представляет собой функцию горутины, которая обрабатывает данные.
func worker(ctx context.Context, log *logrus.Logger, id int, data <-chan int) {
	for {
		select {
		case <-ctx.Done(): // Проверка сигнала на завершение контекста
			log.Infof("Worker %d is stopping\n", id)

			return
		case val := <-data: // Получение данных из канала
			log.Infof("Worker %d received data: %d\n", id, val)
		}
	}
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Гарантия вызова cancel для очистки ресурсов контекста

	data := make(chan int) // Канал для передачи данных воркерам

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, log, i, data)
	}

	// Горутина для генерации и отправки данных воркерам
	go func() {
		for i := 0; ; i++ {
			data <- i

			time.Sleep(dataInterval) // Пауза перед следующей отправкой данных
		}
	}()

	// Настройка перехвата сигналов OS для корректного завершения работы
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // Перехват сигналов SIGINT и SIGTERM
	<-c                                             // Блокировка до получения сигнала
	log.Info("Shutting down gracefully...")
	cancel() // Отправка сигнала всем воркерам для остановки
}
