package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/sirupsen/logrus"
)

var errPointNotFound = errors.New("point not initialised")

// Point структура, представляющая точку в двумерном пространстве.
type Point struct {
	x float64
	y float64
}

// NewPoint конструктор для создания нового экземпляра Point.
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance метод для расчёта расстояния от текущей точки до другой.
// Возвращает расстояние и возможную ошибку, если одна из точек не инициализирована.
func (p *Point) Distance(other *Point) (float64, error) {
	if p == nil || other == nil {
		return 0, fmt.Errorf("func Distance() given a nil pointer: %w", errPointNotFound)
	}

	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2)), nil //nolint:mnd
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Создаем две точки
	p1 := NewPoint(1, 2) //nolint:mnd
	p2 := NewPoint(4, 6) //nolint:mnd

	// Вычисляем и выводим расстояние между ними
	distance, err := p1.Distance(p2)
	if err != nil {
		log.WithError(err).Panic("При вызове Distance с nil аргументом произошла ошибка")
	}

	log.Infof("Расстояние между точками: %.2f\n", distance)

	// Демонстрация ошибки: передача nil в качестве аргумента
	distance, err = p1.Distance(nil)
	if err != nil {
		log.WithError(err).Panic("При вызове Distance с nil аргументом произошла ошибка")
	}

	log.Infof("Расстояние между точками: %.2f\n", distance)
}
