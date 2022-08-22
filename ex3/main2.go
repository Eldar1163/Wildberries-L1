package main

import (
	"fmt"
)

func evalSquare2(num int, arr []int, ind int, ch chan bool) {
	arr[ind] = num * num // Вычисление квадрата числа и сохранение результата
	ch <- true           // Отправка данных в канал, сигнализирующая о готовности результата данного экземпляра горутины
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	partRes := make([]int, len(arr))
	ch := make(chan bool) // создание канала, для реализации блокировки main, пока все горутины не выполнятся
	for ind, num := range arr {
		go evalSquare2(num, partRes, ind, ch) // Создание горутины, выполняющей evalSquare
	}

	for i := 0; i < len(arr); i++ {
		_ = <-ch // Пока горутина не отправит данные в канал, исполнение остановится здесь и так для каждой горутины
	}
	// Суммирования частичных результатов
	sum := 0
	for _, part := range partRes {
		sum += part
	}

	fmt.Printf("Сумма квадратов = %d", sum)
}
