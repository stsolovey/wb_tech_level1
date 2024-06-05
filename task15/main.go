package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	initialSize     = 1024 // 1024 bytes, понятно объясненное значение константы.
	substringLength = 100  // Ясно определённая длина подстроки.
)

var errStringTooSmall = errors.New("string is too small")

// Создание и возврат строки заданной длины
func createHugeString(size int) string {
	var builder strings.Builder

	builder.Grow(size) // Можно сразу создать строку нужного размера.

	for range make([]int, size) {
		builder.WriteByte('a') // Заполнение строки символами 'a' для примера.
	}

	return builder.String()
}

// Проверка на наличие ошибки перед взятием среза и явное копирование строки.
func someFunc() (string, error) {
	v := createHugeString(initialSize)

	if len(v) < substringLength {
		return "", fmt.Errorf("someFunc: %w", errStringTooSmall) // Ошибка, если исходная строка меньше среза.
	}

	return string([]byte(v[:substringLength])), nil // Явное копирование данных для избежания утечки памяти.
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	justString, err := someFunc() // Использование локальной переменной вместо глобальной.
	if err != nil {
		log.WithField("err", err).Panic("Failed to create a new string")
	}

	log.Infoln(justString, len(justString))
}
