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
		rwMutex sync.RWMutex // Использование RWMutex
		wg      sync.WaitGroup
		data    = make(map[int]int)
	)

	// Функция для конкурентной записи данных в map
	writeData := func(key, value int) {
		defer wg.Done() // Обеспечение корректного завершения горутины

		rwMutex.Lock() // Блокировка доступа к map для записи

		data[key] = value // Запись данных в map
		log.Infof("Запись: key - %d, value - %d", key, value)

		rwMutex.Unlock() // Разблокировка доступа к map после записи
	}

	// Функция для конкурентного чтения данных из map
	readData := func(key int) int {
		rwMutex.RLock() // Блокировка доступа к map для чтения

		defer rwMutex.RUnlock() // Разблокировка доступа к map после чтения

		return data[key] // Чтение данных из map
	}

	// Запуск 10 горутин для записи данных в map
	for i := range goNum {
		wg.Add(1) // Увеличение счетчика горутин

		go writeData(i, i*goNum) // Запуск горутины для записи данных
	}

	wg.Wait()

	// Чтение данных
	for i := range goNum {
		log.Infof("Чтение: key - %d, value - %d", i, readData(i))
	}
}
