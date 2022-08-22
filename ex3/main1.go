package main

import (
	"fmt"
)

func evalSquare1(num int, arr []int, ind int) {
	arr[ind] = num * num // Вычисление квадрата числа и сохранение результата
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	partRes := make([]int, len(arr))
	for ind, num := range arr {
		go evalSquare1(num, partRes, ind) // Создание горутины, выполняющей evalSquare
	}

	fmt.Println("Нажмите enter")
	_, err := fmt.Scanln() // Задержка выполнения main, для того, чтобы горутины успели выполниться
	if err != nil {
		return
	}

	// Суммирования частичных результатов
	sum := 0
	for _, part := range partRes {
		sum += part
	}

	fmt.Printf("Сумма квадратов = %d", sum)
}
