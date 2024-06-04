package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{}) // Настройка форматирования логов

	// Определение длительности работы программы.
	const duration = 5 * time.Second

	// Создание канала для передачи целочисленных данных.
	dataChan := make(chan int)

	// Горутина для генерации данных.
	go func() {
		// Начальное значение для данных.
		for i := 0; ; i++ {
			dataChan <- i // Отправка данных в канал.

			time.Sleep(1 * time.Second) // Интервал между отправками составляет 1 секунду.
		}
	}()

	// Таймер для контроля времени выполнения программы.
	timeout := time.After(duration)

	// Цикл для чтения данных из канала до истечения таймера.
	for {
		select {
		case data := <-dataChan: // Чтение данных из канала.
			log.Infoln("Received:", data)
		case <-timeout: // Срабатывание таймера по истечении заданного времени.
			log.Infoln("Time is up!")

			return // Завершение работы программы.
		}
	}
}
