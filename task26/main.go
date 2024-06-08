package main

import (
	"github.com/sirupsen/logrus"
	"unicode"
)

// areCharactersUnique проверяет, содержит ли строка только уникальные символы.
func areCharactersUnique(s string) bool {
	charMap := make(map[rune]any) // Использование map для хранения уникальных символов
	for _, char := range s {
		normalizedChar := unicode.ToLower(char) // Приведение символов к нижнему регистру
		if _, exists := charMap[normalizedChar]; exists {
			return false // Найден повторяющийся символ
		}
		charMap[normalizedChar] = struct{}{}
	}
	return true
}

func main() {
	log := logrus.New()

	examples := []string{"абвгд", "abCdefAaf", "aabcd", "1234567890", "!@#$%^&*()_+"}
	for _, example := range examples {
		log.Infof("%s\t - %t", example, areCharactersUnique(example))
	}
}
