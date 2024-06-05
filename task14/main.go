package main

// 14. Разработать программу,
// которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

func determineType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	default:
		return "unknown"
	}
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	// Примеры переменных различных типов
	var (
		intVar      = 42
		stringVar   = "hello"
		boolVar     = true
		chanIntVar  = make(chan int)
		chanStrVar  = make(chan string)
		chanBoolVar = make(chan bool)
	)

	// Список переменных для проверки
	vars := []interface{}{intVar, stringVar, boolVar, chanIntVar, chanStrVar, chanBoolVar}

	// Метод1: Определение типа каждой переменной, метод switch case type.
	log.Infoln("==== Метод 1 ====")

	for _, v := range vars {
		log.Infof("Тип переменной %v - %s", v, determineType(v))
	}

	// Метод 2: Определение типа каждой переменной с помощью библиотеки reflect.
	log.Infoln("==== Метод 2 ====")

	for _, v := range vars {
		log.Infof("Тип переменной %v - %s", v, reflect.TypeOf(v).String())
	}
}
