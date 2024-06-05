package main

import "github.com/sirupsen/logrus"

// 12. Имеется последовательность строк
// - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	set := createStringSet(strings)

	log.Infof("Исходная последовательность: %v", strings)
	log.Infof("Созданное множество: %v", set)
}

func createStringSet(strings []string) map[string]any {
	set := make(map[string]any) // Для создания множества используем ключи словаря (map).
	for _, str := range strings {
		set[str] = nil // Значения не нужны, потому заполняем их тем, что не требует памяти, например struct{}{}.
	}

	return set
}
