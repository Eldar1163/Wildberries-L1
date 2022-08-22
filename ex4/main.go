package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

/*
В данной программе, в качестве способа остановки работы горутин была выбрано -
завершение работы main, т.к. в данном случае это удобно, из-за того, что
кроме записи в поток в main больше ничего не делается. А так можно было
воспользоваться каналом для организации остановки горутин.
*/
func input(n *int) error {
	fmt.Print("Enter worker count: ")
	_, err := fmt.Scan(n)
	fmt.Println()
	return err
}

func worker(id int, ch chan int) {
	for {
		fmt.Printf("Worker %d, value = %d", id, <-ch)
		fmt.Println()
	}
}

func main() {
	var n int

	for input(&n) != nil { // ввод числа воркеров, пока не будет введено целое число
	}

	ch := make(chan int)                 // канал для передачи случайных чисел в горутины
	osSigCh := make(chan os.Signal, 1)   // канал для обработки комбинации ctrl+C
	signal.Notify(osSigCh, os.Interrupt) // принимает внешние сигналы в канал osSigCh

	for i := 0; i < n; i++ {
		go worker(i, ch) // запуск воркеров
	}

	rnd := rand.New(rand.NewSource(time.Now().Unix())) // инициализация генератора случайных чисел

	for {
		select {
		case <-osSigCh: // обработка ctrl+C
			return
		default:
		}
		ch <- rnd.Intn(65)
		time.Sleep(400 * time.Millisecond)
	}
}
