package main

import (
	"fmt"
	"time"
)

func input(n *int) error {
	fmt.Print("Enter working time in seconds: ")
	_, err := fmt.Scan(n)
	fmt.Println()
	return err
}

func main() {
	var n int

	for input(&n) != nil { // ввод количества секунд работы программы
	}

	start := time.Now() // Начальная временная точка отсчета

	ch := make(chan int)
	go func() {
		for {
			fmt.Println(<-ch)                  // Вывод данных из канала
			time.Sleep(400 * time.Millisecond) // Задежка в 400 миллисекунд
		}
	}()

	series := 1 // Пример последовательных данных

	for ; ; series++ {
		if time.Now().Sub(start) >= time.Duration(n)*time.Second { // Проверка достижения конца времени работы программы
			return
		}
		ch <- series // Отправка данных в канал
	}
}
