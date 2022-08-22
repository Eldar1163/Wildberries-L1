package main

import "fmt"

// Функция удаляет из слайса элемент на позиции index, и ничего не делает, в случае некорректного индекса
func removeByIndex(slice []int, index int) []int {
	// Проверка на корректность индекса
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Удаление элемента из слайса
	if index != len(slice)-1 {
		return append(slice[:index], slice[index+1:]...)
	} else {
		return slice[:index]
	}
}

func main() {
	arr := []int{1, 4, 2, 6, 8, 3, 9, 10}
	// Массив до удаления
	fmt.Printf("Before remove: %v\n", arr)

	// Удаление
	arr = removeByIndex(arr, 3)

	// Массим после удаления
	fmt.Printf("After remove: %v\n", arr)
}
