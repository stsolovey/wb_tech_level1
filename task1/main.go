package main

import (
	"github.com/sirupsen/logrus"
)

// Human структура с произвольным набором полей и методов.
type Human struct {
	log  *logrus.Logger
	Name string
	Age  int
}

// NewHuman конструктор для создания экземпляра Human с логгером.
func NewHuman(name string, age int) Human {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	return Human{
		log:  log,
		Name: name,
		Age:  age,
	}
}

// Greet метод Human для печати приветствия.
func (h Human) Greet() {
	h.log.Infof("Hello, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

// Action структура, которая встраивает Human.
type Action struct {
	Human
}

// PerformAction метод Action для выполнения действия.
func (a Action) PerformAction() {
	a.log.Infof("Performing an action.")
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	const humanAge = 30
	// Создаем экземпляр Human.
	human := NewHuman("John", humanAge)
	// Создаем экземпляр Action с встраиванием Human.
	action := Action{Human: human}

	// Используем метод структуры Human через экземпляр Action.
	action.Greet()         // Наследованный метод от Human.
	action.PerformAction() // Метод структуры Action.
}
