package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

const goNum = 10

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	var (
		mutex sync.Mutex // Использование мьютекса для потокобезопасной работы с map
		wg    sync.WaitGroup

		data = make(map[int]int)
	)

	// Функция для конкурентной записи данных в map
	writeData := func(key, value int) {
		defer wg.Done() // Обеспечение корректного завершения горутины

		mutex.Lock() // Блокировка доступа к map

		data[key] = value // Запись данных в map

		log.Infof("Запись: key - %d, value - %d", key, value)

		mutex.Unlock() // Разблокировка доступа к map
	}

	// Запуск 10 горутин для записи данных в map
	for i := range goNum {
		wg.Add(1) // Увеличение счетчика горутин

		go writeData(i, i*goNum) // Запуск горутины для записи данных
	}

	wg.Wait() // Ожидание завершения всех горутин

	// Вывод данных из map
	mutex.Lock() // Блокировка доступа к map перед чтением

	for key, value := range data {
		log.Infof("Чтение: key - %d, value - %d", key, value)
	}

	mutex.Unlock() // Разблокировка доступа к map после чтения
}
