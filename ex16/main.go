package main

import "fmt"

// Быстрая сортировка arr - сортируемый массив, first - начальный индекс, last - конечный индекс

func QuickSort(arr []int, first int, last int) {
	//Два указателя i И j требуются для разделения массива на 2 части
	i := first
	j := last
	// Опорный элемент
	p := arr[(i+j)/2]

	for {
		// Оставляем все элементы меньшие опорного слева
		for arr[i] < p {
			i++
		}

		// Оставляем все элементы больший опорного справа
		for arr[j] > p {
			j--
		}

		// Обмен найденного большего с меньшим, чтобы не нарушить разделение
		if i <= j {
			buf := arr[i]
			arr[i] = arr[j]
			arr[j] = buf
			// Чтобы не топтаться на месте
			i++
			j--
		}

		// Постусловие выхода из цикла разделения массива на 2 части
		if i >= j {
			break
		}
	}

	// Сортировка левой части
	if first < j {
		QuickSort(arr, first, j)
	}

	// Сортировка правой части
	if i < last {
		QuickSort(arr, i, last)
	}
}

func main() {
	arr := []int{5, 2, 7, 4, 9, 3, 1, 3, 8}
	fmt.Println("Not sorted:\n", arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted:\n", arr)
}
