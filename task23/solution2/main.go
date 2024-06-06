package main

import (
	"github.com/sirupsen/logrus"
)

// removeAtIndexCopy removes an element at index i from a slice s without modifying the original slice.
func removeAtIndexCopy(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return append([]int(nil), s...) // Return a copy of the original slice if index is out of range
	}

	result := make([]int, len(s)-1)

	copy(result, s[:i])       // Copy elements before index i
	copy(result[i:], s[i+1:]) // Copy elements after index i

	return result
}

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})

	slice := []int{1, 2, 3, 4, 5}

	log.Infoln("Original slice:", slice)

	index := 2 // Element to remove (0-based index)
	modifiedSlice := removeAtIndexCopy(slice, index)
	log.Infoln("Modified slice:", modifiedSlice)
}
