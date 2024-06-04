package main

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// funcDelay задает задержку между итерациями работы горутины.
	funcDelay = 1 * time.Second
	// mainDelay задает основную задержку перед остановкой горутины.
	mainDelay = 3 * time.Second
)

// workerWithDone работает до тех пор, пока не получит сигнал на завершение через канал done.
func workerWithDone(done chan bool, log *logrus.Logger) {
	for {
		select {
		case <-done: // Если получен сигнал через канал done, то завершить работу.
			log.Infoln("Worker with done stopped")
			return
		default:
			log.Infoln("Worker with done working...")
			time.Sleep(funcDelay)
		}
	}
}

// workerWithContext работает до тех пор, пока не получит сигнал на завершение через контекст.
func workerWithContext(ctx context.Context, log *logrus.Logger) {
	for {
		select {
		case <-ctx.Done(): // Если произошла отмена контекста, то завершить работу.
			log.Infoln("Worker with context stopped")
			return
		default:
			log.Infoln("Worker with context working...")
			time.Sleep(funcDelay)
		}
	}
}

// workerWithAtomic работает до тех пор, пока атомарная переменная stopFlag не будет установлена в 1.
func workerWithAtomic(stopFlag *atomic.Bool, log *logrus.Logger) {
	for {
		if stopFlag.Load() { // Если атомарная переменная stopFlag == true то завершаем работу
			log.Infoln("Worker with atomic stopped")
			return
		}
		log.Infoln("Worker with atomic working...")
		time.Sleep(funcDelay)
	}
}

// workerWithSignals работает до тех пор, пока не получит сигнал на завершение через канал stop.
func workerWithSignals(stop chan os.Signal, log *logrus.Logger) {
	for {
		select {
		case <-stop: // Если получен системный сигнал через канал stop, завершить работу.
			log.Infoln("Worker with signals stopped")
			return
		default:
			log.Infoln("Worker with signals working...")
			time.Sleep(funcDelay)
		}
	}
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// 1. Использование канала done.
	log.Infoln("Starting worker with done")
	done := make(chan bool)      // Канал для управления остановкой горутины.
	go workerWithDone(done, log) // Запускаем горутину c каналом done.
	time.Sleep(mainDelay)        // Даем горутине поработать.
	done <- true                 // Отправляем сигнал через канал done для остановки горутины.
	time.Sleep(time.Second)

	// 2. То же самое, но используем таймер AfterFunc.
	log.Infoln("Starting worker with timer")
	go workerWithDone(done, log)
	time.AfterFunc(mainDelay, func() { // Используем time.AfterFunc
		done <- true
	})
	time.Sleep(mainDelay + time.Second)

	// 3. Использование контекста.
	log.Infoln("Starting worker with context")
	ctx, cancel := context.WithCancel(context.Background()) // Создаем контекст с возможностью отмены.
	go workerWithContext(ctx, log)                          // Запускаем горутину с контекстом.
	time.Sleep(mainDelay)
	cancel() // Отменяем контекст, чтобы остановить горутину.
	time.Sleep(time.Second + time.Second)

	// 4. Использование атомарной переменной.
	log.Infoln("Starting worker with atomic")
	var stopFlag atomic.Bool            // Атомарная переменная для управления остановкой горутины.
	go workerWithAtomic(&stopFlag, log) // Запускаем горутину с указателем на атомарную переменную.
	time.Sleep(mainDelay)
	stopFlag.Store(true) // Устанавливаем значение stopFlag в 1 для остановки горутины.
	time.Sleep(time.Second + time.Second)

	// 5. Использование системных сигналов.
	log.Infoln("Starting worker with signals")
	stop := make(chan os.Signal, 1)                      // Создаем канал для получения системных сигналов.
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM) // Указываем, какие сигналы следует перехватывать.
	go workerWithSignals(stop, log)                      // Запускаем горутину с каналом stop.
	time.AfterFunc(mainDelay, func() {                   // Используем time.AfterFunc
		stop <- syscall.SIGTERM // Отправляем системный сигнал.
	})
	sig := <-stop // Ожидание получения системного сигнала из канала stop.
	log.Infoln("Received signal:", sig)
}
