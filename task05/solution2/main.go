package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{}) // Настройка форматирования логов

	const duration = 5 * time.Second // Определение времени работы программы

	dataChan := make(chan int)      // Создание канала для передачи данных
	doneChan := make(chan struct{}) // Канал для уведомления о завершении

	timer := time.NewTimer(duration) // Создание таймера для ограничения времени выполнения

	go func() {
		defer close(dataChan) // Гарантируем закрытие канала по завершению горутины

		for i := 0; ; i++ {
			time.Sleep(1 * time.Second) // Пауза между отправками данных

			select {
			case dataChan <- i: // Отправляем данные
				if i < 4 { //nolint:mnd
					if !timer.Stop() {
						<-timer.C
					}

					timer.Reset(duration) // Перезапускаем таймер
					log.Infoln("Таймер сброшен")
				}
			case <-timer.C: // Проверяем таймер
				log.Infoln("Время вышло, работа горутины завершена!")

				return
			case <-doneChan: // Завершаем горутину, если основной поток завершён
				log.Infoln("Получен сигнал завершения, работа горутины завершена!")

				return
			}
		}
	}()

	// Цикл для чтения данных из канала
	for {
		select {
		case data, ok := <-dataChan:
			if !ok {
				log.Infoln("Канал закрыт!")
				close(doneChan) // Закрываем канал - сигнализируем о завершении работы основной горутины

				return
			}

			log.Infoln("Получено:", data)
		case <-timer.C:
			log.Infoln("Время вышло, работа программы завершена!")
			close(doneChan) // Закрываем канал - сигнализируем о завершении работы основной горутины

			return
		}
	}
}
