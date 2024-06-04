package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

// Определение длительности работы программы.
const programWorkDuration = 5 * time.Second

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{}) // Настройка форматирования логов

	// Создание канала для передачи целочисленных данных.
	dataChan := make(chan int)

	// Создание канала для перехвата системных сигналов.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Горутина для генерации данных.
	go func() {
		for i := 0; ; i++ {
			dataChan <- i

			time.Sleep(1 * time.Second)
		}
	}()

	// Таймер для автоматического завершения программы.
	go func() {
		time.Sleep(programWorkDuration)

		sigChan <- syscall.SIGTERM // Имитация получения сигнала завершения работы.
	}()

	// Цикл обработки сообщений и сигналов.
	for {
		select {
		case data := <-dataChan:
			log.Infoln("Received:", data)
		case sig := <-sigChan:
			log.Infoln("Received signal:", sig)

			return
		}
	}
}
