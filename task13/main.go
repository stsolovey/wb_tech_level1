package main

import (
	"github.com/sirupsen/logrus"
)

const (
	aVal           = 5
	bVal           = 10
	channelBufSize = 2
)

func setVals(a, b int) (int, int) {
	return a, b
}

func main() { //nolint:funlen
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	a, b := setVals(aVal, bVal)
	log.Infof("Изначальные значения: a = %d, b = %d", a, b)

	// Метод 1: Использование арифметических операций
	a += b
	b = a - b
	a -= b
	log.Infof("После обмена (метод 1): a = %d, b = %d", a, b)

	// Восстанавливаем значения для следующего примера
	a, b = setVals(aVal, bVal)

	// Метод 2: Использование побитового XOR
	a ^= b
	b = a ^ b
	a ^= b

	log.Infof("После обмена (метод 2): a = %d, b = %d", a, b)

	// Восстанавливаем значения для следующего примера
	a, b = setVals(aVal, bVal)

	// Метод 3: Использование указателей
	swapPointers := func(a, b *int) { *a, *b = *b, *a }

	swapPointers(&a, &b)

	log.Infof("После обмена (метод 3): a = %d, b = %d", a, b)

	// Восстанавливаем значения для следующего примера
	a, b = setVals(aVal, bVal)

	// Метод 4: Использование множественного присваивания

	a, b = b, a
	log.Infof("После обмена (метод 4): a = %d, b = %d", a, b)

	// Восстанавливаем значения для следующего примера
	a, b = setVals(aVal, bVal)

	// Метод 5: Использование умножения и деления

	// Предотвращение деления на ноль
	switch {
	case a == 0:
		a = b
		b = 0
	case b == 0:
		b = a
		a = 0
	default:
		// Использование умножения и деления для обмена значениями
		b *= a
		a = b / a
		b /= a
	}

	log.Infof("После обмена (метод 5): a = %d, b = %d", a, b)

	// Восстанавливаем значения для следующего примера
	a, b = setVals(aVal, bVal)

	// Метод 6: Использование битовых операций

	// то же что a = a + b
	a = (a & b) + (a | b)
	// то же что b = a - b
	b = a + (^b) + 1
	// то же что a = a - b
	a = a + (^b) + 1

	log.Infof("После обмена (метод 6): a = %d, b = %d", a, b)

	// Восстанавливаем значения для следующего примера
	a, b = setVals(aVal, bVal)

	// Метод 7: Использование каналов, метод не эффективен, но показывает возможности и принцип работы каналов.

	ch := make(chan int, channelBufSize)

	// Отправляем значения в канал
	ch <- a
	ch <- b

	// Получаем значения из канала в обратном порядке
	b = <-ch
	a = <-ch

	log.Infof("После обмена (метод 7): a = %d, b = %d", a, b)
}
