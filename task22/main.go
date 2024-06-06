package main

import (
	"math/big"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Инициализация переменных a и b со значениями больше 2^20
	a, okA := new(big.Float).SetString("1048577") // 2^20 + 1
	b, okB := new(big.Float).SetString("1048578") // 2^20 + 2

	// Проверка успешности конвертации строк в big.Float
	if !okA || !okB {
		log.Panic("Ошибка при конвертации входных данных в big.Float")
	}

	// Выполнение операций
	sum := new(big.Float).Add(a, b)        // Сложение
	difference := new(big.Float).Sub(a, b) // Вычитание
	product := new(big.Float).Mul(a, b)    // Умножение
	quotient := new(big.Float).Quo(a, b)   // Деление

	// Вывод результатов
	log.Infof("Сложение: %s\n", sum.Text('f', -1)) // Вывод в формате с плавающей точкой
	log.Infof("Вычитание: %s\n", difference.Text('f', -1))
	log.Infof("Умножение: %s\n", product.Text('f', -1))
	log.Infof("Деление: %s\n", quotient.Text('f', -1))
}
