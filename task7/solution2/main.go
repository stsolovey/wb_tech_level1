package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

const goNum = 10 // Кол-во горутин.

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	type kv struct {
		key   int
		value int
	}

	dataChan := make(chan kv)       // Создание канала для передачи данных.
	stopChan := make(chan struct{}) // Создание канала для остановки горутины.

	var wg sync.WaitGroup

	data := make(map[int]int)

	// Горутина для обработки данных из канала.
	go func() {
		for {
			select {
			case kv := <-dataChan: // Получение данных из канала.
				data[kv.key] = kv.value // Запись данных в map.
				log.Infof("Запсь: key - %d, value - %d", kv.key, kv.value)
			case <-stopChan: // Получение сигнала о завершении.
				return // Завершение горутины.
			}
		}
	}()

	// Функция для записи данных в канал.
	// Принимает ключ и значение, которые необходимо записать в канал dataChan.
	// Использует sync.WaitGroup для синхронизации завершения горутин.
	writeData := func(key, value int) {
		defer wg.Done() // Обеспечение корректного завершения горутины после выполнения тела функции

		// Создание экземпляра структуры kv с переданными ключом и значением
		// и отправка его в канал dataChan для дальнейшей обработки в отдельной горутине.
		dataChan <- kv{key, value}
	}

	// Запуск 10 горутин для записи данных в канал.
	for i := range goNum {
		wg.Add(1) // Увеличение счетчика горутин.

		go writeData(i, i*goNum) // Запуск горутины для записи данных.
	}

	wg.Wait()       // Ожидание завершения всех горутин.
	close(stopChan) // Закрытие канала для остановки горутины.

	// Вывод данных из map.
	for key, value := range data {
		log.Infof("Чтение: key - %d, value - %d", key, value)
	}
}
