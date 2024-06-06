package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{}) // Настройка форматирования логов

	// Устанавливаем таймаут для контекста.
	const timeoutDuration = 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel() // Гарантируем вызов cancel для освобождения ресурсов контекста.
	// Создание канала для передачи целочисленных данных.
	dataChan := make(chan int)
	// Горутина для генерации данных.
	go func() {
		// Начальное значение для данных и счётчик итераций.
		for i := 0; ; i++ {
			select {
			case <-ctx.Done(): // Проверка на отмену контекста.
				return // Выход из горутины, если контекст отменён.
			default:
				dataChan <- i // Отправка данных в канал.

				time.Sleep(1 * time.Second) // Интервал между отправками составляет 1 секунду.
			}
		}
	}()

	// Цикл для чтения данных из канала до истечения контекста.
	for {
		select {
		case data := <-dataChan: // Чтение данных из канала.
			log.Infoln("Received:", data)
		case <-ctx.Done(): // Срабатывание контекста по истечении таймаута или при его отмене.
			log.Infoln("Time is up!")

			return // Завершение работы программы.
		}
	}
}
