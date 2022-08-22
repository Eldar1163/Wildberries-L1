package main

import "fmt"

// Функция возвращает позицию заданного элемента в массиве, если элемент не представлен, то возвращает -1
func position(arr []int, elem int) int {
	for ind, val := range arr {
		if val == elem {
			return ind
		}
	}
	return -1
}

// Функция строит пересечение двух множеств
func intersection(set1 []int, set2 []int) []int {
	//Результирующее множество
	var result []int
	//Для обмена двух массивов
	var buf []int
	//Обмен массивов, для того чтобы итерироваться по более короткому
	if len(set2) < len(set1) {
		buf = set1
		set1 = set2
		set2 = buf
	}
	// Итерируемся по короткому массиву, если его элемент есть и в другом массиве, но при этом
	// его еще нет в результирующем, то осуществить добалвение этого элемента в результирующий
	for _, elem := range set1 {
		if position(set2, elem) >= 0 && position(result, elem) == -1 {
			result = append(result, elem)
		}
	}

	return result
}

func main() {
	set1 := []int{4, 2, 6, 3, 7, 4}
	set2 := []int{8, 4, 3, 7, 9, 5, 1}

	fmt.Println("Set 1 =", set1)
	fmt.Println("Set 2 =", set2)
	intersects1s2 := intersection(set1, set2)
	fmt.Println("Intersection of set1 and set2 =", intersects1s2)
}
