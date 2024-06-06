package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

// customSleep использует time.Ticker для ожидания указанной длительности без блокирования.
func customSleep(d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop() // Очистка ticker при возвращении из функции.

	// Ожидание получения значения из канала ticker.
	// Это произойдет после прохождения длительности 'd'.
	<-ticker.C
}

func main() {
	log := logrus.New()

	const duration = 10 * time.Second

	log.Infof("Start at : %v\n", time.Now())
	customSleep(duration)
	log.Infof("Finish at: %v\n", time.Now())
}
