package main

import (
	"github.com/sirupsen/logrus"
)

// removeAtIndexSwap удаляет элемент на позиции i, заменяя его последним элементом.
// Порядок элементов не сохраняется. Эффективно при больших слайсах, когда порядок не важен.
func removeAtIndexSwap(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s
	}

	s[i] = s[len(s)-1] // Замена удаляемого элемента последним

	return s[:len(s)-1] // Возвращаем слайс без последнего элемента
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	slice := []int{1, 2, 3, 4, 5}
	log.Infoln("Original slice:", slice)

	index := 2 // Индекс удаляемого элемента
	modifiedSlice := removeAtIndexSwap(slice, index)
	log.Infoln("Modified slice:", modifiedSlice)
}
