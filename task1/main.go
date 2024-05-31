package main

import (
	"fmt"
)

// Human структура с произвольным набором полей и методов
type Human struct {
	Name string
	Age  int
}

// Greet метод Human для печати приветствия
func (h Human) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

// Action структура, которая встраивает Human
type Action struct {
	Human
}

// PerformAction метод Action для выполнения действия
func (a Action) PerformAction() {
	fmt.Println("Performing an action.")
}

func main() {
	// Создаем экземпляр Human
	human := Human{Name: "John", Age: 30}
	// Создаем экземпляр Action с встраиванием Human
	action := Action{Human: human}

	// Используем метод структуры Human через экземпляр Action
	action.Greet()         // Наследованный метод от Human
	action.PerformAction() // Метод структуры Action
}
