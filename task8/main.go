package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// setBitToOne устанавливает i-й бит переменной num в 1.
func setBitToOne(num int64, i uint) int64 {
	// 1 << i создает число, у которого только i-й бит равен 1, а остальные 0.
	// Побитовое ИЛИ (|) устанавливает i-й бит в 1, остальные биты остаются без изменений.
	return num | (1 << i)
}

// setBitToZero устанавливает i-й бит переменной num в 0.
func setBitToZero(num int64, i uint) int64 {
	// ^(1 << i) создает число, у которого все биты равны 1, кроме i-го, который равен 0.
	// Побитовое И (AND) (&) устанавливает i-й бит в 0, остальные биты остаются без изменений.
	return num & ^(1 << i)
}

// toBinaryString преобразует целое число в строку с двоичным представлением.
func toBinaryString(num int64) string {
	return fmt.Sprintf("%064b", num)
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	var num int64 = 42 // Пример переменной, для которой будем изменять биты

	var i uint = 1 // Индекс бита, который будем изменять

	log.Infof("Исходное число: %d (%s)\n", num, toBinaryString(num))

	// Устанавливаем i-й бит в 0
	num = setBitToZero(num, i)

	log.Infof("Число после установки %d-го бита в 0: %d (%s)\n", i, num, toBinaryString(num))

	// Устанавливаем i-й бит в 1.
	num = setBitToOne(num, i)

	log.Infof("Число после установки %d-го бита в 1: %d (%s)\n\n", i, num, toBinaryString(num))

	log.Infoln("Demo part")

	// Демонстрация операций с битами.
	var num64 uint64 = 1 << 1 //nolint:ineffassign // беззнаковое представление для упрощения.

	log.Infof("демо  1: %d (%s)\n", 1, toBinaryString(1))

	for c := range 10 {
		num64 = ^(1 << c)

		log.Infof("демо    1 << %d: %s\n", c, toBinaryString(1<<c))
		log.Infof("демо ^(1 << %d): %064b\n", c, num64)
	}
}
