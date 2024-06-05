package main

import (
	"github.com/sirupsen/logrus"
	"sort"
)

// 10. Дана последовательность температурных колебаний:
// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножествах не важна.

const temperatureGroupStep = 10

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Исходные данные.
	temperatures := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Словарь для хранения групп температур.
	groups := make(map[int][]float64)

	// Группировка температур с шагом десять.
	for _, t := range temperatures {
		key := int(t/temperatureGroupStep) * temperatureGroupStep // Округление до 10 для формирования ключа.
		groups[key] = append(groups[key], t)
	}

	// Вывод значений словаря
	keys := make([]int, 0, len(groups)) // Преаллоцируем по требованию линтера.
	for k := range groups {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		log.Infof("%d: %v", k, groups[k])
	}
}
