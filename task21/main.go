package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Data Структура данных.
type Data struct {
	Name  string `json:"name" xml:"name"`
	Value string `json:"value" xml:"value"`
}

// JSONSystem Интерфейс для работы с JSON.
type JSONSystem interface {
	ToJSON(data Data) (string, error)
	FromJSON(input string) (Data, error)
}

// JSONHandler Класс, реализующий работу с JSON.
type JSONHandler struct{}

func (j JSONHandler) ToJSON(data Data) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("main.go ToJSON(...) json.Marshal(data): %w", err)
	}

	return string(jsonData), nil
}

func (j JSONHandler) FromJSON(input string) (Data, error) {
	var data Data

	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return Data{}, fmt.Errorf("main.go FromJSON(...) json.Unmarshal(...): %w", err)
	}

	return data, nil
}

// XMLSystem Интерфейс для работы с XML.
type XMLSystem interface {
	ToXML(data Data) (string, error)
	FromXML(input string) (Data, error)
}

// XMLHandler Класс, реализующий работу с XML.
type XMLHandler struct{}

func (x XMLHandler) ToXML(data Data) (string, error) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("main.go ToXML(...) xml.Marshal(data): %w", err)
	}

	return string(xmlData), nil
}

func (x XMLHandler) FromXML(input string) (Data, error) {
	var data Data

	err := xml.Unmarshal([]byte(input), &data)
	if err != nil {
		return Data{}, fmt.Errorf("main.go FromXML(...) xml.Unmarshal(...): %w", err)
	}

	return data, nil
}

// Adapter Адаптер для преобразования JSON в XML и наоборот.
type Adapter struct {
	jsonHandler JSONSystem
	xmlHandler  XMLSystem
}

func (a Adapter) JSONToXML(jsonStr string) (string, error) {
	data, err := a.jsonHandler.FromJSON(jsonStr)
	if err != nil {
		return "", fmt.Errorf("main.go JSONToXML(...) a.jsonHandler.FromJSON(jsonStr): %w", err)
	}

	xmlStr, err := a.xmlHandler.ToXML(data)
	if err != nil {
		return "", fmt.Errorf("main.go JSONToXML(...) a.xmlHandler.ToXML(data): %w", err)
	}

	return xmlStr, nil
}

func (a Adapter) XMLToJSON(xmlStr string) (string, error) {
	data, err := a.xmlHandler.FromXML(xmlStr)
	if err != nil {
		return "", fmt.Errorf("main.go XMLToJSON(...) a.xmlHandler.FromXML(xmlStr): %w", err)
	}

	jsonStr, err := a.jsonHandler.ToJSON(data)
	if err != nil {
		return "", fmt.Errorf("main.go XMLToJSON(...) a.jsonHandler.ToJSON(data): %w", err)
	}

	return jsonStr, nil
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	jsonHandler := JSONHandler{}
	xmlHandler := XMLHandler{}
	adapter := Adapter{jsonHandler: jsonHandler, xmlHandler: xmlHandler}

	jsonData := `{"name":"example", "value":"1234"}`

	xmlData, err := adapter.JSONToXML(jsonData)
	if err != nil {
		log.Infoln("Ошибка преобразования JSON в XML:", err)
	} else {
		log.Infoln("JSON в XML:", xmlData)
	}

	xmlData = `<Data><name>example</name><value>1234</value></Data>`

	jsonData, err = adapter.XMLToJSON(xmlData)
	if err != nil {
		log.Infoln("Ошибка преобразования XML в JSON:", err)
	} else {
		log.Infoln("XML в JSON:", jsonData)
	}
}
