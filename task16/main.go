package main

import (
	"github.com/sirupsen/logrus"
)

// quickSort реализует алгоритм быстрой сортировки.
// arr - это срез (массив), который нужно отсортировать.
func quickSort(arr []int) []int {
	if len(arr) < 2 { //nolint:mnd
		return arr
	}

	// Инициализация индексов для будущей работы алгоритма
	left, right := 0, len(arr)-1

	pivotIndex := medianOfThree(arr, left, right)

	// Перемещение опорного элемента в конец среза
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Разделение массива на элементы меньше опорного и больше опорного
	// Итерируем через массив от начала до элемента перед опорным
	for i := range right {
		// Если текущий элемент меньше опорного элемента, который находится в позиции 'right',
		// то нужно поместить этот элемент в левую часть массива.
		// Это условие проверяет, является ли элемент меньше опорного.
		if arr[i] < arr[right] {
			// Обмен arr[i] и arr[left] позволяет переместить меньший элемент в 'левую' часть,
			// которая будет содержать все элементы меньше опорного.
			// arr[left] - это граница, до которой уже добавлены все меньшие элементы.
			arr[i], arr[left] = arr[left], arr[i]

			// После обмена увеличиваем 'left', что расширяет границу 'левой' части,
			// теперь она включает только что перенесённый меньший элемент.
			left++
		}
	}

	// Размещение опорного элемента на его окончательное место
	arr[left], arr[right] = arr[right], arr[left]

	// Рекурсивное применение быстрой сортировки к левому и правому подмассивам
	quickSort(arr[:left])   // Сортировка левой части
	quickSort(arr[left+1:]) // Сортировка правой части

	return arr
}

// medianOfThree определяет медианное значение среди первого, среднего и последнего элементов
// и переставляет элементы так, чтобы медиана оказалась на средней позиции. Это помогает
// в выборе более оптимального опорного элемента для быстрой сортировки.
func medianOfThree(arr []int, left int, right int) int {
	// Вычисление индекса среднего элемента
	mid := left + (right-left)/2 //nolint:mnd

	// Сравниваем и меняем местами элементы, если необходимо, чтобы
	// arr[left] стал меньше или равен arr[mid]
	if arr[left] > arr[mid] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}

	// Сравниваем и меняем местами элементы, если необходимо, чтобы
	// arr[left] стал меньше или равен arr[right]
	if arr[left] > arr[right] {
		arr[left], arr[right] = arr[right], arr[left]
	}

	// Сравниваем и меняем местами элементы, если необходимо, чтобы
	// arr[mid] стал меньше или равен arr[right]
	// Это гарантирует, что arr[mid] теперь медиана среди arr[left], arr[mid], arr[right]
	if arr[mid] > arr[right] {
		arr[mid], arr[right] = arr[right], arr[mid]
	}

	// Возвращаем индекс среднего элемента, который теперь содержит медианное значение
	return mid
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	arr := []int{9, -3, 5, 2, 6, 8, -6, 1, 3} // Исходный массив
	log.Infoln("Original array:", arr)        // Вывод исходного массива
	sortedArr := quickSort(arr)               // Сортировка массива
	log.Infoln("Sorted array:  ", sortedArr)  // Вывод отсортированного массива
}
