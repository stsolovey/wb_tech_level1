package main

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// reverseWords переворачивает слова в строке, сохраняя порядок слов.
func reverseWords(s string) string {
	// Разбиваем строку на слова с помощью функции strings.Fields
	words := strings.Fields(s)
	// Определяем длину среза слов
	n := len(words)
	// Создаем новый срез для хранения перевернутых слов
	reversed := make([]string, n)
	// Переворачиваем порядок слов
	for i := range n {
		reversed[i] = words[n-1-i]
	}
	// Соединяем перевернутые слова обратно в строку с пробелами между ними
	return strings.Join(reversed, " ")
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	input := "snow dog sun"

	log.Infof("Исходная: %s", input)
	log.Infof("Перевернутая: %s", reverseWords(input))
}
