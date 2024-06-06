package main

import (
	"github.com/sirupsen/logrus"
)

// reverseString переворачивает переданную строку, корректно обрабатывая Unicode символы.
func reverseString(s string) string {
	// Конвертируем строку в срез рун, чтобы корректно работать с Unicode символами
	runes := []rune(s)
	// Определяем длину среза
	n := len(runes)
	// Создаем новый срез для хранения перевернутых рун
	reversed := make([]rune, n)
	// Переворачиваем руны
	for i := range n {
		reversed[i] = runes[n-1-i]
	}
	// Конвертируем перевернутый срез рун обратно в строку и возвращаем её
	return string(reversed)
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	input := "главрыба"

	log.Infof("Original: %s", input)
	log.Infof("Reversed: %s", reverseString(input))
}
