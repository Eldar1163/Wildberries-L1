package main

import (
	"fmt"
	"sync"
)

func evalSquare3(num int, arr []int, ind int, group *sync.WaitGroup) {
	arr[ind] = num * num // Вычисление квадрата числа и сохранение результата
	group.Done()         // Указываем, что горутина завершила свое выполнение
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	partRes := make([]int, len(arr))
	wg := sync.WaitGroup{} //Создание WaitGroup для реализации задержки на время выполнения всех горутин
	wg.Add(len(arr))
	for ind, num := range arr {
		go evalSquare3(num, partRes, ind, &wg) // Создание горутины, выполняющей evalSquare
	}

	wg.Wait() // Ждем, пока выполнятся все горутины группы

	// Суммирования частичных результатов
	sum := 0
	for _, part := range partRes {
		sum += part
	}

	fmt.Printf("Сумма квадратов = %d", sum)
}
