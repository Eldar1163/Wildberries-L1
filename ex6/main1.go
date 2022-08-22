package main

import (
	"fmt"
	"math/rand"
	"time"
)

func goroutine1(ch chan struct{}) {
	fmt.Println("Goroutine was started")
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		// Если канала закрыт, то завершаем горутину
		_, opened := <-ch
		if !opened {
			fmt.Println("Goroutine was stopped")
			return
		}
		fmt.Println(rnd.Intn(65))
		time.Sleep(400 * time.Millisecond)
	}
}

func main() {
	ch := make(chan struct{})
	go goroutine1(ch)

	time.Sleep(5 * time.Second)
	close(ch) // Закрываем канал, чтобы завершить горутину

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
