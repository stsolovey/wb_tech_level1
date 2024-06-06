package main

import (
	"github.com/sirupsen/logrus"
)

// binarySearch реализует алгоритм бинарного поиска.
// arr - это отсортированный срез, в котором производится поиск,
// key - значение, которое нужно найти.
// Функция возвращает индекс элемента в срезе и булево значение, указывающее, был ли элемент найден.
func binarySearch(arr []int, key int) (int, bool) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		// Находим средний индекс, избегая переполнения.
		mid := low + (high-low)/2 //nolint:mnd

		// Сравниваем ключ с элементом в середине среза
		switch {
		case arr[mid] == key:
			return mid, true
		case arr[mid] > key:
			high = mid - 1 // Искать в левой половине
		case arr[mid] < key:
			low = mid + 1 // Искать в правой половине
		}
	}

	return -1, false // Ключ не найден
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	arr := []int{1, 2, 4, 5, 7, 8, 9} // Пример отсортированного среза
	key := 5                          // Значение, которое нужно найти

	index, found := binarySearch(arr, key)
	if found {
		log.Infof("Element %d found at index %d\n", key, index)
	} else {
		log.Infof("Element %d not found in the array\n", key)
	}
}
