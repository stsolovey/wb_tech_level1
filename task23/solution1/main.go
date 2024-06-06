package main

import (
	"github.com/sirupsen/logrus"
)

// removeAtIndex удаляет элемент по индексу i из слайса s.
// Функция сохраняет порядок оставшихся элементов.
func removeAtIndex(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s // Возвращаем исходный слайс, если индекс находится вне диапазона
	}
	// Используем append для создания нового слайса без i-го элемента
	return append(s[:i], s[i+1:]...)
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{}) // Настройка форматирования логов

	slice := []int{1, 2, 3, 4, 5}
	log.Infoln("Исходный слайс:", slice)

	index := 2 // Индекс удаляемого элемента (индексация с нуля)
	modifiedSlice := removeAtIndex(slice, index)

	log.Infoln("Изменённый слайс:", modifiedSlice)
}
