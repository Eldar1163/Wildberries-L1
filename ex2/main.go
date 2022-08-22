package main

import (
	"fmt"
	"time"
)

func evalSquare(num int) {
	fmt.Println(num * num)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	for _, num := range arr {
		go evalSquare(num) // Создание горутины, выполняющей evalSquare
	}

	_, err := fmt.Scanln() // Задержка выполнения main, для того, чтобы горутины успели выполниться
	if err != nil {
		return
	}
}
