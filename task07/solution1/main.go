package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

const goNum = 10 // Кол-во горутин.

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	var (
		wg   sync.WaitGroup
		data sync.Map // Использование sync.Map для потокобезопасной работы с map.
	)

	// Функция для конкурентной записи данных в sync.Map.
	writeData := func(key, value int) {
		defer wg.Done() // Обеспечение корректного завершения горутины.

		data.Store(key, value) // Запись данных в sync.Map.

		log.Infof("Запсь: key - %d, value - %d", key, value)
	}

	// Запуск 10 горутин для записи данных в sync.Map.
	for i := range goNum {
		wg.Add(1) // Увеличение счетчика горутин.

		go writeData(i, i*goNum) // Запуск горутины для записи данных.
	}

	wg.Wait() // Ожидание завершения всех горутин.

	// Вывод данных из sync.Map.
	data.Range(func(key, value any) bool {
		log.Infof("Чтение: key - %d, value - %d", key, value)

		return true // Продолжение перебора элементов.
	})
}
