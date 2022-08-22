package main

import "fmt"

// Функция создает из последовательности - множество, которое представляется map
func createSet(arr []string) map[string]int {
	// Хранить множество будем в map
	set := make(map[string]int)

	// Запись элементов в map, исключая повторы за счет механизма работы map
	for ind, val := range arr {
		set[val] = ind
	}
	return set
}

// Функция печатает множество
func printSet(set map[string]int) {
	fmt.Print("{ ")
	for key, _ := range set {
		fmt.Print(key + " ")
	}
	fmt.Print("}")
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	printSet(createSet(sequence))
}
