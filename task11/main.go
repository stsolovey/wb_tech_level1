package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

// 11. Реализовать пересечение двух неупорядоченных множеств.

var errRangeTooSmall = errors.New("range too small for the specified set size")

const (
	setSize  = 5
	minValue = 1
	maxValue = 10
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	set1, err := createRandomSet(setSize, minValue, maxValue)
	if err != nil {
		log.WithError(err).Panic("Error creating random set")
	}

	set2, err := createRandomSet(setSize, minValue, maxValue)
	if err != nil {
		log.WithError(err).Panic("Error creating random set")
	}

	intersection := interserctSets(set1, set2)

	log.Infof("Множество 1: %v", set1)
	log.Infof("Множество 2: %v", set2)
	log.Infof("Пересечение множеств: %v", intersection)
}

// Создание множества случайных значений без повторений.
func createRandomSet(size, min, max int) ([]int, error) {
	if max-min+1 < size {
		return nil, fmt.Errorf("createRandomSet: %w", errRangeTooSmall)
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) //nolint:gosec

	// Создание map для хранения уникальных случайных чисел.
	set := make(map[int]any)
	for len(set) < size {
		// Генерация случайного числа.
		num := r.Intn(max-min+1) + min
		// Добавление случайного числа в map (в ключи добавляются только уникальные значения).
		set[num] = struct{}{}
	}

	// Преобразование в слайс для возврата значения.
	result := make([]int, 0)
	for num := range set {
		result = append(result, num)
	}

	return result, nil
}

// Возврат пересечения двух множеств.
func interserctSets(set1, set2 []int) []int {
	setMap := make(map[int]any) // Для уникальных значений set1.

	var intersection []int // Для возвращаемого результата.

	for _, num := range set1 {
		setMap[num] = struct{}{} // Добавляем уникальные значения.
	}

	for _, num := range set2 {
		if _, ok := setMap[num]; ok {
			intersection = append(intersection, num) // Добавляем только пересекающиеся элементы
		}
	}

	return intersection
}
