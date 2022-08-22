package main

import (
	"fmt"
)

// Функция реализует бинарный поиск, он работает только на сортированных данных
func binarySearch(arr []int, elem int) int {
	// Начало и конец рассматриваемой области
	f := 0
	l := len(arr) - 1

	for {
		// Всегда предполагаем, что наш элемент посередине рассматриваемой области
		mid := (f + l) / 2
		if arr[mid] == elem {
			return mid
		} else if elem < arr[mid] {
			// Сдвигаем конец влево
			l = mid - 1
		} else {
			// Сдвигаем начало вправо
			f = mid + 1
		}
		// Условие остановки, когда индексы вышли за пределы
		if f < 0 || l >= len(arr) {
			return -1
		}
	}
}

func main() {
	arr := []int{1, 3, 6, 7, 8, 9, 16, 77}
	item := arr[5]
	fmt.Println(arr)
	fmt.Printf("Item = %d, index = %d", item, binarySearch(arr, 9))
}
